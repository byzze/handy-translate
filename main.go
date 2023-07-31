package main

import (
	"lyzee-translate/log"
	"lyzee-translate/mywindow"
	"lyzee-translate/register"
)

func main() {
	log.Init()

	go register.Hook()

	mywindow.Init()
}
