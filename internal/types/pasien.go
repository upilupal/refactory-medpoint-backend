package types

import (
	"github.com/sev-2/raiden"
)

type Pasien struct {
	raiden.TypeBase
}

func (t *Pasien) Name() string {
	return "_pasien"
}

func (r *Pasien) Format() string {
	return "pasien[]"
}

func (r *Pasien) Enums() []string {
	return []string{}
}

func (r *Pasien) Comment() *string {
	return nil
}

