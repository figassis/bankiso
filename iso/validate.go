package iso

import (
	"fmt"
	//	"reflect"
	"encoding/xml"
	"regexp"
	"strings"

	"github.com/figassis/bankiso/iso20022"
)

var ModelRegistry map[string]interface{} = map[string]interface{}{
	"acmt.007.001.02": &iso20022.ACMT00700102{}, /*
		"acmt.008.001.02": &iso20022.ACMT00800102{},
		"acmt.009.001.02": &iso20022.ACMT00900102{},
		"acmt.010.001.02": &iso20022.ACMT01000102{},
		"acmt.011.001.02": &iso20022.ACMT01100102{},
		"acmt.012.001.02": &iso20022.ACMT01200102{},
		"acmt.013.001.02": &iso20022.ACMT01300102{},
		"acmt.014.001.02": &iso20022.ACMT01400102{},
		"acmt.015.001.02": &iso20022.ACMT01500102{},
		"acmt.016.001.02": &iso20022.ACMT01600102{},
		"acmt.017.001.02": &iso20022.ACMT01700102{},
		"acmt.018.001.02": &iso20022.ACMT01800102{},
		"acmt.019.001.02": &iso20022.ACMT01900102{},
		"acmt.020.001.02": &iso20022.ACMT02000102{},
		"acmt.021.001.02": &iso20022.ACMT02100102{},
		"acmt.022.001.02": &iso20022.ACMT02200102{},
		"acmt.023.001.02": &iso20022.ACMT02300102{},
		"acmt.024.001.02": &iso20022.ACMT02400102{},
		"auth.001.001.01": &iso20022.AUTH00100101{},
		"auth.002.001.01": &iso20022.AUTH00200101{},
		"auth.003.001.01": &iso20022.AUTH00300101{},
		"auth.018.001.01": &iso20022.AUTH01800101{},
		"auth.019.001.01": &iso20022.AUTH01900101{},
		"auth.020.001.01": &iso20022.AUTH02000101{},
		"auth.021.001.01": &iso20022.AUTH02100101{},
		"auth.022.001.01": &iso20022.AUTH02200101{},
		"auth.023.001.01": &iso20022.AUTH02300101{},
		"auth.024.001.01": &iso20022.AUTH02400101{},
		"auth.025.001.01": &iso20022.AUTH02500101{},
		"auth.026.001.01": &iso20022.AUTH02600101{},
		"auth.027.001.01": &iso20022.AUTH02700101{},
		"caaa.001.001.05": &iso20022.CAAA00100105{},
		"caaa.002.001.05": &iso20022.CAAA00200105{},
		"caaa.003.001.05": &iso20022.CAAA00300105{},
		"caaa.004.001.05": &iso20022.CAAA00400105{},
		"caaa.005.001.05": &iso20022.CAAA00500105{},
		"caaa.006.001.05": &iso20022.CAAA00600105{},
		"caaa.007.001.05": &iso20022.CAAA00700105{},
		"caaa.008.001.05": &iso20022.CAAA00800105{},
		"caaa.009.001.05": &iso20022.CAAA00900105{},
		"caaa.010.001.05": &iso20022.CAAA01000105{},
		"caaa.011.001.05": &iso20022.CAAA01100105{},
		"caaa.012.001.05": &iso20022.CAAA01200105{},
		"caaa.013.001.05": &iso20022.CAAA01300105{},
		"caaa.014.001.05": &iso20022.CAAA01400105{},
		"caaa.015.001.05": &iso20022.CAAA01500105{},
		"caaa.016.001.03": &iso20022.CAAA01600103{},
		"caaa.017.001.03": &iso20022.CAAA01700103{},
		"caam.001.001.02": &iso20022.CAAM00100102{},
		"caam.002.001.02": &iso20022.CAAM00200102{},
		"caam.003.001.02": &iso20022.CAAM00300102{},
		"caam.004.001.02": &iso20022.CAAM00400102{},
		"caam.005.001.02": &iso20022.CAAM00500102{},
		"caam.006.001.02": &iso20022.CAAM00600102{},
		"caam.007.001.01": &iso20022.CAAM00700101{},
		"caam.008.001.01": &iso20022.CAAM00800101{},
		"caam.009.001.02": &iso20022.CAAM00900102{},
		"caam.010.001.02": &iso20022.CAAM01000102{},
		"caam.011.001.01": &iso20022.CAAM01100101{},
		"caam.012.001.01": &iso20022.CAAM01200101{},
		"cain.001.001.01": &iso20022.CAIN00100101{},
		"cain.002.001.01": &iso20022.CAIN00200101{},
		"cain.003.001.01": &iso20022.CAIN00300101{},
		"cain.004.001.01": &iso20022.CAIN00400101{},
		"cain.005.001.01": &iso20022.CAIN00500101{},
		"cain.006.001.01": &iso20022.CAIN00600101{},
		"cain.007.001.01": &iso20022.CAIN00700101{},
		"cain.008.001.01": &iso20022.CAIN00800101{},
		"cain.009.001.01": &iso20022.CAIN00900101{},
		"cain.010.001.01": &iso20022.CAIN01000101{},
		"cain.011.001.01": &iso20022.CAIN01100101{},
		"cain.012.001.01": &iso20022.CAIN01200101{},
		"cain.013.001.01": &iso20022.CAIN01300101{},
		"camt.026.001.04": &iso20022.CAMT02600104{},
		"camt.027.001.04": &iso20022.CAMT02700104{},
		"camt.028.001.06": &iso20022.CAMT02800106{},
		"camt.029.001.06": &iso20022.CAMT02900106{},
		"camt.030.001.04": &iso20022.CAMT03000104{},
		"camt.031.001.04": &iso20022.CAMT03100104{},
		"camt.032.001.03": &iso20022.CAMT03200103{},
		"camt.033.001.04": &iso20022.CAMT03300104{},
		"camt.034.001.04": &iso20022.CAMT03400104{},
		"camt.035.001.03": &iso20022.CAMT03500103{},
		"camt.036.001.03": &iso20022.CAMT03600103{},
		"camt.037.001.04": &iso20022.CAMT03700104{},
		"camt.038.001.03": &iso20022.CAMT03800103{},
		"camt.039.001.04": &iso20022.CAMT03900104{},
		"camt.052.001.06": &iso20022.CAMT05200106{},
		"camt.053.001.06": &iso20022.CAMT05300106{},
		"camt.054.001.06": &iso20022.CAMT05400106{},
		"camt.055.001.05": &iso20022.CAMT05500105{},
		"camt.056.001.05": &iso20022.CAMT05600105{},
		"camt.057.001.05": &iso20022.CAMT05700105{},
		"camt.058.001.05": &iso20022.CAMT05800105{},
		"camt.059.001.05": &iso20022.CAMT05900105{},
		"camt.060.001.03": &iso20022.CAMT06000103{},
		"camt.086.001.02": &iso20022.CAMT08600102{},
		"camt.087.001.03": &iso20022.CAMT08700103{},
		"catm.001.001.05": &iso20022.CATM00100105{},
		"catm.002.001.05": &iso20022.CATM00200105{},
		"catm.003.001.05": &iso20022.CATM00300105{},
		"catm.004.001.04": &iso20022.CATM00400104{},
		"catm.005.001.02": &iso20022.CATM00500102{},
		"catm.006.001.02": &iso20022.CATM00600102{},
		"catm.007.001.01": &iso20022.CATM00700101{},
		"catm.008.001.01": &iso20022.CATM00800101{},
		"catp.001.001.02": &iso20022.CATP00100102{},
		"catp.002.001.02": &iso20022.CATP00200102{},
		"catp.003.001.02": &iso20022.CATP00300102{},
		"catp.004.001.02": &iso20022.CATP00400102{},
		"catp.005.001.02": &iso20022.CATP00500102{},
		"catp.006.001.02": &iso20022.CATP00600102{},
		"catp.007.001.02": &iso20022.CATP00700102{},
		"catp.008.001.02": &iso20022.CATP00800102{},
		"catp.009.001.02": &iso20022.CATP00900102{},
		"catp.010.001.02": &iso20022.CATP01000102{},
		"catp.011.001.02": &iso20022.CATP01100102{},
		"catp.012.001.01": &iso20022.CATP01200101{},
		"catp.013.001.01": &iso20022.CATP01300101{},
		"catp.014.001.01": &iso20022.CATP01400101{},
		"catp.015.001.01": &iso20022.CATP01500101{},
		"catp.016.001.01": &iso20022.CATP01600101{},
		"catp.017.001.01": &iso20022.CATP01700101{},
		"head.001.001.01": &iso20022.HEAD00100101{},
		"pacs.002.001.07": &iso20022.PACS00200107{},
		"pacs.003.001.06": &iso20022.PACS00300106{},
		"pacs.004.001.06": &iso20022.PACS00400106{},
		"pacs.007.001.06": &iso20022.PACS00700106{},
		"pacs.008.001.06": &iso20022.PACS00800106{},
		"pacs.009.001.06": &iso20022.PACS00900106{},
		"pacs.010.001.02": &iso20022.PACS01000102{},
		"pain.001.001.07": &iso20022.PAIN00100107{},
		"pain.002.001.07": &iso20022.PAIN00200107{},
		"pain.007.001.06": &iso20022.PAIN00700106{},
		"pain.008.001.06": &iso20022.PAIN00800106{},
	*/"pain.009.001.04": &iso20022.PAIN00900104{}, /*
		"pain.010.001.04": &iso20022.PAIN01000104{},
		"pain.011.001.04": &iso20022.PAIN01100104{},
		"pain.012.001.04": &iso20022.PAIN01200104{},
		"pain.013.001.05": &iso20022.PAIN01300105{},
		"pain.014.001.05": &iso20022.PAIN01400105{},
		"remt.001.001.02": &iso20022.REMT00100102{},
		"remt.002.001.01": &iso20022.REMT00200101{},
	*/
}

func makeInstance(message string) (result ISOMessage, valid bool) {

	val, ok := ModelRegistry[message]
	if !ok {
		return
	}

	valid = true
	result = val.(ISOMessage)
	return
}

//Basic XML version 1.0 validation via header redex
//Indented to use as a mode switch
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

//To call after maxing sure message is calling ValidXML
//returns a boolean and the message type, eg. true, "acmt"
//@TODO: implement full xsd schema validation
func Validate20022(message string) (iso bool, domain, code string) {

	valid := isXML(message)
	if !valid {
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
	valid, err = regexp.MatchString(search, v.Format)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	if !valid {
		return
	}

	//Eg. pain.009.001.04
	code = strings.Split(v.Format, "xsd:")[1]

	//Eg. pain
	domain = strings.Split(code, ".")[0]

	iso = true

	//fmt.Println("Message Code: ", code)
	//fmt.Println("Message Type: ", domain)
	return
}

func Validate8583(message string) (iso bool, domain, code string) {
	return
}

func ValidateISO(message string) (iso bool, std, domain, code string) {

	iso20022, iso20022Domain, iso20022Code := Validate20022(message)
	iso8583, iso8583Domain, iso8583Code := Validate8583(message)

	if iso20022 && iso8583 {
		return
	}

	if iso20022 {
		std = "iso20022"
		domain = iso20022Domain
		code = iso20022Code
		iso = true
	}

	if iso8583 {
		std = "iso8583"
		domain = iso8583Domain
		code = iso8583Code
		iso = true
	}

	return
}

func SelectISO20022(message string) (result ISOMessage, valid bool) {
	v, std, domain, code := ValidateISO(message)

	if !v || std != "iso20022" {
		fmt.Println("SelectISO20022: ", v, std, domain, code)
		return
	}

	result, valid = makeInstance(code)

	return
}

func DecodeISO20022(message string) (result ISOMessage, valid bool) {
	a, ok := SelectISO20022(message)

	if !ok {
		return
	}

	result = a
	valid = true

	return

}
