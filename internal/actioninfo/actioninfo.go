package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	// TODO: добавить методы
	Parse(data string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
	for _, val := range dataset {
		err := dp.Parse(val)
		if err != nil {
			log.Printf("Ошибка парсинга данных '%s': %v", val, err)
			continue
		}

		infoString, err := dp.ActionInfo()
		if err != nil {

			log.Printf("Ошибка формирования информации об активности после парсинга '%s': %v", val, err)
			continue
		}

		fmt.Println(infoString)
	}
}
