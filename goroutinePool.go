package main

import (
"fmt"
"github.com/labstack/gommon/log"
"time"
)

type Pool struct {
	JobQueue    chan Job         // 带处理的任务数据
	WorkerCurrentNum     int       // 当前工作的协程数
	MaxWorker   int             // 最大工作协程数
	Result  chan bool
	FinishCallBack  func() error
}

type Job struct{
	ID int
}

type Worker struct {
	Result  chan bool
}

func (w *Worker) DoJob(job Job){
	fmt.Println(job.ID)
	time.Sleep(time.Millisecond * 200)
	w.Result <- true
}


func (g *Pool) SetFinishCallBack(f func() error) {
	g.FinishCallBack  = f
}

// 往Job任务队列里面放入待处理的job
func (g *Pool) AddJob(job Job) {
	g.JobQueue <- job
}

// 开启协程池
func (g *Pool) Run() {
	go g.stop()
	for {
		if g.WorkerCurrentNum < g.MaxWorker {
			select {
			case job, ok := <-g.JobQueue:
				if ok {
					worker := &Worker{g.Result}
					go worker.DoJob(job)
					g.WorkerCurrentNum ++
				}else{
					log.Info("goroutine pool over")
					return
				}
			}
		}
	}
}

func (g *Pool) stop(){
	for {
		<- g.Result
		g.WorkerCurrentNum --
	}
}

func main() {
	jobQueue := make(chan Job)
	resultQueue := make(chan bool)
	p := &Pool{
		MaxWorker: 5,
		JobQueue: jobQueue,
		Result: resultQueue,
	}
	go func (){
		for i:=0; i<10000000000;i++{
			job := Job{i}
			p.AddJob(job)
		}
		close(p.JobQueue)
	}()
	p.Run()
}