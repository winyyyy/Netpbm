package netpbm

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PBM struct {
	data          [][]bool
	width, height int
	magicNumber   string
}

func ReadPBM(filename string) (*PBM, error) {
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

	var data []bool
	if magicNumber == "P1" {
		for scanner.Scan() {
			line := scanner.Text()
			if strings.HasPrefix(line, "#") {
				continue
			}
			var row []bool
			for _, char := range line {
				if char == '0' {
					row = append(row, false)
				} else if char == '1' {
					row = append(row, true)
				}
			}
			data = append(data, row...)
		}
	} else if magicNumber == "P4" {
		bytes := make([]byte, (width*height+7)/8)
		_, err := file.Read(bytes)
		if err != nil {
			return nil, err
		}
		for _, b := range bytes {
			for i := 7; i >= 0; i-- {
				data = append(data, (b>>uint(i))&1 == 1)
			}
		}
	} else {
		return nil, fmt.Errorf("format PBM non pris en charge: %s", magicNumber)
	}

	return &PBM{
		data:        chunkSlice(data, width),
		width:       width,
		height:      height,
		magicNumber: magicNumber,
	}, nil
}

func (pbm *PBM) Size() (int, int) {
	return pbm.width, pbm.height
}

func (pbm *PBM) At(x, y int) bool {
	return pbm.data[y][x]
}

func (pbm *PBM) Set(x, y int, value bool) {
	pbm.data[y][x] = value
}

func (pbm *PBM) Save(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Fprintf(file, "%s\n%d %d\n", pbm.magicNumber, pbm.width, pbm.height)

	for _, row := range pbm.data {
		for _, pixel := range row {
			if pixel {
				fmt.Fprint(file, "1 ")
			} else {
				fmt.Fprint(file, "0 ")
			}
		}
		fmt.Fprintln(file)
	}

	return nil
}

func (pbm *PBM) Invert() {
	for y := 0; y < pbm.height; y++ {
		for x := 0; x < pbm.width; x++ {
			pbm.data[y][x] = !pbm.data[y][x]
		}
	}
}

func (pbm *PBM) Flip() {
	for y := 0; y < pbm.height; y++ {
		for x := 0; x < pbm.width/2; x++ {
			pbm.data[y][x], pbm.data[y][pbm.width-x-1] = pbm.data[y][pbm.width-x-1], pbm.data[y][x]
		}
	}
}

func (pbm *PBM) Flop() {
	for y := 0; y < pbm.height/2; y++ {
		pbm.data[y], pbm.data[pbm.height-y-1] = pbm.data[pbm.height-y-1], pbm.data[y]
	}
}

func (pbm *PBM) SetMagicNumber(magicNumber string) {
	pbm.magicNumber = magicNumber
}

func chunkSlice(slice []bool, width int) [][]bool {
	var result [][]bool
	for i := 0; i < len(slice); i += width {
		end := i + width
		if end > len(slice) {
			end = len(slice)
		}
		result = append(result, slice[i:end])
	}
	return result
}
