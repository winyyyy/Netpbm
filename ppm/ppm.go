package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Pixel struct {
	R, G, B uint8
}

type PPM struct {
	data          [][]Pixel
	width, height int
	magicNumber   string
	max           uint
}

// ReadPPM reads a PPM image from a file and returns a struct that represents the image.
func ReadPPM(filename string) (*PPM, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read magic number
	scanner.Scan()
	magicNumber := scanner.Text()

	// Skip comments
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			continue
		}
		break
	}

	// Read dimensions
	scanner.Scan()
	dimensions := strings.Fields(scanner.Text())
	width, _ := strconv.Atoi(dimensions[0])
	height, _ := strconv.Atoi(dimensions[1])

	// Read max value
	scanner.Scan()
	maxValue, _ := strconv.Atoi(scanner.Text())

	// Read pixel data
	var data [][]Pixel
	for y := 0; y < height; y++ {
		scanner.Scan()
		line := scanner.Text()
		var row []Pixel
		values := strings.Fields(line)
		for i := 0; i < len(values); i += 3 {
			r, _ := strconv.Atoi(values[i])
			g, _ := strconv.Atoi(values[i+1])
			b, _ := strconv.Atoi(values[i+2])
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
