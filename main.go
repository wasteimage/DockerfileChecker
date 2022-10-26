package main

import (
	"DockerfileChecker/validator"
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
)

type CacheChecker struct{}

func (c *CacheChecker) Check(dockerFileContent []byte) error {
	validCache := []string{
		"FROM ${CACHE}/",
		"FROM ${CI}/${BASE_IMAGE}:${BASE_IMAGE_TAG}",
	}

	for _, cache := range validCache {
		r, err := regexp.Compile(cache)
		if err != nil {
			panic(err)
		}

		isExist := r.Match(dockerFileContent)

		if isExist {
			fmt.Println(fmt.Sprintf("Row '%s' found :)", cache))
			return nil
		}
	}
	return errors.New("cache not found")
}

func main() {

	dockerFileContent, err := ioutil.ReadFile("aboba")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dockerFileContent))
	var cacheChecker CacheChecker
	valid := validator.New(dockerFileContent)
	valid.AddChecker(&cacheChecker)
	err = valid.Validate()
	if err != nil {
		fmt.Println(err)
	}

}
