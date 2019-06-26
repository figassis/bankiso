package pain

import (
	"encoding/xml"

	"github.com/figassis/bankiso/iso20022"
)

type Document01300104 struct {
	XMLName xml.Name                             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.04 Document"`
	Message *CreditorPaymentActivationRequestV04 `xml:"CdtrPmtActvtnReq"`
}

func (d *Document01300104) AddMessage() *CreditorPaymentActivationRequestV04 {
	d.Message = new(CreditorPaymentActivationRequestV04)
	return d.Message
}

// The CreditorPaymentActivationRequest message is sent by the Creditor sending party to the Debtor receiving party, directly or through agents. It is used to initiate a creditor payment activation request.
type CreditorPaymentActivationRequestV04 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *iso20022.GroupHeader45 `xml:"GrpHdr"`

	// Set of characteristics that applies to the debit side of the payment transactions included in the creditor payment initiation.
	PaymentInformation []*iso20022.PaymentInstruction17 `xml:"PmtInf"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CreditorPaymentActivationRequestV04) AddGroupHeader() *iso20022.GroupHeader45 {
	c.GroupHeader = new(iso20022.GroupHeader45)
	return c.GroupHeader
}

func (c *CreditorPaymentActivationRequestV04) AddPaymentInformation() *iso20022.PaymentInstruction17 {
	newValue := new(iso20022.PaymentInstruction17)
	c.PaymentInformation = append(c.PaymentInformation, newValue)
	return newValue
}

func (c *CreditorPaymentActivationRequestV04) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
func ( d *Document01300104 ) String() (result string, ok bool) { return }
