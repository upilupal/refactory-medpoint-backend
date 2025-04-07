package types

import (
	"github.com/sev-2/raiden"
)

type StatusAktif struct {
	raiden.TypeBase
}

func (t *StatusAktif) Name() string {
	return "status_aktif"
}

func (r *StatusAktif) Format() string {
	return "status_aktif"
}

func (r *StatusAktif) Enums() []string {
	return []string{"masuk","libur"}
}

func (r *StatusAktif) Comment() *string {
	return nil
}

