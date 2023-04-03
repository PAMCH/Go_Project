package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//시간 값을 랜덤 시드로 설정, 프로그램 시작 시 마다 랜덤 값 생성
	rand.Seed(time.Now().UnixNano())

	//0~99 사이의 랜덤 값 대입
	n := rand.Intn(100)

	//랜덤 값 출력
	fmt.Println(n)
}
