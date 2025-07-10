package queries

import (
	"aegis/internal/crypto"
	"aegis/internal/mpass"
	"crypto/sha256"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"path/filepath"
)

var masterPass = mpass.GetMasterPass()

func InitDB() *sql.DB {
	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal("Could not get user config dir:", err)
	}

	aegisConfigDir := filepath.Join(configDir, "aegis")

	err = os.MkdirAll(aegisConfigDir, 0700)
	if err != nil {
		log.Fatalf("Failed to create config dir: %v", err)
	}

	dbPath := filepath.Join(aegisConfigDir, "pm.sqlite")

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	return db
}

func CreatePasswordsTable(db *sql.DB) {
	createPasswordsTableSQL := `
	CREATE TABLE IF NOT EXISTS pwds (
		username TEXT PRIMARY KEY,
	    password_hash BLOB NOT NULL,
		password_ciphertext BLOB NOT NULL,
		nonce BLOB NOT NULL,
		salt BLOB NOT NULL,
		created_on DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_on DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := db.Exec(createPasswordsTableSQL)
	if err != nil {
		log.Fatalf("Failed to create table %v", err)
	}
}

func hashPassword(password string) []byte {
	h := sha256.Sum256([]byte(password))
	return h[:]
}

func AddNewPassword(db *sql.DB, username, password string) {
	userPassword := []byte(password)
	passwordHash := hashPassword(password)

	p := crypto.NewPasswordManager(userPassword, masterPass)

	cipherText, nonce, salt, err := p.EncryptPassword()
	if err != nil {
		log.Fatalln(err)
	}

	stmt := `
        INSERT INTO pwds (username, password_hash, password_ciphertext, nonce, salt)
        VALUES (?, ?, ?, ?, ?);
    `
	_, err = db.Exec(stmt, username, passwordHash, cipherText, nonce, salt)
	if err != nil {
		log.Fatalf("Password could not be added in the database: %v", err)
	}
}

func FetchPassword(db *sql.DB, username string) string {
	row := db.QueryRow(`SELECT password_ciphertext, nonce, salt FROM pwds WHERE username = ?`, username)

	var cipherText, nonce, salt []byte
	err := row.Scan(&cipherText, &nonce, &salt)
	if err != nil {
		log.Fatalf("Failed to fetch password data: %v", err)
	}

	p := crypto.NewPasswordManager([]byte{}, masterPass)
	pass, err := p.DecryptPassword(cipherText, nonce, salt)
	if err != nil {
		log.Fatalf("Password could not be decrypted: %v", err)
	}

	return string(pass)
}

func FetchUserData(db *sql.DB) ([]map[string]string, error) {
	rows, err := db.Query(`SELECT username, password_ciphertext, password_hash, created_on, updated_on FROM pwds`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []map[string]string

	for rows.Next() {
		var username, createdOn, updatedOn string
		var passwordCiphertext []byte
		var passwordHash []byte

		err := rows.Scan(&username, &passwordCiphertext, &passwordHash, &createdOn, &updatedOn)
		if err != nil {
			return nil, err
		}

		userMap := make(map[string]string)
		userMap["username"] = username
		userMap["password_hash"] = string(passwordHash)
		userMap["created_on"] = createdOn
		userMap["updated_on"] = updatedOn
		userMap["password_ciphertext"] = string(passwordCiphertext)

		results = append(results, userMap)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func DeleteUserByPasswordHash(db *sql.DB, username string) {
	stmt := `DELETE FROM pwds WHERE username = ?`

	result, err := db.Exec(stmt, username)
	if err != nil {
		log.Printf("failed to delete user by username: %s", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("failed to check rows affected: %s", err)
	}

	if rowsAffected == 0 {
		log.Printf("no user found with the given username")
	}
}

func EditUserPassword(db *sql.DB, newPassword, username string) {
	stmt := `
		UPDATE pwds
		SET password_ciphertext = ?, nonce = ?, salt = ?, password_hash = ?, updated_on = datetime('now')
		WHERE username = ?
	`

	userPassword := []byte(newPassword)

	p := crypto.NewPasswordManager(userPassword, masterPass)

	cipherText, nonce, salt, err := p.EncryptPassword()
	if err != nil {
		log.Fatalln(err)
	}

	newPasswordHash := hashPassword(newPassword)

	result, err := db.Exec(stmt, cipherText, nonce, salt, newPasswordHash, username)
	if err != nil {
		log.Printf("failed to update password by username: %s", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("could not check affected rows: %s", err)
	}

	if rowsAffected == 0 {
		log.Printf("no entry found with the given username")
	}
}
