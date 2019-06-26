package pain

import (
	"encoding/xml"

	"github.com/figassis/bankiso/iso20022"
)

type Document00700102 struct {
	XMLName xml.Name                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.02 Document"`
	Message *CustomerPaymentReversalV02 `xml:"CstmrPmtRvsl"`
}

func (d *Document00700102) AddMessage() *CustomerPaymentReversalV02 {
	d.Message = new(CustomerPaymentReversalV02)
	return d.Message
}

// Scope
// The CustomerPaymentReversal message is sent by the initiating party to the next party in the payment chain. It is used to reverse a payment previously executed.
// Usage
// The CustomerPaymentReversal message is exchanged between a non-financial institution customer and an agent to reverse a CustomerDirectDebitInitiation message that has been settled. The result will be a credit on the debtor account.
// The CustomerPaymentReversal message refers to the original CustomerDirectDebitInitiation message by means of references only or by means of references and a set of elements from the original instruction.
// The CustomerPaymentReversal message can be used in domestic and cross-border scenarios.
type CustomerPaymentReversalV02 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *iso20022.GroupHeader40 `xml:"GrpHdr"`

	// Information concerning the original group of transactions, to which the message refers.
	OriginalGroupInformation *iso20022.OriginalGroupInformation22 `xml:"OrgnlGrpInf"`

	// Information concerning the original payment information, to which the reversal message refers.
	OriginalPaymentInformationAndReversal []*iso20022.OriginalPaymentInformation2 `xml:"OrgnlPmtInfAndRvsl,omitempty"`
}

func (c *CustomerPaymentReversalV02) AddGroupHeader() *iso20022.GroupHeader40 {
	c.GroupHeader = new(iso20022.GroupHeader40)
	return c.GroupHeader
}

func (c *CustomerPaymentReversalV02) AddOriginalGroupInformation() *iso20022.OriginalGroupInformation22 {
	c.OriginalGroupInformation = new(iso20022.OriginalGroupInformation22)
	return c.OriginalGroupInformation
}

func (c *CustomerPaymentReversalV02) AddOriginalPaymentInformationAndReversal() *iso20022.OriginalPaymentInformation2 {
	newValue := new(iso20022.OriginalPaymentInformation2)
	c.OriginalPaymentInformationAndReversal = append(c.OriginalPaymentInformationAndReversal, newValue)
	return newValue
}
func ( d *Document00700102 ) String() (result string, ok bool) { return }
