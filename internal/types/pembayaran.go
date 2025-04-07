package types

import (
	"github.com/sev-2/raiden"
)

type Pembayaran struct {
	raiden.TypeBase
}

func (t *Pembayaran) Name() string {
	return "_pembayaran"
}

func (r *Pembayaran) Format() string {
	return "pembayaran[]"
}

func (r *Pembayaran) Enums() []string {
	return []string{}
}

func (r *Pembayaran) Comment() *string {
	return nil
}

