// 代码生成时间: 2025-08-24 10:48:04
package main

import (
    "fmt"
    "os"
    "os/exec"
    "strings"
    "syscall"
    "time"
)

// Process represents a running process
type Process struct {
    Process *os.Process
    Name    string
    PID     int
}

// NewProcess creates a new process
func NewProcess(name string, cmd *exec.Cmd) (*Process, error) {
    err := cmd.Start()
    if err != nil {
        return nil, err
    }
    return &Process{
        Process: cmd.Process,
        Name:    name,
        PID:     cmd.Process.Pid,
    }, nil
}

// Kill sends a termination signal to the process
func (p *Process) Kill() error {
    err := p.Process.Signal(syscall.SIGTERM)
    if err != nil {
        return err
    }
    return nil
}

// Wait waits for the process to finish
func (p *Process) Wait() error {
    if p.Process == nil {
        return fmt.Errorf("process is not running")
    }
    _, err := os.Wait4(p.Process.Pid, nil, 0, nil)
    if err != nil {
        return err
    }
    return nil
}

// ListProcesses lists all running processes
func ListProcesses() ([]*Process, error) {
    // This is a simplified version and would require actual implementation to interact with the OS
    // to get the list of processes. Here we just simulate it with a hardcoded list.
    processes := []*Process{
        {
            Process: &os.Process{Pid: 1234},
            Name:    "Process1",
            PID:     1234,
        },
        {
            Process: &os.Process{Pid: 5678},
            Name:    "Process2",
            PID:     5678,
        },
    }
    return processes, nil
}

func main() {
    // Example of how to use the Process Manager
    cmd := exec.Command("ls", "-l")
    process, err := NewProcess("LsCommand", cmd)
    if err != nil {
        fmt.Printf("Failed to start process: %s
", err)
        return
    }
    defer process.Wait()

    fmt.Printf("Process %s started with PID %d
", process.Name, process.PID)

    // Simulate doing some work...
    time.Sleep(2 * time.Second)

    // Now we kill the process
    if err := process.Kill(); err != nil {
        fmt.Printf("Failed to kill process: %s
", err)
    } else {
        fmt.Printf("Process %s killed successfully
", process.Name)
    }

    // List all processes (this is a mock implementation and would need actual OS interaction)
    processes, err := ListProcesses()
    if err != nil {
        fmt.Printf("Failed to list processes: %s
", err)
    } else {
        for _, p := range processes {
            fmt.Printf("Process Name: %s, PID: %d
", p.Name, p.PID)
        }
    }
}