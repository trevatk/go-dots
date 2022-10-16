package dots

// Account
type Account struct {
	ID       string `json:"id"`       // string <uuid>
	UserID   string `json:"user_id"`  // string <uuid>
	Currency string `json:"currency"` // string <iso-4217>
	Country  string `json:"country"`  // string <iso-3166>
	Mask     string `json:"mask"`
}

// Delivery
type Delivery struct {
	Method string `json:"method"`
}

// Payee
type Payee struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	CountryCode string `json:"country_code"`
	PhoneNumber string `json:"phone_number"`
}

// PayoutLink
type PayoutLink struct {
	ID             string `json:"id"`
	Link           string `json:"link"`
	OriginalAmount int    `json:"original_amount"`
	Amount         int    `json:"amount"`
	Status         string `json:"status"`
	APIAppName     string `json:"api_app_name"`
	Payee          *Payee `json:"payee"`
}

// PayoutMethods
type PayoutMethods struct {
	ACHAccouns []string
	Paypal     string
	Venmo      string
}

// Plaid
type Plaid struct {
	Balances     *PlaidBalances       `json:"balances"`
	Transactions []*PlaidTransactions `json:"transactions"`
}

// PlaidBalances
type PlaidBalances struct {
	Available              int    `json:"available"`
	Current                int    `json:"current"`
	Limit                  int    `json:"limit"`
	ISOCurrencyCode        string `json:"iso_currency_code"`
	UnofficialCurrencyCode string `json:"unofficial_currency_code"`
}

// PlaidPaymentMeta
type PlaidPaymentMeta struct {
	ByOrderOf        string `json:"by_order_of"`
	Payee            string `json:"payee"`
	Payer            string `json:"payer"`
	PaymentMethod    string `json:"payment_method"`
	PaymentProcessor string `json:"payment_processor"`
	PPDID            string `json:"ppd_id"`
	Reason           string `json:"reason"`
	ReferenceNumber  string `json:"reference_number"`
}

// PlaidLocation
type PlaidLocation struct {
	Address   string `json:"address"`
	City      string `json:"city"`
	State     string `json:"state"`
	Zip       string `json:"zip"`
	Latitude  string `json:"lat"`
	Longitude string `json:"lon"`
}

// PlaidTransactions
type PlaidTransactions struct {
	AccountID              string            `json:"account_id"`
	Amount                 int               `json:"amount"`
	ISOCurrencyCode        string            `json:"iso_currency_code"`
	UnofficialCurrencyCode string            `json:"unofficial_currency_code"`
	Category               []string          `json:"category"`
	CategoryID             string            `json:"category_id"`
	Date                   string            `json:"date"`
	DateTime               string            `json:"datetime"`
	Location               *PlaidLocation    `json:"location"`
	Name                   string            `json:"name"`
	PaymentMeta            *PlaidPaymentMeta `json:"payment_meta"`
	Pending                bool              `json:"pending"`
	PendingTransactionID   string            `json:"pending_transaction_id"`
	TransactionID          string            `json:"transaction_id"`
}

// Receipt
type Receipt struct {
	Items     []*Item    `json:"items"`
	Breakdown *Breakdown `json:"breakdown"`
}

// Transaction
type Transaction struct {
	ID                  int                 `json:"id"`
	Date                string              `json:"date"`
	SourceUsername      string              `json:"source_username"`
	DestinationUsername string              `json:"destination_username"`
	Amount              int                 `json:"amount"`
	Type                TransactionTypeEnum `json:"type"`
	Completed           bool                `json:"completed"`
	Notes               interface{}         `json:"notes"`
	Receipt             *Receipt            `json:"receipt"`
	CreditTransactionID string              `json:"credit_transaction_id"`
}

// Transactions
type Transactions struct {
	UserID  string      `json:"user_id"`
	Amount  int         `json:"amount"`
	Receipt *Receipt    `json:"receipt"`
	Notes   interface{} `json:"notes"`
}

// UserLimited
type UserLimited struct {
	ID        string         `json:"id,omitempty"` // string or null
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Username  string         `json:"username"`
	Email     string         `json:"email"` // string <email>
	Wallet    *WalletLimited `json:"wallet_limited"`
}

// VerifyUser
type VerifyUser struct {
	ID string `json:"id"` // string <uuid>
}

// Wallet
type Wallet struct {
	Amount             int `json:"amount"`              // user's balance in cents
	WithdrawableAmount int `json:"withdrawable_amount"` // user's balance they can withdraw
	CreditBalance      int `json:"credit_balance"`      // user's credit balance
}

// WalletLimited
type WalletLimited struct {
	Amount             int `json:"amount"`
	WithdrawableAmount int `json:"withdrawable_amount"`
}
