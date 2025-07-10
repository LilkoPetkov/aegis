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

func NewPasswordManager(userPassword, masterPassword []byte) *PasswordManager {
	return &PasswordManager{
		uPassword:      userPassword,
		masterPassword: masterPassword,
	}
}

func deriveKey(password, salt []byte) ([]byte, error) {
	return scrypt.Key(password, salt, 1<<15, 8, 1, 32)
}

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

func (p *PasswordManager) EncryptPassword() ([]byte, []byte, []byte, error) {
	return p.encrypt()
}

func (p *PasswordManager) DecryptPassword(ciphertext, nonce, salt []byte) ([]byte, error) {
	return p.decrypt(ciphertext, nonce, salt)
}
