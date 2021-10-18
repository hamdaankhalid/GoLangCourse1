package main
/*
Write a program which reads information from a file and represents it in a slice of structs. 
Assume that there is a text file which contains a series of names. Each line of the text file
has a first name and a last name, in that order, separated by a single space on the line. 

Your program will define a name struct which has two fields, fname for the first name, and
lname for the last name. Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file. 
Your program will successively read each line of the text file and create 
a struct which contains the first and last names found in the file. 
Each struct created will be added to a slice, and after all lines have been read from the
file, your program will have a slice containing one struct for each line in the file.
After reading all lines from the file, your program should iterate through your slice
of structs and print the first and last names found in each struct.
*/

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

type name struct {
	firstName string
	lastName string
}

func main()  {
	var inputFileName string
	reader := bufio.NewReader(os.Stdin)
	var storedValues []name = make([]name, 0)

	fmt.Println("Please input the name of the text file:")
	inputFileName, _ = reader.ReadString('\n')
    inputFileName = strings.Replace(inputFileName, "\n", "", -1)

	var err error = openAndRead(inputFileName, &storedValues)
	
	if err==nil{
		printStore(&storedValues)
	}else{
		fmt.Println("Error", err)
	}
}

func openAndRead(inputFile string, store *[]name) error{

	file, err := os.Open(inputFile)
    if err != nil {
        return err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
		var nameRes name = stringToName(scanner.Text())
		*store = append(*store, nameRes)
	}

    if err := scanner.Err(); err != nil {
		return err
	}
	
	return nil
}

func printStore(store *[]name) {
	for _, val := range *store {
		fmt.Println(val.firstName, " ", val.lastName)
	}
}

func stringToName(nameString string) name{
	var split []string = strings.Split(nameString, " ")
	var res name = name{ firstName: split[0], lastName: split[1]}
	return res
}
