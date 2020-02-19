package main

import (
	"fmt"
	"latihan/main/lib"
	"strconv"
)

func main() {

	var x, y int

	x = 5
	y = 6

	j, k := hitung(x, y)

	fmt.Println("x * y adalah " + strconv.Itoa(j))
	fmt.Println("x + y adalah " + strconv.Itoa(k))

	rs := lib.Cobalib()
	fmt.Println(rs)
}

func hitung(x int, y int) (hasil_kali int, hasil_tambah int) {
	hasil_kali = x * y
	hasil_tambah = x + y

	return hasil_kali, hasil_tambah
}
