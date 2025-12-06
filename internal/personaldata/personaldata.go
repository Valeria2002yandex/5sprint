package personaldata

import "fmt"

type Personal struct {
	// TODO: добавить поля
	Name   string
	Weight int
	Height int
}

func (p Personal) Print() {
	// TODO: реализовать функцию
	fmt.Println("Имя:", p.Name)
	fmt.Println("Вес:", p.Weight)
	fmt.Println("Рост:", p.Height)

}
