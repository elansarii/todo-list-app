package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// addTask("Finish this app", counter)
	// addTask("IDK man", counter)
	// addTask("Drink water", counter)

	// printTasks()

}
func addTask(task string, counter int) {
	counter++
	f, err := os.OpenFile("tasks.csv", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	check(err)
	defer f.Close()
	writer := csv.NewWriter(f)
	defer writer.Flush()
	id := strconv.Itoa(counter)

	currentTime := time.Now().Format(time.RFC3339)

	record := []string{id, task, currentTime, "false"}

	if err := writer.Write(record); err != nil {
		log.Fatalf("Error writing record to CSV: %s", err)
	}

}
