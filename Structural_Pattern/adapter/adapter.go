package main

import "fmt"

/*
适配器：使接口不兼容的对象能够相互合作
1. 转换对象，使其能够与其他对象进行交互
2. 适配器为被封装对象提供不同的接口，代理模式为对象提供相同的接口，装饰为对象提供强化的接口
*/

// Computer 客户端接口（适配的目标接口）
type Computer interface {
	InsertIntoLightingPort()
}

// Adaptee 被适配的目标接口
type Adaptee interface {
	InsertIntoUSBPort()
}

// Client 客户端代码
type Client struct {
}

func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightingPort()
}

// Mac 服务
type Mac struct{}

func (m *Mac) InsertIntoLightingPort() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}

// Windows 需要适配器使得windows也可以使用lightning接口，然后将他转化成usb
type Windows struct{}

func (m *Windows) InsertIntoUSBPort() {
	fmt.Println("usb connector is plugged into windows machine.")
}

// WindowAdapt 适配器
type WindowAdapt struct {
	Adaptee
}

func (w *WindowAdapt) InsertIntoLightingPort() {
	fmt.Println("convert lightning to usb")
	w.InsertIntoUSBPort()
}

func main() {
	client := &Client{}
	mac := &Mac{}
	client.InsertLightningConnectorIntoComputer(mac)

	windows := &Windows{}
	adaptee := &WindowAdapt{windows}
	client.InsertLightningConnectorIntoComputer(adaptee)

}
