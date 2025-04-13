/*
Copyright © 2025 Ben Fleuty <github.com/benfleuty>
*/

package model

import (
	"fmt"
	"time"
)

type Task struct {
	ID          int
	Description string
	Created     time.Time
	Done        bool
}

func (t Task) ToString() string {
	return fmt.Sprintf("%v", t)
}
