package auth

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
)

/*
Implementation of the ISO20022 standard for Authorities and Financial Investigations
https://www.iso20022.org/payments_messages.page - auth

Authorities (auth) transactions are as follows:
1  -  InformationRequestOpeningV01

2  -  InformationRequestResponseV01
3  -  InformationRequestStatusChangeNotificationV01
*/

type AccountHolder struct {
	AccountNumber string
	BankNumber    string
}
