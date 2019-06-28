package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var file *os.File
	name := os.Args[1]
	content := os.Args[2]
	if _, err := os.Stat(name); err == nil {
		fmt.Printf("%v already exits\n", name)
		return
	}

	fmt.Printf("Creating %v...\n", name)
	file, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}

	file.WriteString(content)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for _, line := range lines {
		fmt.Println(line)
	}

	// fmt.Printf("Removing %v...\n", name)
	// os.Remove(name)

}
