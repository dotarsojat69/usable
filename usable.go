package usable

import (
	"fmt"
	"math"
)

// Max mengembalikan nilai maksimum dari dua bilangan
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Min mengembalikan nilai minimum dari dua bilangan
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// IsPrime memeriksa apakah suatu bilangan merupakan bilangan prima
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// PrintHello mencetak pesan hello dengan nama yang diberikan
func PrintHello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}
