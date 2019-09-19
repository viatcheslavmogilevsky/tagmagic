package main

import (
	"bufio"
	"os"
)

// func Loop {

// }

type Context struct {
	searchQuery string
	offset      int
	limit       int
}

func prompt(scanner bufio.Scanner) {

}

func loop() {
	scanner := bufio.NewScanner(os.Stdin)
	for {

		// fmt.Print("Enter Text: ")
		// scanner.Scan()
		// text := scanner.Text()
		// if text != "q" {
		// 	screen.Clear()
		// 	screen.MoveTopLeft()
		// 	fmt.Println(text)
		// 	getTable().Render()
		// } else {
		// 	fmt.Println("Exiting!")
		// 	break
		// }
	}
}
