package iso

import (
	"encoding/xml"
	"fmt"
	"regexp"
	"strings"

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
	//"github.com/figassis/bankiso/iso20022/secl"
	//"github.com/figassis/bankiso/iso20022/semt"
	//"github.com/figassis/bankiso/iso20022/sese"
	//"github.com/figassis/bankiso/iso20022/admi"
	//"github.com/figassis/bankiso/iso20022/colr"
	//"github.com/figassis/bankiso/iso20022/fxtr"
	//"github.com/figassis/bankiso/iso20022/reda"
	//"github.com/figassis/bankiso/iso20022/seev"
	//"github.com/figassis/bankiso/iso20022/setr"
	//"github.com/figassis/bankiso/iso20022/tsin"
	//"github.com/figassis/bankiso/iso20022/tsmt"
	//"github.com/figassis/bankiso/iso20022/tsrv"
)

var ISO20022Registry map[string]interface{} = map[string]interface{}{
	"acmt.007.001.02": &acmt.Document00700102{},
	"acmt.008.001.02": &acmt.Document00800102{},
	"acmt.009.001.02": &acmt.Document00900102{},
	"acmt.010.001.02": &acmt.Document01000102{},
	"acmt.011.001.02": &acmt.Document01100102{},
	"acmt.012.001.02": &acmt.Document01200102{},
	"acmt.013.001.02": &acmt.Document01300102{},
	"acmt.014.001.02": &acmt.Document01400102{},
	"acmt.015.001.02": &acmt.Document01500102{},
	"acmt.016.001.02": &acmt.Document01600102{},
	"acmt.017.001.02": &acmt.Document01700102{},
	"acmt.018.001.02": &acmt.Document01800102{},
	"acmt.019.001.02": &acmt.Document01900102{},
	"acmt.020.001.02": &acmt.Document02000102{},
	"acmt.021.001.02": &acmt.Document02100102{},
	"acmt.022.001.02": &acmt.Document02200102{},
	"acmt.023.001.02": &acmt.Document02300102{},
	"acmt.024.001.02": &acmt.Document02400102{},
	"auth.001.001.01": &auth.Document00100101{},
	"auth.002.001.01": &auth.Document00200101{},
	"auth.003.001.01": &auth.Document00300101{},
	"auth.018.001.01": &auth.Document01800101{},
	"auth.019.001.01": &auth.Document01900101{},
	"auth.020.001.01": &auth.Document02000101{},
	"auth.021.001.01": &auth.Document02100101{},
	"auth.022.001.01": &auth.Document02200101{},
	"auth.023.001.01": &auth.Document02300101{},
	"auth.024.001.01": &auth.Document02400101{},
	"auth.025.001.01": &auth.Document02500101{},
	"auth.026.001.01": &auth.Document02600101{},
	"auth.027.001.01": &auth.Document02700101{},
	"caaa.001.001.05": &caaa.Document00100105{},
	"caaa.002.001.05": &caaa.Document00200105{},
	"caaa.003.001.05": &caaa.Document00300105{},
	"caaa.004.001.05": &caaa.Document00400105{},
	"caaa.005.001.05": &caaa.Document00500105{},
	"caaa.006.001.05": &caaa.Document00600105{},
	"caaa.007.001.05": &caaa.Document00700105{},
	"caaa.008.001.05": &caaa.Document00800105{},
	"caaa.009.001.05": &caaa.Document00900105{},
	"caaa.010.001.05": &caaa.Document01000105{},
	"caaa.011.001.05": &caaa.Document01100105{},
	"caaa.012.001.05": &caaa.Document01200105{},
	"caaa.013.001.05": &caaa.Document01300105{},
	"caaa.014.001.05": &caaa.Document01400105{},
	"caaa.015.001.05": &caaa.Document01500105{},
	"caaa.016.001.03": &caaa.Document01600103{},
	"caaa.017.001.03": &caaa.Document01700103{},
	"caam.001.001.02": &caam.Document00100102{},
	"caam.002.001.02": &caam.Document00200102{},
	"caam.003.001.02": &caam.Document00300102{},
	"caam.004.001.02": &caam.Document00400102{},
	"caam.005.001.02": &caam.Document00500102{},
	"caam.006.001.02": &caam.Document00600102{},
	"caam.007.001.01": &caam.Document00700101{},
	"caam.008.001.01": &caam.Document00800101{},
	"caam.009.001.02": &caam.Document00900102{},
	"caam.010.001.02": &caam.Document01000102{},
	"caam.011.001.01": &caam.Document01100101{},
	"caam.012.001.01": &caam.Document01200101{},
	"cain.001.001.01": &cain.Document00100101{},
	"cain.002.001.01": &cain.Document00200101{},
	"cain.003.001.01": &cain.Document00300101{},
	"cain.004.001.01": &cain.Document00400101{},
	"cain.005.001.01": &cain.Document00500101{},
	"cain.006.001.01": &cain.Document00600101{},
	"cain.007.001.01": &cain.Document00700101{},
	"cain.008.001.01": &cain.Document00800101{},
	"cain.009.001.01": &cain.Document00900101{},
	"cain.010.001.01": &cain.Document01000101{},
	"cain.011.001.01": &cain.Document01100101{},
	"cain.012.001.01": &cain.Document01200101{},
	"cain.013.001.01": &cain.Document01300101{},
	"camt.026.001.04": &camt.Document02600104{},
	"camt.027.001.04": &camt.Document02700104{},
	"camt.028.001.06": &camt.Document02800106{},
	"camt.029.001.06": &camt.Document02900106{},
	"camt.030.001.04": &camt.Document03000104{},
	"camt.031.001.04": &camt.Document03100104{},
	"camt.032.001.03": &camt.Document03200103{},
	"camt.033.001.04": &camt.Document03300104{},
	"camt.034.001.04": &camt.Document03400104{},
	"camt.035.001.03": &camt.Document03500103{},
	"camt.036.001.03": &camt.Document03600103{},
	"camt.037.001.04": &camt.Document03700104{},
	"camt.038.001.03": &camt.Document03800103{},
	"camt.039.001.04": &camt.Document03900104{},
	"camt.052.001.06": &camt.Document05200106{},
	"camt.053.001.06": &camt.Document05300106{},
	"camt.054.001.06": &camt.Document05400106{},
	"camt.055.001.05": &camt.Document05500105{},
	"camt.056.001.05": &camt.Document05600105{},
	"camt.057.001.05": &camt.Document05700105{},
	"camt.058.001.05": &camt.Document05800105{},
	"camt.059.001.05": &camt.Document05900105{},
	"camt.060.001.03": &camt.Document06000103{},
	"camt.086.001.02": &camt.Document08600102{},
	//"camt.087.001.03": &camt.Document08700103{}, //had a problem generating message definition. Got amigous message definition RequestToModifyPaymentV03
	"catm.001.001.05": &catm.Document00100105{},
	"catm.002.001.05": &catm.Document00200105{},
	"catm.003.001.05": &catm.Document00300105{},
	"catm.004.001.04": &catm.Document00400104{},
	"catm.005.001.02": &catm.Document00500102{},
	"catm.006.001.02": &catm.Document00600102{},
	"catm.007.001.01": &catm.Document00700101{},
	"catm.008.001.01": &catm.Document00800101{},
	"catp.001.001.02": &catp.Document00100102{},
	"catp.002.001.02": &catp.Document00200102{},
	"catp.003.001.02": &catp.Document00300102{},
	"catp.004.001.02": &catp.Document00400102{},
	"catp.005.001.02": &catp.Document00500102{},
	"catp.006.001.02": &catp.Document00600102{},
	"catp.007.001.02": &catp.Document00700102{},
	"catp.008.001.02": &catp.Document00800102{},
	"catp.009.001.02": &catp.Document00900102{},
	"catp.010.001.02": &catp.Document01000102{},
	"catp.011.001.02": &catp.Document01100102{},
	"catp.012.001.01": &catp.Document01200101{},
	"catp.013.001.01": &catp.Document01300101{},
	"catp.014.001.01": &catp.Document01400101{},
	"catp.015.001.01": &catp.Document01500101{},
	"catp.016.001.01": &catp.Document01600101{},
	"catp.017.001.01": &catp.Document01700101{},
	"head.001.001.01": &head.Document00100101{},
	"pacs.002.001.07": &pacs.Document00200107{},
	"pacs.003.001.06": &pacs.Document00300106{},
	"pacs.004.001.06": &pacs.Document00400106{},
	"pacs.007.001.06": &pacs.Document00700106{},
	"pacs.008.001.06": &pacs.Document00800106{},
	"pacs.009.001.06": &pacs.Document00900106{},
	"pacs.010.001.02": &pacs.Document01000102{},
	"pain.001.001.07": &pain.Document00100107{},
	"pain.002.001.07": &pain.Document00200107{},
	"pain.007.001.06": &pain.Document00700106{},
	"pain.008.001.06": &pain.Document00800106{},
	"pain.009.001.04": &pain.Document00900104{},
	"pain.010.001.04": &pain.Document01000104{},
	"pain.011.001.04": &pain.Document01100104{},
	"pain.012.001.04": &pain.Document01200104{},
	"pain.013.001.05": &pain.Document01300105{},
	"pain.014.001.05": &pain.Document01400105{},
	"remt.001.001.02": &remt.Document00100102{},
	"remt.002.001.01": &remt.Document00200101{},
}

func makeISO20022(code, message string) (valid bool, result ISOMessage) {

	val, ok := ISO20022Registry[code]
	if !ok {
		return
	}

	err := xml.Unmarshal([]byte(message), &val)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	valid = true
	result = val.(ISOMessage)
	return
}

func makeISO8583(code, message string) (valid bool, result ISOMessage) {
	return
}

//Basic XML version 1.0 validation via header regex
//Eg. if XML, process as iso20022, else, check if it's an iso8583 message
func isXML(message string) bool {
	//Check if string is valid XML
	header1 := regexp.QuoteMeta("<?xml version='1.0' encoding='UTF-8'?>")
	header2 := regexp.QuoteMeta("<?xml version=\"1.0\" encoding=\"UTF-8\"?>")
	valid, err := regexp.MatchString("^("+header1+"|"+header2+")", message)

	if err != nil {
		fmt.Println("xml.isXML: Could not process file. ", err)
		return false
	}

	if valid {
		//@TODO: check if XML is well formed
		return true
	}

	return false
}

//Returns a boolean and the message type, eg. true, "acmt"
//@TODO: implement full xsd schema validation
func ValidateISO20022(message string) (valid bool, domain, code string) {

	isXml := isXML(message)
	if !isXml {
		return
	}

	v := Document{Format: ""}

	//@TODO To improve performance, consider using a string split or regex match instead
	err := xml.Unmarshal([]byte(message), &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	//Match all iso20022 message types
	search := `^urn:iso:std:iso:20022:tech:xsd:(acmt|admi|auth|caaa|caam|camt|catm|catp|colr|fxtr|pacs|pain|reda|remt|secl|seev|semt|sese|setr|tsin|tsmt|tsrv|head)\.[0-9]{3}\.[0-9]{3}\.[0-9]{2}$`
	match, err := regexp.MatchString(search, v.Format)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	if !match {
		return
	}

	//Eg. pain.009.001.04
	code = strings.Split(v.Format, "xsd:")[1]

	//Eg. pain
	domain = strings.Split(code, ".")[0]

	valid = true

	return
}

//Not yet implemented
func ValidateISO8583(message string) (valid bool, domain, code string) {
	return
}

//Determines if string is valid ISO 20022 or 8583
func ValidateISO(message string) (valid bool, standard, domain, code string) {

	iso20022, iso20022Domain, iso20022Code := ValidateISO20022(message)
	iso8583, iso8583Domain, iso8583Code := ValidateISO8583(message)

	if iso20022 && iso8583 {
		return
	}

	if iso20022 {
		standard = "iso20022"
		domain = iso20022Domain
		code = iso20022Code
		valid = true
	}

	if iso8583 {
		standard = "iso8583"
		domain = iso8583Domain
		code = iso8583Code
		valid = true
	}

	return
}

//Return ISO Message struct containing message content as well as
//identifiers to facilitate type casting/assertion and message handling
func Decode(message string) (valid bool, result ISOMessage, standard string, domain string, code string) {

	ok, standard, domain, code := ValidateISO(message)

	if !ok {
		return
	}

	switch standard {
	case "iso20022":
		valid, result = makeISO20022(code, message)
	case "iso8583":
		valid, result = makeISO8583(code, message)
	}
	return
}
