package db

import (
	"context"
	"testing"

	"github.com/BinayRajbanshi/GoBasicBank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T, account Account) Entry {
	entryParams := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}
	entry, err := testStore.CreateEntry(context.Background(), entryParams)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, entry.Amount, entryParams.Amount)
	require.Equal(t, entry.AccountID, entryParams.AccountID)

	require.NotZero(t, entry.CreatedAt)
	require.NotZero(t, entry.AccountID)

	return entry
}

func TestCreateEntry(t *testing.T) {
	account := createRandomAccount(t) // making this independently so that there won't be chain of dependency inside one function
	createRandomEntry(t, account)
}

func TestListEntry(t *testing.T) {
	account := createRandomAccount(t)
	for i := 1; i <= 10; i++ {
		createRandomEntry(t, account)
	}

	// limit 5 and offset 5 means, skip 5 rows and give next 5 rows
	arg := ListEntriesParams{
		AccountID: account.ID,
		Offset:    5,
		Limit:     5,
	}

	// as we create 10 random accounts initially, it is obvious that we will get 5 accounts as a response with offset 5 and limit 5
	entries, err := testStore.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entries := range entries {
		require.NotEmpty(t, entries)
	}
}
