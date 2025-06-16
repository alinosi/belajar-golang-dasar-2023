package main

import "fmt"

type Student struct {
    Name  string
    Grade int
}

func main() {
    s1 := Student{"John", 85}
    s2 := &s1
    s3 := s2
    
    s2.Grade = 90
    s3.Name = "Jane"
    
    fmt.Printf("s1: %+v\n", s1)
    fmt.Printf("s2: %+v\n", *s2)
    fmt.Printf("s3: %+v\n", *s3)
}

// Analysis :