/*
Copyright © 2025 Ben Fleuty <github.com/benfleuty>
*/

package model

import "fmt"

type Task struct {
	ID          int
	Description string
	Created     int
	Done        bool
}

func (t Task) ToString() string {
	return fmt.Sprintf("%v", t)
}
