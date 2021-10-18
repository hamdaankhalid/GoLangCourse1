package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"encoding/json"
)

/*
Write a program which prompts the user to first enter a name, and then enter an address.
Your program should create a map and add the name and address to the map using the keys “name”
 and “address”, respectively. Your program should use Marshal() to create a JSON object from 
 the map, and then your program should print the JSON object.
*/

func main()  {
	var mapStore = make(map[string]string)
	var name string
	var address string

	var reader *bufio.Reader = bufio.NewReader(os.Stdin)

	fmt.Println("Enter Name:")

	name, _ = reader.ReadString('\n')

    name = strings.Replace(name, "\n", "", -1)

	fmt.Println("Enter Address:")

	address, _ = reader.ReadString('\n')

	address = strings.Replace(address, "\n", "", -1)
	
	mapStore["name"] = name
	mapStore["address"] = address
	
	jsonOutput, err := marshalize(&mapStore)

	if err != nil{
		fmt.Println("Error", err)
	}else{
		fmt.Println(jsonOutput)
	}
}

func marshalize(inputMap *map[string]string) (string, error){
	jsonOutput, err := json.Marshal(*inputMap)

	return string(jsonOutput), err
}
