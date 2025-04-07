package models

import (
	"github.com/sev-2/raiden/pkg/db"
)

type Spesialisasi struct {
	db.ModelBase
	Id        int64   `json:"id,omitempty" column:"name:id;type:bigint;primaryKey;autoIncrement;nullable:false"`
	Nama      *string `json:"nama,omitempty" column:"name:nama;type:varchar;nullable"`
	Deskripsi *string `json:"deskripsi,omitempty" column:"name:deskripsi;type:text;nullable"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"spesialisasi" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	DokterSpesialisasis []*Dokter `json:"dokter_spesialisasis,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:spesialisasi_id"`
}
