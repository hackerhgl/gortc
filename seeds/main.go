package main

import (
	"fmt"
	env "gortc/services/env"
	mysql "gortc/services/mysql"
	"os"
)

var CLEAN string = "--clean"
var SEED string = "--seed"

func main() {

	if len(os.Args) < 2 {
		fmt.Println("No command found")
		return
	}

	command := os.Args[1]

	env.Init()
	mysql.Connect()

	if command == CLEAN {
		Clean()
	} else if command == SEED {
		Seed()
	} else {
		fmt.Println("Invalid command")
	}

}
