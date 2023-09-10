package main

import (
	"encoding/json"
	"net/http"
)

func getDetailToko(w http.ResponseWriter, r http.Request) {
	// already hit data from DB, get detail toko

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"name":           "Byyy Store",
		"total_products": "100",
	})
}

func getAllToko() {

}

func main() {

}
