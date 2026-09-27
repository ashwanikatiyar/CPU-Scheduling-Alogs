package main

import (
	"fmt"
	"log"

	"cpu_scheduling_go_project/scheduler"
)

func applyTransientEvent(processes []scheduler.Process) []scheduler.Process {
	modified := make([]scheduler.Process, len(processes))
	copy(modified, processes)

	// Inject 5 sudden burst-heavy jobs at time t = 15
	transientJobs := []scheduler.Process{
		{ID: 101, ArrivalTime: 15, BurstTime: 12, RemainingTime: 12, Priority: 1},
		{ID: 102, ArrivalTime: 15, BurstTime: 15, RemainingTime: 15, Priority: 1},
		{ID: 103, ArrivalTime: 15, BurstTime: 10, RemainingTime: 10, Priority: 2},
		{ID: 104, ArrivalTime: 15, BurstTime: 14, RemainingTime: 14, Priority: 1},
		{ID: 105, ArrivalTime: 15, BurstTime: 11, RemainingTime: 11, Priority: 2},
	}

	// Adjust ongoing/upcoming processes
	for i := range modified {
		if modified[i].ArrivalTime >= 15 {
			modified[i].BurstTime *= 2
			modified[i].RemainingTime = modified[i].BurstTime
		}
	}

	modified = append(modified, transientJobs...)
	return modified
}

func main() {
	processes, err := scheduler.LoadProcesses("data/processes.json")
	if err != nil {
		log.Fatalf("Failed to load processes: %v", err)
	}

	fmt.Println("==================================================")
	fmt.Println("  SIMULATION 1: NORMAL WORKLOAD (48 PROCESSES)")
	fmt.Println("==================================================")

	scheduler.PrintMetrics("1. FCFS", scheduler.RunFCFS(processes))
	scheduler.PrintMetrics("2. SJF (Non-Preemptive)", scheduler.RunSJF(processes))
	// scheduler.PrintMetrics("3. LJF (Non-Preemptive)", scheduler.RunLJF(processes))
	scheduler.PrintMetrics("4. Round Robin (TQ=2)", scheduler.RunRoundRobin(processes, 2))
	scheduler.PrintMetrics("5. Priority (Preemptive)", scheduler.RunPriorityPreemptive(processes))
	scheduler.PrintMetrics("6. SRTF", scheduler.RunSRTF(processes))
	// scheduler.PrintMetrics("7. LRTF", scheduler.RunLRTF(processes))
	// scheduler.PrintMetrics("8. HRRN", scheduler.RunHRRN(processes))

	fmt.Println("\n==================================================")
	fmt.Println("  SIMULATION 2: TRANSIENT EVENT WORKLOAD")
	fmt.Println("==================================================")

	transientProcesses := applyTransientEvent(processes)

	scheduler.PrintMetrics("1. FCFS (Transient)", scheduler.RunFCFS(transientProcesses))
	scheduler.PrintMetrics("2. SJF (Transient)", scheduler.RunSJF(transientProcesses))
	// scheduler.PrintMetrics("3. LJF (Transient)", scheduler.RunLJF(transientProcesses))
	scheduler.PrintMetrics("4. Round Robin (Transient)", scheduler.RunRoundRobin(transientProcesses, 2))
	scheduler.PrintMetrics("5. Priority (Transient)", scheduler.RunPriorityPreemptive(transientProcesses))
	scheduler.PrintMetrics("6. SRTF (Transient)", scheduler.RunSRTF(transientProcesses))
	// scheduler.PrintMetrics("7. LRTF (Transient)", scheduler.RunLRTF(transientProcesses))
	// scheduler.PrintMetrics("8. HRRN (Transient)", scheduler.RunHRRN(transientProcesses))
}