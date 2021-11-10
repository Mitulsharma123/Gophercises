package model

import "fmt"

func CreateTodo(name, todo string) error {
	insertQ, err := con.Query("INSERT INTO TODO VALUES(?,?)", "Mitul", "Coding")
	// to avoid sql injection added ?
	defer insertQ.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func DeleteTodo(name string) error {
	insertQ, err := con.Query("DELETE FROM TODO WHERE name=?)", name)
	defer insertQ.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
