package db

import (
	"context"
	"testing"
	"time"

	"github.com/BinayRajbanshi/GoBasicBank/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, newAccount1 Account, newAccount2 Account) Transfer {
	transferParams := CreateTransferParams{
		FromAccountID: newAccount1.ID,
		ToAccountID:   newAccount2.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testStore.CreateTransfer(context.Background(), transferParams)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, transferParams.FromAccountID, transfer.FromAccountID)
	require.Equal(t, transferParams.ToAccountID, transfer.ToAccountID)
	require.Equal(t, transferParams.Amount, transfer.Amount)

	require.NotZero(t, transfer.CreatedAt)
	require.NotZero(t, transfer.Amount)
	return transfer
}

func TestCreateTransfer(t *testing.T) {
	newAccount1 := createRandomAccount(t)
	newAccount2 := createRandomAccount(t)

	createRandomTransfer(t, newAccount1, newAccount2)
}

func TestListTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	for i := 1; i <= 10; i++ {
		createRandomTransfer(t, account1, account2)
		createRandomTransfer(t, account2, account1)
	}

	arg := ListTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.Balance,
		Limit:         5,
		Offset:        5,
	}

	transfers, err := testStore.ListTransfers(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}

func TestGetTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	transfer1 := createRandomTransfer(t, account1, account2)

	transfer2, err := testStore.GetTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}
