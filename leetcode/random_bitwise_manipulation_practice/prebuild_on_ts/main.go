package main

import "fmt"

const BIT_VECTOR_SIZE = 64

func hash1(n int) int {
	return n * 23 % BIT_VECTOR_SIZE
}
func hash2(n int) int {
	return n * 37 % BIT_VECTOR_SIZE
}
func hash3(n int) int {
	return n * 127 % BIT_VECTOR_SIZE
}

func isContained(arr []int, n int) bool {
	res1 := arr[0]&(1<<hash1(n)) != 0
	res2 := arr[1]&(1<<hash2(n)) != 0
	res3 := arr[2]&(1<<hash3(n)) != 0

	return res1 && res2 && res3
}

func add(arr []int, n int) {
	if !isContained(arr, n) {
		arr[0] |= 1 << hash1(n)
		arr[1] |= 1 << hash2(n)
		arr[2] |= 1 << hash3(n)
	}
}

func main() {
	// Инициализируем наш фильтр (3 числа по 64 бита)
	filter := make([]int, 3)

	// Список ID мест (например, ID из базы данных)
	places := []int{101, 505, 999, 1337}

	fmt.Println("=== Добавление мест ===")
	for _, id := range places {
		fmt.Printf("Добавляем место ID %d\n", id)
		add(filter, id)
	}

	fmt.Println("\n=== Проверка существующих (должно быть true) ===")
	for _, id := range places {
		fmt.Printf("Место %d в истории? %t\n", id, isContained(filter, id))
	}

	fmt.Println("\n=== Проверка новых мест (скорее всего false) ===")
	newPlaces := []int{202, 404, 777}
	for _, id := range newPlaces {
		fmt.Printf("Место %d в истории? %t\n", id, isContained(filter, id))
	}

	fmt.Println("\n=== Состояние фильтра (битовые маски) ===")
	for i, val := range filter {
		fmt.Printf("Массив[%d]: %064b\n", i, uint64(val))
	}
}
