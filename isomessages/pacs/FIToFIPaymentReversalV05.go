package pacs

import (
	"encoding/xml"

	"github.com/figassis/bankiso/iso20022"
)

type Document00700105 struct {
	XMLName xml.Name                  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.007.001.05 Document"`
	Message *FIToFIPaymentReversalV05 `xml:"FIToFIPmtRvsl"`
}

func (d *Document00700105) AddMessage() *FIToFIPaymentReversalV05 {
	d.Message = new(FIToFIPaymentReversalV05)
	return d.Message
}

// Scope
// The FinancialInstitutionToFinancialInstitutionPaymentReversal message is sent by an agent to the next party in the payment chain. It is used to reverse a payment previously executed.
// Usage
// The FIToFIPaymentReversal message is exchanged between agents to reverse a payment message that has been settled. The result will be a credit on the debtor account (when the reversed payment was a Direct Debit) or a debit on the creditor account (when the reversed payment was a Credit Transfer).
// The FIToFIPaymentReversal message may or may not be the follow-up of a payment message.
// The FIToFIPaymentReversal message refers to the original payment message by means of references only or by means of references and a set of elements from the original instruction.
// The FIToFIPaymentReversal message can be used in domestic and cross-border scenarios.
type FIToFIPaymentReversalV05 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *iso20022.GroupHeader57 `xml:"GrpHdr"`

	// Information concerning the original group of transactions, to which the message refers.
	OriginalGroupInformation *iso20022.OriginalGroupHeader3 `xml:"OrgnlGrpInf,omitempty"`

	// Information concerning the original transactions, to which the reversal message refers.
	TransactionInformation []*iso20022.PaymentTransaction51 `xml:"TxInf,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (f *FIToFIPaymentReversalV05) AddGroupHeader() *iso20022.GroupHeader57 {
	f.GroupHeader = new(iso20022.GroupHeader57)
	return f.GroupHeader
}

func (f *FIToFIPaymentReversalV05) AddOriginalGroupInformation() *iso20022.OriginalGroupHeader3 {
	f.OriginalGroupInformation = new(iso20022.OriginalGroupHeader3)
	return f.OriginalGroupInformation
}

func (f *FIToFIPaymentReversalV05) AddTransactionInformation() *iso20022.PaymentTransaction51 {
	newValue := new(iso20022.PaymentTransaction51)
	f.TransactionInformation = append(f.TransactionInformation, newValue)
	return newValue
}

func (f *FIToFIPaymentReversalV05) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	f.SupplementaryData = append(f.SupplementaryData, newValue)
	return newValue
}
func ( d *Document00700105 ) String() (result string, ok bool) { return }
