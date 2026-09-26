package scheduler

func RunSRTF(input []Process) []Process {
	procs := make([]Process, len(input))
	copy(procs, input)

	n := len(procs)
	currentTime := 0
	completed := 0

	for completed < n {
		idx := -1
		minRemaining := 1000000

		for i := 0; i < n; i++ {
			if procs[i].ArrivalTime <= currentTime && procs[i].RemainingTime > 0 {
				if procs[i].RemainingTime < minRemaining {
					minRemaining = procs[i].RemainingTime
					idx = i
				}
			}
		}

		if idx != -1 {
			if procs[idx].StartTime == -1 {
				procs[idx].StartTime = currentTime
			}

			procs[idx].RemainingTime--
			currentTime++

			if procs[idx].RemainingTime == 0 {
				completed++
				procs[idx].CompletionTime = currentTime
				procs[idx].TurnaroundTime = procs[idx].CompletionTime - procs[idx].ArrivalTime
				procs[idx].WaitingTime = procs[idx].TurnaroundTime - procs[idx].BurstTime
				procs[idx].ResponseTime = procs[idx].StartTime - procs[idx].ArrivalTime
			}
		} else {
			currentTime++
		}
	}

	return procs
}