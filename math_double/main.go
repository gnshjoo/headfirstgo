package main

import "fmt"

func double(number int) {
	number *= 2
}

func main() {
	amount := 6
	double(amount)
	fmt.Println(amount)
}

// 매개변수 복사 형태 "pass-by-value" 라고 부르기도한다.
// 함수 내에서 발생한 모든 변경사항은 그 함수 내에서만 유효
