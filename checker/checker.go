package checker

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

type Checker interface {
	Validate()
}

func Validate(dockerFilePath string, rows []string) error {
	//получает пизды, докер файл и условие валидации
	//ведет лог и возвращает ошибку или возвращает нихуя

	for _, row := range rows {

		f, err := os.Open(dockerFilePath)
		if err != nil {
			return err
		}

		scanner := bufio.NewScanner(f)
		isContains := false

		for scanner.Scan() {
			if strings.Contains(scanner.Text(), row) {
				isContains = true
				break
			}
		}

		if err := scanner.Err(); err != nil {
			return err
		}

		if !isContains {
			return errors.New("иди в пизду")
		}

		f.Close()
	}

	return nil
}
