package todo

import (
	"fmt"
	"strconv"
)

type Todo struct {
	Id    string
	title string
	done  bool
}

var id int

func init() {
	id = 0
}

func New(title string) *Todo {
	return &Todo{
		Id:    strconv.Itoa(id),
		title: title,
		done:  false,
	}
}

func (t *Todo) ToString() string {
	isDone := "□"
	if t.done {
		isDone = "☑"
	}

	return fmt.Sprintf("%s %s %s", isDone, t.Id, t.title)
}

func (t *Todo) Update(title string) {
	if len(title) == 0 {
		t.done = !t.done
	} else {
		t.title = title
	}
}
