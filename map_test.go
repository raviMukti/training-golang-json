package traininggolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapDecode(t *testing.T) {

	jsonString := `{"id":"P001","name":"Macbook Pro","image_url":"http://localhost/image"}`

	jsonBytes := []byte(jsonString)

	var result map[string]interface{}

	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["image_url"])
}

func TestMapEncode(t *testing.T) {

	product := map[string]interface{}{
		"id":        "POO1",
		"name":      "Macbook Pro",
		"image_url": "http://localhost/images",
	}

	bytes, _ := json.Marshal(product)

	fmt.Println(string(bytes))
}
