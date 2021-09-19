package main

import (
	"fmt"
	"math/rand"
	"time"
)

//난수 추출된 수의 소수 판정 프로그램 v0.3
//소수는 1과 자기 자신 외에 나누어 떨어지지 않음
func main() {
	seed := time.Now().Unix()
	rand.Seed(seed)

	isPrime := true
	number := rand.Intn(150) + 2
	fmt.Println("임의로 추출된 수 : ", number)

	for i := 2; i < number; i++ {
		if number%i == 0 {
			isPrime = false
			//count++
			//count = count+1
		}
	}

	if isPrime == true {
		fmt.Println(number, "는(은) 소수입니다")
	} else {
		fmt.Println(number, "는(은) 소수가 아닙니다")
	}
}
