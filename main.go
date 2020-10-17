package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const path = "/etc/hosts"

func removeLine(lineNumber int, path string) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	info, _ := os.Stat(path)
	mode := info.Mode()

	array := strings.Split(string(file), "\n")
	array = append(array[:lineNumber], array[lineNumber+1:]...)
	ioutil.WriteFile(path, []byte(strings.Join(array, "\n")), mode)
}

func getLineNo(website, path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	// Splits on newlines by default.
	scanner := bufio.NewScanner(f)

	line := 1
	// https://golang.org/pkg/bufio/#Scanner.Scan
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), website) {
			return line, nil
		}

		line++
	}

	if err := scanner.Err(); err != nil {
		// Handle the error
		fmt.Println("File reader error")
	}

	return 0, errors.New("No line found")
}

func addWebsiteToFile(website, path string) {
	f, err := os.OpenFile(path,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	if _, err := f.WriteString("\n127.0.0.1 " + website + "\n"); err != nil {
		fmt.Println(err)
	}
}

func main() {
	args := os.Args
	var website string = args[2]
	if (website == "") {
		fmt.Println("Please enter a website value")
	}

	if len(args) < 2 {
		fmt.Println("Invalid arguments passed")
		return
	}

	switch(args[1]) {
	case "b":
		addWebsiteToFile(website, path)
		fmt.Println("Blacklisted the website. Good. Now go to work. O_o")
		break
	case "ub":
		ln, err := getLineNo(website, path)
		if err != nil || ln == 0 {
			fmt.Println("Operation failed", err)
			return
		}
		removeLine(ln - 1, path)
		fmt.Println("Whilelisted the website, please don't be unproductive!!!")
		break
	default:
		fmt.Println("Please enter valid option")
	}
}
