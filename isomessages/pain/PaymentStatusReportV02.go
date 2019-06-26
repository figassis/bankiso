package pain

import (
	"encoding/xml"

	"github.com/figassis/bankiso/iso20022"
)

type Document00200102 struct {
	XMLName xml.Name                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.02 Document"`
	Message *PaymentStatusReportV02 `xml:"pain.002.001.02"`
}

func (d *Document00200102) AddMessage() *PaymentStatusReportV02 {
	d.Message = new(PaymentStatusReportV02)
	return d.Message
}

// Scope
// The PaymentStatusReport message is sent by an instructed agent to the previous party in the payment chain. It is used to inform this party about the positive or negative status of an instruction (either single or file). It is also used to report on a pending instruction.
// Usage
// The PaymentStatusReport message is exchanged between an agent and a non-financial institution customer to provide status information on instructions previously sent. Its usage will always be governed by a bilateral agreement between the agent and the non-financial institution customer.
// The PaymentStatusReport message can be used to provide information about the status (e.g. rejection, acceptance) of the initiation of a credit transfer, a direct debit, as well as on the initiation of other customer instructions (e.g. PaymentCancellationRequest).
// The PaymentStatusReport message refers to the original instruction(s) by means of references only or by means of references and a set of elements from the original instruction.
// The PaymentStatusReport message can be used in domestic and cross-border scenarios.
// The PaymentStatusReport message exchanged between agents and non-financial institution customers is identified in the schema as follows: urn:iso:std:iso:20022:tech:xsd:pain.002.001.02
type PaymentStatusReportV02 struct {

	// Set of characteristics shared by all individual transactions included in the status report message.
	GroupHeader *iso20022.GroupHeader5 `xml:"GrpHdr"`

	// Original group information concerning the group of transactions, to which the status report message refers to.
	OriginalGroupInformationAndStatus *iso20022.OriginalGroupInformation1 `xml:"OrgnlGrpInfAndSts"`

	// Information concerning the original transactions, to which the status report message refers.
	TransactionInformationAndStatus []*iso20022.PaymentTransactionInformation1 `xml:"TxInfAndSts,omitempty"`
}

func (p *PaymentStatusReportV02) AddGroupHeader() *iso20022.GroupHeader5 {
	p.GroupHeader = new(iso20022.GroupHeader5)
	return p.GroupHeader
}

func (p *PaymentStatusReportV02) AddOriginalGroupInformationAndStatus() *iso20022.OriginalGroupInformation1 {
	p.OriginalGroupInformationAndStatus = new(iso20022.OriginalGroupInformation1)
	return p.OriginalGroupInformationAndStatus
}

func (p *PaymentStatusReportV02) AddTransactionInformationAndStatus() *iso20022.PaymentTransactionInformation1 {
	newValue := new(iso20022.PaymentTransactionInformation1)
	p.TransactionInformationAndStatus = append(p.TransactionInformationAndStatus, newValue)
	return newValue
}
func ( d *Document00200102 ) String() (result string, ok bool) { return }
