package calc

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func GetMessage(targetURL string) (string, *http.Cookie) {
	resp, err := http.Get(targetURL + "/get_message")
	if err != nil {
		fmt.Println("Get Error")
	}
	defer resp.Body.Close()

	// get cookie from response
	cookie := resp.Cookies()[0]
	body, _ := io.ReadAll(resp.Body)

	return string(body), cookie
}

func PostVerify(targetURL string, address string, signedMessage string, cookie *http.Cookie) string {
	//create url encoded content
	params := url.Values{}
	params.Add("address", address)
	params.Add("signedMessage", signedMessage)

	//create a new io.Reader
	postBody := strings.NewReader(params.Encode())

	// create new POST request
	req, _ := http.NewRequest("POST", targetURL+"/verify", postBody)

	//set header for url encoded-data
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	//add cookie
	cookieBack := &http.Cookie{
		Name:  cookie.Name,
		Value: cookie.Value,
	}
	req.AddCookie(cookieBack)

	//send req
	client := http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	//display verified result from server
	body, _ := io.ReadAll(resp.Body)

	return string(body)
}

func Randmsg() string {
	rander := rand.New(rand.NewSource(time.Now().UnixNano()))
	base := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	message := make([]byte, 32)
	for i := range message {
		message[i] = base[rander.Intn(len(base))]
	}
	return string(message)
}

func HashMessage(message string) []byte {
	hash := crypto.Keccak256([]byte(message))

	return hash
}

func PublickeytoAddress(publickey []byte) common.Address {
	hash := crypto.Keccak256(publickey[1:])
	address := hash[12:]

	return common.HexToAddress(hex.EncodeToString(address))
}

func Getaddress(privatehexkey string) string {
	//gey publickey from privatekey
	privatekey, err := crypto.HexToECDSA(privatehexkey[2:])
	if err != nil {
		fmt.Println("Error! Can't convert to ECDSA private key")
		return "error"
	}
	publickey := privatekey.Public().(*ecdsa.PublicKey)

	//derive address from publickey
	address := crypto.PubkeyToAddress(*publickey)
	fmt.Println("get address from pk:", address)

	return address.Hex()
}

func SignMessage(message string, privatehexkey string) string {
	privatekey, err := crypto.HexToECDSA(privatehexkey[2:])
	if err != nil {
		fmt.Println("Error! Can't convert to ECDSA private key")
		return "error"
	}

	signature, err := crypto.Sign(HashMessage(message), privatekey)
	if err != nil {
		fmt.Println("Error! Can't sign message")
		return "error"
	}

	return hexutil.Encode(signature)

}

func VerifySignature(message string, address string, signature string) bool {
	//recover pk from signature and message
	signaturebyte, _ := hexutil.Decode(signature)
	recpublickey, err := crypto.Ecrecover(HashMessage(message), signaturebyte)
	if err != nil {
		fmt.Println("Error ecrecovering signature")
		return false
	}

	//derive address from pk
	recaddress := PublickeytoAddress(recpublickey)

	// Check derived-address and wallet address
	matches := (recaddress.Hex() == address)
	fmt.Println("Provided Wallet Address: ", address)
	fmt.Println("Signature Derived Address: ", recaddress)
	fmt.Println("Verify: ", matches)

	return matches
}
