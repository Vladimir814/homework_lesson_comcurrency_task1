package main

// Конкурентно порахувати суму усіх слайсів int, та роздрукувати результат.
// Приклад:
// [ [ 4, 6 ], [ 7, 9 ] ]
// Результат друку:
// “result: 26”

import "fmt"


func main() {
	// Розкоментуй мене)
	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}
	// Ваша реалізація
	var (
		sum int
		ch  = make(chan int)
	)

	for _, slice := range n {
		for _, v := range slice {
			go func(ch chan int, v int) {
				ch <- v
			}(ch, v)
			sum += <-ch
		}
	}
	close(ch)
	fmt.Printf("result: %d\n", sum)
}