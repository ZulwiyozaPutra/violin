package route

import (
	"app/shared/router"
	"app/controller"
)

// ROUTES consists of slices of Routes
var Routes = []router.Route{
	router.Route{Path: "/", Handler: controller.Index},
	router.Route{Path: "/select", Handler: controller.UserSelection},
	router.Route{Path: "/selected", Handler: controller.UserSelected},
}