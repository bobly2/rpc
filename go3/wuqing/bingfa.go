package wuqing

import "fmt"

func Bingfa() {

	result := make(chan int)

	go func() {
		sum := 0
		for i := 0; i < 10; i++ {
			sum = sum + i
		}
		result <- sum
	}()
	fmt.Print(<-result)
}
