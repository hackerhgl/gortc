package gortc_service_env

import (
	"fmt"

	e "github.com/Netflix/go-env"
)

var env Environment

func Init() {
	_, err := e.UnmarshalFromEnviron(&env)
	if err != nil {
		fmt.Println("CRASHED")
	}
	fmt.Println("env.MYSQL.HOST")
	fmt.Println(env.MYSQL.HOST)
}

func Env() Environment {
	return env
}
