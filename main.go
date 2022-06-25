package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"students/lib"
	"students/models"
)

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

			id, err := strconv.Atoi(textSlice[1])
			if err != nil {
				fmt.Printf("id[%s] atoi err[%s]\n", textSlice[1], err)
				continue
			}

			if err := models.Add(id, "name", textSlice[2]); err != nil {
				fmt.Printf("add info(id:%d) err: %s\n", id, err.Error())
			}
		case "mod":
			if len(textSlice) < 3 {
				err = errors.New("illegal input")
				fmt.Println(err)
				continue
			}

			id, err := strconv.Atoi(textSlice[1])
			if err != nil {
				if err := models.UpdateByName(textSlice[1], textSlice[2:]...); err != nil {
					fmt.Printf("update by name[%s] err:%s", textSlice[1], err.Error())
				}
				continue
			}

			if err := models.UpdateByID(id, textSlice[2:]...); err != nil {
				fmt.Printf("update by id[%d] err:%s", id, err.Error())
			}
		case "show":
			if len(textSlice) < 2 {
				err = errors.New("illegal input")
				fmt.Println(err)
				continue
			}

			id, err := strconv.Atoi(textSlice[1])
			if err != nil {
				fmt.Printf("id[%s] atoi err[%s]\n", textSlice[1], err)
				continue
			}

			info, err := models.Get(id)
			if err != nil {
				fmt.Printf("get by id(%d) err: %s\n", id, err.Error())
				continue
			}

			for k, v := range info {
				line := k + ": " + v
				fmt.Println(line)
			}
		case "list":
			var key, value string
			resultMaps := make([]map[string]string, 0)
			if len(textSlice) > 1 {
				key = textSlice[1]
				value = textSlice[2]
				resultMaps = models.List(key, value, textSlice[3])
			} else {
				resultMaps = models.List("", "", "")
			}

			for _, r := range resultMaps {
				fmt.Println("---------------")
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

			id, err := strconv.Atoi(textSlice[1])
			if err != nil {
				fmt.Printf("id[%s] atoi err[%s]\n", textSlice[1], err)
				continue
			}

			if err := models.Delete(id); err != nil {
				fmt.Printf("id[%d] not found, err: %s\n", id, err.Error())
			}
		case "help":
			fmt.Println(lib.Usage)
		default:
			continue
		}
	}
}
