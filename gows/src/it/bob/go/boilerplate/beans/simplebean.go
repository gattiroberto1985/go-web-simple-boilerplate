package beans

import "fmt"

type SimpleBean struct {
	Id    int
	value string
}

func (s *SimpleBean) SayHello() {
	fmt.Printf("Hello!")
}
