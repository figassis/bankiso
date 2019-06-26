package acmt

import (
	"encoding/xml"

	"github.com/figassis/bankiso/iso20022"
)

type Document02200102 struct {
	XMLName xml.Name                             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 Document"`
	Message *IdentificationModificationAdviceV02 `xml:"IdModAdvc"`
}

func (d *Document02200102) AddMessage() *IdentificationModificationAdviceV02 {
	d.Message = new(IdentificationModificationAdviceV02)
	return d.Message
}

// Scope
// The IdentificationModificationAdvice message is sent by an assigner to an assignee. The message is used to advice on the correct party and/or account identification information.
// Usage
// The IdentificationModificationAdvice message is sent after the receipt of one or several transaction messages that included no longer valid party and/or account identification information.
// The IdentificationModificationAdvice message is exchanged between financial institutions and between financial institutions and non financial institutions and can contain one or more modification advises.
// There is no time limit on the time between the sending of an IdentificationModificationAdvice message and the receipt of the transaction messages that the IdentificationModificationAdvice refers to.
// The IdentificationModificationAdvice includes the correct party and/or account identification information, the IdentificationModificationAdvice or the included information may be forwarded to the initiating party of the transaction messages.
type IdentificationModificationAdviceV02 struct {

	// Identifies the identification assignment.
	Assignment *iso20022.IdentificationAssignment2 `xml:"Assgnmt"`

	// Provides reference information on the original message.
	OriginalTransactionReference *iso20022.OriginalTransactionReference18 `xml:"OrgnlTxRef,omitempty"`

	// Information concerning the identification data that is advised to be modified.
	Modification []*iso20022.IdentificationModification2 `xml:"Mod"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (i *IdentificationModificationAdviceV02) AddAssignment() *iso20022.IdentificationAssignment2 {
	i.Assignment = new(iso20022.IdentificationAssignment2)
	return i.Assignment
}

func (i *IdentificationModificationAdviceV02) AddOriginalTransactionReference() *iso20022.OriginalTransactionReference18 {
	i.OriginalTransactionReference = new(iso20022.OriginalTransactionReference18)
	return i.OriginalTransactionReference
}

func (i *IdentificationModificationAdviceV02) AddModification() *iso20022.IdentificationModification2 {
	newValue := new(iso20022.IdentificationModification2)
	i.Modification = append(i.Modification, newValue)
	return newValue
}

func (i *IdentificationModificationAdviceV02) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	i.SupplementaryData = append(i.SupplementaryData, newValue)
	return newValue
}
func ( d *Document02200102 ) String() (result string, ok bool) { return }
