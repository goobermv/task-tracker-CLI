package models

import "errors"

var (
	ErrNoTitle          = errors.New("title is empty, please input a title for the task")
	ErrNoDescription    = errors.New("Description is empty, please input a description for the task")
	ErrNoTaskWithThisID = errors.New("No task with this ID")
)
