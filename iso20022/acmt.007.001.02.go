package iso20022

import (
	"errors"
	//	"fmt"
)

//Defines an ISO 20022 Message
//Implements ISOMessage interface
type ACMT00700102 struct {
	Name string
}

//Converts message struct to valid ISO 20022 XML message
func (message *ACMT00700102) String() (result string, err error) {
	err = errors.New("Function not yet implemented")
	return
}
