package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"text/tabwriter"
	"time"
)

var width int
var length int

var Col = map[int]string{1: "red", 2: "green", 3: "blue", 4: "orange", 5: "violet", 6: "yellow", 7: "purple", 8: "tartan"}

var chart = make(map[int]string)

func main() {

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Welcome to the Quilting program")
	fmt.Println("Use this program to work out a quilting pattern where no two colours are next to each other")

	fmt.Println("Input the width of your quilt, in panels: ")
	fmt.Scan(&width)

	fmt.Println("Input the length of your quilt, in panels: ")
	fmt.Scan(&length)

	layout(width, length) //send the dimensions to layout func

	printMap(width) // print the layout

	cuttingGuide(chart) // print the cutting list

}

func layout(w int, l int) {

	c := len(Col)
	// fmt.Println("c=:", c) //debugging
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
		n := rand.Intn(max) + min
		c := Col[n]
		// fmt.Println("Rand=", n, "Colour= ", c) //debugging Println
		return c
	}
}

func getRandomColourWithBlacklist(min int, max int, blacklisted []string) string {
	//fmt.Println("RwB Min:", min, "RwB Max:", max, "RwB Blaclisted:", blacklisted) //debugging print
	excluded := map[string]bool{}
	for _, x := range blacklisted {
		excluded[x] = true
	}

	for {
		n := rand.Intn(max) + min
		c := Col[n]
		//fmt.Println(n, c) //debugging
		if !excluded[c] {
			return c
		}
	}

}

func printMap(w int) {
	// fmt.Println(Col) //debugging

	keys := make([]int, 0, len(chart))
	for k := range chart {
		keys = append(keys, k)
	}
	// Sort the keys so that the output is in order
	sortInts(keys)
	// print header
	fmt.Println("")
	fmt.Println("-------------------------------------------------------")
	fmt.Println("##################### Quilt Map #######################")
	fmt.Println("-------------------------------------------------------")
	fmt.Println("")

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

func cuttingGuide(cha map[int]string) {
	//Create a   dictionary of values for each element
	dict := make(map[string]int)
	for _, num := range cha {
		dict[num] = dict[num] + 1
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Println("")
	fmt.Println("------------------------------------")
	fmt.Fprintln(w, "Colour:\t", "Number Required:\t")
	for key, value := range dict {

		fmt.Fprintln(w, key, "\t", value, "\t")
	}
	w.Flush()

}
