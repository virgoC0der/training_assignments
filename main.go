package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"students/handlers"
	"students/skiplist"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// init skip list
	skiplist.Init()
	handlers.Init()
	for {
		fmt.Print(">> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		text = strings.TrimSuffix(text, "\n")
		textSlice := strings.Split(text, " ")

		args := make([]string, 0)
		if len(textSlice) > 1 {
			args = textSlice[1:]
		}
		err = handlers.Handle(textSlice[0], args...)
		if err != nil {
			fmt.Printf("handle commands err: %s\n", err.Error())
			continue
		}
	}
}
