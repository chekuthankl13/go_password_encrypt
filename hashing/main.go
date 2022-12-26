package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {

	///////// hashing /////

	// 	password := []byte("")
	// 	// password1 := []byte("")

	// 	hashedpassword, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	fmt.Println(string(hashedpassword))
	// // err =
	// 	err = bcrypt.CompareHashAndPassword(hashedpassword, password)

	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}

	///////////// end ////////////

	////// sha256 encrpty ///////////////////

	s := "-------"

	h := sha256.New()

	h.Write([]byte(s))

	bs := h.Sum(nil)

	fmt.Println(s)

	fmt.Printf("%x\n", bs)

	///////////// end /////////
}
