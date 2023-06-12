package main

/*
抽象工厂：定义了创建不同产品的接口，将实际创建工作交给具体工厂类
1. 创建了一系列相关的对象，无需指定具体类
*/

type ISportFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
}

func GetSportsFactory() ISportFactory {}
