package pacs

import (
	"encoding/xml"

	"github.com/figassis/bankiso/iso20022"
)

type Document00400102 struct {
	XMLName xml.Name          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02 Document"`
	Message *PaymentReturnV02 `xml:"PmtRtr"`
}

func (d *Document00400102) AddMessage() *PaymentReturnV02 {
	d.Message = new(PaymentReturnV02)
	return d.Message
}

// Scope
// The PaymentReturn message is sent by an agent to the previous agent in the payment chain to undo a payment previously settled.
// Usage
// The PaymentReturn message is exchanged between agents to return funds after settlement of credit transfer instructions (i.e. FIToFICustomerCreditTransfer message and FinancialInstitutionCreditTransfer message) or direct debit instructions (FIToFICustomerDirectDebit message).
// The PaymentReturn message should not be used between agents and non-financial institution customers. Non-financial institution customers will be informed about a debit or a credit on their account(s) through a BankToCustomerDebitCreditNotification message ('notification') and/or BankToCustomerAccountReport/BankToCustomerStatement message ('statement').
// The PaymentReturn message can be used to return single instructions or multiple instructions from one or different files.
// The PaymentReturn message can be used in domestic and cross-border scenarios.
// The PaymentReturn message refers to the original instruction(s) by means of references only or by means of references and a set of elements from the original instruction.
type PaymentReturnV02 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *iso20022.GroupHeader38 `xml:"GrpHdr"`

	// Information concerning the original group of transactions, to which the message refers.
	OriginalGroupInformation *iso20022.OriginalGroupInformation21 `xml:"OrgnlGrpInf,omitempty"`

	// Information concerning the original transactions, to which the return message refers.
	TransactionInformation []*iso20022.PaymentTransactionInformation27 `xml:"TxInf,omitempty"`
}

func (p *PaymentReturnV02) AddGroupHeader() *iso20022.GroupHeader38 {
	p.GroupHeader = new(iso20022.GroupHeader38)
	return p.GroupHeader
}

func (p *PaymentReturnV02) AddOriginalGroupInformation() *iso20022.OriginalGroupInformation21 {
	p.OriginalGroupInformation = new(iso20022.OriginalGroupInformation21)
	return p.OriginalGroupInformation
}

func (p *PaymentReturnV02) AddTransactionInformation() *iso20022.PaymentTransactionInformation27 {
	newValue := new(iso20022.PaymentTransactionInformation27)
	p.TransactionInformation = append(p.TransactionInformation, newValue)
	return newValue
}
func ( d *Document00400102 ) String() (result string, ok bool) { return }
