package models

import (
	"github.com/sev-2/raiden/pkg/db"
)

type Dokter struct {
	db.ModelBase
	Id             int64   `json:"id,omitempty" column:"name:id;type:bigint;primaryKey;autoIncrement;nullable:false"`
	SpesialisasiId int64   `json:"spesialisasi_id,omitempty" column:"name:spesialisasi_id;type:bigint;nullable:false"`
	Nama           *string `json:"nama,omitempty" column:"name:nama;type:varchar;nullable"`
	NoTelepon      *string `json:"no_telepon,omitempty" column:"name:no_telepon;type:varchar;nullable;unique"`
	Biografi       *string `json:"biografi,omitempty" column:"name:biografi;type:text;nullable"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"dokter" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	JadwalDokterDokters []*JadwalDokter `json:"jadwal_dokter_dokters,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:dokter_id"`
	Spesialisasi        *Spesialisasi   `json:"spesialisasi,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:spesialisasi_id"`
}
