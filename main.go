package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"strings"
)

var (
	validIDHash = map[string]string{"eth": "/ETH", "btc": "/BTC", "all": ""}
	url         = "https://api.hitbtc.com/api/2/public/currency"
	port        = "8090"
)

// currency
func currency(w http.ResponseWriter, req *http.Request) {

	// Checking if Request Method  is Get or not
	if req.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "HTTP Method  Not Supported : "+req.Method)
		return
	}

	//Parsing ID
	id := path.Base(req.RequestURI)

	//Checking for Valid ID Option
	val, ok := validIDHash[strings.ToLower(id)]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid Currency ID: "+id)
		return
	}

	// Creating HTTP Get Request
	req, err := http.NewRequest(http.MethodGet, url+val, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Unable to Create HTTP Request")
		return
	}
	// Making HTTP Request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Unable to Make HTTP Request")
		return
	}
	// Reading Request Body
	dataByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Unable to Read Resp Body")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(dataByte))
}

// Usage

// curl http://localhost:8090/currency/eth
// curl http://localhost:8090/currency/btc
// curl http://localhost:8090/currency/all

func main() {

	// Adding Currency Endpoint
	http.HandleFunc("/currency/", currency)
	fmt.Println("Starting HTTP Server :", port)
	// Starting HTTP SERVER at Port 8090
	http.ListenAndServe(":"+port, nil)
}
