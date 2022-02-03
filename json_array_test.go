package traininggolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonArray(t *testing.T) {
	customer := Customer{
		FirstName:  "Ravi",
		MiddleName: "Mukti",
		LastName:   "Hartadi",
		Hobbies:    []string{"Gaming", "Internet", "Coding"},
	}

	bytes, _ := json.Marshal(customer)

	fmt.Println(string(bytes))
}

func TestJsonArrayDecode(t *testing.T) {

	jsonString := `{"FirstName":"Ravi","MiddleName":"Mukti","LastName":"Hartadi","Hobbies":["Gaming","Internet","Coding"]}`

	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	json.Unmarshal(jsonBytes, customer)

	fmt.Println(customer)
}

func TestJsonArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName:  "Ravi",
		MiddleName: "Mukti",
		LastName:   "Hartadi",
		Hobbies:    []string{"Gaming", "Internet", "Coding"},
		Addresses: []Address{
			{
				Street:     "Jalan Satu",
				Country:    "Indonesia",
				PostalCode: "1",
			},
			{
				Street:     "Jalan Dua",
				Country:    "Indonesia",
				PostalCode: "2",
			},
		},
	}

	bytes, _ := json.Marshal(customer)

	fmt.Println(string(bytes))
}

func TestJsonArrayComplexDecode(t *testing.T) {

	jsonString := `{"FirstName":"Ravi","MiddleName":"Mukti","LastName":"Hartadi","Hobbies":["Gaming","Internet","Coding"],"Addresses":[{"Street":"Jalan Satu","Country":"Indonesia","PostalCode":"1"},{"Street":"Jalan Dua","Country":"Indonesia","PostalCode":"2"}]}`

	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	json.Unmarshal(jsonBytes, customer)

	fmt.Println(customer)
}

func TestOnlyJsonArrayComplexDecode(t *testing.T) {

	jsonString := `[{"Street":"Jalan Satu","Country":"Indonesia","PostalCode":"1"},{"Street":"Jalan Dua","Country":"Indonesia","PostalCode":"2"}]`

	jsonBytes := []byte(jsonString)

	address := &[]Address{}

	json.Unmarshal(jsonBytes, address)

	fmt.Println(address)
}
