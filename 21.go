package main

import "fmt"

// Native infrastructure. У нас есть линуксер, который только умееет пользоваться терминалами
// На его дистрибутиве установлен только один терминал - st

type LinuxUser struct{}

func (l LinuxUser) viewDirectory(t Terminal) {
	fmt.Println("Listing directory")
	fmt.Println("Directory content is: ", t.ls())
}

type Terminal interface {
	ls() string
}

type SucklessTerminal struct{}

func (s SucklessTerminal) ls() string {
	return "directory content"
}

// Foreign tool. А ещё у нас есть винда, которую компания закупила, но не обучила линуксеров ею пользоваться

type WindowsExplorer struct{}

func (w WindowsExplorer) openFolder() string {
	return "directory content"
}

// Adapter. Я написал адаптер, который гуи просмотрщик конвертирует в понятный терминал

type WindowsToLinuxAdapter struct {
	explorer WindowsExplorer
}

func (w WindowsToLinuxAdapter) ls() string {
	return w.explorer.openFolder()
}

func main() {
	client := LinuxUser{}

	fmt.Printf("Using adapter:\n\n")

	explorer := WindowsExplorer{}
	terminalEmulator := WindowsToLinuxAdapter{explorer: explorer}
	client.viewDirectory(terminalEmulator)

	fmt.Printf("\nUsing native tools:\n\n")

	terminal := SucklessTerminal{}
	client.viewDirectory(terminal)
}
