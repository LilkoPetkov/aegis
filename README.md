# 🔐 Aegis — Lightweight Password Manager (Fyne GUI + SQLite)

**Aegis** is a secure, minimalistic, and cross-platform password manager built in [Go](https://golang.org) using the [Fyne GUI toolkit](https://fyne.io/). It stores encrypted passwords locally in a SQLite database and offers a sleek graphical interface for managing entries.

> 🧩 No cloud. No sync. Just fast, local encryption.

---

## ✨ Features

- ✅ Simple graphical user interface (Fyne)
- ✅ Local AES-GCM encryption of stored passwords
- ✅ Password entries stored in `~/.config/aegis/pm.sqlite`
- ✅ Add, edit, and delete password entries
- ✅ Master password-based encryption
- ✅ Dynamic window resizing based on number of entries
- ✅ Cross-platform support: Linux, macOS, Windows

---

## 📦 Installation

### 1. Clone the repository

```bash
git clone https://github.com/LilkoPetkov/aegis.git
cd aegis
