# Aegis Password Manager

A secure, cross-platform password manager built with Go and Fyne, featuring AES-256-GCM encryption and a modern graphical user interface.

## 🔐 Features

### Core Functionality

- **Secure Password Storage**: All passwords are encrypted using AES-256-GCM encryption
- **Master Password Protection**: Single master password protects all stored credentials
- **User-Friendly GUI**: Modern interface built with Fyne framework
- **Cross-Platform**: Runs on Windows, macOS, and Linux

### Password Management

- Add new password entries with username/password pairs
- View stored passwords (with password masking)
- Edit existing password entries
- Delete password entries
- Copy passwords to clipboard with one click

### Import/Export

- **CSV Export**: Export all password data to CSV format
- **CSV Import**: Import passwords from CSV files
- Backup and restore functionality

## 🏗️ Architecture

### Project Structure

```
aegis/
├── internal/
│   ├── crypto/          # Encryption/decryption logic
│   ├── queries/         # Database operations
│   ├── mpass/           # Master password handling
│   ├── pass_import/     # CSV import functionality
│   ├── pass_export/     # CSV export functionality
│   └── ui/              # User interface components
```

### Security Architecture

#### Encryption

- **Algorithm**: AES-256-GCM (Galois/Counter Mode)
- **Key Derivation**: Scrypt with the following parameters:
  - N: 32768 (2^15)
  - r: 8
  - p: 1
  - Key length: 32 bytes (256 bits)
- **Salt**: 16-byte random salt per password entry
- **Nonce**: Unique nonce for each encryption operation

#### Password Security

- **Master Password**: Retrieved from `AEGIS_MASTER_PASS` environment variable
- **Password Hashing**: SHA-256 for password verification
- **Secure Storage**: All sensitive data encrypted at rest

## 🚀 Installation & Setup

### Prerequisites

- Go 1.19 or higher
- Required Go modules:
  - `fyne.io/fyne/v2`
  - `github.com/mattn/go-sqlite3`
  - `golang.org/x/crypto/scrypt`

### Environment Setup

Set your master password as an environment variable:

```bash
export AEGIS_MASTER_PASS="your_secure_master_password"
```

### Database Storage

- **Location**: `~/.config/aegis/pm.sqlite` (Linux/macOS) or equivalent on Windows
- **Type**: SQLite3 database
- **Auto-creation**: Database and tables are created automatically on first run

## 📊 Database Schema

```sql
CREATE TABLE pwds (
    username TEXT PRIMARY KEY,
    password_hash BLOB NOT NULL,
    password_ciphertext BLOB NOT NULL,
    nonce BLOB NOT NULL,
    salt BLOB NOT NULL,
    created_on DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_on DATETIME DEFAULT CURRENT_TIMESTAMP
);
```

## 🖥️ User Interface

### Main Window Features

- **Modern Design**: Gradient backgrounds and intuitive layout
- **Password Cards**: Each stored password displayed as an individual card
- **Action Buttons**: Copy, Edit, and Delete options for each entry
- **Toolbar**: Import, Export, and Add New Password buttons

### Window Components

- **Main View**: Scrollable list of password cards
- **Add Password Dialog**: Form for creating new entries
- **Edit Password Dialog**: Update existing password entries
- **Import/Export Dialogs**: File selection for CSV operations

## 🔄 Import/Export Format

### CSV Structure

The CSV files contain the following columns:

- `username`: Account username/identifier
- `password_hash`: SHA-256 hash of the password
- `password_ciphertext`: AES-encrypted password data
- `nonce`: Encryption nonce (as byte array)
- `salt`: Scrypt salt (as byte array)
- `created_on`: Timestamp of creation
- `updated_on`: Timestamp of last modification

## ⚠️ Disclaimer

This password manager is designed for educational and personal use. While it implements strong cryptographic practices, any password manager should undergo thorough security auditing before use with sensitive data. Always maintain secure backups of your password data.
