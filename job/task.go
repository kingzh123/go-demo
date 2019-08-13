package job

import "fmt"

type Task struct {

}

func (t *Task) Do()  {
	fmt.Println("do job task")
}
