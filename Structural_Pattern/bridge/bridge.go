package main

import "fmt"

/*
桥接模式：
1. 将一个大类或一系列相关的类拆分为抽象和实现两个独立的层次结构
2. 抽象部分包含对实现部分的引用
3. 抽象部分的调用委派给实现部分的对象
*/

// Computer 抽象部分
type Computer interface {
	Print()
	SetPrinter(Printer)
}

// Printer 实施层
type Printer interface {
	PrintFile()
}

type Mac struct {
	printer Printer
}

func (m *Mac) Print() {
	fmt.Println("Print request for Mac")
	m.printer.PrintFile()
	fmt.Println()
}

func (m *Mac) SetPrinter(printer Printer) {
	m.printer = printer
}

type Windows struct {
	printer Printer
}

func (w *Windows) Print() {
	fmt.Println("Print request for windows")
	w.printer.PrintFile()
	fmt.Println()
}

func (w *Windows) SetPrinter(printer Printer) {
	w.printer = printer
}

type Epson struct {
}

func (e *Epson) PrintFile() {
	fmt.Println("Printing by a EPSON Printer")
}

type Hp struct {
}

func (h *Hp) PrintFile() {
	fmt.Println("Printing by a HP Printer")
}

func GetComputer(comType string) Computer {
	switch comType {
	case "mac":
		return &Mac{}
	case "win":
		return &Windows{}
	}
	return nil
}

func main() {
	hpPrinter := &Hp{}
	epsonPrinter := &Epson{}

	macComputer := GetComputer("mac")
	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()

	winComputer := GetComputer("win")
	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()
	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()
}
