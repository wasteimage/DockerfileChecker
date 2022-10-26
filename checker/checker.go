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
	if len(rows) < 1 {
		fmt.Println("Please add at least one row!")
		return nil
	}

	for _, row := range rows {

		b, err := ioutil.ReadFile(dockerFilePath)
		if err != nil {
			panic(err)
		}
		fmt.Println("Reading file...")

		isExist, err := regexp.Match(row, b)
		if err != nil {
			panic(err)
		}
		if !isExist {
			return errors.New(fmt.Sprintf("row '%s' not exist", row))
		}
		fmt.Println(fmt.Sprintf("Row '%s' found :)", row))
	}

	fmt.Println("All rows founded! Dockerfile is in the format you need")

	return nil
}
