package scheduler

import (
	"encoding/json"
	"fmt"
	"os"
)

type Process struct {
	ID             int `json:"id"`
	ArrivalTime    int `json:"arrival_time"`
	BurstTime      int `json:"burst_time"`
	RemainingTime  int `json:"-"`
	Priority       int `json:"priority"` // Lower number = Higher priority
	CompletionTime int `json:"-"`
	TurnaroundTime int `json:"-"`
	WaitingTime    int `json:"-"`
	ResponseTime   int `json:"-"`
	StartTime      int `json:"-"`
}

func LoadProcesses(filePath string) ([]Process, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error in reading file:", filePath, "Error:", err)
		return nil, err
	}
	var processes []Process
	err = json.Unmarshal(data, &processes)
	if err != nil {
		fmt.Println("Error in unmarshaling file data:", string(data), "Error:", err)
		return nil, err
	}

	for i := range processes {
		processes[i].RemainingTime = processes[i].BurstTime
		processes[i].StartTime = -1
	}
	return processes, nil
}

func PrintMetrics(name string, processes []Process) {
	var totalTAT, totalWT, totalRT float64
	n := float64(len(processes))

	fmt.Printf("\n--- Results for %s ---\n", name)
	fmt.Printf("%-6s %-8s %-6s %-6s %-6s %-6s\n", "PID", "Arrival", "Burst", "CT", "TAT", "WT")

	for _, p := range processes {
		totalTAT += float64(p.TurnaroundTime)
		totalWT += float64(p.WaitingTime)
		totalRT += float64(p.ResponseTime)
		fmt.Printf("%-6d %-8d %-6d %-6d %-6d %-6d\n", p.ID, p.ArrivalTime, p.BurstTime, p.CompletionTime, p.TurnaroundTime, p.WaitingTime)
	}

	fmt.Printf("Average Turnaround Time: %.2f\n", totalTAT/n)
	fmt.Printf("Average Waiting Time:    %.2f\n", totalWT/n)
	fmt.Printf("Average Response Time:   %.2f\n", totalRT/n)
}
