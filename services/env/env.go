package gortc_service_env

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	e "github.com/Netflix/go-env"
)

var env Environment

func Init() {
	loadEnvFromFile()

	_, err := e.UnmarshalFromEnviron(&env)
	if err != nil {
		fmt.Println("CRASHED")
	}
	fmt.Println(env.MYSQL)
}

func E() Environment {
	return env
}

func loadEnvFromFile() {
	file, err := ioutil.ReadFile(".env")
	if err != nil {
		fmt.Println("unable to read file")
	}
	content := string(file)
	if content == "" {
		return
	}
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		kv := strings.Split(line, "=")
		for index, val := range kv {
			kv[index] = strings.TrimSpace(val)
		}
		os.Setenv(kv[0], kv[1])
	}
}
