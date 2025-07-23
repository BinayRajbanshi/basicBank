package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/BinayRajbanshi/GoBasicBank/db/mock"
	db "github.com/BinayRajbanshi/GoBasicBank/db/sqlc"
	"github.com/BinayRajbanshi/GoBasicBank/util"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestGetAccountAPI(t *testing.T) {
	account := randomAccount()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)

	// build stubs (Stubs return fixed values when called, regardless of the input.)
	store.EXPECT().
		GetAccount(gomock.Any(), gomock.Eq(account.ID)).
		Times(1).Return(account, nil)

	// in the above line of code
	//     EXPECT() is a recorder: it sets up expected method calls on the mock.
	// Think of it as saying:
	// “I expect the following method to be called with these parameters, exactly this many times, and here's what it should return.”

	// start test server and send request
	server := newTestServer(t, store)

	recorder := httptest.NewRecorder()
	url := fmt.Sprintf("/api/v1/accounts/%d", account.ID)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	server.router.ServeHTTP(recorder, request)
	// check the response
	require.Equal(t, http.StatusOK, recorder.Code)
	requireBodyMatchAccount(t, recorder.Body, account)
}

func randomAccount() db.Account {
	return db.Account{
		ID:       util.RandInt(1, 1000),
		Owner:    util.RandOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}

func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, account db.Account) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotAccounts []db.Account
	err = json.Unmarshal(data, &gotAccounts)
	require.NoError(t, err)
	require.Equal(t, account, gotAccounts)
}
