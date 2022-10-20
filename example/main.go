package main

import (
	"context"
	"fmt"
	"log"

	"github.com/trevatk/go-dots"
)

func main() {

	var (
		clientID = "xyx" // retrieve from dots.env
		apiKey   = "zxz" // retrieve from dots.env
	)

	ctx := context.TODO()

	api := dots.New(clientID, apiKey, true)

	// new user parameters
	nup := &dots.InputCreateUserParams{
		Email:       "johndoe@gmail.com",
		CountryCode: "1", // united states
		PhoneNumber: "8675309",
		FirstName:   "John",
		LastName:    "Doe",
		Username:    "john.doe",
	}

	// make create user api call
	// verificationID is required to send user a verification code
	cup, err := api.CreateUser(ctx, nup)
	if err != nil {
		log.Fatal(err)
	}

	// pass verification token into parameters
	svt := &dots.InputSendVerificationTokenParams{
		VerificationID: cup.VerificationID,
	}

	// make send verification code api call
	_, err = api.SendVerificationToken(ctx, svt)
	if err != nil {
		log.Fatal(err)
	}

	// the end user will receive a text message with a code
	// the next api call will verify the user using that code

	// pass the verificationID from above and code provided by end user
	vut := &dots.InputVerifyUserTokenParams{
		VerificationID:    cup.VerificationID,
		VerificationToken: "505123",
	}

	// make api call
	vur, err := api.VerifyUserToken(ctx, vut)
	if err != nil {
		log.Fatal(err)
	}

	// YAY!
	fmt.Printf("successfully verified user %s", vur.User.ID)
}
