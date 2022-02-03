package traininggolangjson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamingDecoder(t *testing.T) {

	reader, _ := os.Open("Customer.json")

	decoder := json.NewDecoder(reader)

	customer := &Customer{}

	decoder.Decode(customer)

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
}
