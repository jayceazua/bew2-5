package main

import "fmt"

func main() {
	// option 1 -> var colors map[string]string
	/*option ->*/
	// colors := make(map[string]string)
	/*option 3 -> */
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#F45656",
		"white": "#ffffff",
	}
	// colors["white"] = "#ffffff"
	// delete(colors, "white")
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

/*
Difference between structs and Maps:
Map:
1. all keys must be the same type
2. all values must be the same type
3. keys are indexed - we can iterate over them
4. represent collection of related properties
5. no need to know all the keys st compile time
6. reference type

Struct:
1. values can be of different type
2. kets don't support indexing
3. need to know all the different fields at compile time
4. represent a "thing" with a lot of different properties
5. value type

func printMap(argumentName typeOfMap) {
	for key, value := range argumentName {
		
	}
}
*/