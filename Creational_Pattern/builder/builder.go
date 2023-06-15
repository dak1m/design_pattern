package main

import "fmt"

/*
生成器：使用相同的创建代码生成不同类型和形式的对象
1. 主管类：封装一些产品构造顺序，用主管类构造产品，使客户端直接与主管关联，增加复用性
2. 为每种产品创建具体生成类，并实现其构造步骤；且不需要有通用接口，所以相同过程创建不同产品
3. 与工厂区别，工厂直接返回产品，生成器则额外执行一些步骤在返回产品
*/

type IBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
}

func NewDirector(builder IBuilder) *Director {
	return &Director{
		builder: builder,
	}
}

type Director struct {
	builder IBuilder
}

func (d *Director) setBuilder(builder IBuilder) {
	d.builder = builder
}

func (d *Director) Build() {
	d.builder.setWindowType()
	d.builder.setDoorType()
	d.builder.setNumFloor()
}

type House struct {
	windowType string
	doorType   string
	floor      int
}

func NewHouse() *House {
	return &House{}
}

func (h *House) setWindowType() {
	h.windowType = "落地窗"
}

func (h *House) setDoorType() {
	h.doorType = "铁门"
}

func (h *House) setNumFloor() {
	h.floor = 3
}

func (h *House) getHouse() *House {
	return h
}

var _ IBuilder = (*House)(nil)

func NewHouseBlueprint() *HouseBlueprint {
	return &HouseBlueprint{}
}

type HouseBlueprint struct {
	windowType string
	doorType   string
	floor      string
}

func (h *HouseBlueprint) setWindowType() {
	h.windowType = "窗户类型：落地窗"
}

func (h *HouseBlueprint) setDoorType() {
	h.doorType = "门的材质：铁门"
}

func (h *HouseBlueprint) setNumFloor() {
	h.floor = "房间楼层：2楼"
}

func (h *HouseBlueprint) getHouseBlueprint() *HouseBlueprint {
	return h
}

var _ IBuilder = (*HouseBlueprint)(nil)

func main() {
	house := NewHouse()
	houseBlueprint := NewHouseBlueprint()

	director := NewDirector(house)
	director.Build()
	fmt.Printf("%+v", house.getHouse())

	director.setBuilder(houseBlueprint)
	director.Build()
	fmt.Printf("%+v", houseBlueprint.getHouseBlueprint())

}
