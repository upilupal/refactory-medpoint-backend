package types

import (
	"github.com/sev-2/raiden"
)

type StatusReservasi struct {
	raiden.TypeBase
}

func (t *StatusReservasi) Name() string {
	return "_status_reservasi"
}

func (r *StatusReservasi) Format() string {
	return "status_reservasi[]"
}

func (r *StatusReservasi) Enums() []string {
	return []string{}
}

func (r *StatusReservasi) Comment() *string {
	return nil
}

