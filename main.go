package main

import (
    "bufio"
    "fmt"
	"os"
	"github.com/inancgumus/screen"
)


func main() {
	// tm.Clear()
	// tm.MoveCursor(0, 0)

	// chart := tm.NewLineChart(100, 20)
	// data := new(tm.DataTable)
	// data.AddColumn("Time")
	// data.AddColumn("Sin(x)")
	// data.AddColumn("Cos(x+1)")

	// for i := 0.1; i < 10; i += 0.1 {
	// 	data.AddRow(i, math.Sin(i), math.Cos(i+1))
	// }

	// tm.Println(chart.Draw(data))
	// tm.Flush()

    // To create dynamic array
    // arr := make([]string, 0)
    scanner := bufio.NewScanner(os.Stdin)
    for {
        fmt.Print("Enter Text: ")
        // Scans a line from Stdin(Console)
		scanner.Scan()
        // Holds the string that scanned
        text := scanner.Text()
        if text != "q" {
			screen.Clear()
			screen.MoveTopLeft()
            fmt.Println(text)
            // arr = append(arr, text)
        } else {
			fmt.Println("Exiting!")
            break
        }

    }
    // Use collected inputs
    // fmt.Println(arr)
}