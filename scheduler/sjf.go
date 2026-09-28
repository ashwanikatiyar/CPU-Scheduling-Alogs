package scheduler

func RunSJF(input []Process) []Process {
	procs := make([]Process, len(input))
	copy(procs, input)

	n := len(procs)
	completed := 0
	currentTime := 0
	isCompleted := make([]bool, n)

	for completed < n {
		idx := -1
		minBurst := 1000000

		for i := 0; i < n; i++ {
			if procs[i].ArrivalTime <= currentTime && !isCompleted[i] {
				if procs[i].BurstTime < minBurst {
					minBurst = procs[i].BurstTime
					idx = i
				}
			}
		}

		if idx != -1 {
			procs[idx].StartTime = currentTime
			procs[idx].CompletionTime = currentTime + procs[idx].BurstTime
			procs[idx].TurnaroundTime = procs[idx].CompletionTime - procs[idx].ArrivalTime
			procs[idx].WaitingTime = procs[idx].TurnaroundTime - procs[idx].BurstTime
			procs[idx].ResponseTime = procs[idx].StartTime - procs[idx].ArrivalTime

			currentTime = procs[idx].CompletionTime
			isCompleted[idx] = true
			completed++
		} else {
			currentTime++
		}
	}

	return procs
}
