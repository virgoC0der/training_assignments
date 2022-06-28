# employee_management
## Project structure
```
.
├── LICENSE
├── Makefile
├── README.md
├── cover.out
├── employee_management
├── go.mod
├── go.sum
├── lib
│   ├── common.go
│   └── errors.go
├── main.go
├── models
│   ├── add.go
│   ├── add_test.go
│   ├── delete.go
│   ├── delete_test.go
│   ├── get.go
│   ├── get_test.go
│   ├── init.go
│   ├── init_test.go
│   ├── update.go
│   └── update_test.go
├── run.jpg
└── skiplist
└── skiplist.go
```

## Compile
1. Compile binary
```shell
make build 
```

2. Unittest
```shell
make test 
```

## How To Use
```shell
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
if you want to select a key and sort, list [key] [value] [sort_key] 
list name abc id
help
show function that the system can do
exit
exit the system
```


