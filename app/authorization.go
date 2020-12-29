package app

import (
	"github.com/thetogi/YReserve2/model"
	"github.com/thetogi/YReserve2/util"
)

func (a *App) UserHasPermissionTo(permissionId string) bool {
	userRoles := a.UserSession.Roles
	roleNames := util.StringToStringArray(userRoles)
	for _, roleName := range roleNames {
		role := model.GetRole(roleName)

		for _, permission := range role.Permissions {
			if permissionId == permission {
				return true
			}
		}
	}
	return false
}
