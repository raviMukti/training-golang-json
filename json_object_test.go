package traininggolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Hobbies    []string
	Addresses  []Address
}

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

func TestJsonObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Ravi",
		MiddleName: "Mukti",
		LastName:   "Hartadi",
	}

	bytes, _ := json.Marshal(customer)

	fmt.Println(string(bytes))
}
