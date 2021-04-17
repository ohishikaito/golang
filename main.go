package main

import (
	"config"
	"fmt"
)

func main() {
	fmt.Println(config.Config.Apikey)
	fmt.Println(config.Config.ApiSecret)
}

func init() {
	fmt.Println("-------------------------- init! --------------------------")
}
