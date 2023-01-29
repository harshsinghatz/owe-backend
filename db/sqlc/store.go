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



// Create a transaction
// 1. Create a transaction from the sender to the reciever  
// 2. Sender balance ++
// 3. Reciever balance --
func (store *Store)RequestTx(ctx context.Context, arg ) error{
	
}