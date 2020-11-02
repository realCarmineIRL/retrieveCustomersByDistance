package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestGetInfoFromFile(t *testing.T) {
	expected := Customers{}

	expected = append(expected, Customer{12, "Christina McArdle", "52.986375", "-6.043701"})
	expected = append(expected, Customer{1, "Alice Cahill", "51.92893", "-10.27699"})
	expected = append(expected, Customer{2, "Ian McArdle", "51.8856167", "-10.4240951"})

	results := GetInfoFromFile("customers_test.txt")

	for i := 0; i < len(results); i++ {
		if results[i] != expected[i] {
			t.Errorf("Error Reading File")
		}
	}
}

func TestFilterCustomerByDistance(t *testing.T) {
	cs := Customers{}
	expected := Customers{}

	cs = append(cs, Customer{12, "Christina McArdle", "52.986375", "-6.043701"})
	cs = append(cs, Customer{1, "Alice Cahill", "51.92893", "-10.27699"})
	cs = append(cs, Customer{2, "Ian McArdle", "51.8856167", "-10.4240951"})

	expected = append(expected, Customer{12, "Christina McArdle", "52.986375", "-6.043701"})

	results := FilterCustomerByDistance(cs, 100)

	for i := 0; i < len(results); i++ {
		if results[i] != expected[i] {
			t.Errorf("Expected id: %v, Got: %v", expected[i].UserID, results[i].UserID)
		}
	}

}

func TestSort(t *testing.T) {
	cs := Customers{}
	expected := Customers{}

	cs = append(cs, Customer{12, "Christina McArdle", "52.986375", "-6.043701"})
	cs = append(cs, Customer{1, "Alice Cahill", "51.92893", "-10.27699"})
	cs = append(cs, Customer{2, "Ian McArdle", "51.8856167", "-10.4240951"})

	expected = append(expected, Customer{1, "Alice Cahill", "51.92893", "-10.27699"})
	expected = append(expected, Customer{2, "Ian McArdle", "51.8856167", "-10.4240951"})
	expected = append(expected, Customer{12, "Christina McArdle", "52.986375", "-6.043701"})

	cs.Sort()

	for i := 0; i < len(cs); i++ {
		if cs[i] != expected[i] {
			t.Errorf("Expected id: %v, Got: %v", expected[i].UserID, cs[i].UserID)
		}
	}
}

func TestPrint(t *testing.T) {

	os.Remove("_test_output.txt")

	cs := Customers{}

	cs = append(cs, Customer{1, "Alice Cahill", "51.92893", "-10.27699"})
	cs = append(cs, Customer{2, "Ian McArdle", "51.8856167", "-10.4240951"})
	cs = append(cs, Customer{12, "Christina McArdle", "52.986375", "-6.043701"})

	expectedRows := []string{"Alice Cahill,1", "Ian McArdle,2", "Christina McArdle,12"}

	cs.Print("_test_output.txt")

	bs, _ := ioutil.ReadFile("_test_output.txt")

	rows := strings.Split(string(bs), "\n")

	for i := 0; i < len(rows); i++ {
		if rows[i] != expectedRows[i] {
			t.Errorf("Expected id: %v, Got: %v", expectedRows[i], rows[i])
		}
	}

	os.Remove("_test_output.txt")
}
