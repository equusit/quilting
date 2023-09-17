package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var width int
var length int

var Col = map[int]string{1: "red", 2: "green", 3: "blue", 4: "orange", 5: "violet", 6: "yellow", 7: "purple"}

var chart = make(map[int]string)

func main() {

	// Seed the random number generator only once
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Welcome to the Quilting program")
	fmt.Println("Use this program to work out a quilting pattern where no two colours are next to each other")

	fmt.Println("Input the width of your quilt, in panels: ")
	fmt.Scan(&width)

	fmt.Println("Input the length of your quilt, in panels: ")
	fmt.Scan(&length)

	layout(width, length)

	printMap(width)

}

func layout(w int, l int) {

	c := len(Col)
	n := w * l
	for i := 1; i <= n; i++ {

		chart[i] = getRandomColour(1, c)

		if i > 1 && i < w && chart[i] == chart[i-1] {
			s := []string{chart[i]}
			chart[i] = getRandomColourWithBlacklist(1, c, s)
		} else if i > w && chart[i] == chart[i-1] || chart[i] == chart[i-w] {
			s := []string{chart[i], chart[i-w]}
			chart[i] = getRandomColourWithBlacklist(1, c, s)
		}
	}
}
func getRandomColour(min int, max int) string {
	//fmt.Println("Random min;", min, "Random max", max)
	for {
		n := rand.Intn(max-min) + min
		c := Col[n]
		fmt.Println("Rand=", n, "Colour= ", c)
		return c
	}
}

func getRandomColourWithBlacklist(min int, max int, blacklisted []string) string {
	fmt.Println("RwB Min:", min, "RwB Max:", max, "RwB Blaclisted:", blacklisted)
	excluded := map[string]bool{}
	for _, x := range blacklisted {
		excluded[x] = true
	}

	for {
		n := rand.Intn(max-min) + min
		c := Col[n]
		fmt.Println(n, c)
		if !excluded[c] {
			return c
		}
	}

}

func printMap(w int) {
	fmt.Println(Col)

	keys := make([]int, 0, len(chart))
	for k := range chart {
		keys = append(keys, k)
	}
	// Sort the keys so that the output is in order
	sortInts(keys)

	// Create rows with a fixed number of columns
	for i := 0; i < len(keys); i += w {
		end := i + w
		if end > len(keys) {
			end = len(keys)
		}
		row := make([]string, w)
		for j, k := range keys[i:end] {
			row[j] = chart[k]
		}
		fmt.Println(strings.Join(row, "\t"))
	}
}

func sortInts(slice []int) {
	for i := 0; i < len(slice)-1; i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
}
