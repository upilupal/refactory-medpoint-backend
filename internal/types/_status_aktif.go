package types

import (
	"github.com/sev-2/raiden"
)

type StatusAktif struct {
	raiden.TypeBase
}

func (t *StatusAktif) Name() string {
	return "_status_aktif"
}

func (r *StatusAktif) Format() string {
	return "status_aktif[]"
}

func (r *StatusAktif) Enums() []string {
	return []string{}
}

func (r *StatusAktif) Comment() *string {
	return nil
}

