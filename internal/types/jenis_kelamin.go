package types

import (
	"github.com/sev-2/raiden"
)

type JenisKelamin struct {
	raiden.TypeBase
}

func (t *JenisKelamin) Name() string {
	return "jenis_kelamin"
}

func (r *JenisKelamin) Format() string {
	return "jenis_kelamin"
}

func (r *JenisKelamin) Enums() []string {
	return []string{"laki_laki","perempuan"}
}

func (r *JenisKelamin) Comment() *string {
	return nil
}

