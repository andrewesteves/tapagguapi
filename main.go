package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/andrewesteves/tapagguapi/models"
	"github.com/andrewesteves/tapagguapi/transformations"
)

func main() {
	http.HandleFunc("/receipt", func(w http.ResponseWriter, r *http.Request) {
		var receipt models.ReceiptXML
		data, err := getURL(r.URL.Query().Get("url"))
		if err != nil {
			log.Printf("Failed to get XML: %v", err)
		}
		xml.Unmarshal(data, &receipt)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(transformations.ReceiptToJSON(receipt))
	})
	log.Fatal(http.ListenAndServe(":9090", nil))
}

func getURL(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("STATUS error: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("READ BODY error: %v", err)
	}

	return data, nil
}
