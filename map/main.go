package main

import "fmt"

func main() {
	colors := map[string]string{
		"red": "#FF0000",
		"green": "#00FF00",
		"white": "#FFFFFF",
	}

	colors["black"] = "#000000"

	printMap(colors)

	fmt.Println(colors["red"])
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("hex code for", color, "is", hex)
	}
}
