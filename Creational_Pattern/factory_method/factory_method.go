package main

import (
	"errors"
	"fmt"
)

/*
简单工厂：在不指定具体类的情况下，创建具体产品
1. 在父类中提供一个创建对象的方法，允许子类决定实例化对象
2. 产品具有相同的接口，子类才能返回不同产品
*/

// IGun 产品接口
type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

// Gun 具体产品的父类
type Gun struct {
	name  string
	power int
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) setPower(power int) {
	g.power = power
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) getPower() int {
	return g.power
}

// AK47 具体产品
type AK47 struct {
	Gun
}

func newAK47() IGun {
	return &AK47{
		Gun: Gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}

// 具体产品
type tommy struct {
	Gun
}

func newTommy() IGun {
	return &tommy{
		Gun: Gun{
			name:  "Tommy gun",
			power: 2,
		},
	}
}

// 工厂
func getGun(gunType string) (IGun, error) {
	if gunType == "AK47" {
		return newAK47(), nil
	}
	if gunType == "Tommy" {
		return newTommy(), nil
	}
	return nil, errors.New("Unknown gun type")
}

func main() {
	ak47, _ := getGun("AK47")
	tommy, _ := getGun("Tommy")

	fmt.Printf("Gun:%s, Power:%d\n", ak47.getName(), ak47.getPower())
	fmt.Printf("Gun:%s, Power:%d\n", tommy.getName(), tommy.getPower())
}
