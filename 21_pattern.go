package main

import "fmt"

// USB - ожидаемый интерфейс с методом ConnectWithUSB
type USB interface {
	ConnectWithUSB()
}

// MicroUSB - существующий класс с методом ConnectWithMicroUSB
type MicroUSB struct{}

func (m *MicroUSB) ConnectWithMicroUSB() {
	fmt.Println("Connected with MicroUSB")
}

// Adapter - адаптер, который делает интерфейс MicroUSB совместимым с USB
type Adapter struct {
	microUSB *MicroUSB
}

func (a *Adapter) ConnectWithUSB() {
	// Адаптируем вызов ConnectWithUSB к ConnectWithMicroUSB
	a.microUSB.ConnectWithMicroUSB()
}

func main() {
	// Создаем объект MicroUSB
	microUSBDevice := &MicroUSB{}

	// Используем адаптер, чтобы сделать MicroUSB совместимым с USB
	var usbDevice USB = &Adapter{microUSBDevice}

	// Теперь мы можем использовать метод ConnectWithUSB, который вызывает ConnectWithMicroUSB внутри адаптера
	usbDevice.ConnectWithUSB()
}
