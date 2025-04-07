package models

import (
	"github.com/sev-2/raiden/pkg/db"
	"github.com/sev-2/raiden/pkg/postgres"
	"medpoinraiden/internal/types"
)

type JadwalDokter struct {
	db.ModelBase
	Id          int64              `json:"id,omitempty" column:"name:id;type:bigint;primaryKey;autoIncrement;nullable:false"`
	DokterId    *int64             `json:"dokter_id,omitempty" column:"name:dokter_id;type:bigint;nullable"`
	Hari        *string            `json:"hari,omitempty" column:"name:hari;type:varchar;nullable"`
	JamMulai    *postgres.DateTime `json:"jam_mulai,omitempty" column:"name:jam_mulai;type:time;nullable"`
	JamSelesai  *postgres.DateTime `json:"jam_selesai,omitempty" column:"name:jam_selesai;type:time;nullable"`
	StatusAktif types.StatusAktif  `json:"status_aktif,omitempty" column:"name:status_aktif;type:status_aktif;nullable"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"jadwal_dokter" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	Dokter                              *Dokter      `json:"dokter,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:dokter_id"`
	PasiensThroughReservasiJadwalDokter []*Pasien    `json:"pasiens_through_reservasi_jadwal_dokter,omitempty" join:"joinType:manyToMany;through:reservasi;sourcePrimaryKey:id;sourceForeignKey:jadwal_dokter_id;targetPrimaryKey:id;targetForeign:jadwal_dokter_id"`
	ReservasiJadwals                    []*Reservasi `json:"reservasi_jadwals,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:jadwal_dokter_id"`
}
