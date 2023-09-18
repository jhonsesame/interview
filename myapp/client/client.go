package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"myapp/calc"
)

func main() {
	targetURL := "http://127.0.0.1:1323"

	// generate private key and address
	defaultprivKey := "0xfad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19"
	defaultaddress := calc.Getaddress(defaultprivKey)
	privatekey := flag.String("sk", defaultprivKey, "please provide secret key")
	address := flag.String("addr", defaultaddress, "please provide wallet address")
	flag.Parse()

	// get random message and grab session cookie
	resp_json, cookie := calc.GetMessage(targetURL) // get random message
	fmt.Print("GET /get_message: " + resp_json)

	// Pull out message from json response
	var resp_map map[string]interface{}
	json.Unmarshal([]byte(resp_json), &resp_map)
	message := resp_map["message"].(string)

	// sign message
	signature := calc.SignMessage(message, *privatekey)

	// verify signature, maintain session using cookie
	result := calc.PostVerify(targetURL, *address, signature, cookie)
	fmt.Print("POST /verify: " + result + "\n")

	return
}
