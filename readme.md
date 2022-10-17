# Go-Dots

## Installation

```bash
$ go get github.com/trevatk/go-dots
```

## Simple Workflow

```golang

ctx := context.TODO()

// create new dots api instance
api := New(clientID, apiKey, true)

// new user parameters
nup := &InputCreateUserParams{
    // add fields
}

// create user
api.CreateUser(context.TODO(), nup)

// send verification token to user

// submit verification token provided from user

```