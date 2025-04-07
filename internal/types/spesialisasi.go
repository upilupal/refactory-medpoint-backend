package types

import (
	"github.com/sev-2/raiden"
)

type Spesialisasi struct {
	raiden.TypeBase
}

func (t *Spesialisasi) Name() string {
	return "_spesialisasi"
}

func (r *Spesialisasi) Format() string {
	return "spesialisasi[]"
}

func (r *Spesialisasi) Enums() []string {
	return []string{}
}

func (r *Spesialisasi) Comment() *string {
	return nil
}

