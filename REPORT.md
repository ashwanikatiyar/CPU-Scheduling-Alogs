# CPU Scheduling Simulation & Transient Event Analysis
**Course:** Advanced Operating Systems (UMKC)  
**Implementation Language:** Go (Golang)  
**Workload Scale:** 48 Processes (Based on Student ID: 48)  

---

## Executive Summary

CPU scheduling is a foundational responsibility of modern operating systems, determining how process execution requests are multiplexed across available processing cores. The primary goal of a CPU scheduler is to balance system throughput, fairness, latency, and resource utilization. 

This project implements a high-performance simulation framework in Go to evaluate five prominent CPU scheduling policies:
1. **First-Come, First-Served (FCFS)**
2. **Shortest Job First (SJF - Non-Preemptive)**
3. **Round Robin (RR - Time Quantum = 2)**
4. **Priority Scheduling (Preemptive)**
5. **Shortest Remaining Time First (SRTF - Preemptive SJF)**

The evaluation is conducted across two scenario models:
- **Baseline Workload:** A sequence of 48 processes arriving deterministically with varying CPU burst durations and priority levels.
- **Transient Event Workload:** A dynamic event simulating an asynchronous I/O completion at $t = 15$, where a high-priority, short-burst process ($\text{P}48$: Arrival $= 15$, Burst $= 1$, Priority $= 0$) unexpectedly enters the Ready Queue.

Empirical metrics—including Average Turnaround Time ($\overline{TAT}$), Average Waiting Time ($\overline{WT}$), Average Response Time ($\overline{RT}$), and individual transient latency—are analyzed to evaluate how non-preemptive, time-sliced, and preemptive algorithms respond to sudden workload shifts.

---

## 1. System Architecture & Workload Specification

### 1.1 Workload Parameters ($N = 48$)
Per project requirements, the total process count $N = 48$ is derived from the last two digits of the student ID. The process workload dataset (`data/processes.json`) defines process arrival times ($AT \in [0, 47]$), CPU burst times ($BT \in [1, 9]$), and static priority assignments ($\text{Priority} \in [1, 5]$, where smaller integers signify higher priority).

### 1.2 Mathematical Metrics Definitions
For any process $P_i$:
- **Completion Time ($CT_i$):** Timestamp when $P_i$ completes execution.
- **Turnaround Time ($TAT_i$):** Total elapsed time from arrival to completion:
  $$\text{TAT}_i = CT_i - AT_i$$
- **Waiting Time ($WT_i$):** Total time spent waiting in the Ready Queue:
  $$\text{WT}_i = TAT_i - BT_i$$
- **Response Time ($RT_i$):** Time elapsed between process arrival and its first CPU execution slice:
  $$\text{RT}_i = \text{StartTime}_i - AT_i$$

---

## 2. Description of Evaluated CPU Schedulers

### 2.1 First-Come, First-Served (FCFS)
- **Policy:** Non-preemptive, strict FIFO queue ordering.
- **Mechanics:** Processes are dispatched in the exact order of their arrival. Once a process occupies the CPU, it retains control until its burst finishes.
- **Characteristics:** Simple, zero scheduling overhead, but highly susceptible to the **Convoy Effect**, where short processes are delayed behind long-running processes.

### 2.2 Shortest Job First (SJF - Non-Preemptive)
- **Policy:** Non-preemptive, burst-length priority.
- **Mechanics:** When the CPU becomes idle, the scheduler inspects all currently ready processes and selects the process with the shortest CPU burst time.
- **Characteristics:** Minimizes average waiting time among non-preemptive policies. However, it requires prior knowledge of burst times and can cause starvation for long processes.

### 2.3 Round Robin (RR - Time Quantum $q = 2$)
- **Policy:** Preemptive, circular time-sharing.
- **Mechanics:** Ready processes are assigned CPU execution slices up to $q = 2$ time units. If a process does not finish within its quantum, it is preempted and re-queued at the back of the Ready Queue.
- **Characteristics:** Ensures strict fairness and low response times for interactive applications. Performance depends heavily on quantum length selection.

### 2.4 Priority Scheduling (Preemptive)
- **Policy:** Preemptive, priority-driven.
- **Mechanics:** At every tick, the scheduler selects the process with the highest priority (numerical value $0$ being supreme). If a newly arrived process has a higher priority than the currently running process, the running process is immediately preempted.
- **Characteristics:** Ideal for real-time systems where urgent tasks demand immediate CPU access. Risk of starvation for lower-priority jobs unless aging mechanisms are applied.

### 2.5 Shortest Remaining Time First (SRTF)
- **Policy:** Preemptive SJF.
- **Mechanics:** At every time unit, the scheduler compares remaining burst times of all ready processes. The process with the minimum remaining CPU time is executed.
- **Characteristics:** Mathematically optimal for minimizing average turnaround time. Higher context switching overhead compared to non-preemptive SJF.

---

## 3. Transient Event Definition & Mechanics

### 3.1 Transient Event Concept
In real-world operating systems, workloads are rarely static. Interrupts, user interactions, and I/O device completions introduce sudden, high-priority demands into the CPU queue. 

### 3.2 Simulation Invalidation Model
At $t = 15$, process $\text{P}48$ undergoes an asynchronous I/O completion event:
- **Original State (Normal Workload):** Arrives at $t = 47$, Burst Time $= 3$, Priority $= 1$.
- **Transient State (Event Workload):** Arrives at $t = 15$, Burst Time $= 1$, Priority $= 0$.

This transient event tests how rapidly each scheduling algorithm detects, prioritizes, and executes a critical, short-burst task arriving mid-stream.

```
Baseline P48:  t=47 [Burst = 3, Priority = 1]
Transient P48: t=15 [Burst = 1, Priority = 0] <--- High-Priority I/O Return
```

---

## 4. Experimental Results & Performance Analysis

### 4.1 Summary of Average Metrics

| Scheduling Algorithm | Simulation Mode | Avg Turnaround Time ($\overline{TAT}$) | Avg Waiting Time ($\overline{WT}$) | Avg Response Time ($\overline{RT}$) |
| :--- | :--- | :---: | :---: | :---: |
| **FCFS** | Baseline | 92.04 | 87.46 | 87.46 |
| **FCFS** | Transient Event | 90.38 | 85.83 | 85.83 |
| **SJF (Non-Preemptive)** | Baseline | 61.75 | 57.17 | 57.17 |
| **SJF (Non-Preemptive)** | Transient Event | 60.75 | 56.21 | 56.21 |
| **Round Robin ($q=2$)** | Baseline | 117.04 | 112.46 | 37.94 |
| **Round Robin ($q=2$)** | Transient Event | 114.15 | 109.60 | 37.56 |
| **Priority (Preemptive)** | Baseline | 89.69 | 85.10 | 82.92 |
| **Priority (Preemptive)** | Transient Event | 88.15 | 83.60 | 81.60 |
| **SRTF (Preemptive SJF)**| Baseline | 61.65 | 57.06 | 56.79 |
| **SRTF (Preemptive SJF)**| Transient Event | 60.62 | 56.08 | 55.81 |

---

### 4.2 Transient Event Reaction: Detailed Analysis of Process $\text{P}48$

The key indicator of an algorithm's adaptability is how it processes $\text{P}48$ upon its arrival at $t = 15$.

| Algorithm | $\text{P}48$ Arrival ($AT$) | Start Time ($ST$) | Completion ($CT$) | Turnaround ($TAT$) | Waiting ($WT$) | Response ($RT$) | Reaction Grade |
| :--- | :---: | :---: | :---: | :---: | :---: | :---: | :---: |
| **Priority (Preemptive)** | 15 | 15 | 16 | 1 | 0 | 0 | **Instantaneous** |
| **SRTF (Preemptive)** | 15 | 15 | 16 | 1 | 0 | 0 | **Instantaneous** |
| **SJF (Non-Preemptive)** | 15 | 17 | 18 | 3 | 2 | 2 | **Rapid (Bounded Wait)** |
| **Round Robin ($q=2$)** | 15 | 44 | 45 | 30 | 29 | 29 | **Moderate** |
| **FCFS** | 15 | 76 | 77 | 62 | 61 | 61 | **Poor (Convoy Delayed)** |

#### Detailed Observations:
1. **Preemptive Priority & SRTF (Immediate Response):** Both algorithms achieve an immediate context switch at $t = 15$. Because $\text{P}48$ has Priority $0$ and remaining time $1$, it preempts the running process instantly, yielding $\text{WT} = 0$ and $\text{RT} = 0$.
2. **SJF (Non-Preemptive Rapid Selection):** SJF cannot preempt the process running at $t = 15$. However, as soon as that process completes at $t = 17$, SJF selects $\text{P}48$ because its burst ($1$) is shorter than any other ready job. $\text{P}48$ waits only $2$ time units.
3. **Round Robin (Fairness over Urgency):** RR places $\text{P}48$ into the circular queue. $\text{P}48$ must wait for preceding quantum rotations, suffering a waiting time of $29$ units despite its high priority and short burst.
4. **FCFS (Convoy Obstruction):** FCFS forces $\text{P}48$ to wait behind all 15 processes already queued before $t = 15$, leading to a massive waiting time of $61$ units.

---

## 5. Comprehensive Comparison & Pros/Cons Matrix

### 5.1 First-Come, First-Served (FCFS)
- **Pros:** 
  - Extremely simple to implement using a standard FIFO queue.
  - Zero scheduling and preemption overhead.
  - No risk of starvation; every process eventually executes.
- **Cons:**
  - High average turnaround and waiting times ($\overline{TAT} = 92.04$).
  - Vulnerable to the Convoy Effect.
  - Cannot handle transient or real-time events efficiently ($\text{WT} = 61$ for $\text{P}48$).

### 5.2 Shortest Job First (SJF - Non-Preemptive)
- **Pros:**
  - Dramatically improves turnaround efficiency ($\overline{TAT} = 61.75$, a $32.9\%$ reduction over FCFS).
  - Low context switching overhead compared to preemptive policies.
  - Quickly absorbs short transient bursts once the CPU becomes free ($\text{WT} = 2$ for $\text{P}48$).
- **Cons:**
  - Requires apriori knowledge of CPU burst times.
  - Risk of starvation for long processes under heavy workloads.
  - Cannot preempt long-running processes during urgent transient events.

### 5.3 Round Robin (RR - $q = 2$)
- **Pros:**
  - Excellent responsiveness for interactive processes ($\overline{RT} = 37.94$, the lowest among baseline algorithms).
  - Guarantees strict CPU fairness; no starvation possible.
  - Prevents single process monopolization.
- **Cons:**
  - Higher average turnaround time ($\overline{TAT} = 117.04$) due to frequent context switching.
  - Ignores process priority and burst length when scheduling transient events ($\text{WT} = 29$ for $\text{P}48$).

### 5.4 Priority Scheduling (Preemptive)
- **Pros:**
  - Perfect for real-time and mission-critical applications.
  - Instantaneous preemption for high-priority transient events ($\text{WT} = 0$, $\text{RT} = 0$ for $\text{P}48$).
  - Flexible policy control via priority assignment.
- **Cons:**
  - Severe risk of starvation for low-priority processes.
  - Overhead from frequent preemptive context switches.
  - Requires priority aging mechanisms in production environments.

### 5.5 Shortest Remaining Time First (SRTF)
- **Pros:**
  - Achieves the optimal average turnaround time ($\overline{TAT} = 61.65$) and waiting time ($\overline{WT} = 57.06$).
  - Combines the benefits of shortest-burst prioritization with immediate preemption.
  - Instant response to short transient bursts ($\text{WT} = 0$ for $\text{P}48$).
- **Cons:**
  - High scheduling overhead due to continuous monitoring of remaining burst times.
  - Starvation risk for longer jobs.
  - Requires accurate runtime burst estimations.

---

## 6. Implementation Code Base

### 6.1 Process Data Structure (`scheduler/process.go`)
```go
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
		return nil, err
	}
	var processes []Process
	if err = json.Unmarshal(data, &processes); err != nil {
		return nil, err
	}
	for i := range processes {
		processes[i].RemainingTime = processes[i].BurstTime
		processes[i].StartTime = -1
	}
	return processes, nil
}
```

### 6.2 Preemptive Priority Scheduler (`scheduler/priority.go`)
```go
package scheduler

func RunPriorityPreemptive(input []Process) []Process {
	procs := make([]Process, len(input))
	copy(procs, input)

	n := len(procs)
	currentTime := 0
	completed := 0

	for completed < n {
		idx := -1
		highestPriority := 1000000

		for i := 0; i < n; i++ {
			if procs[i].ArrivalTime <= currentTime && procs[i].RemainingTime > 0 {
				if procs[i].Priority < highestPriority {
					highestPriority = procs[i].Priority
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
```

### 6.3 Shortest Remaining Time First (`scheduler/srtf.go`)
```go
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
```

### 6.4 Round Robin Scheduler (`scheduler/rr.go`)
```go
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
```

---

## 7. Conclusion

This project successfully implemented and evaluated five fundamental CPU scheduling algorithms in Go under both baseline and transient event conditions with a dataset of 48 processes.

Key conclusions include:
1. **SRTF** delivers the overall best average turnaround time ($\overline{TAT} = 61.65$), making it optimal for overall execution efficiency.
2. **Preemptive Priority Scheduling** and **SRTF** provide superior responsiveness to transient high-priority events, yielding immediate response times ($\text{RT} = 0$).
3. **Round Robin** excels at maintaining interactive fairness ($\overline{RT} = 37.94$), though it struggles to prioritize urgent transient tasks without priority extensions.
4. **FCFS** is unsuitable for dynamic operating systems due to severe convoy delays when handling transient I/O interrupts.

Modern general-purpose operating systems (such as Linux's Completely Fair Scheduler - CFS) synthesize these findings by combining time-slice fairness with dynamic priority and preemption capabilities.
