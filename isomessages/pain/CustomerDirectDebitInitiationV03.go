package pain

import (
	"encoding/xml"

	"github.com/figassis/bankiso/iso20022"
)

type Document00800103 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.03 Document"`
	Message *CustomerDirectDebitInitiationV03 `xml:"CstmrDrctDbtInitn"`
}

func (d *Document00800103) AddMessage() *CustomerDirectDebitInitiationV03 {
	d.Message = new(CustomerDirectDebitInitiationV03)
	return d.Message
}

// Scope
// The CustomerDirectDebitInitiation message is sent by the initiating party to the forwarding agent or creditor agent. It is used to request single or bulk collection(s) of funds from one or various debtor's account(s) for a creditor.
// Usage
// The CustomerDirectDebitInitiation message can contain one or more direct debit instructions.
// The message can be used in a direct or a relay scenario:
// - In a direct scenario, the message is sent directly to the creditor agent. The creditor agent is the account servicer of the creditor.
// - In a relay scenario, the message is sent to a forwarding agent. The forwarding agent acts as a concentrating financial institution. It will forward the CustomerDirectDebitInitiation message to the creditor agent.
// The message can also be used by an initiating party that has authority to send the message on behalf of the creditor. This caters for example for the scenario of a payments factory initiating all payments on behalf of a large corporate.
// The CustomerDirectDebitInitiation message can be used in domestic and cross-border scenarios.
// The CustomerDirectDebitInitiation may or may not contain mandate related information, i.e. extracts from a mandate, such as MandateIdentification or DateOfSignature. The CustomerDirectDebitInitiation message must not be considered as a mandate.
// The CustomerDirectDebitInitiation message must not be used by the creditor agent to execute the direct debit instruction(s). The FIToFICustomerDirectDebit message must be used instead.
type CustomerDirectDebitInitiationV03 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *iso20022.GroupHeader55 `xml:"GrpHdr"`

	// Set of characteristics that apply to the credit side of the payment transactions included in the direct debit transaction initiation.
	PaymentInformation []*iso20022.PaymentInstruction7 `xml:"PmtInf"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CustomerDirectDebitInitiationV03) AddGroupHeader() *iso20022.GroupHeader55 {
	c.GroupHeader = new(iso20022.GroupHeader55)
	return c.GroupHeader
}

func (c *CustomerDirectDebitInitiationV03) AddPaymentInformation() *iso20022.PaymentInstruction7 {
	newValue := new(iso20022.PaymentInstruction7)
	c.PaymentInformation = append(c.PaymentInformation, newValue)
	return newValue
}

func (c *CustomerDirectDebitInitiationV03) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
func ( d *Document00800103 ) String() (result string, ok bool) { return }
