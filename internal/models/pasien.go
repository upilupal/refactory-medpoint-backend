package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
	"github.com/sev-2/raiden/pkg/postgres"
	"medpoinraiden/internal/types"
)

type Pasien struct {
	db.ModelBase
	Id            int64              `json:"id,omitempty" column:"name:id;type:bigint;primaryKey;autoIncrement;nullable:false"`
	Nama          *string            `json:"nama,omitempty" column:"name:nama;type:varchar;nullable"`
	TanggalLahir  *postgres.DateTime `json:"tanggal_lahir,omitempty" column:"name:tanggal_lahir;type:date;nullable"`
	JenisKelamin  types.JenisKelamin `json:"jenis_kelamin,omitempty" column:"name:jenis_kelamin;type:jenis_kelamin;nullable"`
	Alamat        *string            `json:"alamat,omitempty" column:"name:alamat;type:varchar;nullable"`
	NoTelepon     *string            `json:"no_telepon,omitempty" column:"name:no_telepon;type:varchar;nullable;unique"`
	TanggalDaftar *postgres.DateTime `json:"tanggal_daftar,omitempty" column:"name:tanggal_daftar;type:date;nullable"`
	CreatedBy     *uuid.UUID         `json:"created_by,omitempty" column:"name:created_by;type:uuid;nullable"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"pasien" rlsEnable:"true" rlsForced:"true"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	JadwalDoktersThroughReservasiPasien []*JadwalDokter `json:"jadwal_dokters_through_reservasi_pasien,omitempty" join:"joinType:manyToMany;through:reservasi;sourcePrimaryKey:id;sourceForeignKey:pasien_id;targetPrimaryKey:id;targetForeign:pasien_id"`
	ReservasiPasiens                    []*Reservasi    `json:"reservasi_pasiens,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:pasien_id"`
}
