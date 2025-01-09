package middleware

import (
	"gin-sample-framework/errors"
	"gin-sample-framework/pkg/permission"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func HasPermission() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			fuallPath = c.FullPath()
			prefix    = "/api/v1"
		)
		if fuallPath == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, H(errors.BadRequest, nil, errors.BadRequest.String()))
			return
		}
		result := strings.TrimPrefix(fuallPath, prefix)
		router := extractBeforeSecondSlash(result)
		if len(router) == 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, H(errors.BadRequest, nil, errors.BadRequest.String()))
			return
		}

		if per := permission.Permission.GetGroup(router); per != nil {
			child, ok := per.RoutePerm[strings.TrimPrefix(result, result)]
			if !ok {
				c.AbortWithStatusJSON(http.StatusOK, H(errors.NotPermission, nil, errors.NotPermission.String()))
				return
			}
			var perm permission.ApiPermissions
			for _, node := range child {
				if node.Method != c.Request.Method {
					continue
				}
				perm = node.Permissions
				if perm == 0 {
					c.AbortWithStatusJSON(http.StatusOK, H(errors.NotPermission, nil, errors.NotPermission.String()))
					return
				}
				var storePerm = permission.All
				if (perm & storePerm) != perm {
					c.AbortWithStatusJSON(http.StatusOK, H(errors.NotPermission, nil, errors.NotPermission.String()))
					return
				}
			}
		}
		c.Next()

	}
}

func extractBeforeSecondSlash(input string) string {
	firstSlashIndex := strings.Index(input, "/")

	if firstSlashIndex != -1 {
		secondSlashIndex := strings.Index(input[firstSlashIndex+1:], "/")

		if secondSlashIndex != -1 {
			return input[:firstSlashIndex+secondSlashIndex+1]
		}
	}

	return input
}
