package types

import (
	"github.com/sev-2/raiden"
)

type MetodePembayaran struct {
	raiden.TypeBase
}

func (t *MetodePembayaran) Name() string {
	return "metode_pembayaran"
}

func (r *MetodePembayaran) Format() string {
	return "metode_pembayaran"
}

func (r *MetodePembayaran) Enums() []string {
	return []string{"tunai","transfer","e_wallet"}
}

func (r *MetodePembayaran) Comment() *string {
	return nil
}

