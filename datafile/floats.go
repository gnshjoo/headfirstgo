// Package datafile은 파일로부터 샘플데이터를 읽어옵니다

package datafile

import (
	"bufio"
	"os"
	"strconv"
)

//GetFloats는 파일의 각 줄을 flat64 타입으로 읽어 옵니다.
func GetFloats(fileName string)  ([3]float64, error) {
	var numbers [3]float64
	file, err := os.Open(fileName)
	if err != nil {
		return numbers, err
	}
	i :=0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}
		i++
	}
	err = file.Close()
	if err != nil {
		return numbers, err
	}
	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}
	return numbers, nil
}