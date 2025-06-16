package main

import "fmt"

func changePointer(/* parameter di sini */) {
    // Buat pointer menunjuk ke nilai baru (88)
    newValue := 88
    // Ubah pointer agar menunjuk ke newValue
}

func main() {
    x := 42
    ptr := &x
    
    fmt.Println("Sebelum:", *ptr)
    changePointer(/* argumen di sini */)
    fmt.Println("Sesudah:", *ptr)
}