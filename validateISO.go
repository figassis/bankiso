package iso

import (
	"encoding/xml"
	"fmt"
	"regexp"
	"strings"
)

const Example = `<?xml version="1.0" encoding="UTF-8"?><Document xmlns="urn:iso:std:iso:20022:tech:xsd:pain.009.001.04"><MndtInitnReq><GrpHdr><MsgId>BBBB654322</MsgId><CreDtTm>2013-06-10T11:00:00</CreDtTm><InitgPty><Nm>Jersey Mobile Phone</Nm><PstlAdr><StrtNm>Virginia Lane</StrtNm><BldgNb>36</BldgNb><PstCd>NJ 07311</PstCd><TwnNm>Jersey City</TwnNm><Ctry>US</Ctry></PstlAdr></InitgPty></GrpHdr><Mndt><MndtReqId>Johns/005</MndtReqId><Ocrncs><SeqTp>RCUR</SeqTp><Frqcy><Tp>MNTH</Tp></Frqcy><FrstColltnDt>2013-06-25</FrstColltnDt></Ocrncs><Cdtr><Nm>Jersey Mobile Phone</Nm></Cdtr><CdtrAcct><Id><Othr><Id>76543</Id></Othr></Id></CdtrAcct><CdtrAgt><FinInstnId><BICFI>DDDDUS31</BICFI></FinInstnId></CdtrAgt><Dbtr><Nm>Johnson</Nm></Dbtr><DbtrAcct><Id><Othr><Id>5544732</Id></Othr></Id></DbtrAcct><DbtrAgt><FinInstnId><BICFI>FFFFUS91</BICFI></FinInstnId></DbtrAgt><RfrdDoc><Tp><CdOrPrtry><Cd>DISP</Cd></CdOrPrtry></Tp><Nb>JMP/24653</Nb><RltdDt>2013-06-11</RltdDt></RfrdDoc></Mndt></MndtInitnReq></Document>`

//const MessageTypes = "acmt|admi|auth|caaa|caam|camt|catm|catp|colr|fxtr|pacs|pain|reda|remt|secl|seev|semt|sese|setr|tsin|tsmt|tsrv|head"

type Document struct {
	Format string `xml:"xmlns,attr"`
}

//Basix XML version 1.0 validation via header redex
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

func main() {

	fmt.Println(ValidateISO(Example))

}
