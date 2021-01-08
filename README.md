# GoQuick
GoQuick is a vendor package to quickly generate SQL packages based and a SQL table structure

GoQuick is not production ready and is still in development. You are encouraged to log issues so that this gets betetr

### Download
[darwin amd64](https://github.com/andrewwormald/goquick/raw/master/bin/gowuick_amd64_darwin)

#### Basic implementation 
```golang
//go:generate ../bin/goquick -schema_path=schema.sql -package_name=example
```
#### Only generate for specific tables using a comma separated list
```go
//go:generate ../bin/goquick -schema_path=schema.sql -package_name=example -tables=transactions
```

#### Example schema.sql file that is provided to goquick
```sql
create table users (
    id bigint not null,
    name varchar(255)not null,
);

create table transactions (
    id bigint auto increment primary,
    type int,
);
```

#### Example of the output of just one SQL table
```go
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
```
