package pacs

import (
	"encoding/xml"

	"github.com/figassis/bankiso/iso20022"
)

type Document00600101 struct {
	XMLName xml.Name                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.006.001.01 Document"`
	Message *PaymentCancellationRequestV01 `xml:"pacs.006.001.01"`
}

func (d *Document00600101) AddMessage() *PaymentCancellationRequestV01 {
	d.Message = new(PaymentCancellationRequestV01)
	return d.Message
}

// Scope
// The PaymentCancellationRequest message is sent by the initiating party or any agent, to the next party in the payment chain. It is used to request the cancellation of an instruction previously sent.
// Usage
// The PaymentCancellationRequest message is exchanged between agents to request the cancellation of a payment message previously sent (i.e. FIToFICustomerCreditTransfer, FIToFICustomerDirectDebit, and FinancialInstitutionCreditTransfer).
// The PaymentCancellationRequest message can be used to request the cancellation of single instructions or multiple instructions, from one or multiple files.
// The PaymentCancellationRequest message can be used in domestic and cross-border scenarios.
// The PaymentCancellationRequest message refers to the original instruction(s) by means of references only or by means of references and a set of elements from the original instruction.
// The PaymentCancellationRequest message exchanged between agents is identified in the schema as follows:
// urn:iso:std:iso:20022:tech:xsd:pacs.006.001.01
type PaymentCancellationRequestV01 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *iso20022.GroupHeader7 `xml:"GrpHdr"`

	// Information concerning the original group of transactions, to which the message refers.
	OriginalGroupInformation *iso20022.OriginalGroupInformation4 `xml:"OrgnlGrpInf"`

	// Information concerning the original transactions, to which the cancellation request message refers.
	TransactionInformation []*iso20022.PaymentTransactionInformation3 `xml:"TxInf,omitempty"`
}

func (p *PaymentCancellationRequestV01) AddGroupHeader() *iso20022.GroupHeader7 {
	p.GroupHeader = new(iso20022.GroupHeader7)
	return p.GroupHeader
}

func (p *PaymentCancellationRequestV01) AddOriginalGroupInformation() *iso20022.OriginalGroupInformation4 {
	p.OriginalGroupInformation = new(iso20022.OriginalGroupInformation4)
	return p.OriginalGroupInformation
}

func (p *PaymentCancellationRequestV01) AddTransactionInformation() *iso20022.PaymentTransactionInformation3 {
	newValue := new(iso20022.PaymentTransactionInformation3)
	p.TransactionInformation = append(p.TransactionInformation, newValue)
	return newValue
}
func ( d *Document00600101 ) String() (result string, ok bool) { return }
