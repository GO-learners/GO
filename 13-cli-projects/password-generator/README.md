# Password Generator

A CLI tool to generate secure random passwords using Go.

## Usage

```bash
go run main.go [flags]
```

## Flags

- `-length` — password length (default: 16)
- `-count` — number of passwords (default: 1)
- `-no-upper` — exclude uppercase letters
- `-no-lower` — exclude lowercase letters
- `-no-digits` — exclude digits
- `-no-special` — exclude special characters

## Examples

```bash
go run main.go
go run main.go -length 24 -count 5
go run main.go -no-special
```
