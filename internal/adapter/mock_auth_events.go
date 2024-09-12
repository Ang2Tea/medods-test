package adapter

import (
	"context"
	"log"

	"Ang2Tea/medods-test/internal/usecase"
)

var _ usecase.IAuthEvent = (*mockAuthEvent)(nil)

type mockAuthEvent struct{}

func NewMockAuthEvent() *mockAuthEvent {
	return &mockAuthEvent{}
}

func (m *mockAuthEvent) IPAddressChanged(ctx context.Context, oldIPAddress string, newIPAddress string) {
	log.Println("INFO", "IPAddressChanged", oldIPAddress, newIPAddress)
}
