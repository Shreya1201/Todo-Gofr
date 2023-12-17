package main

import (
	"database/sql"
	"fmt"
	"log"

	"gofr.dev/pkg/gofr"

	_ "github.com/go-sql-driver/mysql"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func main() {

	// db connect
	username := "root"
	password := "password"
	host := "localhost"
	port := "3306"
	dbName := "Todo"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database!")

	app := gofr.New()

	// get API to check if it's working or not
	app.GET("/", func(ctx *gofr.Context) (interface{}, error) {
		return "Welcome to todo_list api!", nil
	})

	// post todo for creating a todo item using the POST method
	app.POST("/todo/add", func(ctx *gofr.Context) (interface{}, error) {
		//title := ctx.PathParam("title")
		status := ctx.Param("status")
		title := ctx.Param("title")
		//status := "ctx.Param(status)"

		_, err := db.ExecContext(ctx, "INSERT INTO todos (title, status) VALUES (?, ?)", title, status)

		if err != nil {
			return nil, err
		}

		return nil, nil
	})

	// get all todos for getting all todo items using the GET method
	app.GET("/todos", func(ctx *gofr.Context) (interface{}, error) {
		var todos []Todo

		rows, err := db.QueryContext(ctx, "SELECT * FROM todos")
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var todo Todo
			if err := rows.Scan(&todo.ID, &todo.Title, &todo.Status); err != nil {
				return nil, err
			}

			todos = append(todos, todo)
		}

		return todos, nil
	})

	// put todo
	app.PUT("/todo/update", func(ctx *gofr.Context) (interface{}, error) {
		id := ctx.Param("id")
		status := ctx.Param("status")

		_, err := db.ExecContext(ctx, "UPDATE todos SET status=? WHERE id=?", status, id)
		if err != nil {
			return nil, err
		}

		return nil, nil
	})

	// delete todo by ID using the DELETE method
	app.DELETE("/todo/delete", func(ctx *gofr.Context) (interface{}, error) {
		id := ctx.Param("id")

		_, err := db.ExecContext(ctx, "DELETE FROM todos WHERE id=?", id)
		if err != nil {
			return nil, err
		}

		return nil, nil
	})

	app.Start()
}
