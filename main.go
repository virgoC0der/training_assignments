package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var (
	studentIDMap   map[string]map[string]string
	studentNameMap map[string][]map[string]string
)

const usage = `
Usage of this system:
  add
    add an employee into the system, eg: add id [name]
    add 0001 jack 2022-06-05 security software-engineer
  mod
    modify the employee info by id, eg: mod id [date:xxxx-xx-xx]
    mod id date:2022-06-06
  del
    del employee by id, eg: del id
    del 0001
  show
    checkout employee info by id, eg: show id|[name:xxxxx]
    show name:jack
  list
    checkout all employees in the system, eg: list
  help
    show function that the system can do
  exit
	exit the system`

func main() {
	reader := bufio.NewReader(os.Stdin)
	studentIDMap = make(map[string]map[string]string)
	studentNameMap = make(map[string][]map[string]string)
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
			addStudent(textSlice[1], textSlice[2])
		case "mod":
			if len(textSlice) < 3 {
				err = errors.New("illegal input")
				fmt.Println(err)
				continue
			}
			modStudent(textSlice[1], textSlice[2:]...)
		case "show":
			if len(textSlice) < 2 {
				err = errors.New("illegal input")
				fmt.Println(err)
				continue
			}
			showStudent(textSlice[1])
		case "list":
			listStudent()
		case "del":
			if len(textSlice) < 2 {
				err = errors.New("illegal input")
				fmt.Println(err)
				continue
			}

			deleteStudent(textSlice...)
		case "help":
			fmt.Println(usage)
		default:
			continue
		}
	}
}

func addStudent(id, name string) {
	if _, ok := studentIDMap[id]; ok {
		fmt.Println("student already exists")
		return
	}

	studentIDMap[id] = make(map[string]string)
	studentIDMap[id] = map[string]string{
		"name": name,
	}

	if _, ok := studentNameMap[name]; !ok {
		studentNameMap[name] = make([]map[string]string, 0)
	}
	studentNameMap[name] = append(studentNameMap[name], studentIDMap[id])
}

func modStudent(id string, args ...string) {
	if _, ok := studentIDMap[id]; !ok {
		fmt.Println("student does not exist")
		return
	}

	for _, arg := range args {
		kv := strings.Split(arg, ":")
		if len(kv) < 2 {
			continue
		}
		studentIDMap[id][kv[0]] = kv[1]
	}

	name := studentIDMap[id]["name"]
	for i := 0; i < len(studentNameMap[name]); i++ {
		studentNameMap[name][i] = studentIDMap[id]
	}
}

func showStudent(key string) {
	if student, ok := studentIDMap[key]; ok {
		for k, v := range student {
			line := k + ": " + v
			fmt.Println(line)
		}
		return
	}

	if students, ok := studentNameMap[key]; ok {
		for _, student := range students {
			for k, v := range student {
				line := k + ": " + v
				fmt.Println(line)
			}
		}
	}
}

func listStudent() {
	for _, student := range studentIDMap {
		fmt.Println("---------------")
		for k, v := range student {
			line := k + ": " + v
			fmt.Println(line)
		}
	}
}

func deleteStudent(args ...string) {
	if len(args) == 2 {
		delete(studentIDMap, args[1])
		return
	}

	if len(args) == 3 {
		delete(studentIDMap[args[1]], args[2])
	}
}
