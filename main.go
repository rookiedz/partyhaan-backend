package main

import (
	"crypto/sha256"
	"fmt"
	"partyhaan/backend/services"

	"golang.org/x/crypto/pbkdf2"
)

func main() {
	salt, err := services.GenerateSalt(4)
	if err != nil {
		fmt.Println(err)
	}
	password, err := services.HashPassword(`7eb7db7d1ac8dfcfd17b4769@#$%^&*()\_+!7778bfa982f5752ea0e0fa846de5e38ca0642c79`, salt)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pbkdf2.Key([]byte(`7eb7db7d1ac8dfcfd17b4769@#$%^&*()\_+!7778bfa982f5752ea0e0fa846de5e38ca0642c79`), []byte(salt), 4096, 32, sha256.New))
	fmt.Println(salt)
	fmt.Println(password)
}
