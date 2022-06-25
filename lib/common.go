package lib

const Usage = `
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
	if you want to sort by a key,
	list name
  help
    show function that the system can do
  exit
	exit the system`
