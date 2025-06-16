package main

import "fmt"

type Node struct {
    Value int
    Next  *Node
}

func createList() *Node {
    // Buat linked list: 1 -> 2 -> 3
    // Return pointer ke node pertama
}

func printList(head *Node) {
    // Print semua nilai dalam linked list
}

func main() {
    list := createList()
    printList(list)
}