package main

import (
	"fmt"
	"github.com/emirpasic/gods/sets/hashset"
	"os/exec"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello, World!")
	}
	set := hashset.New()
	set.Add(1)
	fmt.Printf("set=%v\n", set)
	out, _ := exec.Command("./go mod tidy").Output()
	fmt.Printf("out=%v\n", out)
	out2, _ := exec.Command("./go install github.com/emirpasic/gods@latest").Output()
	fmt.Printf("out=%v\n", out2)
}
