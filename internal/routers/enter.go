package routers

import "go-ecomerce-backend-api/internal/routers/manager"
import "go-ecomerce-backend-api/internal/routers/user"

type RouterGroup struct {
	User    user.UserRouterGroup
	Manager manager.ManagerRouterGroup
}

var RouterGroupApp = new(RouterGroup)
