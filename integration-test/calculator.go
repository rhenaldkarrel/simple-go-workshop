package main

import "fmt"

func Add(a int, b int) int {
  return a + b
}

func Multiply(a int, b int) int {
  return a * b
}

func main() {
  fmt.Println("Addition: ", Add(3, 4))
  fmt.Println("Multiplication: ", Multiply(3, 4))
}