package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"

	"github.com/davecgh/go-spew/spew"
	"github.com/figassis/bankiso/iso"
	"github.com/figassis/bankiso/isomessages/acmt"
	"github.com/figassis/bankiso/isomessages/head"
	"github.com/figassis/bankiso/isomessages/pacs"
	"github.com/figassis/bankiso/isomessages/pain"
)

func main() {
	var file string
	flag.StringVar(&file, "f", "", "Specify a path to the iso20022 xml message file.")
	flag.Parse()

	if file == "" {
		fmt.Println("Usage: bankiso -f path-to-xml-file")
		os.Exit(0)
	}

	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	example := string(data)

	message, domain, code, err := iso.Decode(example)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Data Type: %v, Business Domain: %v, Message Code: %v\n", reflect.TypeOf(message), domain, code)

	switch message.(type) {
	case *acmt.Document00700102:
		content := message.(*acmt.Document00700102)
		spew.Dump(content)
	case *acmt.Document00800102:
		content := message.(*acmt.Document00800102)
		spew.Dump(content)
	case *acmt.Document00900102:
		content := message.(*acmt.Document00900102)
		spew.Dump(content)
	case *acmt.Document01000102:
		content := message.(*acmt.Document01000102)
		spew.Dump(content)
	case *acmt.Document01100102:
		content := message.(*acmt.Document01100102)
		spew.Dump(content)
	case *acmt.Document01200102:
		content := message.(*acmt.Document01200102)
		spew.Dump(content)
	case *acmt.Document01300102:
		content := message.(*acmt.Document01300102)
		spew.Dump(content)
	case *acmt.Document01400102:
		content := message.(*acmt.Document01400102)
		spew.Dump(content)
	case *acmt.Document01500102:
		content := message.(*acmt.Document01500102)
		spew.Dump(content)
	case *acmt.Document01600102:
		content := message.(*acmt.Document01600102)
		spew.Dump(content)
	case *acmt.Document01700102:
		content := message.(*acmt.Document01700102)
		spew.Dump(content)
	case *acmt.Document01800102:
		content := message.(*acmt.Document01800102)
		spew.Dump(content)
	case *acmt.Document01900102:
		content := message.(*acmt.Document01900102)
		spew.Dump(content)
	case *acmt.Document02000102:
		content := message.(*acmt.Document02000102)
		spew.Dump(content)
	case *acmt.Document02100102:
		content := message.(*acmt.Document02100102)
		spew.Dump(content)
	case *acmt.Document02200102:
		content := message.(*acmt.Document02200102)
		spew.Dump(content)
	case *acmt.Document02300102:
		content := message.(*acmt.Document02300102)
		spew.Dump(content)
	case *acmt.Document02400102:
		content := message.(*acmt.Document02400102)
		spew.Dump(content)
	case *head.Document00100101:
		content := message.(*head.Document00100101)
		spew.Dump(content)
	case *pacs.Document00200107:
		content := message.(*pacs.Document00200107)
		spew.Dump(content)
	case *pacs.Document00300106:
		content := message.(*pacs.Document00300106)
		spew.Dump(content)
	case *pacs.Document00400106:
		content := message.(*pacs.Document00400106)
		spew.Dump(content)
	case *pacs.Document00700106:
		content := message.(*pacs.Document00700106)
		spew.Dump(content)
	case *pacs.Document00800106:
		content := message.(*pacs.Document00800106)
		spew.Dump(content)
	case *pacs.Document00900106:
		content := message.(*pacs.Document00900106)
		spew.Dump(content)
	case *pacs.Document01000102:
		content := message.(*pacs.Document01000102)
		spew.Dump(content)
	case *pain.Document00100107:
		content := message.(*pain.Document00100107)
		spew.Dump(content)
	case *pain.Document00200107:
		content := message.(*pain.Document00200107)
		spew.Dump(content)
	case *pain.Document00700106:
		content := message.(*pain.Document00700106)
		spew.Dump(content)
	case *pain.Document00800106:
		content := message.(*pain.Document00800106)
		spew.Dump(content)
	case *pain.Document00900104:
		content := message.(*pain.Document00900104)
		spew.Dump(content)
	case *pain.Document01000104:
		content := message.(*pain.Document01000104)
		spew.Dump(content)
	case *pain.Document01100104:
		content := message.(*pain.Document01100104)
		spew.Dump(content)
	case *pain.Document01200104:
		content := message.(*pain.Document01200104)
		spew.Dump(content)
	case *pain.Document01300105:
		content := message.(*pain.Document01300105)
		spew.Dump(content)
	case *pain.Document01400105:
		content := message.(*pain.Document01400105)
		spew.Dump(content)
	default:
		fmt.Println("Unknown Message format")
		os.Exit(0)
	}

}
