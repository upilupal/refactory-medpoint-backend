package types

import (
	"github.com/sev-2/raiden"
)

type Reservasi struct {
	raiden.TypeBase
}

func (t *Reservasi) Name() string {
	return "_reservasi"
}

func (r *Reservasi) Format() string {
	return "reservasi[]"
}

func (r *Reservasi) Enums() []string {
	return []string{}
}

func (r *Reservasi) Comment() *string {
	return nil
}

