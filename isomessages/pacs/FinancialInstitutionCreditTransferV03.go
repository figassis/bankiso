package pacs

import (
	"encoding/xml"

	"github.com/figassis/bankiso/iso20022"
)

type Document00900103 struct {
	XMLName xml.Name                               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.03 Document"`
	Message *FinancialInstitutionCreditTransferV03 `xml:"FinInstnCdtTrf"`
}

func (d *Document00900103) AddMessage() *FinancialInstitutionCreditTransferV03 {
	d.Message = new(FinancialInstitutionCreditTransferV03)
	return d.Message
}

// Scope
// The FinancialInstitutionCreditTransfer message is sent by a debtor financial institution to a creditor financial institution, directly or through other agents and/or a payment clearing and settlement system.
// It is used to move funds from a debtor account to a creditor, where both debtor and creditor are financial institutions.
// Usage
// The FinancialInstitutionCreditTransfer message is exchanged between agents and can contain one or more credit transfer instructions where debtor and creditor are both financial institutions.
// The FinancialInstitutionCreditTransfer message does not allow for grouping: a CreditTransferTransactionInformation block must be present for each credit transfer transaction.
// The FinancialInstitutionCreditTransfer message can be used in domestic and cross-border scenarios.
type FinancialInstitutionCreditTransferV03 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *iso20022.GroupHeader51 `xml:"GrpHdr"`

	// Set of elements providing information specific to the individual credit transfer(s).
	CreditTransferTransactionInformation []*iso20022.CreditTransferTransaction4 `xml:"CdtTrfTxInf"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (f *FinancialInstitutionCreditTransferV03) AddGroupHeader() *iso20022.GroupHeader51 {
	f.GroupHeader = new(iso20022.GroupHeader51)
	return f.GroupHeader
}

func (f *FinancialInstitutionCreditTransferV03) AddCreditTransferTransactionInformation() *iso20022.CreditTransferTransaction4 {
	newValue := new(iso20022.CreditTransferTransaction4)
	f.CreditTransferTransactionInformation = append(f.CreditTransferTransactionInformation, newValue)
	return newValue
}

func (f *FinancialInstitutionCreditTransferV03) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	f.SupplementaryData = append(f.SupplementaryData, newValue)
	return newValue
}
func ( d *Document00900103 ) String() (result string, ok bool) { return }
