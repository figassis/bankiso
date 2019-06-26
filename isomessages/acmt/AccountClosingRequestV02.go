package acmt

import (
	"encoding/xml"

	"github.com/figassis/bankiso/iso20022"
)

type Document01900102 struct {
	XMLName xml.Name                  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.019.001.02 Document"`
	Message *AccountClosingRequestV02 `xml:"AcctClsgReq"`
}

func (d *Document01900102) AddMessage() *AccountClosingRequestV02 {
	d.Message = new(AccountClosingRequestV02)
	return d.Message
}

// The AccountClosingRequest message is sent from an organisation to a financial institution as part of the account closing process. It is the initial request message to close an account.
type AccountClosingRequestV02 struct {

	// Set of elements for the identification of the message and related references.
	References *iso20022.References4 `xml:"Refs"`

	// Identifies the business sender of the message, if it is not the account owner or account servicing financial institution. OrganisationIdentification6
	From *iso20022.OrganisationIdentification8 `xml:"Fr,omitempty"`

	// Unique and unambiguous identification of the account between the account owner and the account servicer.
	AccountIdentification *iso20022.AccountForAction2 `xml:"AcctId"`

	// Unique and unambiguous identifier of a financial institution, as assigned under an internationally recognised or proprietary identification scheme.
	AccountServicerIdentification *iso20022.BranchAndFinancialInstitutionIdentification5 `xml:"AcctSvcrId"`

	// Identification of the organisation requesting the change.
	OrganisationIdentification *iso20022.Organisation14 `xml:"OrgId"`

	// Specifies target dates.
	ContractDates *iso20022.AccountContract4 `xml:"CtrctDts,omitempty"`

	// Identification of the account to which the remaining positive balance of the account to be closed must be transferred or account from which funds can be moved to the account to be closed and which balance is negative. This account must be held in the same financial institution as the account to be closed if the transfer account is used to compensate a negative balance. For a positive balance to be transferred, an account in another financial institution might be used. In that case the account servicer is mandatory.
	BalanceTransferAccount *iso20022.AccountForAction1 `xml:"BalTrfAcct,omitempty"`

	// Unique and unambiguous identifier of a financial institution, as assigned under an internationally recognised or proprietary identification scheme, that is the servicer of the transfer account.
	TransferAccountServicerIdentification *iso20022.BranchAndFinancialInstitutionIdentification5 `xml:"TrfAcctSvcrId,omitempty"`

	// Contains the signature with its components, namely signed info, signature value, key info and the object.
	DigitalSignature []*iso20022.PartyAndSignature2 `xml:"DgtlSgntr,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AccountClosingRequestV02) AddReferences() *iso20022.References4 {
	a.References = new(iso20022.References4)
	return a.References
}

func (a *AccountClosingRequestV02) AddFrom() *iso20022.OrganisationIdentification8 {
	a.From = new(iso20022.OrganisationIdentification8)
	return a.From
}

func (a *AccountClosingRequestV02) AddAccountIdentification() *iso20022.AccountForAction2 {
	a.AccountIdentification = new(iso20022.AccountForAction2)
	return a.AccountIdentification
}

func (a *AccountClosingRequestV02) AddAccountServicerIdentification() *iso20022.BranchAndFinancialInstitutionIdentification5 {
	a.AccountServicerIdentification = new(iso20022.BranchAndFinancialInstitutionIdentification5)
	return a.AccountServicerIdentification
}

func (a *AccountClosingRequestV02) AddOrganisationIdentification() *iso20022.Organisation14 {
	a.OrganisationIdentification = new(iso20022.Organisation14)
	return a.OrganisationIdentification
}

func (a *AccountClosingRequestV02) AddContractDates() *iso20022.AccountContract4 {
	a.ContractDates = new(iso20022.AccountContract4)
	return a.ContractDates
}

func (a *AccountClosingRequestV02) AddBalanceTransferAccount() *iso20022.AccountForAction1 {
	a.BalanceTransferAccount = new(iso20022.AccountForAction1)
	return a.BalanceTransferAccount
}

func (a *AccountClosingRequestV02) AddTransferAccountServicerIdentification() *iso20022.BranchAndFinancialInstitutionIdentification5 {
	a.TransferAccountServicerIdentification = new(iso20022.BranchAndFinancialInstitutionIdentification5)
	return a.TransferAccountServicerIdentification
}

func (a *AccountClosingRequestV02) AddDigitalSignature() *iso20022.PartyAndSignature2 {
	newValue := new(iso20022.PartyAndSignature2)
	a.DigitalSignature = append(a.DigitalSignature, newValue)
	return newValue
}

func (a *AccountClosingRequestV02) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}
func ( d *Document01900102 ) String() (result string, ok bool) { return }
