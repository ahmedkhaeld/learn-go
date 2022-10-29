package datafile

import (
	"bufio"
	"os"
)

func ReadLines(fn string) ([]string, error) {
	var lines []string
	file, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, err
	}
	return lines, nil
}
