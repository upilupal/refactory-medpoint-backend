package roles

import (
	"github.com/sev-2/raiden"
)

type User1 struct {
	raiden.RoleBase
}

func (r *User1) Name() string {
	return "user1"
}

func (r *User1) CanLogin() bool {
	return true
}

