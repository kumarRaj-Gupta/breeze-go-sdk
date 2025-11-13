package main

import (
	"breeze-go-client/breeze"
	"fmt"
)

func main() {
	breezeClient := breeze.NewBreezeClient("app_key", "secret_key")
	fmt.Println("Breeze Go Client Initialized.")
	fmt.Println(breezeClient.GetLoginURL())
}
