package core

type User struct {
	username string
	password string
	tasks    []Task
}

type Task struct {
	title       string
	description string
	owner       User
}
