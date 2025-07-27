package pass_import

import (
	"aegis/internal/queries"
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// ImportPasswordsCsv imports passwords from a CSV file into the database.
//
// Args:
//
//	filePath: The path to the CSV file to be imported.
//
// Returns:
//
//	An error if one occurred.
func ImportPasswordsCsv(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("Cannot open CSV: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("Failed to read CSV: %w", err)
	}
	if len(records) < 2 {
		return fmt.Errorf("CSV must contain at least 1 row of data")
	}

	preparedStmt, err := queries.InsertNewPasswordsFromFile()
	if err != nil {
		return fmt.Errorf("Error preparing while preparing statement: %w", err)
	}
	defer preparedStmt.Close()

	writeRecords(records, preparedStmt)

	return nil
}

// writeRecords writes password records to the database.
//
// Args:
//
//	records: A 2D string slice containing the password records.
//	stmt: The prepared statement to use for inserting the records.
//
// Returns:
//
//	An error if one occurred.
func writeRecords(records [][]string, stmt *sql.Stmt) error {
	for _, row := range records[1:] {
		username := row[0]

		hash, err := parseByteArray(row[1])
		if err != nil {
			log.Printf("invalid hash: %v", err)
			continue
		}
		cipher, err := parseByteArray(row[2])
		if err != nil {
			log.Printf("invalid cipher: %v", err)
			continue
		}
		nonce, err := parseByteArray(row[3])
		if err != nil {
			log.Printf("invalid nonce: %v", err)
			continue
		}
		salt, err := parseByteArray(row[4])
		if err != nil {
			log.Printf("invalid salt: %v", err)
			continue
		}

		_, err = stmt.Exec(username, hash, cipher, nonce, salt)
		if err != nil {
			log.Printf("insert error: %v", err)
		}
	}

	return nil
}

// parseByteArray parses a string representation of a byte array into a byte slice.
//
// Args:
//
//	s: The string to parse.
//
// Returns:
//
//	The parsed byte slice and an error if one occurred.
func parseByteArray(s string) ([]byte, error) {
	s = strings.Trim(s, "[]")
	if s == "" {
		return []byte{}, nil
	}

	parts := strings.Fields(s)
	result := make([]byte, len(parts))
	for i, part := range parts {
		n, err := strconv.Atoi(part)
		if err != nil {
			return nil, fmt.Errorf("invalid byte value %q: %w", part, err)
		}
		result[i] = byte(n)
	}
	return result, nil
}
