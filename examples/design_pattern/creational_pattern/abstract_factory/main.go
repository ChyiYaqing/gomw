package main

import (
	"fmt"
	"runtime"
)

/*
To implement the Abstract Factory Pattern in Golang.
	1. first define an interface that defines methods for creating related objects.
	2. second implementations of this interface, which are called concrete factories.
*/

type Window string
type Button string
type Menu string

type WidgetFactory interface {
	CreateWindow() Window
	CreateButton() Button
	CreateMenu() Menu
}

type WinFactory struct{}

func (w *WinFactory) CreateWindow() Window {
	return "WinWindow"
}

func (w *WinFactory) CreateButton() Button {
	return "WinButton"
}

func (w *WinFactory) CreateMenu() Menu {
	return "WinMenu"
}

type MacFactory struct{}

func (m *MacFactory) CreateWindow() Window {
	return "MacWindow"
}

func (m *MacFactory) CreateButton() Button {
	return "MacButton"
}

func (m *MacFactory) CreateMenu() Menu {
	return "MacMenu"
}

func main() {
	var factory WidgetFactory
	if runtime.GOOS == "windows" {
		factory = &WinFactory{}
	} else {
		factory = &MacFactory{}
	}
	window := factory.CreateWindow()
	button := factory.CreateButton()
	menu := factory.CreateMenu()

	fmt.Println(window, button, menu)

}
