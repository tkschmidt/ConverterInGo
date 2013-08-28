package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func visit(path string, f os.FileInfo, err error) error {
	if strings.Contains(path, "dat") {
		fmt.Printf(path + "\n")
		getConvert(path, f)
	}
	//fmt.Printf("Visited: %s\n", path)
	return nil
}

func getConvert(path string, f os.FileInfo) {
	file, _ := os.Open(path)
	scanner := bufio.NewScanner(file)
	var firstLine bool = true
	var lineContext string = ""
	fileb, _ := os.Create(strings.Replace(f.Name(), ".dat", "", -1) + ".csv")
	writer := bufio.NewWriter(fileb)
	for scanner.Scan() {
		if firstLine == false {
			lineContext = strings.Replace(strings.Replace(scanner.Text(), " ", "", -1), "2013", "2013 ", -1)
			writer.WriteString(lineContext + "\n")
		} else {
			firstLine = false
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	writer.Flush()
	file.Close()
}

func main() {
	err := filepath.Walk(".", visit)
	fmt.Printf("filepath.Walk() returned %v\n", err)

}
