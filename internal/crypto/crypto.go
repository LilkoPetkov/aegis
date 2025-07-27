package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"golang.org/x/crypto/scrypt"
	"io"
)

type PasswordEncryption interface {
	EncryptPassword() ([]byte, []byte, []byte, error)
	DecryptPassword(ciphertext, nonce, salt []byte) ([]byte, error)
}

type PasswordManager struct {
	uPassword      []byte
	masterPassword []byte
}

// NewPasswordManager creates a new PasswordManager instance.
//
// Args:
//
//	userPassword: The user's password to be encrypted.
//	masterPassword: The master password for encryption.
//
// Returns:
//
//	A new PasswordManager instance.
func NewPasswordManager(userPassword, masterPassword []byte) *PasswordManager {
	return &PasswordManager{
		uPassword:      userPassword,
		masterPassword: masterPassword,
	}
}

// deriveKey derives a key from a password and salt using scrypt.
//
// Args:
//
//	password: The password to derive the key from.
//	salt: The salt to use for key derivation.
//
// Returns:
//
//	The derived key and an error if one occurred.
func deriveKey(password, salt []byte) ([]byte, error) {
	return scrypt.Key(password, salt, 1<<15, 8, 1, 32)
}

// encrypt encrypts the user's password using AES-256-GCM.
//
// Returns:
//
//	The ciphertext, nonce, salt, and an error if one occurred.
func (p PasswordManager) encrypt() ([]byte, []byte, []byte, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return nil, nil, nil, err
	}

	key, err := deriveKey(p.masterPassword, salt)
	if err != nil {
		return nil, nil, nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, nil, nil, err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, nil, nil, err
	}

	ciphertext := aesGCM.Seal(nil, nonce, p.uPassword, nil)
	return ciphertext, nonce, salt, nil
}

// decrypt decrypts the user's password using AES-256-GCM.
//
// Args:
//
//	ciphertext: The encrypted password.
//	nonce: The nonce used for encryption.
//	salt: The salt used for key derivation.
//
// Returns:
//
//	The decrypted password and an error if one occurred.
func (p PasswordManager) decrypt(ciphertext, nonce, salt []byte) ([]byte, error) {
	key, err := deriveKey(p.masterPassword, salt)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

// EncryptPassword encrypts the user's password.
//
// Returns:
//
//	The ciphertext, nonce, salt, and an error if one occurred.
func (p *PasswordManager) EncryptPassword() ([]byte, []byte, []byte, error) {
	return p.encrypt()
}

// DecryptPassword decrypts the user's password.
//
// Args:
//
//	ciphertext: The encrypted password.
//	nonce: The nonce used for encryption.
//	salt: The salt used for key derivation.
//
// Returns:
//
//	The decrypted password and an error if one occurred.
func (p *PasswordManager) DecryptPassword(ciphertext, nonce, salt []byte) ([]byte, error) {
	return p.decrypt(ciphertext, nonce, salt)
}
