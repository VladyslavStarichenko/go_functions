package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	// Створюємо файл з квадратними коренями чисел від 1 до 10
	err := CreateSqrtFile(10, "sqrt.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Знаходимо суму квадратів чисел у файлі
	sum, err := SumOfSquares("sqrt.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Сума квадратів чисел у файлі: %.2f\n", sum)
}

// Функція для створення файлу з квадратними коренями цілих чисел від 1 до n

func CreateSqrtFile(n int, filename string) error {
	// Відкриваємо файл на запис
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	// Записуємо квадратні корені цілих чисел від 1 до n у файл
	for i := 1; i <= n; i++ {
		sqrt := math.Sqrt(float64(i))
		fmt.Fprintf(f, "%.2f\n", sqrt)
	}

	return nil
}

// Функція для знаходження суми квадратів чисел у файлі

func SumOfSquares(filename string) (float64, error) {
	// Відкриваємо файл на читання
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	// Читаємо числа з файлу та обчислюємо суму їх квадратів
	scanner := bufio.NewScanner(f)
	sum := 0.0
	for scanner.Scan() {
		num := 0.0
		_, err := fmt.Sscanf(scanner.Text(), "%f", &num)
		if err != nil {
			return 0, err
		}
		sum += math.Pow(num, 2)
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return sum, nil
}
