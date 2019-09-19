package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/inancgumus/screen"
	"github.com/olekukonko/tablewriter"
)

func getTable() *tablewriter.Table {
	data := [][]string{
		[]string{"A", "The Good", "500"},
		[]string{"B", "The Very very Bad Man", "288"},
		[]string{"C", "The Ugly", "120"},
		[]string{"D", "The Gopher", "800"},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Sign", "Rating"})

	for _, v := range data {
		table.Append(v)
	}

	return table
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter Text: ")
		scanner.Scan()
		text := scanner.Text()
		if text != "q" {
			screen.Clear()
			screen.MoveTopLeft()
			fmt.Println(text)
			getTable().Render()
		} else {
			fmt.Println("Exiting!")
			break
		}
	}
}
