package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00000",
		"white": "#ffffff",
	}
	colorsNew := make(map[string]string)
	colorsNew["white"] = "#ffff"
	fmt.Println(colors)
	fmt.Println(colorsNew)

	delete(colorsNew, "white")

	fmt.Println(colorsNew)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
