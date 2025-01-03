package main

import (
	"fmt"
	"os"

	"frontendmasters.com/go/files/fileutils"
)



func main() {
	rootPath, _ := os.Getwd()
	filePath := rootPath + "/data/text.txt"

	c, err := fileutils.ReadTextFile(filePath)
	if err == nil {
		fmt.Println(c)
		newContent := fmt.Sprintf("Original: %v \n Double original: %v%v", c, c, c,)
		fileutils.WriteToFile(filePath + "output.txt", newContent)
	} else {
		fmt.Printf("ERROR PANIC!! %v", err)
	}
}