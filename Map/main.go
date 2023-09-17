package main

import "fmt"

func main() {
	// Empty maps variants

	// var colors map[string]string
	// colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#4BF745",
		"black": "#000000",
		"blue":  "#32423S",
	}

	colors["white"] = "#FFFFFF"

	delete(colors, "white")

	fmt.Println(colors)
	printMap(colors)
}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
