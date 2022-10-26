package checker

import (
	"bufio"
	"os"
	"strings"
)

type Checker interface {
	Validate()
}

func Validate(dockerFilePath string, rows []string) error {
	//получает пизды, докер файл и условие валидации
	//ведет лог и возвращает ошибку или возвращает нихуя
	f, err := os.Open(dockerFilePath)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, row := range rows {

		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			if strings.Contains(scanner.Text(), row) {
				break
			}
		}

		if err := scanner.Err(); err != nil {
			return err
		}

	}

	return nil
}
