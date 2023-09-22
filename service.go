package main

import (
	"context"
	"net/http"
	"time"

	"connectrpc.com/connect"
	emailsv1 "github.com/bufbuild/gophercon-2023-workshop/gen/emails/v1"
	"github.com/bufbuild/gophercon-2023-workshop/gen/emails/v1/emailsv1connect"
)

func NewService(store *Storage) *EmailService {
	return &EmailService{
		store: store,
	}
}

type EmailService struct {
	store *Storage
}

func (e *EmailService) Run(ctx context.Context) error {
	mux := http.NewServeMux()
	mux.Handle(emailsv1connect.NewEmailServiceHandler(e))
	srv := &http.Server{
		Addr:    ServiceAddress,
		Handler: mux,
	}
	go func() {
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		_ = srv.Shutdown(shutdownCtx)
	}()
	return srv.ListenAndServe()
}

func (e *EmailService) GetEmail(ctx context.Context, req *connect.Request[emailsv1.GetEmailRequest]) (*connect.Response[emailsv1.GetEmailResponse], error) {
	record, err := e.store.GetEmail(req.Msg.GetUserId())
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&emailsv1.GetEmailResponse{
		UserEmail: record,
	}), nil
}

func (e *EmailService) UpdateEmail(ctx context.Context, req *connect.Request[emailsv1.UpdateEmailRequest]) (*connect.Response[emailsv1.UpdateEmailResponse], error) {
	if err := e.store.UpdateEmail(req.Msg.GetUserId(), req.Msg.GetNewAddress()); err != nil {
		return nil, err
	}
	return connect.NewResponse(&emailsv1.UpdateEmailResponse{}), nil
}
