type Schema struct {
	Extension	[]Extension	`xml:"complexType>simpleContent>extension"`
	Attribute	[]Attribute	`xml:"complexType>simpleContent>extension>attribute"`
	Enumeration	[][]Enumeration	`xml:"simpleType>restriction>enumeration"`
	MinInclusive	[]MinInclusive	`xml:"simpleType>restriction>minInclusive"`
	Length	[]Length	`xml:"simpleType>restriction>length"`
	Xmlns	string	`xml:"xmlns,attr"`
	ElementChoiceComplexType	[][]Element	`xml:"complexType>choice>element"`
	Any	[]Any	`xml:"complexType>sequence>any"`
	SimpleType	[]SimpleType	`xml:"simpleType"`
	FractionDigits	[]FractionDigits	`xml:"simpleType>restriction>fractionDigits"`
	TotalDigits	[]TotalDigits	`xml:"simpleType>restriction>totalDigits"`
	Pattern	[]Pattern	`xml:"simpleType>restriction>pattern"`
	MaxLength	[]MaxLength	`xml:"simpleType>restriction>maxLength"`
	Element	Element	`xml:"element"`
	ElementSequenceComplexType	[][]Element	`xml:"complexType>sequence>element"`
	MinLength	[]MinLength	`xml:"simpleType>restriction>minLength"`
	Xs	string	`xml:"xs,attr"`
	ElementFormDefault	string	`xml:"elementFormDefault,attr"`
	ComplexType	[]ComplexType	`xml:"complexType"`
	Restriction	[]Restriction	`xml:"simpleType>restriction"`
	TargetNamespace	string	`xml:"targetNamespace,attr"`
}

type ElementSequenceComplexType struct {
	MinOccurs	string	`xml:"minOccurs,attr"`
	Name	string	`xml:"name,attr"`
	Type	string	`xml:"type,attr"`
	MaxOccurs	string	`xml:"maxOccurs,attr"`
}
type Attribute struct {
	Name	string	`xml:"name,attr"`
	Type	string	`xml:"type,attr"`
	Use	string	`xml:"use,attr"`
}
type Enumeration struct {
	Value	string	`xml:"value,attr"`
}
type Any struct {
	ProcessContents	string	`xml:"processContents,attr"`
	Namespace	string	`xml:"namespace,attr"`
}
type TotalDigits struct {
	Value	string	`xml:"value,attr"`
}
type MinLength struct {
	Value	string	`xml:"value,attr"`
}
type MaxLength struct {
	Value	string	`xml:"value,attr"`
}
type ElementChoiceComplexType struct {
	Name	string	`xml:"name,attr"`
	Type	string	`xml:"type,attr"`
}
type SimpleType struct {
	Name	string	`xml:"name,attr"`
}
type FractionDigits struct {
	Value	string	`xml:"value,attr"`
}
type Length struct {
	Value	string	`xml:"value,attr"`
}
type Pattern struct {
	Value	string	`xml:"value,attr"`
}
type Element struct {
	Type	string	`xml:"type,attr"`
	Name	string	`xml:"name,attr"`
}
type ComplexType struct {
	Name	string	`xml:"name,attr"`
}
type Extension struct {
	Base	string	`xml:"base,attr"`
}
type Restriction struct {
	Base	string	`xml:"base,attr"`
}
type MinInclusive struct {
	Value	string	`xml:"value,attr"`
}

