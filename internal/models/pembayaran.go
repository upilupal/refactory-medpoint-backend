package models

import (
	"github.com/sev-2/raiden/pkg/db"
	"github.com/sev-2/raiden/pkg/postgres"
	"medpoinraiden/internal/types"
)

type Pembayaran struct {
	db.ModelBase
	Id               int64                  `json:"id,omitempty" column:"name:id;type:bigint;primaryKey;autoIncrement;nullable:false"`
	CreatedAt        postgres.DateTime      `json:"created_at,omitempty" column:"name:created_at;type:timestampz;nullable:false;default:now()"`
	ReservasiId      int64                  `json:"reservasi_id,omitempty" column:"name:reservasi_id;type:bigint;nullable:false"`
	JumlahBayar      *float64               `json:"jumlah_bayar,omitempty" column:"name:jumlah_bayar;type:numeric;nullable"`
	StatusPembayaran types.StatusPembayaran `json:"status_pembayaran,omitempty" column:"name:status_pembayaran;type:status_pembayaran;nullable:false"`
	UpdatedAt        *postgres.DateTime     `json:"updated_at,omitempty" column:"name:updated_at;type:timestamp;nullable"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"pembayaran" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	Reservasi *Reservasi `json:"reservasi,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:reservasi_id"`
}
