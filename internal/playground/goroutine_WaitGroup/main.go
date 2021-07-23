package main

import (
	"time"
	"sync"
	"log"
)

func init() {
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds)
}

type testConcurrency struct {
	min int
	max int
	country string
}

func printCountry(test *testConcurrency, groupTest *sync.WaitGroup) {
	for i :=test.max ; i>test.min; i-- {
		time.Sleep(1*time.Millisecond)
		log.Println(test.country)
	}

	log.Println("Done printing country:", test.country)
	groupTest.Done()
}

func  main() {
	groupTest := new(sync.WaitGroup)

	japan := new(testConcurrency)
	china := new(testConcurrency)
	india := new(testConcurrency)

	japan.country = "Japan"
	japan.min = 0
	japan.max = 5

	china.country = "China"
	china.min = 0
	china.max = 5

	india.country = "India"
	india.min = 0
	india.max = 5

	go printCountry(japan, groupTest)
	go printCountry(china, groupTest)
	go printCountry(india, groupTest)

	log.Println("Waiting START")
	groupTest.Add(3)
	groupTest.Wait()
	log.Println("Waiting END")
}
