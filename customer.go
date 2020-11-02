package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

// Customers represent an slice of customer
type Customers []Customer

// Print receives an slice of customers and print to file Name and UserID separated by "," .
func (cs Customers) Print(filename string) error {
	o := output{}
	for _, customer := range cs {
		c := fmt.Sprintf("%v,%v", customer.Name, customer.UserID)
		o = append(o, c)
	}

	return ioutil.WriteFile(filename, []byte(o.toString()), 0666)
}

// Sort receives an slice of customers and sort the by userID
func (cs Customers) Sort() {
	sort.Slice(cs, func(i, j int) bool {
		return cs[i].UserID < cs[j].UserID
	})
}

// Customer struct represent a customer
type Customer struct {
	UserID    int    `json:"user_id"`
	Name      string `json:"name"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type output []string

func (o output) toString() string {
	return strings.Join([]string(o), "\n")
}

// FilterCustomerByDistance remove customers if not match the expected distance
func FilterCustomerByDistance(cs Customers, distance float64) Customers {
	customersInRange := Customers{}

	for _, customer := range cs {
		if cd, err := CalcDistance(customer.Latitude, customer.Longitude); cd <= distance && err == nil {
			customersInRange = append(customersInRange, customer)
		}
	}

	return customersInRange
}

// GetInfoFromFile read customers from file
func GetInfoFromFile(filename string) Customers {

	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	rows := strings.Split(string(bs), "\n")

	cs := Customers{}
	for _, record := range rows {
		c := Customer{}
		err := json.Unmarshal([]byte(record), &c)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		cs = append(cs, c)
	}

	return cs
}
