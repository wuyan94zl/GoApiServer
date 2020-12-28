package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wuyan94zl/api/app/controllers/admin"
	"github.com/wuyan94zl/api/pkg/utils"
)

// 注册路由列表
func ApiRouter(router *gin.RouterGroup) {
	utils.AddRoute(router, "POST", "/admin/login", admin.Login)

	//// start admin
	//utils.AddRoute(router, "POST", "/admin/create", admin.Create)
	//utils.AddRoute(router, "POST", "/admin/update", admin.Update)
	//utils.AddRoute(router, "GET", "/admin/delete", admin.Delete)
	//utils.AddRoute(router, "GET", "/admin/info", admin.Info)
	//utils.AddRoute(router, "POST", "/admin/paginate", admin.Paginate)
	//// end admin
}
