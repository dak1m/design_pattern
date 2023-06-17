package main

import "fmt"

/*
原型：复制已有对象，并不与他们耦合
1. 克隆过程委派给被克隆对象，得到新的实例，并包含部分预设定配置
2. 所有原型类需要一个通用接口，使得未知具体类的情况也能复制对象
*/

// Inode 原型接口
type Inode interface {
	Print(indentation string)
	Clone() Inode
}

// PrototypeManager 原型管理器
type PrototypeManager struct {
	prototypes map[string]Inode
}

func NewPrototypeManager() *PrototypeManager {
	return &PrototypeManager{prototypes: make(map[string]Inode)}
}

func (p *PrototypeManager) Get(name string) Inode {
	return p.prototypes[name].Clone()
}

func (p *PrototypeManager) Set(name string, node Inode) {
	p.prototypes[name] = node
}

// File 具体原型
type File struct {
	name string
}

func (f *File) Print(indentation string) {
	fmt.Println(indentation + f.name)
}

func (f *File) Clone() Inode {
	return &File{name: f.name + "_clone"}
}

type Folder struct {
	children []Inode
	name     string
}

func (f *Folder) Print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, v := range f.children {
		v.Print(indentation + indentation)
	}
}

func (f *Folder) Clone() Inode {
	cloneFolder := &Folder{name: f.name + "_clone"}
	var tempChildren []Inode
	for _, v := range f.children {
		tempChildren = append(tempChildren, v.Clone())
	}
	cloneFolder.children = tempChildren
	return cloneFolder
}

func main() {
	manager := NewPrototypeManager()

	file1 := &File{name: "file1"}
	file2 := &File{name: "file2"}
	file3 := &File{name: "file3"}

	manager.Set("file1", file1)
	manager.Set("file2", file2)
	manager.Set("file3", file3)

	folder := &Folder{
		children: []Inode{file1, file2, file3},
		name:     "folder",
	}
	manager.Set("folder", folder)

	file1Clone := manager.Get("file1")
	file2Clone := manager.Get("file2")
	file3Clone := manager.Get("file3")
	folderClone := manager.Get("folder")
	for _, clone := range []Inode{file1Clone, file2Clone, file3Clone, folderClone} {
		clone.Print(" ")
	}
}
