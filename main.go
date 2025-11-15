package main

import (
	"breeze-go-client/breeze"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	AppKey := "9Lws0945644628f5PW394813990390h2"
	SecretKey := strings.TrimSpace("956K898494)0p5YF4xfN7bu45S897!M6")
	breezeClient := breeze.NewBreezeClient(AppKey, SecretKey)
	fmt.Println("Breeze Go Client Initialized.")

	loginURL := breezeClient.GetLoginURL()

	// Reading from os.Stdin
	fmt.Println("Open this link 👍")
	fmt.Println(loginURL)
	fmt.Println("Paste the session key named apisession here:")
	cmdlineReader := bufio.NewReader(os.Stdin)
	inputValid := false
	for !inputValid {
		input, err := cmdlineReader.ReadString('\n')
		if err != nil {
			fmt.Println("Couldn't process your input properly. Please retry")
			continue
		}
		input = strings.TrimSpace(input)
		fmt.Println("The session key your provided:", input)
		err = breezeClient.CompleteLogin(input)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Couldn't Process this session key. Please Retry 😥 ")
			continue
		}
		fmt.Println("Retrieved the Session Token Successfully. ")
		inputValid = true
	}
	fmt.Println("Your Holdings:")
	fmt.Println(breezeClient.GetHoldings())
}
