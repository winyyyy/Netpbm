// Package netpbm provides functions to work with PPM, PGM, and PBM image formats.
package netpbm

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// PPM represents a PPM image.
type PPM struct {
	data        [][]Pixel
	width, height int
	magicNumber string
	max          uint
}

// Pixel represents a pixel with R, G, and B components.
type Pixel struct {
	R, G, B uint8
}

// ReadPPM reads a PPM image from a file and returns a struct that represents the image.
func ReadPPM(filename string) (*PPM, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	magicNumber := scanner.Text()

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			continue
		}
		break
	}

	scanner.Scan()
	dimensions := strings.Fields(scanner.Text())
	width, _ := strconv.Atoi(dimensions[0])
	height, _ := strconv.Atoi(dimensions[1])

	scanner.Scan()
	maxValue, _ := strconv.Atoi(scanner.Text())

	var data [][]Pixel
	for y := 0; y < height; y++ {
		scanner.Scan()
		line := scanner.Text()
		var row []Pixel
		for i := 0; i < len(line); i += 3 {
			r, _ := strconv.Atoi(line[i : i+1])
			g, _ := strconv.Atoi(line[i+1 : i+2])
			b, _ := strconv.Atoi(line[i+2 : i+3])
			row = append(row, Pixel{uint8(r), uint8(g), uint8(b)})
		}
		data = append(data, row)
	}

	return &PPM{
		data:        data,
		width:       width,
		height:      height,
		magicNumber: magicNumber,
		max:         uint(maxValue),
	}, nil
}

// Size returns the width and height of the image.
func (ppm *PPM) Size() (int, int) {
	return ppm.width, ppm.height
}

// At returns the value of the pixel at (x, y).
func (ppm *PPM) At(x, y int) Pixel {
	return ppm.data[y][x]
}

// Set sets the value of the pixel at (x, y).
func (ppm *PPM) Set(x, y int, value Pixel) {
	ppm.data[y][x] = value
}

// Save saves the PPM image to a file and returns an error if there was a problem.
func (ppm *PPM) Save(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Fprintf(file, "%s\n%d %d\n%d\n", ppm.magicNumber, ppm.width, ppm.height, ppm.max)

	for _, row := range ppm.data {
		for _, pixel := range row {
			fmt.Fprintf(file, "%d%d%d", pixel.R, pixel.G, pixel.B)
		}
		fmt.Fprintln(file)
	}

	return nil
}

// Invert inverts the colors of the PPM image.
func (ppm *PPM) Invert() {
	for y := 0; y < ppm.height; y++ {
		for x := 0; x < ppm.width; x++ {
			ppm.data[y][x].R = ppm.max - ppm.data[y][x].R
			ppm.data[y][x].G = ppm.max - ppm.data[y][x].G
			ppm.data[y][x].B = ppm.max - ppm.data[y][x].B
		}
	}
}

// ToPGM converts the PPM image to PGM.
func (ppm *PPM) ToPGM() *PGM {
	pgmData := make([][]uint8, ppm.height)

	for y := 0; y < ppm.height; y++ {
		pgmData[y] = make([]uint8, ppm.width)

		for x := 0; ppm.width; x++ {
			grayValue := (uint(ppm.data[y][x].R) + uint(ppm.data[y][x].G) + uint(ppm.data[y][x].B)) / 3
			pgmData[y][x] = uint8(grayValue)
		}
	}

	return &PGM{
		data:        pgmData,
		width:       ppm.width,
		height:      ppm.height,
		magicNumber: "P2", // Choose the appropriate PGM magic number.
		max:         uint8(ppm.max),
	}
}

// ToPBM converts the PPM image to PBM.
func (ppm *PPM) ToPBM() *PBM {
	pbmData := make([][]bool, ppm.height)

	for y := 0; y < ppm.height; y++ {
		pbmData[y] = make([]bool, ppm.width)

		for x := 0; x < ppm.width; x++ {
			threshold := ppm.max / 2
			pbmData[y][x] = (uint(ppm.data[y][x].R) + uint(ppm.data[y][x].G) + uint(ppm.data[y][x].B)) / 3 > threshold
		}
	}

	return &PBM{
		data:        pbmData,
		width:       ppm.width,
		height:      ppm.height,
		magicNumber: "P1", // Choose the appropriate PBM magic number.
	}
}

// DrawSierpinskiTriangle draws a Sierpinski triangle.
func (ppm *PPM) DrawSierpinskiTriangle(n int, start Point, width int, color Pixel) {
	if n <= 0 {
		return
	}

	height := int(float64(width) * math.Sqrt(3) / 2)
	p1 := Point{start.X, start.Y}
	p2 := Point{start.X + width, start.Y}
	p3 := Point{start.X + width / 2, start.Y - height}

	m1 := Point{(p1.X + p2.X) / 2, (p1.Y + p2.Y) / 2}
	m2 := Point{(p2.X + p3.X) / 2, (p2.Y + p3.Y) / 2}
	m3 := Point{(p3.X + p1.X) / 2, (p3.Y + p1.Y) / 2}

	ppm.DrawLine(m1, m2, color)
	ppm.DrawLine(m2, m3, color)
	ppm.DrawLine(m3, m1, color)

	ppm.DrawSierpinskiTriangle(n-1, start, width/2, color)
	ppm.DrawSierpinski
