# 🔐 Password Generator

A simple CLI tool to generate cryptographically secure random passwords, built with Go.

## Features

- Uses `crypto/rand` for secure randomness
- Configurable password length
- Generate multiple passwords at once
- Toggle character types (uppercase, lowercase, digits, special)

## Usage

```bash
go run main.go [flags]
```

### Flags

| Flag          | Default | Description                    |
|---------------|---------|--------------------------------|
| `-length`     | `16`    | Length of the password          |
| `-count`      | `1`     | Number of passwords to generate|
| `-no-upper`   | `false` | Exclude uppercase letters      |
| `-no-lower`   | `false` | Exclude lowercase letters      |
| `-no-digits`  | `false` | Exclude digits                 |
| `-no-special` | `false` | Exclude special characters     |

### Examples

```bash
# Generate a 16-character password with all character types
go run main.go

# Generate a 24-character password
go run main.go -length 24

# Generate 5 passwords at once
go run main.go -count 5

# Alphanumeric only (no special characters)
go run main.go -no-special

# Digits only, 6 characters (like a PIN)
go run main.go -no-upper -no-lower -no-special -length 6
```

## Build

```bash
go build -o password-generator .
./password-generator -length 20
```
