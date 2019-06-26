package acmt

import (
	"encoding/xml"

	"github.com/figassis/bankiso/iso20022"
)

type Document02200101 struct {
	XMLName xml.Name                             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.01 Document"`
	Message *IdentificationModificationAdviceV01 `xml:"IdModAdvc"`
}

func (d *Document02200101) AddMessage() *IdentificationModificationAdviceV01 {
	d.Message = new(IdentificationModificationAdviceV01)
	return d.Message
}

// Scope
// The IdentificationModificationAdvice message is sent by an assigner to an assignee. The message is used to advice on the correct party and/or account identification information.
// Usage
// The IdentificationModificationAdvice message is sent after the receipt of one or several transaction messages that included no longer valid party and/or account identification information.
// The IdentificationModificationAdvice message is exchanged between financial institutions and between financial institutions and non financial institutions and can contain one or more modification advises.
// There is no time limit on the time between the sending of an IdentificationModificationAdvice message and the receipt of the transaction messages that the IdentificationModificationAdvice refers to.
// The IdentificationModificationAdvice includes the correct party and/or account identification information, the IdentificationModificationAdvice or the included information may be forwarded to the initiating party of the transaction messages.
type IdentificationModificationAdviceV01 struct {

	// Identifies the identification assignment.
	Assignment *iso20022.IdentificationAssignment1 `xml:"Assgnmt"`

	// Provides reference information on the original message.
	OriginalTransactionReference *iso20022.OriginalTransactionReference14 `xml:"OrgnlTxRef,omitempty"`

	// Information concerning the identification data that is advised to be modified.
	Modification []*iso20022.IdentificationModification1 `xml:"Mod"`
}

func (i *IdentificationModificationAdviceV01) AddAssignment() *iso20022.IdentificationAssignment1 {
	i.Assignment = new(iso20022.IdentificationAssignment1)
	return i.Assignment
}

func (i *IdentificationModificationAdviceV01) AddOriginalTransactionReference() *iso20022.OriginalTransactionReference14 {
	i.OriginalTransactionReference = new(iso20022.OriginalTransactionReference14)
	return i.OriginalTransactionReference
}

func (i *IdentificationModificationAdviceV01) AddModification() *iso20022.IdentificationModification1 {
	newValue := new(iso20022.IdentificationModification1)
	i.Modification = append(i.Modification, newValue)
	return newValue
}
func ( d *Document02200101 ) String() (result string, ok bool) { return }
