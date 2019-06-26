package iso

import (
	"encoding/xml"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/figassis/bankiso/isomessages/acmt"
	"github.com/figassis/bankiso/isomessages/head"
	"github.com/figassis/bankiso/isomessages/pacs"
	"github.com/figassis/bankiso/isomessages/pain"
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
}

func makeISO20022(code, message string) (result Message, err error) {

	val, ok := ISO20022Registry[code]
	if !ok {
		err = fmt.Errorf("Invalid ISO20022 code %s", code)
		return
	}

	if err = xml.Unmarshal([]byte(message), &val); err != nil {
		return
	}

	if result, ok = val.(Message); !ok {
		err = errors.New("Invalid ISO20022 message")
	}
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
func ValidateISO20022(message string) (domain, code string, err error) {

	isXml := isXML(message)
	if !isXml {
		return "", "", errors.New("message is not valid XML")
	}

	v := Document{Format: ""}

	//@TODO To improve performance, consider using a string split or regex match instead
	if err = xml.Unmarshal([]byte(message), &v); err != nil {
		return
	}

	//Match all iso20022 message types
	search := `^urn:iso:std:iso:20022:tech:xsd:(acmt|admi|auth|caaa|caam|camt|catm|catp|colr|fxtr|pacs|pain|reda|remt|secl|seev|semt|sese|setr|tsin|tsmt|tsrv|head)\.[0-9]{3}\.[0-9]{3}\.[0-9]{2}$`
	match, err := regexp.MatchString(search, v.Format)
	if err != nil {
		return
	}

	if !match {
		return "", "", errors.New("Input is not a valid iso20022 message")
	}

	//Eg. pain.009.001.04
	splitCode := strings.Split(v.Format, "xsd:")
	if len(splitCode) < 2 {
		return "", "", fmt.Errorf("%s is not a valid iso20022 message format", v.Format)
	}
	code = splitCode[1]

	splitCode = strings.Split(code, ".")
	if len(splitCode) != 4 {
		return "", "", fmt.Errorf("%s is not a valid iso20022 message code", code)
	}

	//Eg. pain
	domain = splitCode[0]

	return
}

//Return ISO Message struct containing message content as well as
//identifiers to facilitate type casting/assertion and message handling
func Decode(message string) (result Message, domain, code string, err error) {

	if domain, code, err = ValidateISO20022(message); err != nil {
		return
	}

	if result, err = makeISO20022(code, message); err != nil {
		return
	}

	return
}
