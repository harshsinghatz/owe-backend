package db

import (
	"context"
	"testing"

	"github.com/harshsinghatz/owe-backend/util"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T){
	arg:= CreateAccountParams{
		Name: util.RandomName(),
		PhoneNumber: "234324423",
		Currency: "Rupees",
		Balance: 10,
	}

	account, err := testQueries.CreateAccount(context.Background(),arg)

	require.NoError(t, err)
	require.NotEmpty(t,account)

	require.Equal(t,arg.Name,account.Name)
	require.Equal(t,arg.PhoneNumber,account.PhoneNumber)
	require.Equal(t,arg.Currency,account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt);
}