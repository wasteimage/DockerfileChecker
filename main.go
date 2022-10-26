package main

import (
	"DockerfileChecker/checker"
	"fmt"
)

func main() {
	err := checker.Validate("uebok", []string{"huila", "pedik"})
	if err != nil {
		fmt.Println(err)
	}
}
