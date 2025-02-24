package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/mergestat/timediff"
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

func printTasks() {
	data, err := os.ReadFile("tasks.csv")
	check(err)
	writer := tabwriter.NewWriter(os.Stdout, 0, 2, 4, ' ', 0)
	writer.Write([]byte("ID\tTask\tCreated on\tStatus\tInterval"))

	lines := strings.Split(string(data), "\n")
	for _, line := range lines[1:] {
		token := strings.Split(line, ",")
		// Added check to ensure token has at least 4 entries
		if len(token) < 4 {
			continue
		}
		id := token[0]
		task := token[1]
		createdOn := token[2]
		status := token[3]
		parsedTime, err := time.Parse(time.RFC3339, createdOn)
		check(err)

		interval := timediff.TimeDiff(parsedTime)

		fmt.Fprintf(writer, "\n%s\t%s\t%s\t%s\t%s", id, task, createdOn, status, interval)

	}
	writer.Flush()
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
