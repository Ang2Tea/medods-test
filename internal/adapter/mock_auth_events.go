package adapter

import (
	"context"
	"log"
	"time"

	"Ang2Tea/medods-test/internal/usecase"
)

var _ usecase.IAuthEvent = (*mockAuthEvent)(nil)

type mockAuthEvent struct{}

func (m *mockAuthEvent) IPAddressChanged(ctx context.Context, oldIPAddress string, newIPAddress string) {
	log.Println("INFO", time.Now(), "IPAddressChanged", oldIPAddress, newIPAddress)
}
