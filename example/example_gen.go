// Code generated by go generate; DO NOT EDIT.

package example

import (
	"context"
	"database/sql"
)

type Transaction struct {
	Id   int64
	Type int32
}

func listTransactionsAfterFrom(ctx context.Context, dbc *sql.DB, afterFromStatement string) ([]Transaction, error) {
	rows, err := dbc.QueryContext(ctx, "select id, type from transactions "+afterFromStatement+";")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ls []Transaction
	for rows.Next() {
		var l Transaction
		err := rows.Scan(
			&l.Id,
			&l.Type,
		)
		if err != nil {
			return nil, err
		}

		ls = append(ls, l)
	}

	return ls, nil
}

func listTransactionsWhere(ctx context.Context, dbc *sql.DB, where string) ([]Transaction, error) {
	return listTransactionAfterFrom(ctx, dbc, where)
}

func lookupTransactionAfterFrom(ctx context.Context, dbc *sql.DB, afterFromStatement string) (Transaction, error) {
	rows, err := dbc.QueryContext(ctx, "select id, type from transactions "+afterFromStatement+";")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var l Transaction
	for rows.Next() {
		err := rows.Scan(
			&l.Id,
			&l.Type,
		)
		if err != nil {
			return nil, err
		}
	}

	return l, nil
}

func lookupTransactionWhere(ctx context.Context, dbc *sql.DB, where string) (Transaction, error) {
	return lookupTransactionAfterFrom(ctx, dbc, where)
}

func InsertTransaction(ctx context.Context, dbc *sql.DB, strct Transaction) (int64, error) {
	res, err := dbc.ExecContext(ctx, "insert into transactions "+
		"set id=? ,type=?",
		strct.Id,
		strct.Type,
	)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func DeleteTransaction(ctx context.Context, dbc *sql.DB, where string) error {
	_, err := dbc.ExecContext(ctx, "delete from transactions "+where)
	return err
}

func UpdateTransaction(ctx context.Context, dbc *sql.DB, where, set string, args ...interface{}) error {
	_, err = dbc.ExecContext(ctx, "update transactions set "+set+where, args)
	return err
}