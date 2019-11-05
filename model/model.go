package model

import (
	"database/sql"

	"github.com/davepmiller/todomvc/types"
)

func GetAll(db *sql.DB) types.Todos {
	sql := "SELECT * FROM public.todo"
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	result := types.Todos{}
	for rows.Next() {
		todo := types.Todo{}
		err2 := rows.Scan(&todo.Id, &todo.Description, &todo.Active)
		if err2 != nil {
			panic(err2)
		}

		result.Todos = append(result.Todos, todo)
	}

	return result
}

func Put(db *sql.DB, description string) (int64, error) {
	sql := "INSERT INTO public.todo(description) VALUES(?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()
	result, err2 := stmt.Exec(description)
	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()
}

func Delete(db *sql.DB, id int) (int64, error) {
	sql := "DELETE FROM public.todo WHERE id = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}

	result, err2 := stmt.Exec(id)
	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()
}
