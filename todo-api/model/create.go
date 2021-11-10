package model

import "fmt"

func CreateTodo() error {
	insertQ, err := con.Query("INSERT INTO TODO VALUES(?,?)", "Mitul", "Coding")
	defer insertQ.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
