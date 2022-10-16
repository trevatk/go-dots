package dots

import "fmt"

// ErrEmailExits
type ErrEmailExists struct {
	Email     string
	Operation string
}

// Error
func (ee *ErrEmailExists) Error() string {
	return fmt.Sprintf("dots api %s err email %s already exists", ee.Operation, ee.Email)
}

// ErrUsernameExists
type ErrUsernameExists struct {
	Username  string
	Operation string
}

// Error
func (ue *ErrUsernameExists) Error() string {
	return fmt.Sprintf("dots api %s err username %s already exists", ue.Operation, ue.Username)
}

// ErrInvalidInput
type ErrInvalidInput struct {
	Field string
}

// ErrInvalidInput
func (ii *ErrInvalidInput) Error() string {
	return fmt.Sprintf("dots api invalid input field %s", ii.Field)
}

// ErrInvalidSSN
type ErrInvalidSSN struct {
	Last4 string
}

// ErrInvalidSSN
func (is *ErrInvalidSSN) Error() string {
	return fmt.Sprintf("dots api invalid SSN %s", is.Last4[6:10])
}

// ErrInvalidUsername
type ErrInvalidUsername struct {
	Username string
}

// Error
func (iu *ErrInvalidUsername) Error() string {
	return fmt.Sprintf("dots api invalid username %s. Usernames may only contain letters, digits and '-'. They must not contain swear words or protected words.", iu.Username)
}

// ErrInvalidPhoneNumberVOIP
type ErrInvalidPhoneNumberVOIP struct {
	Phone string
}

// Error
func (ip *ErrInvalidPhoneNumberVOIP) Error() string {
	return fmt.Sprintf("dots api invalid phone number %s. VOIP phone number not allowed.", ip.Phone)
}

// ErrInvalidInvoiceRequestedInformation
type ErrInvalidInvoiceRequestedInformation struct {
	// TODO:
	// add field throwing error
}

// Error
func (iiri *ErrInvalidInvoiceRequestedInformation) Error() string {
	return fmt.Sprintf("dots api invalid value for requested_information")
}

// ErrInvalidInvoiceBreakdown
type ErrInvalidInvoiceBreakdown struct{}

// Error
func (iib *ErrInvalidInvoiceBreakdown) Error() string {
	return fmt.Sprintf("dots api invalid invoice breakdown")
}

// ErrInvalidInvoiceItem
type ErrInvalidInvoiceItem struct {
	// TODO:
	// add invalid invoice item
}

// Error
func (iii *ErrInvalidInvoiceItem) Error() string {
	return fmt.Sprintf("dots api invalid invoice item")
}

// ErrInvalidInvoiceItemTotal
type ErrInvalidInvoiceItemTotal struct {
	// TODO:
	// add expected value field
	// add actual value field
}

// Error
func (iiit *ErrInvalidInvoiceItemTotal) Error() string {
	return fmt.Sprintf("dots api item totals add up incorrectly")
}

// ErrInvalidTransacion
type ErrInvalidTransaction struct{}

// Error
func (it *ErrInvalidTransaction) Error() string {
	return fmt.Sprintf("dots api invalid transaction")
}

// ErrServiceUnavailable
type ErrServiceUnavailable struct{}

// Error
func (se *ErrServiceUnavailable) Error() string {
	return fmt.Sprintf("dots api a service wasa unavailable")
}

// ErrUserInfoMissing
type ErrUserInfoMissing struct{}

// Error
func (uim *ErrUserInfoMissing) Error() string {
	return fmt.Sprintf("dots api user has missing or invalid info")
}

// ErrUserInfoMissingCashApp
type ErrUserInfoMissingCashApp struct{}

// Error
func (uimca *ErrUserInfoMissingCashApp) Error() string {
	return fmt.Sprintf("dots api user missing cash app account info")
}

// ErrUserInfoMissingContact
type ErrUserInfoMissingContact struct{}

// Error
func (uimc *ErrUserInfoMissingContact) Error() string {
	return fmt.Sprintf("dots api user missing contact info")
}

// ErrUserInfoMissingCard
type ErrUserInfoMissingCard struct{}

// Error
func (uimc *ErrUserInfoMissingCard) Error() string {
	return fmt.Sprintf("dots api user midding credit card info")
}

// ErrAppInfoMissing
type ErrAppInfoMissing struct{}

// Error
func (aim *ErrAppInfoMissing) Error() string {
	return fmt.Sprintf("dots api app has missing or invalid info")
}

// ErrTransactionInsufficientFunds
type ErrTransactionInsufficientFunds struct{}

// Error
func (tif *ErrTransactionInsufficientFunds) Error() string {
	return fmt.Sprintf("dots api not enough funds in wallet")
}

// ErrResourceBusy
type ErrResourceBusy struct{}

// Error
func (rb *ErrResourceBusy) Error() string {
	return fmt.Sprintf("dots api this is unavailable because one or more previous requests are still being processed")
}
