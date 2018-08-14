package main

import "fmt"
import (
	"bufio"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strconv"
)

func main() {
	var password = "-"
	var cost = 12

	// Use password from args, if provided
	if len(os.Args) > 1 {
		password = os.Args[1]
	}

	// Use cost from args, if provided
	if len(os.Args) > 2 {
		cost = convertStringToInt(os.Args[2])
	}

	if password == "-" {
		password = readFromStdin()
	}

	var hash = bcryptPassword(password, cost)
	fmt.Println(hash)
}

func bcryptPassword(password string, cost int) string {
	var hash, err = bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}

func convertStringToInt(input string) int {
	var i, err = strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return i
}

func readFromStdin() string {
	var reader = bufio.NewReader(os.Stdin)
	fmt.Print("Enter password: ")
	var input, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	// Remove last char, as ReadString returns the input with the delimiter
	return input[:len(input)-1]
}
