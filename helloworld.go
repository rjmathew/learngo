package main
import "fmt"

func main1() {
	fmt.Println("hello world")
    p := new(person)
	p.walk()
//	makewalk(p)
	runwalker(p)
}

func runwalker(w walker) {
	w.walk()
}
func makewalk(b behaviors) {
	b.walk()
}

type person struct {
	name string
}

func (p *person) walk() {
	fmt.Println("person walking")
}

type walker interface {
	walk()
}

type behaviors interface {
    walk()
	talk()
}