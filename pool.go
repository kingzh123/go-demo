package main

import (
	"fmt"
	"go-demo/job"
)

type Jobs struct{
	Id int //任务id
	task interface{}
}

type Workers struct {
	Result chan bool //工作状态队列
}

//工作者执行任务
func (w *Workers) DoJob(j Jobs){
	fmt.Printf("job:%d\n", j.Id)
	//验证 task 类型
	switch j.task.(type) {
	case job.Task:
		fmt.Println("doing job task")
		//interface 转 task
		task, ok := (j.task).(job.Task)
		if ok {
			fmt.Printf("%+v\n", task)
		}
	}
	//工作者任务状态推送到协程池的任务状态队列
	w.Result <- true
}

type pool struct {
	JobQueue chan Jobs //协程池任务队列
	MaxWorks int //最多工作者
	CurrentWorks int //工作者计数器
	Result chan bool //工作状态队列
}

//协程池添加任务
func (p *pool) AddJob(j Jobs) {
	p.JobQueue <- j
}

//停止执行者
func (p *pool) StopWorker() {
	for{
		<- p.Result //释放工作者状态队列
		p.CurrentWorks-- //设置工作者计数器
	}
}

func (p *pool)Run(){
	//开启任务执行监听
	go p.StopWorker()
	//循环接听发布的任务
	for {
		//判断工作者是否超出预计范围
		if p.CurrentWorks < p.MaxWorks {
			//接收队列的任务
			select {
			case j, ok := <-p.JobQueue:
				//验证队列是否还有任务
				if ok {
					//初始化工作者 并关联协程池的result（目的是方便统一管理记录工作者状态）
					w := &Workers{p.Result}
					//设置工作者累加器
					p.CurrentWorks++
					//执行任务
					go w.DoJob(j)
				} else {
					return
				}
			}
		} else {
			fmt.Println("资源不足！")
		}
	}
}

func main() {
	//协程池任务队列
	jobQueue := make(chan Jobs)
	//协程池结果队列
	result := make(chan bool)
	//初始化协程池
	p := &pool{
		MaxWorks: 5,
		JobQueue: jobQueue,
		Result: result,
	}
	//发布任务
	go func() {
		for i:=0;i<10000;i++ {
			t := job.Task{}
			p.AddJob(Jobs{i, t})
		}
		close(p.JobQueue)
	}()
	//运行协程池
	p.Run()
}
