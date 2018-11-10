package main

import (
	"fmt"
)

func main() {
	var (
		kapital float64
		zinssatz float64
	)

	fmt.Println("hello zinsrechner")

	fmt.Print("Bitte Kapital eingeben: ")
	fmt.Scan(&kapital)

	fmt.Print("Bitte Zinssatz eingeben: ")
	fmt.Scan(&zinssatz)

	fmt.Printf("\nDer Zins beträgt %f.\n", (kapital * zinssatz / 100))
}
