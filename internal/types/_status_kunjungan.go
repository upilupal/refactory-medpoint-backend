package types

import (
	"github.com/sev-2/raiden"
)

type StatusKunjungan struct {
	raiden.TypeBase
}

func (t *StatusKunjungan) Name() string {
	return "_status_kunjungan"
}

func (r *StatusKunjungan) Format() string {
	return "status_kunjungan[]"
}

func (r *StatusKunjungan) Enums() []string {
	return []string{}
}

func (r *StatusKunjungan) Comment() *string {
	return nil
}

