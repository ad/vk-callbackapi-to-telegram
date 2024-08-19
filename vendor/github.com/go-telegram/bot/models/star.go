package models

import (
	"encoding/json"
	"fmt"
)

type TransactionPartnerType string

const (
	TransactionPartnerTypeFragment TransactionPartnerType = "fragment"
	TransactionPartnerTypeUser     TransactionPartnerType = "user"
	TransactionPartnerTypeOther    TransactionPartnerType = "other"
)

// TransactionPartner https://core.telegram.org/bots/api#transactionpartner
type TransactionPartner struct {
	Type TransactionPartnerType

	Fragment *TransactionPartnerFragment `json:"fragment,omitempty"`
	User     *TransactionPartnerUser     `json:"user,omitempty"`
	Other    *TransactionPartnerOther    `json:"other,omitempty"`
}

func (m *TransactionPartner) UnmarshalJSON(data []byte) error {
	v := struct {
		Type TransactionPartnerType `json:"type"`
	}{}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}

	switch v.Type {
	case TransactionPartnerTypeFragment:
		m.Type = TransactionPartnerTypeFragment
		m.Fragment = &TransactionPartnerFragment{}
		return json.Unmarshal(data, m.Fragment)
	case TransactionPartnerTypeUser:
		m.Type = TransactionPartnerTypeUser
		m.User = &TransactionPartnerUser{}
		return json.Unmarshal(data, m.User)
	case TransactionPartnerTypeOther:
		m.Type = TransactionPartnerTypeOther
		m.Other = &TransactionPartnerOther{}
		return json.Unmarshal(data, m.Other)
	}

	return fmt.Errorf("unsupported TransactionPartner type")
}

// TransactionPartnerFragment https://core.telegram.org/bots/api#transactionpartnerfragment
type TransactionPartnerFragment struct {
	Type            TransactionPartnerType  `json:"type"`
	WithdrawalState *RevenueWithdrawalState `json:"withdrawal_state,omitempty"`
}

// TransactionPartnerUser https://core.telegram.org/bots/api#transactionpartneruser
type TransactionPartnerUser struct {
	Type           TransactionPartnerType `json:"type"`
	User           User                   `json:"user"`
	InvoicePayload string                 `json:"invoice_payload,omitempty"`
	PaidMedia      []*PaidMedia           `json:"paid_media,omitempty"`
}

// TransactionPartnerOther https://core.telegram.org/bots/api#transactionpartnerother
type TransactionPartnerOther struct {
	Type TransactionPartnerType `json:"type"`
}

type RevenueWithdrawalStateType string

const (
	RevenueWithdrawalStateTypePending   RevenueWithdrawalStateType = "pending"
	RevenueWithdrawalStateTypeSucceeded RevenueWithdrawalStateType = "succeeded"
	RevenueWithdrawalStateTypeFailed    RevenueWithdrawalStateType = "failed"
)

// RevenueWithdrawalState https://core.telegram.org/bots/api#revenuewithdrawalstate
type RevenueWithdrawalState struct {
	Type RevenueWithdrawalStateType `json:"type"`

	Pending   *RevenueWithdrawalStatePending   `json:"pending,omitempty"`
	Succeeded *RevenueWithdrawalStateSucceeded `json:"succeeded,omitempty"`
	Failed    *RevenueWithdrawalStateFailed    `json:"failed,omitempty"`
}

func (m *RevenueWithdrawalState) UnmarshalJSON(data []byte) error {
	v := struct {
		Type RevenueWithdrawalStateType `json:"type"`
	}{}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}

	switch v.Type {
	case RevenueWithdrawalStateTypePending:
		m.Type = RevenueWithdrawalStateTypePending
		m.Pending = &RevenueWithdrawalStatePending{}
		return json.Unmarshal(data, m.Pending)
	case RevenueWithdrawalStateTypeSucceeded:
		m.Type = RevenueWithdrawalStateTypeSucceeded
		m.Succeeded = &RevenueWithdrawalStateSucceeded{}
		return json.Unmarshal(data, m.Succeeded)
	case RevenueWithdrawalStateTypeFailed:
		m.Type = RevenueWithdrawalStateTypeFailed
		m.Failed = &RevenueWithdrawalStateFailed{}
		return json.Unmarshal(data, m.Failed)
	}

	return fmt.Errorf("unsupported RevenueWithdrawalState type")
}

// RevenueWithdrawalStatePending https://core.telegram.org/bots/api#revenuewithdrawalstatepending
type RevenueWithdrawalStatePending struct {
	Type RevenueWithdrawalStateType `json:"type"`
}

// RevenueWithdrawalStateSucceeded https://core.telegram.org/bots/api#revenuewithdrawalstatesucceeded
type RevenueWithdrawalStateSucceeded struct {
	Type RevenueWithdrawalStateType `json:"type"`
	Date int                        `json:"date"`
	URL  string                     `json:"url"`
}

// RevenueWithdrawalStateFailed https://core.telegram.org/bots/api#revenuewithdrawalstatefailed
type RevenueWithdrawalStateFailed struct {
	Type RevenueWithdrawalStateType `json:"type"`
}

// StarTransaction https://core.telegram.org/bots/api#startransaction
type StarTransaction struct {
	ID       string              `json:"id"`
	Amount   int                 `json:"amount"`
	Date     int                 `json:"date"`
	Source   *TransactionPartner `json:"source,omitempty"`
	Receiver *TransactionPartner `json:"receiver,omitempty"`
}

// StarTransactions https://core.telegram.org/bots/api#startransactions
type StarTransactions struct {
	Transactions []StarTransaction `json:"transactions"`
}

// TransactionPartnerTelegramAds https://core.telegram.org/bots/api#transactionpartnertelegramads
type TransactionPartnerTelegramAds struct {
	Type string `json:"type"`
}
