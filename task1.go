package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sortRandNumSlice(randNumSlice []int, s uint8) []int {
	for i := 1; i < len(randNumSlice); i++ {
		t := randNumSlice[i]
		j := i
		switch s {
		case 1:
			for j > 0 && randNumSlice[j-1] > t {
				randNumSlice[j] = randNumSlice[j-1]
				j = j - 1
			}
		case 2:
			for j > 0 && randNumSlice[j-1] < t {
				randNumSlice[j] = randNumSlice[j-1]
				j = j - 1
			}
		default:
			break
		}
		randNumSlice[j] = t
	}
	return randNumSlice
}

func main() {
	var l int
	var randNumSlice []int

	fmt.Printf("\nВведите количество целых чисел = ")
	_, err := fmt.Scanln(&l)
	if err != nil {
		return
	} else {
		rand.Seed(time.Now().UnixNano())
		for i := 0; i < l; i++ {
			randNumSlice = append(randNumSlice, rand.Intn(l))
		}
	}

	fmt.Println("Набор случайных чисел из", l, "элементов")
	fmt.Println(randNumSlice)

	if len(randNumSlice) < 2 {
		fmt.Println("Айдын, нет смысла сортировки (") // Извините, так просто для позитива
	} else {
		var s uint8
		fmt.Printf("\nЗадача принято, а как отсортировать (1 - по возрастанию, 2 - по убыванию)? ")
		_, errS := fmt.Scanln(&s)
		if errS != nil || s > 2 || s == 0 {
			return
		} else {
			switch s {
			case 1:
				fmt.Printf("\nОтсортированные числа по возрастанию:\n")
			case 2:
				fmt.Printf("\nОтсортированные числа по убыванию:\n")
			default:
				break
			}

			fmt.Println(sortRandNumSlice(randNumSlice, s))
			fmt.Printf("\n")
		}
	}
}
