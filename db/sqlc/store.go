package db

import (
	"context"
	"database/sql"
	"fmt"
)

//composition
type Store struct{
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store{
	return &Store{
		Queries: New(db),
		db: db,
	}
}

func (store *Store) execTx(ctx context.Context,fn func(*Queries)error)error{
	tx, err:= store.db.BeginTx(ctx,nil)

	if err!=nil{
		return err
	}

	q:=New(tx)
	qErr:=fn(q)

	if qErr!=nil{
		if rbErr:=tx.Rollback(); rbErr!=nil{
			return fmt.Errorf("tx err: %v, rb err: %v",err,rbErr)
		}
		return qErr
	}

	return tx.Commit()
}

// Create a debt request 
// 1. Create a transaction with reciever_id = id and sender_id as the account id of the account you are requesting debt status="pending" and type="debt"
// 2. When the sender accepts the request following things happen
// 		1. The tranaction status changes from "pending" to "accepted"
// 		2. sender_account balance -> account balance + transaction amount
// 		3. reciever_account balance -> account balance - transaction amount
type TxParams struct {
	TransactionId int64
}

type TxResults struct {
	Transaction Transaction
	FromAccount Account
	ToAccount Account
}

func (store *Store)AcceptDebtTransaction(ctx context.Context, arg TxParams ) (TxResults,error){
	var result TxResults

	err:= store.execTx(ctx,func(q *Queries)error{
		// The tranaction status changes from "pending" to "accepted"
		upErr:=q.UpdateTransactionStatus(ctx,UpdateTransactionStatusParams {
			ID: arg.TransactionId,
			Status: "accepted",
		})
		
		if upErr!=nil {
			return upErr
		}

		transaction,getTransErr:= q.GetTranaction(ctx,arg.TransactionId)

		if getTransErr!=nil{
			return getTransErr
		}

		// sender_account balance -> account balance + transaction amount
		senderAccount,senderAccountErr:=q.GetAccount(ctx,transaction.SenderID)

		if senderAccountErr!=nil{
			return senderAccountErr
		}

		updateSenderBalErr:=q.UpdateAccountBalance(ctx,UpdateAccountBalanceParams{
			ID: senderAccount.ID,
			Balance: senderAccount.Balance + transaction.Amount,
		})

		if updateSenderBalErr!=nil{
			return updateSenderBalErr
		}

		//reciever_account balance -> account balance - transaction amount
		recieverAccount,recieverAccountErr:=q.GetAccount(ctx,transaction.SenderID)

		if recieverAccountErr!=nil{
			return recieverAccountErr
		}

		updateRecieverBalErr:=q.UpdateAccountBalance(ctx,UpdateAccountBalanceParams{
			ID: recieverAccount.ID,
			Balance: recieverAccount.Balance - transaction.Amount,
		})

		if updateRecieverBalErr!=nil{
			return updateRecieverBalErr
		}
	
		return nil
	})

	return result, err
}

// Create a payment request 
// 1. Update the transaction type = "pay" and status = "pending"
// 2. When the sender accepts the request following things happen
// 		1. The tranaction status changes from "pending" to "accepted"
// 		2. sender_account balance -> account balance - transaction amount
// 		3. reciever_account balance -> account balance + transaction amount
//		4. change the transaction balance -> transaction balance - payingAmount

type TxAcceptParams struct {
	TransactionId int64
	PayingAmount int64
}
func (store *Store)AcceptPayTransaction(ctx context.Context, arg TxAcceptParams ) (TxResults,error){
	var result TxResults

	err:= store.execTx(ctx,func(q *Queries)error{
		// The tranaction status changes from "pending" to "accepted"
		upErr:=q.UpdateTransactionStatus(ctx,UpdateTransactionStatusParams {
			ID: arg.TransactionId,
			Status: "accepted",
		})
		
		if upErr!=nil {
			return upErr
		}

		transaction,getTransErr:= q.GetTranaction(ctx,arg.TransactionId)

		if getTransErr!=nil{
			return getTransErr
		}

		// sender_account balance -> account balance + transaction amount
		senderAccount,senderAccountErr:=q.GetAccount(ctx,transaction.SenderID)

		if senderAccountErr!=nil{
			return senderAccountErr
		}

		updateSenderBalErr:=q.UpdateAccountBalance(ctx,UpdateAccountBalanceParams{
			ID: senderAccount.ID,
			Balance: senderAccount.Balance - arg.PayingAmount,
		})

		if updateSenderBalErr!=nil{
			return updateSenderBalErr
		}

		//reciever_account balance -> account balance - transaction amount
		recieverAccount,recieverAccountErr:=q.GetAccount(ctx,transaction.SenderID)

		if recieverAccountErr!=nil{
			return recieverAccountErr
		}

		updateRecieverBalErr:=q.UpdateAccountBalance(ctx,UpdateAccountBalanceParams{
			ID: recieverAccount.ID,
			Balance: recieverAccount.Balance + arg.PayingAmount,
		})

		if updateRecieverBalErr!=nil{
			return updateRecieverBalErr
		}

		//change the transaction balance -> transaction balance - payingAmount
		updateTransErr:=q.UpdateTransactionAmount(ctx,UpdateTransactionAmountParams{
			Amount: transaction.Amount - arg.PayingAmount,
		})
	
		if updateTransErr!=nil{
			return updateTransErr
		}
		
		return nil
	})

	return result, err
}