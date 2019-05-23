package main

import (
	//"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Open the file
	csvfile, err := os.Open("tmp.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)

	file, err := os.Create("result.csv")
	if err != nil {}

	writer := csv.NewWriter(file)
    defer writer.Flush()

	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
	   fmt.Printf("%s : %s\n", record[0], record[1])

	   data:= [][]string{{record[1]}}

	   err2 := writer.WriteAll(data)
	    if err2 != nil {}

	}

}
