package netpbm

import (
	"fmt"
	"testing"
)

func TestPGM(t *testing.T) {
	filename := "images/p1.pbm"

	pgm, err := ReadPGM(filename)
	if err != nil {
		t.Fatalf("Erreur lors de la lecture du fichier PGM: %v", err)
	}

	width, height := pgm.Size()
	fmt.Printf("Largeur: %d, Hauteur: %d\n", width, height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			fmt.Printf("%d ", pgm.At(x, y))
		}
		fmt.Println()
	}

	pgm.Invert()

	err = pgm.Save("images/p1_inverted.pbm")
	if err != nil {
		t.Fatalf("Erreur lors de la sauvegarde du fichier PGM inversé: %v", err)
	}
	fmt.Println("Image inversée sauvegardée avec succès.")
}
