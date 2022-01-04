package domain

import "strings"

type RolePermissions struct {
	rolePermissions map[string][]string
}

func GetRolePermissions() RolePermissions {
	return RolePermissions{map[string][]string{
		"admin": {"GetAllCustomers", "GetCustomer", "NewAccount", "NewTransaction"},
		"user":  {"GetCustomer", "NewTransaction"},
	}}
}

func (rolePermission RolePermissions) IsAuthorizedFor(role string, routeName string) bool {
	perms := rolePermission.rolePermissions[role]

	for _, role := range perms {
		if role == strings.TrimSpace(routeName) {
			return true
		}
	}

	return false
}
