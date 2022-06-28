package handlers

import (
	"fmt"
	"strconv"

	"students/lib"
	"students/models"
)

type (
	employeeHandlers map[string]func(args ...string) error
)

var handlers employeeHandlers

// Init init handlers
func Init() {
	handlers = make(map[string]func(args ...string) error)

	handlers.Register("add", Add)
	handlers.Register("mod", Mod)
	handlers.Register("show", Show)
	handlers.Register("list", List)
	handlers.Register("delete", Delete)
	handlers.Register("help", Help)
}

// Register registers handler to employeeHandlers
func (h employeeHandlers) Register(key string, handler func(args ...string) error) {
	h[key] = handler
}

// Handle handles logics
func Handle(key string, args ...string) error {
	handler, ok := handlers[key]
	if !ok {
		return lib.ErrHandlerNotFound
	}

	return handler(args...)
}

// Add adds a new user
func Add(args ...string) error {
	if len(args) < 3 {
		fmt.Println(lib.ErrIllegalInput)
		return lib.ErrIllegalInput
	}

	id, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Printf("id[%s] atoi err[%s]\n", args[1], err)
		return err
	}

	if err := models.Add(id, "name", args[2]); err != nil {
		fmt.Printf("add info(id:%d) err: %s\n", id, err.Error())
		return err
	}

	return nil
}

// Mod updates user by id or name
func Mod(args ...string) error {
	if len(args) < 3 {
		fmt.Println(lib.ErrIllegalInput)
		return lib.ErrIllegalInput
	}

	id, err := strconv.Atoi(args[1])
	if err != nil {
		if err := models.UpdateByName(args[1], args[2:]...); err != nil {
			fmt.Printf("update by name[%s] err:%s", args[1], err.Error())
			return err
		}
		return nil
	}

	if err := models.UpdateByID(id, args[2:]...); err != nil {
		fmt.Printf("update by id[%d] err:%s", id, err.Error())
		return err
	}

	return nil
}

// Show displays user information
func Show(args ...string) error {
	if len(args) < 2 {
		fmt.Println(lib.ErrIllegalInput)
		return lib.ErrIllegalInput
	}

	id, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Printf("id[%s] atoi err[%s]\n", args[1], err)
		return err
	}

	info, err := models.Get(id)
	if err != nil {
		fmt.Printf("get by id(%d) err: %s\n", id, err.Error())
		return err
	}

	for k, v := range info {
		line := k + ": " + v
		fmt.Println(line)
	}

	return nil
}

// List displays all users' information
func List(args ...string) error {
	result := make([]map[string]string, 0)

	if len(args) > 1 {
		result = models.List(args[1], args[2], args[3])
	} else {
		result = models.List("", "", "")
	}

	for _, r := range result {
		fmt.Println("---------------")
		for k, v := range r {
			line := k + ": " + v
			fmt.Println(line)
		}
	}

	return nil
}

// Delete removes a user
func Delete(args ...string) error {
	if len(args) < 2 {
		fmt.Println(lib.ErrIllegalInput)
		return lib.ErrIllegalInput
	}

	id, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Printf("id[%s] atoi err[%s]\n", args[1], err)
		return nil
	}

	if err := models.Delete(id); err != nil {
		fmt.Printf("id[%d] not found, err: %s\n", id, err.Error())
		return err
	}

	return nil
}

// Help displays user guide
func Help(args ...string) error {
	fmt.Println(lib.Usage)
	return nil
}
