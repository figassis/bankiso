package pacs

import (
	"encoding/xml"

	"github.com/figassis/bankiso/iso20022"
)

type Document00300103 struct {
	XMLName xml.Name                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.03 Document"`
	Message *FIToFICustomerDirectDebitV03 `xml:"FIToFICstmrDrctDbt"`
}

func (d *Document00300103) AddMessage() *FIToFICustomerDirectDebitV03 {
	d.Message = new(FIToFICustomerDirectDebitV03)
	return d.Message
}

// Scope
// The FinancialInstitutionToFinancialInstitutionCustomerDirectDebit message is sent by the creditor agent to the debtor agent, directly or through other agents and/or a payment clearing and settlement system.
// It is used to collect funds from a debtor account for a creditor.
// Usage
// The FItoFICustomerDirectDebit message can contain one or more customer direct debit instructions.
// The FIToFICustomerDirectDebit message does not allow for grouping: the PaymentInformation block must be present once per occurrence of a DirectDebitTransactionInformation block.
// The FItoFICustomerDirectDebit message may or may not contain mandate related information, i.e. extracts from a mandate, such as the MandateIdentification or DateOfSignature. The FIToFICustomerDirectDebit message must not be considered as a mandate.
// The FItoFICustomerDirectDebit message can be used in domestic and cross-border scenarios.
type FIToFICustomerDirectDebitV03 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *iso20022.GroupHeader50 `xml:"GrpHdr"`

	// Set of elements providing information specific to the individual direct debit(s).
	DirectDebitTransactionInformation []*iso20022.DirectDebitTransactionInformation12 `xml:"DrctDbtTxInf"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (f *FIToFICustomerDirectDebitV03) AddGroupHeader() *iso20022.GroupHeader50 {
	f.GroupHeader = new(iso20022.GroupHeader50)
	return f.GroupHeader
}

func (f *FIToFICustomerDirectDebitV03) AddDirectDebitTransactionInformation() *iso20022.DirectDebitTransactionInformation12 {
	newValue := new(iso20022.DirectDebitTransactionInformation12)
	f.DirectDebitTransactionInformation = append(f.DirectDebitTransactionInformation, newValue)
	return newValue
}

func (f *FIToFICustomerDirectDebitV03) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	f.SupplementaryData = append(f.SupplementaryData, newValue)
	return newValue
}
func ( d *Document00300103 ) String() (result string, ok bool) { return }
