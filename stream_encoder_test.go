package traininggolangjson

import (
	"encoding/json"
	"os"
	"testing"
)

func TestStreamingEncoder(t *testing.T) {

	writer, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "Ravi",
		MiddleName: "Mukti",
		LastName:   "Hartadi",
	}

	encoder.Encode(customer)
}
