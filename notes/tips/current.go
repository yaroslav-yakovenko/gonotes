package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	pwd := "password"
	p, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(p))

}
