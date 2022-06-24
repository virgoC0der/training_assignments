package models

func Delete(id int) error {
	return skipList.Delete(id)
}
