package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Jobs struct {
	Jobs []Job `json:"jobs"`
}

type Job struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Result struct {
	Job    Job
	status string
}

func execute(job Job, channel chan Result) {
	log.Println("executing ", job.Name)
	channel <- Result{Job: job, status: "SUCCESS"}
}

func main() {
	jsonFile, err := os.Open("jobs.json")
	channel := make(chan Result)
	if err != nil {
		log.Println(err)
	}
	log.Println("Successfully Opened jobs.json")

	byteValue, _ := io.ReadAll(jsonFile)

	var jobs Jobs

	json.Unmarshal(byteValue, &jobs)

	for i := 0; i < len(jobs.Jobs); i++ {
		go execute(jobs.Jobs[i], channel)
	}

	var data []Result
	for i := 0; i < len(jobs.Jobs); i++ {
		data = append(data, <-channel)
	}
	log.Println(data)
}
