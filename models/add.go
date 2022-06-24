package models

func Add(id int, k, v string) error {
	return skipList.Add(id, k, v)
}
