package main

func main() {
	// var colors map[string]string

	// colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4df745",
	}

	colors["blue"] = "#0000ff"

	delete(colors, "red")

	// fmt.Println(colors)
	printMap(colors)
}

func printMap(m map[string]string) {
	for k, v := range m {
		println(k, v)
	}
}
