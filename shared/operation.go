package shared

import (
	"time"

	"github.com/lxc/lxd/shared/api"
)

type Operation struct {
	Id         string              `json:"id"`
	Class      string              `json:"class"`
	CreatedAt  time.Time           `json:"created_at"`
	UpdatedAt  time.Time           `json:"updated_at"`
	Status     string              `json:"status"`
	StatusCode api.StatusCode      `json:"status_code"`
	Resources  map[string][]string `json:"resources"`
	Metadata   *Jmap               `json:"metadata"`
	MayCancel  bool                `json:"may_cancel"`
	Err        string              `json:"err"`
}
