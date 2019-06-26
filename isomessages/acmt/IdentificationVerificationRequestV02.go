package acmt

import (
	"encoding/xml"

	"github.com/figassis/bankiso/iso20022"
)

type Document02300102 struct {
	XMLName xml.Name                              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.023.001.02 Document"`
	Message *IdentificationVerificationRequestV02 `xml:"IdVrfctnReq"`
}

func (d *Document02300102) AddMessage() *IdentificationVerificationRequestV02 {
	d.Message = new(IdentificationVerificationRequestV02)
	return d.Message
}

// Scope
// The IdentificationVerificationRequest message is sent by an assigner to an assignee. It is used to request the verification of party and/or account identification information.
// Usage
// The IdentificationVerificationRequest message is sent before the sending of one or several transactions messages.
// The IdentificationVerificationRequest message can contain one or more verification requests.
type IdentificationVerificationRequestV02 struct {

	// Identifies the identification assignment.
	Assignment *iso20022.IdentificationAssignment2 `xml:"Assgnmt"`

	// Information concerning the identification data that is requested to be verified.
	Verification []*iso20022.IdentificationVerification2 `xml:"Vrfctn"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (i *IdentificationVerificationRequestV02) AddAssignment() *iso20022.IdentificationAssignment2 {
	i.Assignment = new(iso20022.IdentificationAssignment2)
	return i.Assignment
}

func (i *IdentificationVerificationRequestV02) AddVerification() *iso20022.IdentificationVerification2 {
	newValue := new(iso20022.IdentificationVerification2)
	i.Verification = append(i.Verification, newValue)
	return newValue
}

func (i *IdentificationVerificationRequestV02) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	i.SupplementaryData = append(i.SupplementaryData, newValue)
	return newValue
}
func ( d *Document02300102 ) String() (result string, ok bool) { return }
