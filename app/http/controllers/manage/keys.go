package manage

import (
	"github.com/goal-web/example/app/models"
)

func GetKeys() any {
	groups := models.Keys().Get().ToArray()
	groups = append(groups, models.Key{Id: 0, Name: "创建新的密钥"})
	return groups
}
