package traininggolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecode(t *testing.T) {

	jsonString := `{"FirstName":"Ravi","MiddleName":"Mukti","LastName":"Hartadi"}`

	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	json.Unmarshal(jsonBytes, customer)

	fmt.Println(customer)
}
