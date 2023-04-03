package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// 입력 스트림
var stdin = bufio.NewReader(os.Stdin)

// 값을 입력하는 함수
func InputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)

	// 에러 발생 시 입력 스트림을 비움
	if err != nil {
		stdin.ReadString('\n')
	}

	// 입력 값과 에러 반환
	return n, err
}

func main() {
	// 시간 값을 랜덤 시드로 설정, 프로그램 시작 시 마다 랜덤 값 생성
	rand.Seed(time.Now().UnixNano())

	// 0~99 사이의 랜덤 값 생성
	r := rand.Intn(100)

	// 입력 시도 횟수 확인을 위한 변수 생성
	cnt := 1

	// 숫자 맞추기 게임을 위한 반복문
	for {
		fmt.Printf("숫자 값을 입력하세요 : ")
		n, err := InputIntValue()

		if err != nil {
			// 에러 발생 시 출력
			fmt.Println("숫자만 입력하세요.")
		} else {
			if n < r {
				fmt.Println("더 큰 숫자 입니다.")
			} else if n > r {
				fmt.Println("더 작은 숫자 입니다.")
			} else {
				fmt.Println("맞췄습니다! 시도한 횟수 : ", cnt)
				break
			}
			cnt++
		}
	}
}
