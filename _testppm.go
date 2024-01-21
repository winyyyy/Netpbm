package main

import (
	"fmt"
	"log"
	"netpbm"
)

func main() {

	image, err := netpbm.ReadPPM("exemple.ppm")
	if err != nil {
		log.Fatal(err)
	}

	width, height := image.Size()
	fmt.Printf("Taille de l'image : %d x %d\n", width, height)

	image.Invert()

	err = image.Save("inverse.ppm")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Image inversée enregistrée avec succès.")
}
