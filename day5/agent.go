package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type Jobs struct {
	Jobs []Job `json:"jobs"`
}

type Job struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func execute(job Job) {
	log.Println("executing ", job.Name)
	//return result{id: job.Id, status: "SUCCESS"}
}

func main() {
	jsonFile, err := os.Open("jobs.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened jobs.json")

	byteValue, _ := io.ReadAll(jsonFile)

	var jobs Jobs

	json.Unmarshal(byteValue, &jobs)

	for i := 0; i < len(jobs.Jobs); i++ {
		go execute(jobs.Jobs[i])
		time.Sleep(time.Millisecond * 5)
	}
}
