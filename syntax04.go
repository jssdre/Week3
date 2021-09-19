package main

import (
	"fmt"
	"math/rand"
	"time"
)

//난수 추출된 수의 소수 판정 프로그램 v0.1
//소수는 1과 자기 자신 외에 나누어 떨어지지 않음
func main() {
	// seed 설정
	seed := time.Now().Unix()
	rand.Seed(seed)

	count := 0
	number := rand.Intn(150) + 2 // 0, 1 제외 . 2~152
	fmt.Println("임의로 추출된 수 : ", number)

	for i := 1; i <= number; i++ {
		if number%i == 0 {
			count++
		}
	}

	if count == 2 {
		fmt.Println(number, "는(은) 소수입니다")
	} else {
		fmt.Println(number, "는(은) 소수가 아닙니다")
	}
}
