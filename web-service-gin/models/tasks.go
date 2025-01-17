package models

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var sqliteDatabase, _ = sql.Open("sqlite3", "./taskManager.db")

type Task struct {
	Id          int    `json:"id"`          // Ignored in payload if not needed
	Title       string `json:"title"`       // Mapped from JSON payload
	Description string `json:"description"` // Mapped from JSON payload
	Created_on  string `json:"created"`     // Can be generated in the backend
	Due_Date    string `json:"due_date"`    // Mapped from JSON payload
	Status      string `json:"status"`      // Mapped from JSON payload
}

func GetAllTasks() ([]Task, error) {

	row, err := sqliteDatabase.Query("SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}
	defer row.Close()

	tasks := make([]Task, 0)
	for row.Next() {
		singletask := Task{}
		var Id int
		var Title string
		var Description string
		var Created_on string
		var Due_Date string
		var Status string

		row.Scan(&Id, &Title, &Description, &Created_on, &Due_Date, &Status)
		singletask.Id = Id
		singletask.Title = Title
		singletask.Description = Description
		singletask.Created_on = Created_on
		singletask.Due_Date = Due_Date
		singletask.Status = Status
		// log.Println("Tasks: ", Id, " ", Title, " ", Description, " ", Created_on, "", Due_Date, "", Status, "")
		tasks = append(tasks, singletask)
	}

	fmt.Println(tasks)
	err = row.Err()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func GetTask(iid int) ([]Task, error) {

	row, err := sqliteDatabase.Query(`SELECT * FROM tasks where id=?`, iid)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	tasks := make([]Task, 0)
	for row.Next() {
		singletask := Task{}
		var Id int
		var Title string
		var Description string
		var Created_on string
		var Due_Date string
		var Status string

		row.Scan(&Id, &Title, &Description, &Created_on, &Due_Date, &Status)
		singletask.Id = Id
		singletask.Title = Title
		singletask.Description = Description
		singletask.Created_on = Created_on
		singletask.Due_Date = Due_Date
		singletask.Status = Status
		tasks = append(tasks, singletask)
	}

	fmt.Println(tasks)
	err = row.Err()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

// AddTask adds a new task to the database
func AddTask(title string, description string, dueDate string, status string) error {

	query := "INSERT INTO tasks (title, description, due_date, status) VALUES (?, ?, ?, ?)"
	_, err := sqliteDatabase.Exec(query, title, description, dueDate, status)

	return err
}

func UpdateTask(task_id int, task Task) error {

	query := "UPDATE tasks SET title = ? , description = ?, due_date = ?, status = ? WHERE id = ?"
	_, err := sqliteDatabase.Exec(query, task.Title, task.Description, task.Due_Date, task.Status, task_id)

	return err
}

func DeleteTask(iid int) error {

	query := "Delete From tasks Where id = ? "
	_, err := sqliteDatabase.Exec(query, iid)

	return err
}
