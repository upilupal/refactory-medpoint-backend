package types

import (
	"github.com/sev-2/raiden"
)

type JadwalDokter struct {
	raiden.TypeBase
}

func (t *JadwalDokter) Name() string {
	return "_jadwal_dokter"
}

func (r *JadwalDokter) Format() string {
	return "jadwal_dokter[]"
}

func (r *JadwalDokter) Enums() []string {
	return []string{}
}

func (r *JadwalDokter) Comment() *string {
	return nil
}

