package scheduler

func RunLJF(input []Process) []Process {
	procs := make([]Process, len(input))
	copy(procs, input)

	n := len(procs)
	completed := 0
	currentTime := 0
	isCompleted := make([]bool, n)

	for completed < n {
		idx := -1
		maxBurst := -1

		for i := 0; i < n; i++ {
			if procs[i].ArrivalTime <= currentTime && !isCompleted[i] {
				if procs[i].BurstTime > maxBurst {
					maxBurst = procs[i].BurstTime
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
