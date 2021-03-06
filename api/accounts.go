package api

import (
	"context"
	"encoding/json"
	"net/http"

	"payment-service/models"
	"payment-service/services"

	"github.com/go-chi/chi"
	"github.com/go-kit/kit/endpoint"
	uuid "github.com/satori/go.uuid"

	httptransport "github.com/go-kit/kit/transport/http"
)

func accountsEndpoints(r *chi.Mux, as services.Accounts) {
	listAccountsHandler := httptransport.NewServer(
		accountsEndpoint(as),
		decodeOffsetLimitReq,
		encodeResponseOK,
	)
	r.Method(http.MethodGet, "/accounts", listAccountsHandler)
	accountByIDHandler := httptransport.NewServer(
		accountByIDEndpoint(as),
		decodeAccountByIDRequest,
		encodeResponseOK,
	)
	r.Method(http.MethodGet, "/accounts/{accountID}", accountByIDHandler)
	addAccountHandler := httptransport.NewServer(
		addAccountEndpoint(as),
		decodeAddAccountReq,
		encodeResponseCreated,
	)
	r.Method(http.MethodPost, "/accounts", addAccountHandler)
}

func accountByIDEndpoint(svc services.Accounts) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(uuid.UUID)
		account, err := svc.AccountByID(ctx, req)
		return errResp(account, err)
	}
}

func decodeAccountByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	id, err := uuid.FromString(chi.URLParam(r, "accountID"))
	if err != nil {
		return nil, err
	}
	return id, nil
}

func addAccountEndpoint(svc services.Accounts) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(services.NewAccount)
		id, err := svc.AddAccount(ctx, req)
		return errResp(IDResp{ID: id}, err)
	}
}

func decodeAddAccountReq(_ context.Context, r *http.Request) (interface{}, error) {
	defer r.Body.Close()
	var acc services.NewAccount
	err := json.NewDecoder(r.Body).Decode(&acc)
	return acc, err
}

func accountsEndpoint(svc services.Accounts) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(models.OffsetLimit)
		accs, err := svc.ListOfAccounts(ctx, req)
		return errResp(accs, err)
	}
}
