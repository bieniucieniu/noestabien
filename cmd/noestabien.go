package main

import (
	"log"

	"github.com/bieniucieniu/noestabien/web/handlers"
)

func main() {
	app := handlers.Root()

	A := handlers.A()

	app.Mount("/A/", A)

	log.Println(A.MountPath())

	app.Listen(":3000")
}
