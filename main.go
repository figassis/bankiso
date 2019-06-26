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
	"github.com/figassis/bankiso/iso20022/acmt"
	"github.com/figassis/bankiso/iso20022/auth"
	"github.com/figassis/bankiso/iso20022/caaa"
	"github.com/figassis/bankiso/iso20022/caam"
	"github.com/figassis/bankiso/iso20022/cain"
	"github.com/figassis/bankiso/iso20022/camt"
	"github.com/figassis/bankiso/iso20022/catm"
	"github.com/figassis/bankiso/iso20022/catp"
	"github.com/figassis/bankiso/iso20022/head"
	"github.com/figassis/bankiso/iso20022/pacs"
	"github.com/figassis/bankiso/iso20022/pain"
	"github.com/figassis/bankiso/iso20022/remt"
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
	case *auth.Document00100101:
		content := message.(*auth.Document00100101)
		spew.Dump(content)
	case *auth.Document00200101:
		content := message.(*auth.Document00200101)
		spew.Dump(content)
	case *auth.Document00300101:
		content := message.(*auth.Document00300101)
		spew.Dump(content)
	case *auth.Document01800101:
		content := message.(*auth.Document01800101)
		spew.Dump(content)
	case *auth.Document01900101:
		content := message.(*auth.Document01900101)
		spew.Dump(content)
	case *auth.Document02000101:
		content := message.(*auth.Document02000101)
		spew.Dump(content)
	case *auth.Document02100101:
		content := message.(*auth.Document02100101)
		spew.Dump(content)
	case *auth.Document02200101:
		content := message.(*auth.Document02200101)
		spew.Dump(content)
	case *auth.Document02300101:
		content := message.(*auth.Document02300101)
		spew.Dump(content)
	case *auth.Document02400101:
		content := message.(*auth.Document02400101)
		spew.Dump(content)
	case *auth.Document02500101:
		content := message.(*auth.Document02500101)
		spew.Dump(content)
	case *auth.Document02600101:
		content := message.(*auth.Document02600101)
		spew.Dump(content)
	case *auth.Document02700101:
		content := message.(*auth.Document02700101)
		spew.Dump(content)
	case *caaa.Document00100105:
		content := message.(*caaa.Document00100105)
		spew.Dump(content)
	case *caaa.Document00200105:
		content := message.(*caaa.Document00200105)
		spew.Dump(content)
	case *caaa.Document00300105:
		content := message.(*caaa.Document00300105)
		spew.Dump(content)
	case *caaa.Document00400105:
		content := message.(*caaa.Document00400105)
		spew.Dump(content)
	case *caaa.Document00500105:
		content := message.(*caaa.Document00500105)
		spew.Dump(content)
	case *caaa.Document00600105:
		content := message.(*caaa.Document00600105)
		spew.Dump(content)
	case *caaa.Document00700105:
		content := message.(*caaa.Document00700105)
		spew.Dump(content)
	case *caaa.Document00800105:
		content := message.(*caaa.Document00800105)
		spew.Dump(content)
	case *caaa.Document00900105:
		content := message.(*caaa.Document00900105)
		spew.Dump(content)
	case *caaa.Document01000105:
		content := message.(*caaa.Document01000105)
		spew.Dump(content)
	case *caaa.Document01100105:
		content := message.(*caaa.Document01100105)
		spew.Dump(content)
	case *caaa.Document01200105:
		content := message.(*caaa.Document01200105)
		spew.Dump(content)
	case *caaa.Document01300105:
		content := message.(*caaa.Document01300105)
		spew.Dump(content)
	case *caaa.Document01400105:
		content := message.(*caaa.Document01400105)
		spew.Dump(content)
	case *caaa.Document01500105:
		content := message.(*caaa.Document01500105)
		spew.Dump(content)
	case *caaa.Document01600103:
		content := message.(*caaa.Document01600103)
		spew.Dump(content)
	case *caaa.Document01700103:
		content := message.(*caaa.Document01700103)
		spew.Dump(content)
	case *caam.Document00100102:
		content := message.(*caam.Document00100102)
		spew.Dump(content)
	case *caam.Document00200102:
		content := message.(*caam.Document00200102)
		spew.Dump(content)
	case *caam.Document00300102:
		content := message.(*caam.Document00300102)
		spew.Dump(content)
	case *caam.Document00400102:
		content := message.(*caam.Document00400102)
		spew.Dump(content)
	case *caam.Document00500102:
		content := message.(*caam.Document00500102)
		spew.Dump(content)
	case *caam.Document00600102:
		content := message.(*caam.Document00600102)
		spew.Dump(content)
	case *caam.Document00700101:
		content := message.(*caam.Document00700101)
		spew.Dump(content)
	case *caam.Document00800101:
		content := message.(*caam.Document00800101)
		spew.Dump(content)
	case *caam.Document00900102:
		content := message.(*caam.Document00900102)
		spew.Dump(content)
	case *caam.Document01000102:
		content := message.(*caam.Document01000102)
		spew.Dump(content)
	case *caam.Document01100101:
		content := message.(*caam.Document01100101)
		spew.Dump(content)
	case *caam.Document01200101:
		content := message.(*caam.Document01200101)
		spew.Dump(content)
	case *cain.Document00100101:
		content := message.(*cain.Document00100101)
		spew.Dump(content)
	case *cain.Document00200101:
		content := message.(*cain.Document00200101)
		spew.Dump(content)
	case *cain.Document00300101:
		content := message.(*cain.Document00300101)
		spew.Dump(content)
	case *cain.Document00400101:
		content := message.(*cain.Document00400101)
		spew.Dump(content)
	case *cain.Document00500101:
		content := message.(*cain.Document00500101)
		spew.Dump(content)
	case *cain.Document00600101:
		content := message.(*cain.Document00600101)
		spew.Dump(content)
	case *cain.Document00700101:
		content := message.(*cain.Document00700101)
		spew.Dump(content)
	case *cain.Document00800101:
		content := message.(*cain.Document00800101)
		spew.Dump(content)
	case *cain.Document00900101:
		content := message.(*cain.Document00900101)
		spew.Dump(content)
	case *cain.Document01000101:
		content := message.(*cain.Document01000101)
		spew.Dump(content)
	case *cain.Document01100101:
		content := message.(*cain.Document01100101)
		spew.Dump(content)
	case *cain.Document01200101:
		content := message.(*cain.Document01200101)
		spew.Dump(content)
	case *cain.Document01300101:
		content := message.(*cain.Document01300101)
		spew.Dump(content)
	case *camt.Document02600104:
		content := message.(*camt.Document02600104)
		spew.Dump(content)
	case *camt.Document02700104:
		content := message.(*camt.Document02700104)
		spew.Dump(content)
	case *camt.Document02800106:
		content := message.(*camt.Document02800106)
		spew.Dump(content)
	case *camt.Document02900106:
		content := message.(*camt.Document02900106)
		spew.Dump(content)
	case *camt.Document03000104:
		content := message.(*camt.Document03000104)
		spew.Dump(content)
	case *camt.Document03100104:
		content := message.(*camt.Document03100104)
		spew.Dump(content)
	case *camt.Document03200103:
		content := message.(*camt.Document03200103)
		spew.Dump(content)
	case *camt.Document03300104:
		content := message.(*camt.Document03300104)
		spew.Dump(content)
	case *camt.Document03400104:
		content := message.(*camt.Document03400104)
		spew.Dump(content)
	case *camt.Document03500103:
		content := message.(*camt.Document03500103)
		spew.Dump(content)
	case *camt.Document03600103:
		content := message.(*camt.Document03600103)
		spew.Dump(content)
	case *camt.Document03700104:
		content := message.(*camt.Document03700104)
		spew.Dump(content)
	case *camt.Document03800103:
		content := message.(*camt.Document03800103)
		spew.Dump(content)
	case *camt.Document03900104:
		content := message.(*camt.Document03900104)
		spew.Dump(content)
	case *camt.Document05200106:
		content := message.(*camt.Document05200106)
		spew.Dump(content)
	case *camt.Document05300106:
		content := message.(*camt.Document05300106)
		spew.Dump(content)
	case *camt.Document05400106:
		content := message.(*camt.Document05400106)
		spew.Dump(content)
	case *camt.Document05500105:
		content := message.(*camt.Document05500105)
		spew.Dump(content)
	case *camt.Document05600105:
		content := message.(*camt.Document05600105)
		spew.Dump(content)
	case *camt.Document05700105:
		content := message.(*camt.Document05700105)
		spew.Dump(content)
	case *camt.Document05800105:
		content := message.(*camt.Document05800105)
		spew.Dump(content)
	case *camt.Document05900105:
		content := message.(*camt.Document05900105)
		spew.Dump(content)
	case *camt.Document06000103:
		content := message.(*camt.Document06000103)
		spew.Dump(content)
	case *camt.Document08600102:
		content := message.(*camt.Document08600102)
		//	case *camt.Document08700103:
		//		fmt.Println("Not yet implemented. See https://github.com/figassis/isogen/issues/1")
		//		os.Exit(0)
		spew.Dump(content)
	case *catm.Document00100105:
		content := message.(*catm.Document00100105)
		spew.Dump(content)
	case *catm.Document00200105:
		content := message.(*catm.Document00200105)
		spew.Dump(content)
	case *catm.Document00300105:
		content := message.(*catm.Document00300105)
		spew.Dump(content)
	case *catm.Document00400104:
		content := message.(*catm.Document00400104)
		spew.Dump(content)
	case *catm.Document00500102:
		content := message.(*catm.Document00500102)
		spew.Dump(content)
	case *catm.Document00600102:
		content := message.(*catm.Document00600102)
		spew.Dump(content)
	case *catm.Document00700101:
		content := message.(*catm.Document00700101)
		spew.Dump(content)
	case *catm.Document00800101:
		content := message.(*catm.Document00800101)
		spew.Dump(content)
	case *catp.Document00100102:
		content := message.(*catp.Document00100102)
		spew.Dump(content)
	case *catp.Document00200102:
		content := message.(*catp.Document00200102)
		spew.Dump(content)
	case *catp.Document00300102:
		content := message.(*catp.Document00300102)
		spew.Dump(content)
	case *catp.Document00400102:
		content := message.(*catp.Document00400102)
		spew.Dump(content)
	case *catp.Document00500102:
		content := message.(*catp.Document00500102)
		spew.Dump(content)
	case *catp.Document00600102:
		content := message.(*catp.Document00600102)
		spew.Dump(content)
	case *catp.Document00700102:
		content := message.(*catp.Document00700102)
		spew.Dump(content)
	case *catp.Document00800102:
		content := message.(*catp.Document00800102)
		spew.Dump(content)
	case *catp.Document00900102:
		content := message.(*catp.Document00900102)
		spew.Dump(content)
	case *catp.Document01000102:
		content := message.(*catp.Document01000102)
		spew.Dump(content)
	case *catp.Document01100102:
		content := message.(*catp.Document01100102)
		spew.Dump(content)
	case *catp.Document01200101:
		content := message.(*catp.Document01200101)
		spew.Dump(content)
	case *catp.Document01300101:
		content := message.(*catp.Document01300101)
		spew.Dump(content)
	case *catp.Document01400101:
		content := message.(*catp.Document01400101)
		spew.Dump(content)
	case *catp.Document01500101:
		content := message.(*catp.Document01500101)
		spew.Dump(content)
	case *catp.Document01600101:
		content := message.(*catp.Document01600101)
		spew.Dump(content)
	case *catp.Document01700101:
		content := message.(*catp.Document01700101)
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
	case *remt.Document00100102:
		content := message.(*remt.Document00100102)
		spew.Dump(content)
	case *remt.Document00200101:
		content := message.(*remt.Document00200101)
		spew.Dump(content)
	default:
		fmt.Println("Unknown Message format")
		os.Exit(0)
	}

}
