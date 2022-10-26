package main

import (
	"DockerfileChecker/checker"
	"fmt"
)

func main() {
	//вызывает метод валидате с условиями хуевиями мать ебал
	err := checker.Validate("uebok", []string{"huila", "pedik"})
	if err != nil {
		fmt.Println(err)
	}
}
