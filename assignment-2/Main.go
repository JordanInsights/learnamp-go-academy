package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	forenameReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your forename: ")
	forename, _ := forenameReader.ReadString('\n')
	// forename = strings.TrimSuffix(forename, "\n")

	middleNameReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your middle name: ")
	middleName, _ := middleNameReader.ReadString('\n')

	surnameReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your surname: ")
	surname, _ := surnameReader.ReadString('\n')

	var finalString string = "Full name: " + forename + middleName + surname

	newLineRegex := regexp.MustCompile(`\r?\n`)
	finalString = newLineRegex.ReplaceAllString(finalString, " ")

	fmt.Printf(finalString)
}
