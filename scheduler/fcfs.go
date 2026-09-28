package scheduler

import "sort"

func RunFCFS(input []Process) []Process {
	procs := make([]Process, len(input))
	copy(procs, input)

	sort.Slice(procs, func(i, j int) bool {
		if procs[i].ArrivalTime == procs[j].ArrivalTime {
			return procs[i].ID < procs[j].ID
		}
		return procs[i].ArrivalTime < procs[j].ArrivalTime
	})

	currentTime := 0
	for i := range procs {
		if currentTime < procs[i].ArrivalTime {
			currentTime = procs[i].ArrivalTime
		}
		procs[i].StartTime = currentTime
		procs[i].CompletionTime = currentTime + procs[i].BurstTime
		procs[i].TurnaroundTime = procs[i].CompletionTime - procs[i].ArrivalTime
		procs[i].WaitingTime = procs[i].TurnaroundTime - procs[i].BurstTime
		procs[i].ResponseTime = procs[i].StartTime - procs[i].ArrivalTime
		currentTime = procs[i].CompletionTime
	}

	return procs
}
