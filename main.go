package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"

	"github.com/figassis/bankiso/iso"
	"github.com/figassis/bankiso/iso20022/pain"
)

func main() {
	gopath := os.Getenv("GOPATH")
	file := gopath + "/src/github.com/figassis/bankiso/msg-examples/payments/pain.009.001.04/Business example1 pain.009.001.04.xml"
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}

	example := string(data)

	valid, message, standard, domain, code := iso.Decode(example)

	fmt.Printf("Valid: %v, Data Type: %v, Standard: %v, Business Domain: %v, Message Code: %v\n", valid, reflect.TypeOf(message), standard, domain, code)

	//Type assertion to use in specific function
	content := message.(*pain.Document00900104)

	fmt.Println("Message content: ", content)

}
