package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"log"
)

func checkExt(filePath string) bool {

	splitted := strings.Split(filePath, ".")
	return splitted[len(splitted) -1] == "bf"
}

func getBfContent(filePath string) string {

	data, err := ioutil.ReadFile(filePath)

	if err != nil {
		panic(err)
	}

	return string(data)
}

func optimize(fileCont string) string {

	var reg = regexp.MustCompile(`[^\.|,|\+|<|>|\-|\[|\]]`)
	s := reg.ReplaceAllString(fileCont, "")

	return s
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("No Brainfuck file provided")
		os.Exit(1)
	}

	programPath := os.Args[1]

	if (!checkExt(programPath)) {
		fmt.Println("Please provide a valid Brainfuck file")
		os.Exit(1)
	}

	fileCont := getBfContent(programPath)

	result := optimize(fileCont)

	file, err := os.Create("optimized.bf")

	if err != nil {
		log.Fatal("Could not write optimized file.", err)
	}
	defer file.Close()

	fmt.Fprintf(file, result)
}