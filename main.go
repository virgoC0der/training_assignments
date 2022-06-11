package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"students/models"
)

const usage = `
Usage of this system:
  add
    add an employee into the system, eg: add id [name]
    add 0001 jack 2022-06-05 security software-engineer
  mod
    modify the employee info by id, eg: mod id [date:YYYY-MM-DD]
    mod id date:2022-06-06
  del
    del employee by id, eg: del id
    del 0001
  show
    checkout employee info by id, eg: show id|[name:alice]
    show name:jack
  list
    checkout all employees in the system, eg: list
  help
    show function that the system can do
  exit
	exit the system`

func main() {
	reader := bufio.NewReader(os.Stdin)

	// init skip list
	models.New()
	for {
		fmt.Print(">> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		text = strings.TrimSuffix(text, "\n")
		textSlice := strings.Split(text, " ")
		switch textSlice[0] {
		case "exit":
			return
		case "add":
			if len(textSlice) < 3 {
				err = errors.New("illegal input")
				fmt.Println(err)
				continue
			}

			id, _ := strconv.Atoi(textSlice[1])
			models.Add(id, "name", textSlice[2])
		case "mod":
			if len(textSlice) < 3 {
				err = errors.New("illegal input")
				fmt.Println(err)
				continue
			}

			id, err := strconv.Atoi(textSlice[1])
			if err != nil {
				models.UpdateByName(textSlice[1], textSlice[2:]...)
				continue
			}
			models.UpdateByID(id, textSlice[2:]...)
		case "show":
			if len(textSlice) < 2 {
				err = errors.New("illegal input")
				fmt.Println(err)
				continue
			}

			id, _ := strconv.Atoi(textSlice[1])
			info := models.Get(id)

			for k, v := range info {
				line := k + ": " + v
				fmt.Println(line)
			}
		case "list":
			var key string
			if len(textSlice) > 1 {
				key = textSlice[1]
			}
			resultMaps := models.List(key)
			for i, r := range resultMaps {
				fmt.Println("---------------", i)
				for k, v := range r {
					line := k + ": " + v
					fmt.Println(line)
				}
			}
		case "del":
			if len(textSlice) < 2 {
				err = errors.New("illegal input")
				fmt.Println(err)
				continue
			}

			id, _ := strconv.Atoi(textSlice[1])
			models.Delete(id)
		case "help":
			fmt.Println(usage)
		default:
			continue
		}
	}
}
