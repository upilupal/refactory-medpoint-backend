package roles

import (
	"github.com/sev-2/raiden"
)

type Dokter struct {
	raiden.RoleBase
}

func (r *Dokter) Name() string {
	return "dokter"
}

