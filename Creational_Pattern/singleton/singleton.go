package main

import (
	"fmt"
	"sync"
)

/*
单例模式：保证一个类只有一个实例，并提供一个全局访问节点
*/

// Singleton 单例模式接口
// 通过这个保证包私有类型的指针被导出
type Singleton interface {
	kim()
}

// 包私有
type singleton struct {
}

func (s *singleton) kim() {

}

var (
	instance *singleton
	once     sync.Once
)

// GetInstance 返回实例
func GetInstance() Singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

func main() {
	instance1 := GetInstance()
	instance2 := GetInstance()
	fmt.Println(instance1 == instance2)

	start := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(10)
	sigInstance := [10]Singleton{}
	for i := 0; i < 10; i++ {
		go func(index int) {
			defer wg.Done()
			// 阻塞，根据信号同时开始
			<-start
			sigInstance[index] = GetInstance()
		}(i)
	}
	close(start)
	wg.Wait()
	for i := 1; i < 10; i++ {
		fmt.Println(sigInstance[i] == sigInstance[i-1])
	}
}
