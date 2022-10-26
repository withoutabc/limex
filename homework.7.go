package main

import "fmt"

type Activity1 interface {
	walk()
	crouch() //蹲下
	eat(a int)
}
type Activity2 interface {
	bark()
	Activity1
}
type Animal struct {
	Age           int
	Name          string
	FavouriteFood string
}

func (d *Animal) walk() {
	fmt.Println(d.Name + "开始走路")
}
func (d *Animal) bark() {
	fmt.Println(d.Name + "叫了起来")
}
func (d *Animal) crouch() {
	fmt.Println(d.Name + "蹲下了")
}
func (d *Animal) eat(a int) {
	fmt.Printf("%v吃了%d条%v", d.Name, a, d.FavouriteFood)
}

func main() {
	dog := Animal{Name: "小龙", Age: 6, FavouriteFood: "骨头"}
	dog.bark()
	dog.crouch()

	var cat Activity1 = &Animal{
		Name:          "银东",
		Age:           5,
		FavouriteFood: "鱼"}
	cat.crouch()
	cat.eat(5)

}
