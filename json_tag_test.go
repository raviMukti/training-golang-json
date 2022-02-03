package traininggolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func TestJsonTagEncode(t *testing.T) {

	product := Product{
		Id:       "P001",
		Name:     "Macbook Pro",
		ImageURL: "http://localhost/image",
	}

	bytes, _ := json.Marshal(product)

	fmt.Println(string(bytes))
}

func TestJsonTagDecode(t *testing.T) {

	jsonString := `{"id":"P001","name":"Macbook Pro","image_url":"http://localhost/image"}`

	jsonBytes := []byte(jsonString)

	product := &Product{}

	json.Unmarshal(jsonBytes, product)

	fmt.Println(product)
}
