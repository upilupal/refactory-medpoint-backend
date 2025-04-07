package models

import (
	"github.com/sev-2/raiden/pkg/db"
	"github.com/sev-2/raiden/pkg/postgres"
	"medpoinraiden/internal/types"
)

type Reservasi struct {
	db.ModelBase
	Id               int64                  `json:"id,omitempty" column:"name:id;type:bigint;primaryKey;autoIncrement;nullable:false"`
	PasienId         *int64                 `json:"pasien_id,omitempty" column:"name:pasien_id;type:bigint;nullable"`
	JadwalDokterId   *int64                 `json:"jadwal_dokter_id,omitempty" column:"name:jadwal_dokter_id;type:bigint;nullable"`
	TanggalReservasi *postgres.DateTime     `json:"tanggal_reservasi,omitempty" column:"name:tanggal_reservasi;type:date;nullable"`
	WaktuReservasi   *postgres.DateTime     `json:"waktu_reservasi,omitempty" column:"name:waktu_reservasi;type:time;nullable"`
	Antrian          *string                `json:"antrian,omitempty" column:"name:antrian;type:varchar;nullable"`
	StatusReservasi  types.StatusReservasi  `json:"status_reservasi,omitempty" column:"name:status_reservasi;type:status_reservasi;nullable"`
	MetodePembayaran types.MetodePembayaran `json:"metode_pembayaran,omitempty" column:"name:metode_pembayaran;type:metode_pembayaran;nullable"`
	CreatedAt        postgres.DateTime      `json:"created_at,omitempty" column:"name:created_at;type:timestampz;nullable:false;default:now()"`
	UpdatedAt        *postgres.DateTime     `json:"updated_at,omitempty" column:"name:updated_at;type:timestamp;nullable"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"reservasi" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	JadwalDokterJadwal   *JadwalDokter `json:"jadwal_dokter_jadwal,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:jadwal_dokter_id"`
	KunjunganReservasis  []*Kunjungan  `json:"kunjungan_reservasis,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:reservasi_id"`
	Pasien               *Pasien       `json:"pasien,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:pasien_id"`
	PembayaranReservasis []*Pembayaran `json:"pembayaran_reservasis,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:reservasi_id"`
}
