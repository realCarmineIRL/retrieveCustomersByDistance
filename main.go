package main

import (
	"log"
)

const distance float64 = 100

func main() {

	//Read customers from file
	customerInfo := GetInfoFromFile("customers.txt")

	// filtered customers by range
	customersInRange := FilterCustomerByDistance(customerInfo, distance)

	// Sorting filtered customers by userID ascending
	customersInRange.Sort()

	// output results in file output.txt.
	// e. Ian Kehoe,4
	//    Nora Dempsey,5
	//    Theresa Enright,6
	err := customersInRange.Print("output.txt")
	if err != nil {
		log.Fatal(err)
	}
}
