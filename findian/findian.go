package main

// Write a program which prompts the user to enter a string. The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’. The program should print “Found!” if the entered string starts with the character ‘i’, ends with the character ‘n’, and contains the character ‘a’. The program should print “Not Found!” otherwise. The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.

// Examples: The program should print “Found!” for the following example entered strings, “ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”. The program should print “Not Found!” for the following strings, “ihhhhhn”, “ina”, “xian”.

import (
	"os"
	"fmt"
	"strings"
	"bufio"
)

func main(){
	var input string
	reader := bufio.NewReader(os.Stdin)

	

	fmt.Println("Enter a string")

	input, _ = reader.ReadString('\n')

    input = strings.Replace(input, "\n", "", -1)

	var result = searchString(&input)

	fmt.Println(result)
}

func searchString(in *string) string{
	// search if 'in' contains i at start, contains a, and last char n
	temp := *in
	if((string(temp[0])=="i" || string(temp[0])=="I")  && (strings.Contains(temp, "a") || strings.Contains(temp, "A"))  && (string(temp[len(temp)-1])=="n" || string(temp[len(temp)-1])=="N") ){
		return "Found!"
	} else {
		return "Not found!"
	}
}

