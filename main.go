package main

import (
	"fmt"
	"log"
	"syscall"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	pwd := getPwd()
	hash := hashAndSalt(pwd)
	fmt.Printf("\nHash: %s\nCompare: %t\n", hash, compare(hash, pwd))
}

func getPwd() []byte {
	fmt.Print("Enter your password: ")

	pwd, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Println(err)
		return nil
	}

	return pwd
}

func hashAndSalt(pwd []byte) []byte {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return hash
}

func compare(hash, pwd []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, pwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
