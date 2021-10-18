package main

import (
	"fmt"
	"strconv"
	"sort"
)

/*
Write a program which prompts the user to enter integers and stores the integers in a sorted slice. 
The program should be written as a loop. Before entering the loop, the program should create an empty
integer slice of size (length) 3. During each pass through the loop, the program prompts the user to 
enter an integer to be added to the slice. The program adds the integer to the slice, sorts the slice,
 and prints the contents of the slice in sorted order. The slice must grow in size to accommodate any
 number of integers which the user decides to enter. The program should only quit (exiting the loop) 
 when the user enters the character ‘X’ instead of an integer.
*/
func main() {
	var storedValues []int = make([]int, 0)
	var userInput string

	for {
		fmt.Println("Enter an Integer or enter X to quit:")

		_, err := fmt.Scan(&userInput)
		
		if(err!=nil){
			fmt.Printf("Error occurred %s", err)
		}

		if userInput == "X"{
			break
		}

		err = updateStore(&storedValues, userInput)
		if err != nil{
			fmt.Println(err)
			continue
		}

		fmt.Println(storedValues)
	}
}

func updateStore(array *[]int, val string) error{
	i, err := strconv.Atoi(val)
	if err != nil {
		return err
	}

	*array = append(*array, i)
	sort.Ints(*array)

	return nil
}