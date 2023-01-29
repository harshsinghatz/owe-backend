package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestAcceptPayTransaction(t *testing.T){
	store:= NewStore(testDB)

	accArg1:=CreateAccountParams{
	Name: "Harsh Singh",
		PhoneNumber: "23489329",
		Currency: "ssdif",
		Balance: 0,
	}

	accArg2:=CreateAccountParams{
		Name: "Alex",
		PhoneNumber: "23232",
		Currency: "ssdif",
		Balance: 0,
	}
	account1,err:= testQueries.CreateAccount(context.Background(),accArg1)
	account2,err:= testQueries.CreateAccount(context.Background(),accArg2)

	arg:= CreateTransactionParams{
		RecieverID: account1.ID,
    	SenderID:   account2.ID,
    	Currency:   "sdfisdf",
    	Amount:     10,
    	Message:    sql.NullString{String:"", Valid: true},
    	Deadline:   sql.NullTime{Time:time.Now(), Valid: true},
    	Status: "pending",
		Type: "debt",
	}

	tranaction, err := testQueries.CreateTransaction(context.Background(),arg)
	fmt.Println(tranaction.RecieverID,tranaction.SenderID)

	res,err:=store.AcceptDebtTransaction(context.Background(),TxParams{
		TransactionId: tranaction.ID,
	})

	require.NoError(t,err);

	require.Equal(t,res.Transaction.Status,"accepted")
	require.Equal(t,res.Transaction.Status,"accepted")
	require.Equal(t,res.FromAccount.Balance,int64(10))
	require.Equal(t,res.ToAccount.Balance,int64(-10))
	// run n concurrent transfer transaction

	testQueries.UpdateTransactionType(context.Background(),UpdateTransactionTypeParams{
		ID: tranaction.ID,
		Type: "pay",
	})

	testQueries.UpdateTransactionStatus(context.Background(),UpdateTransactionStatusParams{
		ID: tranaction.ID,
		Status: "pending",
	})

	result,err:=store.AcceptPayTransaction(context.Background(),TxAcceptParams{
		TransactionId: tranaction.ID,
		PayingAmount: 5,
	})
	
	require.Equal(t,result.Transaction.Status,"accepted")
	require.Equal(t,result.FromAccount.Balance,int64(5))
	require.Equal(t,result.ToAccount.Balance,int64(-5))
}