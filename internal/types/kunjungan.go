package types

import (
	"github.com/sev-2/raiden"
)

type Kunjungan struct {
	raiden.TypeBase
}

func (t *Kunjungan) Name() string {
	return "_kunjungan"
}

func (r *Kunjungan) Format() string {
	return "kunjungan[]"
}

func (r *Kunjungan) Enums() []string {
	return []string{}
}

func (r *Kunjungan) Comment() *string {
	return nil
}

