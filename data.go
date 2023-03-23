package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

func createData() {
	log.Println("[MESSAGE] creating json file..")
	proctime := time.Now()

	var myArr []*Person
	for i := 1; i <= 1000; i++ {
		myArr = append(myArr, &Person{
			ID:        uint(i),
			Customer:  fmt.Sprintf("John Smith %d", i),
			Quantity:  10,
			Price:     float64(i * 100),
			Timestamp: time.Now().Format(time.RFC3339),
		})
	}

	file, err := json.MarshalIndent(&myArr, "", "   ")
	if err != nil {
		log.Fatalln(err)
	}
	err = ioutil.WriteFile("person.json", file, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("[MESSAGE] done! processed in %dms", time.Since(proctime).Milliseconds())
}
