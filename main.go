package main

import (
  "fmt"
)

func hello() {
  fmt.Println("hi")
}

func bye() {
    fmt.Println("bye")
}

func main() {
  fmt.Println(bye())
}
