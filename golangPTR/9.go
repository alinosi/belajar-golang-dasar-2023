package main

import (
	"fmt"
	"reflect"
)

func changePointer(ptr **int) {
	// Buat pointer menunjuk ke nilai baru (88)
	newValue := 88
	*ptr = &newValue
	// Ubah pointer agar menunjuk ke newValue
}

func main() {
	x := 42
	ptr := &x
	a := &ptr


	fmt.Println(reflect.TypeOf(a))

	fmt.Println("Sebelum:", *ptr)
	changePointer(&ptr)
	fmt.Println("Sesudah:", *ptr)
}

// untuk soal ini apakah walau parameter yang diterima adalah pointer-pointer 
// compiler akan langsung mengenali walau kita mengirim hanya pointer? 