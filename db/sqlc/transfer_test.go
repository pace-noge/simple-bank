package db

import (
	"context"
	"testing"

	"github.com/pace-noge/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, account1, account2 Account) Transfer {

	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.Equal(t, transfer.FromAccountID, account1.ID)
	require.Equal(t, transfer.ToAccountID, account2.ID)
	require.Equal(t, transfer.Amount, arg.Amount)

	return transfer
}

func createTransferAccount(t *testing.T) (Account, Account) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	return account1, account2
}

func TestCreateTransfer(t *testing.T) {
	account1, account2 := createTransferAccount(t)
	createRandomTransfer(t, account1, account2)
}

func TestGetTransfer(t *testing.T) {
	account1, account2 := createTransferAccount(t)
	transfer1 := createRandomTransfer(t, account1, account2)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)
	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.Equal(t, transfer2.FromAccountID, account1.ID)
	require.Equal(t, transfer2.ToAccountID, account2.ID)

}

func TestListTransfers(t *testing.T) {
	fromAccount, toAccount := createTransferAccount(t)
	for i := 0; i < 5; i++ {
		createRandomTransfer(t, fromAccount, toAccount)
		createRandomTransfer(t, fromAccount, toAccount)
	}

	arg := ListTransfersParams{
		Limit:         5,
		Offset:        5,
		FromAccountID: fromAccount.ID,
		ToAccountID:   toAccount.ID,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.True(t, transfer.FromAccountID == fromAccount.ID || transfer.ToAccountID == toAccount.ID)
	}

}
