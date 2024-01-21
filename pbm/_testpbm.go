package _testpbm.go 

import (
	"fmt"

	"github.com/winyyyy/Netpbm/netpbm"
)

func main() {
	filename := "images/p1.pbm"

	image, err := netpbm.ReadPBM(filename)
	if err != nil {
		fmt.Println("Erreur lors de la lecture de l'image PBM:", err)
		return
	}

	width, height := image.Size()
	fmt.Printf("Taille de l'image: %d x %d\n", width, height)

	value := image.At(1, 1)
	fmt.Printf("Valeur du pixel à la position (1, 1): %t\n", value)

	image.Invert()

	err = image.Save("chemin/vers/votre/image_inverse.pbm")
	if err != nil {
		fmt.Println("Erreur lors de la sauvegarde de l'image inversée:", err)
		return
	}

	fmt.Println("L'image inversée a été sauvegardée avec succès.")
}
