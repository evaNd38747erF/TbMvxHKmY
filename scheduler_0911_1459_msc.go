// 代码生成时间: 2025-09-11 14:59:16
// scheduler.go 定时任务调度器

package main

import (
    "fmt"
    "log"
    "time"
    "github.com/robfig/cron/v3"
)

// Scheduler 定时任务调度器结构体
type Scheduler struct {
    c *cron.Cron
}

// NewScheduler 创建一个新的调度器实例
func NewScheduler() *Scheduler {
    return &Scheduler{
        c: cron.New(cron.WithSeconds()), // 启用秒级别的调度
    },
}

// AddJob 添加一个定时任务
func (s *Scheduler) AddJob(when string, job func()) error {
    _, err := s.c.AddFunc(when, job)
    return err
}

// Start 开始执行调度器
func (s *Scheduler) Start() {
    s.c.Start()
}

// Stop 停止调度器
func (s *Scheduler) Stop() {
    s.c.Stop()
}

// JobExample 定时任务示例函数
func JobExample() {
    fmt.Println("Executing job at", time.Now())
}

func main() {
    scheduler := NewScheduler()
    defer scheduler.Stop()

    // 添加一个每分钟执行一次的定时任务
    if err := scheduler.AddJob("* * * * *", JobExample); err != nil {
        log.Fatalln("Failed to schedule job: ", err)
    }

    // 模拟长时间运行的程序
    for {
        time.Sleep(time.Hour) // 模拟程序长时间运行
    }
}
