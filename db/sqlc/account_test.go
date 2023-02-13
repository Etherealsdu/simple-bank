package db

import (
	"context"
	"database/sql"
	"github.com/Etherealsdu/simplebank/util"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"time"
)

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)

}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      account1.ID,
		Balance: util.RandomMoney(),
	}

	account2, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, arg.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)

}

func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}

func TestNew(t *testing.T) {
	type args struct {
		db DBTX
	}
	tests := []struct {
		name string
		args args
		want *Queries
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueries_CreateAccount(t *testing.T) {
	type fields struct {
		db DBTX
	}
	type args struct {
		ctx context.Context
		arg CreateAccountParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Account
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queries{
				db: tt.fields.db,
			}
			got, err := q.CreateAccount(tt.args.ctx, tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateAccount() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueries_CreateEntry(t *testing.T) {
	type fields struct {
		db DBTX
	}
	type args struct {
		ctx context.Context
		arg CreateEntryParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Entry
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queries{
				db: tt.fields.db,
			}
			got, err := q.CreateEntry(tt.args.ctx, tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateEntry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateEntry() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueries_CreateTransfer(t *testing.T) {
	type fields struct {
		db DBTX
	}
	type args struct {
		ctx context.Context
		arg CreateTransferParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Transfer
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queries{
				db: tt.fields.db,
			}
			got, err := q.CreateTransfer(tt.args.ctx, tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTransfer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTransfer() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueries_DeleteAccount(t *testing.T) {
	type fields struct {
		db DBTX
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queries{
				db: tt.fields.db,
			}
			if err := q.DeleteAccount(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteAccount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestQueries_GetAccount(t *testing.T) {
	type fields struct {
		db DBTX
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Account
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queries{
				db: tt.fields.db,
			}
			got, err := q.GetAccount(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAccount() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueries_GetEntry(t *testing.T) {
	type fields struct {
		db DBTX
	}
	type args struct {
		ctx       context.Context
		accountID int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Entry
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queries{
				db: tt.fields.db,
			}
			got, err := q.GetEntry(tt.args.ctx, tt.args.accountID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEntry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEntry() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueries_GetTransfer(t *testing.T) {
	type fields struct {
		db DBTX
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Transfer
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queries{
				db: tt.fields.db,
			}
			got, err := q.GetTransfer(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTransfer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTransfer() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueries_ListAccounts(t *testing.T) {
	type fields struct {
		db DBTX
	}
	type args struct {
		ctx context.Context
		arg ListAccountsParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Account
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queries{
				db: tt.fields.db,
			}
			got, err := q.ListAccounts(tt.args.ctx, tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListAccounts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListAccounts() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueries_ListEntries(t *testing.T) {
	type fields struct {
		db DBTX
	}
	type args struct {
		ctx context.Context
		arg ListEntriesParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Entry
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queries{
				db: tt.fields.db,
			}
			got, err := q.ListEntries(tt.args.ctx, tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListEntries() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListEntries() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueries_ListTransfers(t *testing.T) {
	type fields struct {
		db DBTX
	}
	type args struct {
		ctx context.Context
		arg ListTransfersParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Transfer
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queries{
				db: tt.fields.db,
			}
			got, err := q.ListTransfers(tt.args.ctx, tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListTransfers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListTransfers() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueries_UpdateAccount(t *testing.T) {
	type fields struct {
		db DBTX
	}
	type args struct {
		ctx context.Context
		arg UpdateAccountParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Account
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queries{
				db: tt.fields.db,
			}
			got, err := q.UpdateAccount(tt.args.ctx, tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateAccount() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueries_WithTx(t *testing.T) {
	type fields struct {
		db DBTX
	}
	type args struct {
		tx *sql.Tx
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Queries
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queries{
				db: tt.fields.db,
			}
			if got := q.WithTx(tt.args.tx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithTx() = %v, want %v", got, tt.want)
			}
		})
	}
}
