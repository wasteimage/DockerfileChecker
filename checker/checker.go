package checker

import (
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
)

type Checker interface {
	Validate()
}

func Validate(dockerFilePath string, rows []string) error {
	for _, row := range rows {

		b, err := ioutil.ReadFile(dockerFilePath)
		if err != nil {
			panic(err)
		}
		fmt.Println("reading file...")

		isExist, err := regexp.Match(row, b)
		if err != nil {
			panic(err)
		}
		if !isExist {
			return errors.New(fmt.Sprintf("row '%s' not exist", row))
		}
		fmt.Println(fmt.Sprintf("row '%s' found :)", row))
	}

	fmt.Println("all rows founded! Dockerfile is in the format you need")

	return nil
}
