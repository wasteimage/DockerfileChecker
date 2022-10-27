package main

import (
	"DockerfileChecker/validator"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"regexp"
)

type CacheChecker struct {
	cache []*regexp.Regexp
}

func NewCacheChecker() *CacheChecker {
	var cacheList []*regexp.Regexp
	validCache := []string{
		`FROM\ +\${CACHE}\/?`,
		`FROM\ +\${CI}\/\${BASE_IMAGE}:\${BASE_IMAGE_TAG}`,
	}

	for _, cache := range validCache {
		r, err := regexp.Compile(cache)
		if err != nil {
			panic(err)
		}

		cacheList = append(cacheList, r)
	}
	return &CacheChecker{cache: cacheList}
}

func (c *CacheChecker) Check(dockerFileContent []byte) error {
	for _, cache := range c.cache {

		isExist := cache.Match(dockerFileContent)

		if isExist {
			fmt.Printf("Row '%v' matches exclude condition", cache)
			return nil
		}
	}

	return errors.New("cache storage usage not found :)")
}

func main() {
	flagDockerFile := flag.String("docker_file_path", "Dockerfile", "Specify docker file path")
	flag.Parse()

	dockerFileContent, err := ioutil.ReadFile(*flagDockerFile)
	if err != nil {
		panic(err)
	}
	cacheChecker := NewCacheChecker()
	valid := validator.New(dockerFileContent)
	valid.AddChecker(cacheChecker)
	err = valid.Validate()
	if err != nil {
		fmt.Println(err)
	}
}
