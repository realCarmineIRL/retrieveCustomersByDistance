package main

import (
	"log"
)

func main() {
	customerInfo := GetInfoFromFile("customers.txt")

	customersInRange := FilterCustomerByDistance(customerInfo, 100)

	customersInRange.Sort()

	err := customersInRange.Print("output.txt")
	if err != nil {
		log.Fatal(err)
	}
}
