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
