package iso

const MessageTypes = "acmt|admi|auth|caaa|caam|camt|catm|catp|colr|fxtr|pacs|pain|reda|remt|secl|seev|semt|sese|setr|tsin|tsmt|tsrv|head"

//Basic struct used to test XML format
type Document struct {
	Format string `xml:"xmlns,attr"`
}

//Defines an ISO Message
//Can either be ISO 8583 or 20022
type ISOMessage interface {
	String() (result string, err error)
}
