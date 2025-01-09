package permission

import "strings"

var Permission = &permission{perm: make(map[string]*controllerPermission)}

type permission struct {
	perm map[string]*controllerPermission
}

func (p *permission) MakeGroup(route, name string) *controllerPermission {
	route = "/" + strings.ReplaceAll(route, "/", "/")
	if c, ok := Permission.perm[route]; ok {
		return c
	}
	c := &controllerPermission{
		Name:      name,
		Router:    route,
		RoutePerm: make(map[string][]*RoutePerm),
	}
	Permission.perm[route] = c
	return c
}

func (p *permission) GetGroup(route string) *controllerPermission {
	return p.perm[route]
}
