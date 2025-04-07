package types

import (
	"github.com/sev-2/raiden"
)

type Dokter struct {
	raiden.TypeBase
}

func (t *Dokter) Name() string {
	return "_dokter"
}

func (r *Dokter) Format() string {
	return "dokter[]"
}

func (r *Dokter) Enums() []string {
	return []string{}
}

func (r *Dokter) Comment() *string {
	return nil
}

