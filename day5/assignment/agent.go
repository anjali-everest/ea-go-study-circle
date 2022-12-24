package main

func main() {
	jobs := readJobs()

	channel := make(chan Result)

	for jobIndex := 0; jobIndex < len(jobs.Jobs); jobIndex++ {
		go execute(jobs.Jobs[jobIndex], channel)
	}

	var data []Result
	for jobIndex := 0; jobIndex < len(jobs.Jobs); jobIndex++ {
		data = append(data, <-channel)
	}

	writeSuccessfulJobData(data)
}
