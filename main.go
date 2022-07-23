package main

// Конкурентно порахувати суму кожного слайсу int, та роздрукувати результат.
// Потрібно використовувати WaitGroup.
// Приклад:
// [ [ 4, 6 ], [ 7, 9 ] ]
// Результат друку:
// Порядок друку не важливий.
// “slice 1: 10”
// “slice 2: 16”

import (
	"fmt"
	"sync"
)

func main() {
	// Розкоментуй мене)
	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}

	// Ваша реалізація
	var wg sync.WaitGroup
	for i := 0; i < len(n); i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			sum(i,n[i])
		}(i)
	}

	wg.Wait()
}

func sum(i int,s []int)  {
	sumOfElements := 0
	for _,v := range s {
		sumOfElements += v
	}
	fmt.Printf("slice %d: %d.\n", i, sumOfElements)
}