package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.Create("./plan.txt")
	file.WriteString("I’m not afraid of difficulties and insist on learning programming")
	buf, _ := os.ReadFile("./plan.txt")
	fmt.Printf("%s\n", buf)
}
