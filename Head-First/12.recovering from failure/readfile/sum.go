package readfile

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//OpenFile opens the file and return a pointer to it,
//along with any error encountered
func OpenFile(fn string) (*os.File, error) {
	fmt.Println("Opening", fn)
	return os.Open(fn)
}

//CloseFile closes a file
func CloseFile(f *os.File) {
	fmt.Println("Closing file")
	f.Close()
}

func GetFloats(fn string) ([]float64, error) {
	var nums []float64
	file, err := OpenFile(fn)
	if err != nil {
		return nil, err
	}
	defer CloseFile(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		nums = append(nums, number)
	}
	if scanner.Err() != nil {
		return nil, err
	}
	return nums, nil
}
