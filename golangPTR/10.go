package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

func createList() *Node {
	address := &Node{}
	value1 := Node{1, address}

    return &value1
}

func printList(head *Node) {

}

func main() {
	list := createList()
	printList(list)
}

// saya yakin address tidak merujuk pada node yang sama dengan value1
// mungkin address merujuk pada node asli, tetapi value1 merupakan instance dari struct
// jadi value1 merupakan copy-an dari struct asli
// saya tidak tau Node yang seharusnya dikirim pada saat pembuatan objek adalah alamat dirinya sendiri(ptr variable)
// atau memang alamat dari node utamanya(ptr strutct)

// lalu saya juga tidak tahu cara pembuatan dari structur data linked-list
// biasanya saya akan menggunakan package built-in dari golagn, yaitu package list
// agar mampu mengimplementasikan sebuah sturktur data dalam pemrograman, apa yang harus dikusai terlebih dahulu?
// mungkin saya sudah mempelajarinya tetapi tidak begitu memahaminya sehingga saya tidak dapat membuat implementasi
// secara manual. Saya butuh saran anda agar saya tidak menjadi text book.
