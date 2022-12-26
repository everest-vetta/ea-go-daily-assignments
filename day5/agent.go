package main

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

type Job struct {
	Id   int
	Name string
}

type Result struct {
	Id     int
	Status string
}

func Agent() {
	ch := make(chan Result)

	content, err := os.ReadFile("./jobs.json")
	if err != nil {
		log.Fatal("Error while opening file: ", err)
	}

	var jobs []Job
	err = json.Unmarshal(content, &jobs)
	if err != nil {
		log.Fatal("Error occured while unmarshalling : ", err)
	}

	for i := range jobs {
		go execute_job(ch, jobs[i])
	}
	go readResultAndSave(ch, "job_result")
	time.Sleep(time.Millisecond * 30)
	close(ch)
}

func execute_job(ch chan Result, job Job) {
	log.Printf("Executing job with id :%v", job.Id)
	result := Result{Id: job.Id, Status: "Success"}
	sendJobToChannel(ch, result)
}

func sendJobToChannel(ch chan Result, result Result) {
	ch <- result
}

func readResultAndSave(ch chan Result, fileName string) {
	for {
		result := <-ch
		saveResultToJsonFile(result, fileName)
	}
}

func saveResultToJsonFile(result Result, fileName string) {
	fileName = fileName + ".json"
	marshal_result, _ := json.Marshal(result)

	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := f.Write([]byte(marshal_result)); err != nil {
		log.Fatal(err)
	}

	if _, err = f.WriteString(",\n"); err != nil {
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
