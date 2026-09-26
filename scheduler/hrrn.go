package scheduler

func RunHRRN(input []Process) []Process {
	procs := make([]Process, len(input))
	copy(procs, input)

	n := len(procs)
	currentTime := 0
	completed := 0
	isCompleted := make([]bool, n)

	for completed < n {
		idx := -1
		maxRatio := -1.0

		for i := 0; i < n; i++ {
			if procs[i].ArrivalTime <= currentTime && !isCompleted[i] {
				waitingTime := currentTime - procs[i].ArrivalTime
				responseRatio := float64(waitingTime+procs[i].BurstTime) / float64(procs[i].BurstTime)

				if responseRatio > maxRatio {
					maxRatio = responseRatio
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