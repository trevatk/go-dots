package dots

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {

	api := New(clientID, apiKey, true)

	p := &InputCreateTransactionParams{}

	r, e := api.CreateTransaction(context.TODO(), p)
	if e != nil {
		t.Log(e)
		t.FailNow()
	}

	assert.Equal(t, true, r.Success)
}

func TestCreateTransactionUnverifiedUser(t *testing.T) {

	api := New(clientID, apiKey, true)

	p := &InputCreateTransactionUnverifiedUserParams{}

	r, e := api.CreateTransactionUnverifiedUser(context.TODO(), p)
	if e != nil {
		t.Log(e)
		t.FailNow()
	}

	assert.Equal(t, true, r.Success)
}

func TestCreateTransactions(t *testing.T) {

	api := New(clientID, apiKey, true)

	p := &InputCreateTransactionsParams{}

	r, e := api.CreateTransactions(context.TODO(), p)
	if e != nil {
		t.Log(e)
		t.FailNow()
	}

	assert.Equal(t, true, r.Success)
}

func TestGetTransactionsBatchStatus(t *testing.T) {

	api := New(clientID, apiKey, true)

	r, e := api.GetTransactionsBatchStatus(context.TODO(), "")
	if e != nil {
		t.Log(e)
		t.FailNow()
	}

	assert.Equal(t, true, r.Success)
}

func TestListUserTransactions(t *testing.T) {

	api := New(clientID, apiKey, true)

	p := &InputListUserTransactions{}

	r, e := api.ListUserTransactions(context.TODO(), p)
	if e != nil {
		t.Log(e)
		t.FailNow()
	}

	assert.Equal(t, true, r.Success)
}

func TestGetTransactionByID(t *testing.T) {

	api := New(clientID, apiKey, true)

	r, e := api.GetTransactionByID(context.TODO(), "")
	if e != nil {
		t.Log(e)
		t.FailNow()
	}

	assert.Equal(t, true, r.Success)
}

func TestCreditUser(t *testing.T) {

	api := New(clientID, apiKey, true)

	p := &InputCreditUserParams{}

	r, e := api.CreditUser(context.TODO(), p)
	if e != nil {
		t.Log(e)
		t.FailNow()
	}

	assert.Equal(t, true, r.Success)
}

func TestRemoveCredit(t *testing.T) {

	api := New(clientID, apiKey, true)

	p := &InputRemoveCreditParams{}

	r, e := api.RemoveCredit(context.TODO(), p)
	if e != nil {
		t.Log(e)
		t.FailNow()
	}

	assert.Equal(t, true, r.Success)
}

func TestCreateACHPayment(t *testing.T) {

	api := New(clientID, apiKey, true)

	p := &InputCreateACHPaymentParams{}

	r, e := api.CreateACHPayment(context.TODO(), p)
	if e != nil {
		t.Log(e)
		t.FailNow()
	}

	assert.Equal(t, true, r.Success)
}
