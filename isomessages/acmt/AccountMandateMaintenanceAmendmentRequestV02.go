package acmt

import (
	"encoding/xml"

	"github.com/figassis/bankiso/iso20022"
)

type Document01800102 struct {
	XMLName xml.Name                                      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.018.001.02 Document"`
	Message *AccountMandateMaintenanceAmendmentRequestV02 `xml:"AcctMndtMntncAmdmntReq"`
}

func (d *Document01800102) AddMessage() *AccountMandateMaintenanceAmendmentRequestV02 {
	d.Message = new(AccountMandateMaintenanceAmendmentRequestV02)
	return d.Message
}

// The AccountMandateMaintenanceAmendmentRequest message is sent from an organisation to a financial institution as part of the account maintenance process. It is sent in response to a request from the financial institution to send additional information. Usage: this update is only about mandate information.
// If modification codes are not used: the organisation will specify under the “Mandate” and “Group” tags the complete information as it should be in the financial institution’s records after processing the update request.
// If modification codes are used (in that case, they must be used everywhere): the organisation will specify under the “Mandate” and “Group” tags which elements must be added, deleted, modified, or if they are unchanged.
// It is not possible to update the account characteristics or organisation information with this message.
//
type AccountMandateMaintenanceAmendmentRequestV02 struct {

	// Set of elements for the identification of the message and related references.
	References *iso20022.References4 `xml:"Refs"`

	// Identifies the business sender of the message, if it is not the account owner or account servicing financial institution.
	From *iso20022.OrganisationIdentification8 `xml:"Fr,omitempty"`

	// Specifies target dates.
	ContractDates *iso20022.AccountContract2 `xml:"CtrctDts,omitempty"`

	// Account contract established between the organisation or the group to which the organisation belongs, and the account servicer. This contract has to be applied for the new account to be opened and maintained.
	UnderlyingMasterAgreement *iso20022.ContractDocument1 `xml:"UndrlygMstrAgrmt,omitempty"`

	// Unique and unambiguous identification of the account between the account owner and the account servicer.
	AccountIdentification []*iso20022.AccountForAction1 `xml:"AcctId"`

	// Unique and unambiguous identifier of a financial institution, as assigned under an internationally recognised or proprietary identification scheme.
	//
	AccountServicerIdentification *iso20022.BranchAndFinancialInstitutionIdentification5 `xml:"AcctSvcrId"`

	// Identification of the organisation requesting the change.
	OrganisationIdentification *iso20022.OrganisationIdentification8 `xml:"OrgId"`

	// Information specifying the account mandate.
	Mandate []*iso20022.OperationMandate3 `xml:"Mndt,omitempty"`

	// Definition of a group of parties.
	Group []*iso20022.Group2 `xml:"Grp,omitempty"`

	// Contains the signature with its components, namely signed info, signature value, key info and the object.
	DigitalSignature []*iso20022.PartyAndSignature2 `xml:"DgtlSgntr,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AccountMandateMaintenanceAmendmentRequestV02) AddReferences() *iso20022.References4 {
	a.References = new(iso20022.References4)
	return a.References
}

func (a *AccountMandateMaintenanceAmendmentRequestV02) AddFrom() *iso20022.OrganisationIdentification8 {
	a.From = new(iso20022.OrganisationIdentification8)
	return a.From
}

func (a *AccountMandateMaintenanceAmendmentRequestV02) AddContractDates() *iso20022.AccountContract2 {
	a.ContractDates = new(iso20022.AccountContract2)
	return a.ContractDates
}

func (a *AccountMandateMaintenanceAmendmentRequestV02) AddUnderlyingMasterAgreement() *iso20022.ContractDocument1 {
	a.UnderlyingMasterAgreement = new(iso20022.ContractDocument1)
	return a.UnderlyingMasterAgreement
}

func (a *AccountMandateMaintenanceAmendmentRequestV02) AddAccountIdentification() *iso20022.AccountForAction1 {
	newValue := new(iso20022.AccountForAction1)
	a.AccountIdentification = append(a.AccountIdentification, newValue)
	return newValue
}

func (a *AccountMandateMaintenanceAmendmentRequestV02) AddAccountServicerIdentification() *iso20022.BranchAndFinancialInstitutionIdentification5 {
	a.AccountServicerIdentification = new(iso20022.BranchAndFinancialInstitutionIdentification5)
	return a.AccountServicerIdentification
}

func (a *AccountMandateMaintenanceAmendmentRequestV02) AddOrganisationIdentification() *iso20022.OrganisationIdentification8 {
	a.OrganisationIdentification = new(iso20022.OrganisationIdentification8)
	return a.OrganisationIdentification
}

func (a *AccountMandateMaintenanceAmendmentRequestV02) AddMandate() *iso20022.OperationMandate3 {
	newValue := new(iso20022.OperationMandate3)
	a.Mandate = append(a.Mandate, newValue)
	return newValue
}

func (a *AccountMandateMaintenanceAmendmentRequestV02) AddGroup() *iso20022.Group2 {
	newValue := new(iso20022.Group2)
	a.Group = append(a.Group, newValue)
	return newValue
}

func (a *AccountMandateMaintenanceAmendmentRequestV02) AddDigitalSignature() *iso20022.PartyAndSignature2 {
	newValue := new(iso20022.PartyAndSignature2)
	a.DigitalSignature = append(a.DigitalSignature, newValue)
	return newValue
}

func (a *AccountMandateMaintenanceAmendmentRequestV02) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}
func ( d *Document01800102 ) String() (result string, ok bool) { return }
