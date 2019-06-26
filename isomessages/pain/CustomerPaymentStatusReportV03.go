package pain

import (
	"encoding/xml"

	"github.com/figassis/bankiso/iso20022"
)

type Document00200103 struct {
	XMLName xml.Name                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.03 Document"`
	Message *CustomerPaymentStatusReportV03 `xml:"CstmrPmtStsRpt"`
}

func (d *Document00200103) AddMessage() *CustomerPaymentStatusReportV03 {
	d.Message = new(CustomerPaymentStatusReportV03)
	return d.Message
}

// Scope
// The CustomerPaymentStatusReport message is sent by an instructed agent to the previous party in the payment chain. It is used to inform this party about the positive or negative status of an instruction (either single or file). It is also used to report on a pending instruction.
// Usage
// The CustomerPaymentStatusReport message is exchanged between an agent and a non-financial institution customer to provide status information on instructions previously sent. Its usage will always be governed by a bilateral agreement between the agent and the non-financial institution customer.
// The CustomerPaymentStatusReport message can be used to provide information about the status (e.g. rejection, acceptance) of the initiation of a credit transfer, a direct debit, as well as on the initiation of other customer instructions.
// The CustomerPaymentStatusReport message refers to the original instruction(s) by means of references only or by means of references and a set of elements from the original instruction.
// The CustomerPaymentStatusReport message can be used in domestic and cross-border scenarios.
type CustomerPaymentStatusReportV03 struct {

	// Set of characteristics shared by all individual transactions included in the status report message.
	GroupHeader *iso20022.GroupHeader36 `xml:"GrpHdr"`

	// Original group information concerning the group of transactions, to which the status report message refers to.
	OriginalGroupInformationAndStatus *iso20022.OriginalGroupInformation20 `xml:"OrgnlGrpInfAndSts"`

	// Information concerning the original payment information, to which the status report message refers.
	OriginalPaymentInformationAndStatus []*iso20022.OriginalPaymentInformation1 `xml:"OrgnlPmtInfAndSts,omitempty"`
}

func (c *CustomerPaymentStatusReportV03) AddGroupHeader() *iso20022.GroupHeader36 {
	c.GroupHeader = new(iso20022.GroupHeader36)
	return c.GroupHeader
}

func (c *CustomerPaymentStatusReportV03) AddOriginalGroupInformationAndStatus() *iso20022.OriginalGroupInformation20 {
	c.OriginalGroupInformationAndStatus = new(iso20022.OriginalGroupInformation20)
	return c.OriginalGroupInformationAndStatus
}

func (c *CustomerPaymentStatusReportV03) AddOriginalPaymentInformationAndStatus() *iso20022.OriginalPaymentInformation1 {
	newValue := new(iso20022.OriginalPaymentInformation1)
	c.OriginalPaymentInformationAndStatus = append(c.OriginalPaymentInformationAndStatus, newValue)
	return newValue
}
func ( d *Document00200103 ) String() (result string, ok bool) { return }
