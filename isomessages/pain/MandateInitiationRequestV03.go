package pain

import (
	"encoding/xml"

	"github.com/figassis/bankiso/iso20022"
)

type Document00900103 struct {
	XMLName xml.Name                     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.03 Document"`
	Message *MandateInitiationRequestV03 `xml:"MndtInitnReq"`
}

func (d *Document00900103) AddMessage() *MandateInitiationRequestV03 {
	d.Message = new(MandateInitiationRequestV03)
	return d.Message
}

// Scope
// The MandateInitiationRequest message is sent by the initiator of the request to his agent. The initiator can either be the debtor or the creditor.
// The MandateInitiationRequest message is forwarded by the agent of the initiator to the agent of the counterparty.
// The MandateInitiationRequest message is used to set-up the instruction that allows the debtor agent to accept instructions from the creditor, through the creditor agent, to debit the account of the debtor.
// Usage
// The MandateInitiationRequest message can contain one or more request(s) to set-up a specific mandate.
// The messages can be exchanged between creditor and creditor agent or debtor and debtor agent and between creditor agent and debtor agent.
// The message can also be used by an initiating party that has authority to send the message on behalf of the creditor or debtor.
// The MandateInitiationRequest message can be used in domestic and cross-border scenarios.
type MandateInitiationRequestV03 struct {

	// Set of characteristics to identify the message and parties playing a role in the mandate initiation, but which are not part of the mandate.
	GroupHeader *iso20022.GroupHeader47 `xml:"GrpHdr"`

	// Set of elements used to provide the details of the mandate signed between the (ultimate) creditor and the (ultimate) debtor.
	Mandate []*iso20022.Mandate4 `xml:"Mndt"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (m *MandateInitiationRequestV03) AddGroupHeader() *iso20022.GroupHeader47 {
	m.GroupHeader = new(iso20022.GroupHeader47)
	return m.GroupHeader
}

func (m *MandateInitiationRequestV03) AddMandate() *iso20022.Mandate4 {
	newValue := new(iso20022.Mandate4)
	m.Mandate = append(m.Mandate, newValue)
	return newValue
}

func (m *MandateInitiationRequestV03) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	m.SupplementaryData = append(m.SupplementaryData, newValue)
	return newValue
}
func ( d *Document00900103 ) String() (result string, ok bool) { return }
