package types

import (
	"github.com/sev-2/raiden"
)

type StatusPembayaran struct {
	raiden.TypeBase
}

func (t *StatusPembayaran) Name() string {
	return "status_pembayaran"
}

func (r *StatusPembayaran) Format() string {
	return "status_pembayaran"
}

func (r *StatusPembayaran) Enums() []string {
	return []string{"belum","menunggu","selesai"}
}

func (r *StatusPembayaran) Comment() *string {
	return nil
}

