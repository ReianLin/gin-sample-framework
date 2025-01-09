package permission

import (
	"github.com/gin-gonic/gin"
)

type ApiPermissions int

const (
	Create ApiPermissions = 1
	Delete ApiPermissions = 2
	Update ApiPermissions = 4
	Read   ApiPermissions = 8
)

const All ApiPermissions = 1<<31 - 1

var ApiPermissionsMap = map[ApiPermissions]string{
	Create: "Create",
	Delete: "Delete",
	Update: "Update",
	Read:   "Read",
	All:    "All",
}

type RoutePerm struct {
	Router      string         `json:"router"`
	Value       string         `json:"value"`
	Method      string         `json:"method"`
	Permissions ApiPermissions `json:"permissions"`
}

type RoutePermHandle struct {
	*RoutePerm
	handleFunc []gin.HandlerFunc `json:"handle_func"`
}

type controllerPermission struct {
	Name      string                  `json:"name"`
	Router    string                  `json:"router"`
	RoutePerm map[string][]*RoutePerm `json:"func_permissions"`
}

func NewPerm(route, method string, perm ApiPermissions, handle ...gin.HandlerFunc) RoutePermHandle {
	return RoutePermHandle{
		RoutePerm: &RoutePerm{Router: route,
			Method:      method,
			Permissions: perm,
			Value:       ApiPermissionsMap[perm]},
		handleFunc: handle}
}

func (c *controllerPermission) Append(routes gin.IRoutes, f ...RoutePermHandle) {
	for _, handle := range f {
		if _, ok := c.RoutePerm[handle.Router]; !ok {
			c.RoutePerm[handle.Router] = append(c.RoutePerm[handle.Router], handle.RoutePerm)
		} else {
			for _, value := range c.RoutePerm[handle.Router] {
				if value.Method != handle.Method {
					c.RoutePerm[handle.Router] = append(c.RoutePerm[handle.Router], handle.RoutePerm)
				}
			}
		}

		routes.Handle(handle.Method, handle.Router, handle.handleFunc...)
	}
}
