package model

import "gin-sample-framework/pkg/permission"

var MenuResourcesList = make([]MenuResources, 0)

type ResourcesModel struct {
	Name      string                    `json:"name"`
	Operation permission.ApiPermissions `json:"operation"`
}

type MenuResources struct {
	ParentName   string           `json:"parent_name"`
	ParentRouter string           `json:"parent_router"`
	Name         string           `json:"name"`
	Router       string           `json:"router"`
	Resources    []ResourcesModel `json:"resources"`
}

// AppendMenuResourcesList 添加菜单
func AppendMenuResourcesList(ParentMenuName, ParentMenuRouter, menuName, router string, apiPermissions ...permission.ApiPermissions) {
	MenuResourcesList = append(MenuResourcesList, MenuResources{
		ParentName:   ParentMenuName,
		ParentRouter: ParentMenuRouter,
		Name:         menuName,
		Router:       router,
		Resources: func() []ResourcesModel {
			resourcesModels := make([]ResourcesModel, 0)
			for _, per := range apiPermissions {
				resourcesModels = append(resourcesModels, ResourcesModel{
					Name:      permission.ApiPermissionsMap[per],
					Operation: per,
				})
			}
			if len(apiPermissions) > 0 {
				resourcesModels = append(resourcesModels, ResourcesModel{
					Name:      permission.ApiPermissionsMap[permission.All],
					Operation: permission.All,
				})
			}
			return resourcesModels
		}(),
	})
}
