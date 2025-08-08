// 代码生成时间: 2025-08-09 01:02:00
package main

import (
    "fmt"
    "log"
    "time"
    "github.com/robfig/cron/v3"
)

// Scheduler struct to hold cron scheduler
type Scheduler struct {
    Cron *cron.Cron
}

// NewScheduler creates and returns a new scheduler with the specified schedule
func NewScheduler(schedule string) (*Scheduler, error) {
    cron, err := cron.ParseStandard(schedule)
    if err != nil {
        return nil, err
    }
    return &Scheduler{Cron: cron}, nil
}

// Start starts the scheduler
func (s *Scheduler) Start() {
    s.Cron.Start()
    fmt.Println("Scheduler started")
}

// Stop stops the scheduler
func (s *Scheduler) Stop() error {
    return s.Cron.Stop()
}

// AddJob adds a job to the scheduler
func (s *Scheduler) AddJob(schedule string, cmd func() error) error {
    _, err := s.Cron.AddFunc(schedule, cmd)
    return err
}

// ExampleJob is an example job function that logs a message
func ExampleJob() error {
    log.Println("Example job is running.")
    return nil
}

func main() {
    s, err := NewScheduler("0 * * * * *") // Run every hour
    if err != nil {
        log.Fatalf("Failed to create scheduler: %v", err)
    }

    if err := s.AddJob("0 * * * * *", ExampleJob); err != nil {
        log.Fatalf("Failed to add job to scheduler: %v", err)
    }

    s.Start()
    
    // To gracefully stop the scheduler, you can use a context or a signal handler in a real application
    // For example purposes, we will just wait for 5 minutes and then stop the scheduler
    time.Sleep(5 * time.Minute)
    if err := s.Stop(); err != nil {
        log.Fatalf("Failed to stop scheduler: %v", err)
    }
    fmt.Println("Scheduler stopped")
}