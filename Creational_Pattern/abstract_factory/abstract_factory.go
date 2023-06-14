package main

import "fmt"

/*
抽象工厂：定义了创建不同产品的接口，将实际创建工作交给具体工厂类
1. 创建了一系列相关的对象，无需指定具体类
*/

// ISportFactory 抽象工厂接口
type ISportFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
}

func GetSportsFactory(brand string) ISportFactory {
	if brand == "adidas" {
		return &Adidas{}
	}
	if brand == "nike" {
		return &Nike{}
	}
	return &Adidas{}
}

// Adidas 具体工厂
type Adidas struct {
}

func (a Adidas) makeShoe() IShoe {
	return &AdidasShoe{Shoe{
		logo: "adidas",
		size: 10,
	}}
}

func (a Adidas) makeShirt() IShirt {
	return &AdidasShirt{Shirt{
		logo: "adidas",
		size: 9,
	}}
}

var _ ISportFactory = (*Adidas)(nil)

// Nike 具体工厂
type Nike struct {
}

func (n Nike) makeShoe() IShoe {
	return &NikeShoe{Shoe{
		logo: "nike",
		size: 10,
	}}
}

func (n Nike) makeShirt() IShirt {
	return &NikeShirt{Shirt{
		logo: "nike",
		size: 9,
	}}
}

var _ ISportFactory = (*Nike)(nil)

// IShoe 抽象产品
type IShoe interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type Shoe struct {
	logo string
	size int
}

func (s *Shoe) setLogo(logo string) {
	s.logo = logo
}

func (s *Shoe) setSize(size int) {
	s.size = size
}

func (s *Shoe) getLogo() string {
	return s.logo
}

func (s *Shoe) getSize() int {
	return s.size
}

var _ IShoe = (*Shoe)(nil)

// IShirt 抽象产品
type IShirt interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type Shirt struct {
	logo string
	size int
}

func (s *Shirt) setLogo(logo string) {
	s.logo = logo
}

func (s *Shirt) setSize(size int) {
	s.size = size
}

func (s *Shirt) getLogo() string {
	return s.logo
}

func (s *Shirt) getSize() int {
	return s.size
}

var _ IShirt = (*Shirt)(nil)

// AdidasShoe 具体产品
type AdidasShoe struct {
	Shoe
}

// NikeShoe 具体产品
type NikeShoe struct {
	Shoe
}

// AdidasShirt 具体产品
type AdidasShirt struct {
	Shirt
}

// NikeShirt 具体产品
type NikeShirt struct {
	Shirt
}

func main() {
	adidasFactory := GetSportsFactory("adidas")
	nikeFactory := GetSportsFactory("nike")

	adidasShoe := adidasFactory.makeShoe()
	adidasShirt := adidasFactory.makeShirt()

	nikeShoe := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirt()

	fmt.Println(adidasShoe.getLogo())
	fmt.Println(adidasShoe.getSize())
	fmt.Println(nikeShoe.getLogo())
	fmt.Println(nikeShoe.getSize())

	fmt.Println(adidasShirt.getLogo())
	fmt.Println(adidasShirt.getSize())
	fmt.Println(nikeShirt.getLogo())
	fmt.Println(nikeShirt.getSize())

	adidasShoe.setLogo("adidas plus")
	adidasShoe.setSize(99)
	nikeShoe.setLogo("nike plus")
	nikeShoe.setSize(99)

	adidasShirt.setLogo("adidas plus")
	adidasShirt.setSize(99)
	nikeShirt.setLogo("nike plus")
	nikeShirt.setSize(99)

	fmt.Println(adidasShoe.getLogo())
	fmt.Println(adidasShoe.getSize())
	fmt.Println(nikeShoe.getLogo())
	fmt.Println(nikeShoe.getSize())

	fmt.Println(adidasShirt.getLogo())
	fmt.Println(adidasShirt.getSize())
	fmt.Println(nikeShirt.getLogo())
	fmt.Println(nikeShirt.getSize())
}
