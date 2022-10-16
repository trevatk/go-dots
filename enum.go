package dots

type EntityTypeEnum string
type CountryEnum string
type StateEnum string
type InvoiceStatusEnum string
type PayoutMethodEnum string
type ACHAccountTypeEnum string
type TransactionTypeEnum string

const (
	Individual EntityTypeEnum = "individual"
	Business   EntityTypeEnum = "business"

	AU CountryEnum = "AU"
	AT CountryEnum = "AT"
	BE CountryEnum = "BE"
	BR CountryEnum = "BR"
	BG CountryEnum = "BG"
	CA CountryEnum = "CA"
	CY CountryEnum = "CY"
	CZ CountryEnum = "CZ"
	DK CountryEnum = "DK"
	EE CountryEnum = "EE"
	FI CountryEnum = "FI"
	FR CountryEnum = "FR"
	DE CountryEnum = "DE"
	GR CountryEnum = "GR"
	HK CountryEnum = "HK"
	HU CountryEnum = "HU"
	IE CountryEnum = "IE"
	IT CountryEnum = "IT"
	JP CountryEnum = "JP"
	LV CountryEnum = "LV"
	LT CountryEnum = "LT"
	LU CountryEnum = "LU"
	MT CountryEnum = "MT"
	NL CountryEnum = "NL"
	NC CountryEnum = "NC"
	NO CountryEnum = "NO"
	PL CountryEnum = "PL"
	PT CountryEnum = "PT"
	RO CountryEnum = "RO"
	SG CountryEnum = "SG"
	SK CountryEnum = "SK"
	SI CountryEnum = "SI"
	ES CountryEnum = "ES"
	SE CountryEnum = "SE"
	CH CountryEnum = "CH"
	AE CountryEnum = "AE"
	GB CountryEnum = "GB"
	US CountryEnum = "US"

	AL         StateEnum = "AL"
	AK         StateEnum = "AK"
	AS         StateEnum = "AS"
	AZ         StateEnum = "AZ"
	AR         StateEnum = "AR"
	California StateEnum = "CA"
	CO         StateEnum = "CO"
	CT         StateEnum = "CT"
	Deleware   StateEnum = "DE"
	DC         StateEnum = "DC"
	FM         StateEnum = "FM"
	FL         StateEnum = "FL"
	GA         StateEnum = "GA"
	GU         StateEnum = "GU"
	HI         StateEnum = "HI"
	ID         StateEnum = "ID"
	IL         StateEnum = "IL"
	IN         StateEnum = "IN"

	Created   InvoiceStatusEnum = "created"
	Completed InvoiceStatusEnum = "completed"
	Expired   InvoiceStatusEnum = "expired"

	Paypal  PayoutMethodEnum = "paypal"
	Venmo   PayoutMethodEnum = "venmo"
	ACH     PayoutMethodEnum = "ach"
	CashApp PayoutMethodEnum = "cash_app"

	Checking ACHAccountTypeEnum = "checking"
	Savings  ACHAccountTypeEnum = "savings"

	Credit     TransactionTypeEnum = "credit"
	WalletEnum TransactionTypeEnum = "wallet"
)
