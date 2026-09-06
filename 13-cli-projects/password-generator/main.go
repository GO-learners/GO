package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"os"
	"strings"
)

const (
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits    = "0123456789"
	special   = "!@#$%^&*()-_=+[]{}|;:',.<>?/"
)

func main() {
	length := flag.Int("length", 16, "Length of the password")
	count := flag.Int("count", 1, "Number of passwords to generate")
	noUpper := flag.Bool("no-upper", false, "Exclude uppercase letters")
	noLower := flag.Bool("no-lower", false, "Exclude lowercase letters")
	noDigits := flag.Bool("no-digits", false, "Exclude digits")
	noSpecial := flag.Bool("no-special", false, "Exclude special characters")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Password Generator - Generate secure random passwords\n\n")
		fmt.Fprintf(os.Stderr, "Usage:\n  password-generator [flags]\n\nFlags:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  password-generator                    # 16-char password with all character types\n")
		fmt.Fprintf(os.Stderr, "  password-generator -length 24         # 24-char password\n")
		fmt.Fprintf(os.Stderr, "  password-generator -count 5           # Generate 5 passwords\n")
		fmt.Fprintf(os.Stderr, "  password-generator -no-special        # No special characters\n")
	}

	flag.Parse()

	// Build the character pool based on flags
	var pool strings.Builder
	if !*noLower {
		pool.WriteString(lowercase)
	}
	if !*noUpper {
		pool.WriteString(uppercase)
	}
	if !*noDigits {
		pool.WriteString(digits)
	}
	if !*noSpecial {
		pool.WriteString(special)
	}

	charset := pool.String()
	if len(charset) == 0 {
		fmt.Fprintln(os.Stderr, "Error: all character types excluded, nothing to generate")
		os.Exit(1)
	}

	if *length < 1 {
		fmt.Fprintln(os.Stderr, "Error: password length must be at least 1")
		os.Exit(1)
	}

	for i := 0; i < *count; i++ {
		password, err := generatePassword(*length, charset)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error generating password: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(password)
	}
}

// generatePassword creates a cryptographically secure random password.
func generatePassword(length int, charset string) (string, error) {
	result := make([]byte, length)
	charsetLen := big.NewInt(int64(len(charset)))

	for i := 0; i < length; i++ {
		idx, err := rand.Int(rand.Reader, charsetLen)
		if err != nil {
			return "", fmt.Errorf("crypto/rand failed: %w", err)
		}
		result[i] = charset[idx.Int64()]
	}

	return string(result), nil
}
