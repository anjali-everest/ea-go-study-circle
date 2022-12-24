package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
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
	Status string
}

func execute(job Job, channel chan Result) {
	log.Println("executing ", job.Name)
	channel <- Result{Job: job, Status: "SUCCESS"}
}

func readJobs() Jobs {
	jsonFile, err := os.Open("io/jobs.json")
	if err != nil {
		log.Println(err)
	}
	log.Println("Successfully Opened jobs.json")
	byteValue, _ := io.ReadAll(jsonFile)
	var jobs Jobs
	json.Unmarshal(byteValue, &jobs)
	return jobs
}

func writeSuccessfulJobData(data []Result) {
	marshaled_data, _ := json.MarshalIndent(data, "", "  ")
	_ = ioutil.WriteFile("io/status.json", marshaled_data, 0644)
}
