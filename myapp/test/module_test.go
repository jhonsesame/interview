package tests

import (
	"fmt"
	"myapp/calc"
	"testing"
)

func TestRandmsg(t *testing.T) {
	randSeq := calc.Randmsg()
	if len(randSeq) != 32 {
		t.Error("Message length error!!!")
	}
}

func TestSignVerify(t *testing.T) {
	instance_privatekey := "0xfad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19"
	//instance_publickey := "0x049a7df67f79246283fdc93af76d4f8cdd62c4886e8cd870944e817dd0b97934fdd7719d0810951e03418205868a5c1b40b192451367f28e0088dd75e15de40c05"
	instance_address := "0x96216849c49358B10257cb55b28eA603c874b05E"
	instance_plainmessage := "f5rg5r7gH468gr"

	getaddress := calc.Getaddress(instance_privatekey)
	fmt.Println("get", getaddress)

	signedmsg := calc.SignMessage(instance_plainmessage, instance_privatekey)
	fmt.Println("signedmsg:", signedmsg)

	ve := calc.VerifySignature(instance_plainmessage, instance_address, signedmsg)
	if ve != true {
		t.Error("Verified fail!!!")
	}
}
