package internal

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

type tokenBalanceResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  string `json:"result"`
}

//GetTokenBalance ...
func GetTokenBalance(walletAddress, contractAddress string) float64 {
	apiURL := "https://api.etherscan.io/api"
	queryParams := map[string]string {
		"module": "account",
		"action": "tokenbalance",
		"tag": "latest",
		"apiKey": "4B31K46SJICNNA8CU2KMVQZ8GYPQFVH3GM",
		"address": walletAddress,
		"contractaddress": contractAddress,
	}

	request, _ := http.NewRequest("GET", apiURL, nil)
	query := request.URL.Query()
	for k,v := range queryParams {
		query.Add(k, v)
	}
	request.URL.RawQuery = query.Encode()

	response, _ := http.DefaultClient.Do(request)

	defer response.Body.Close()
	all, _ := ioutil.ReadAll(response.Body)

	var responseJson tokenBalanceResponse
	json.Unmarshal(all, &responseJson)
	float, _ := strconv.ParseFloat(responseJson.Result, 64)

	return float/ float64(1000000000)
}
