package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

var logFile, csvFile string

var lineCount int = 0
var discardCount int = 0

/* type FilteredItem struct {
	Date         string
	Time         string
	Size         string
	ClientIp     string
	Host         string
	Endpoint     string
	Status       string
	UserAgent    string
	ResponseTime string
}

*/

// function to check given string is in array or not
func ContainsInSlice(sl []string, name string) bool {
	// iterate over the array and compare given string to each element
	for _, value := range sl {
		if strings.Contains(name, value) {
			return true
		}
	}
	return false
}

func readLog() {
	// Open the file for reading
	file, err := os.Open(logFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// list of discard/filter item from logs
	discardList := []string{".js", ".css", ".gif", ".json", ".bmp", ".exe",
		".git", "/static/", ".ico", ".png", ".php", ".sql", ".zip", ".tar", ".txt"}

	// csv file
	csvFileObject, err := os.Create(csvFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a new CSV writer
	writer := csv.NewWriter(csvFileObject)

	// Write the header row to the CSV file
	err = writer.Write([]string{"Date", "Time", "Size", "ClientIp",
		"Host", "Endpoint", "Status", "UserAgent", "ResponseTime"})
	if err != nil {
		panic(err)
	}

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lineCount += 1
		// Process each line of the file
		line := scanner.Text()
		extractedData := strings.Split(line, "\t")
		// total 33 elements
		if len(extractedData) == 33 {

			if !ContainsInSlice(discardList, extractedData[7]) {
				/*
					Date:         extractedData[0],
					Time:         extractedData[1],
					Size:         extractedData[3],
					ClientIp:     extractedData[4],
					Host:         extractedData[15],
					Endpoint:     extractedData[7],
					Status:       extractedData[8],
					UserAgent:    extractedData[10],
					ResponseTime: extractedData[18]
				*/
				err := writer.Write([]string{extractedData[0], extractedData[1],
					extractedData[3], extractedData[4], extractedData[15],
					extractedData[7], extractedData[8], extractedData[10],
					extractedData[18]})
				if err != nil {
					panic(err)
				}
			} else {
				discardCount += 1
			}
		} else {
			fmt.Printf("Omitting->\n%s\n--end---\n", line)
			discardCount += 1
		}

	}

	// Check for any errors that occurred while scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		return
	}
}

func main() {
	start := time.Now()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occurred. Err: %s", err)
	}

	logFile = os.Getenv("LOG_FILE")
	csvFile = os.Getenv("CSV_FILE")

	readLog()

	fmt.Printf("Total lines %d\n", lineCount)
	fmt.Printf("Total discarded lines %d\n", discardCount)
	fmt.Printf("Total lines after filtering %d\n", lineCount-discardCount)

	elapsed := time.Since(start)
	fmt.Printf("Time elapsed: %v\n", elapsed)
}
