package auth

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
)

/*
ISO 20022 Implementation
Business Domain: Payments
Message Set: Authorities and Financial Investigations (auth)
Message Definitions: https://www.iso20022.org/payments_messages.page
Message Schemas: https://github.com/figassis/iso20022/tree/master/msg-schemas/payments/auth

Messages
1.	InformationRequestOpeningV01 - auth.001.001.01
2.	InformationRequestResponseV01 - auth.002.001.01
3.	InformationRequestStatusChangeNotificationV01 - auth.003.001.01
*/

type AccountHolder struct {
	AccountNumber string
	BankNumber    string
}
