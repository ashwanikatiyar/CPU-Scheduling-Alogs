package scheduler

func RunRoundRobin(input []Process, timeQuantum int) []Process {
	procs := make([]Process, len(input))
	copy(procs, input)

	n := len(procs)
	currentTime := 0
	completed := 0

	readyQueue := make([]int, 0)
	inQueue := make([]bool, n)

	for completed < n {
		for i := 0; i < n; i++ {
			if procs[i].ArrivalTime <= currentTime && procs[i].RemainingTime > 0 && !inQueue[i] {
				readyQueue = append(readyQueue, i)
				inQueue[i] = true
			}
		}

		if len(readyQueue) == 0 {
			currentTime++
			continue
		}

		idx := readyQueue[0]
		readyQueue = readyQueue[1:]

		if procs[idx].StartTime == -1 {
			procs[idx].StartTime = currentTime
		}

		execTime := timeQuantum
		if procs[idx].RemainingTime < timeQuantum {
			execTime = procs[idx].RemainingTime
		}

		procs[idx].RemainingTime -= execTime
		currentTime += execTime

		for i := 0; i < n; i++ {
			if procs[i].ArrivalTime <= currentTime && procs[i].RemainingTime > 0 && !inQueue[i] && i != idx {
				readyQueue = append(readyQueue, i)
				inQueue[i] = true
			}
		}

		if procs[idx].RemainingTime > 0 {
			readyQueue = append(readyQueue, idx)
		} else {
			completed++
			procs[idx].CompletionTime = currentTime
			procs[idx].TurnaroundTime = procs[idx].CompletionTime - procs[idx].ArrivalTime
			procs[idx].WaitingTime = procs[idx].TurnaroundTime - procs[idx].BurstTime
			procs[idx].ResponseTime = procs[idx].StartTime - procs[idx].ArrivalTime
		}
	}

	return procs
}
