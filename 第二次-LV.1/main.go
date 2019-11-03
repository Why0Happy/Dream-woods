package main

import "fmt"

type person struct {
	name string
	age int
	gender string
}

type dove interface {
	gugugu()
}

type repeater interface {
	repeat(string)
}

type lemondemon interface {
	lueluelue()
}

type nicemonster interface {
	aima()
}


func (p*person) gugugu() {
	fmt.Println(p.name,"又鸽了")
}

func (p*person) repeat(word string) {
	fmt.Println(word)
}

func(p*person) lueluelue(){
	fmt.Println(p.name,"酸")
}

func(p*person) aima(){
	fmt.Println("艾玛，真香")
}

func main(){
	p:= &person{
		name: "abc",
		age:22,
		gender: "man",
	}
	p.gugugu()
	p.repeat("ignb")
	p.lueluelue()
	p.aima()
}