package main

import (
	"fmt"
	"math/rand"
)

type playera struct {
	name string
	hp   float32
	mp   int
	sta  float32
}
type playerb struct {
	name string
	hp   float32
	mp   int
	sta  float32
}

func main() {
	a := &playera{
		"A角色",
		100,
		0,
		10,
	}
	b := &playerb{
		"B角色",
		300,
		0,
		20,
	}
	for i := 0; i < 10; i++ {
		a.ta(b)
		b.tb(a)
	}
}

func (p *playera) aa(b *playerb) {
	b.hp -= p.sta
}

func (p *playerb) ab(b *playera) {
	b.hp -= p.sta
	if p.hp < 50 {
		p.hp += 10
	}
}

func (p *playera) sa(b *playerb) {
	i := rand.Intn(2)
	if i == 0 {
		p.sa(b)
	}
}

func (p *playerb) sb(b *playera) {
	if p.hp < 50 {
		println("我没有足够的法力值")
	} else {
		b.sta = 0.9 * b.sta
	}
}

func (p *playera) ua() {
	fmt.Printf(
		"*------------------------------* \n"+
			"%v hp:%v mp:%v sta:%v   \n"+
			"*-------------------------------*", p.name, p.hp, p.mp, p.sta)
}
func (p *playerb) ub() {
	fmt.Printf(
		"*------------------------------* \n"+
			"%v hp:%v mp:%v sta:%v   \n"+
			"*-------------------------------*", p.name, p.hp, p.mp, p.sta)
}

func (p *playera) ta(b *playerb) {
	fmt.Printf("请选择你的行动:  1:普攻")
	var i int
	fmt.Scanf("%d", &i)
	switch i {
	case 1:
		p.aa(b)
	}
	p.sa(b)
}

func (p *playerb) tb(b *playera) {
	fmt.Printf("请选择你的行动:  1:普攻; 2:技能")
	var i int
	fmt.Scanf("%d", &i)
	switch i {
	case 1:
		p.ab(b)
	case 2:
		p.sb(b)
	}
}
