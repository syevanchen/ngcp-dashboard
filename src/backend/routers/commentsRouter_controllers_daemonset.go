package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/Qihoo360/wayne/src/backend/controllers/daemonset:DaemonSetController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/wayne/src/backend/controllers/daemonset:DaemonSetController"],
		beego.ControllerComments{
			Method:           "List",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/wayne/src/backend/controllers/daemonset:DaemonSetController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/wayne/src/backend/controllers/daemonset:DaemonSetController"],
		beego.ControllerComments{
			Method:           "Create",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/wayne/src/backend/controllers/daemonset:DaemonSetController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/wayne/src/backend/controllers/daemonset:DaemonSetController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id([0-9]+)`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/wayne/src/backend/controllers/daemonset:DaemonSetController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/wayne/src/backend/controllers/daemonset:DaemonSetController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/:id([0-9]+)`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/wayne/src/backend/controllers/daemonset:DaemonSetController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/wayne/src/backend/controllers/daemonset:DaemonSetController"],
		beego.ControllerComments{
			Method:           "Update",
			Router:           `/:id([0-9]+)`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/wayne/src/backend/controllers/daemonset:DaemonSetController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/wayne/src/backend/controllers/daemonset:DaemonSetController"],
		beego.ControllerComments{
			Method:           "GetNames",
			Router:           `/names`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/wayne/src/backend/controllers/daemonset:DaemonSetController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/wayne/src/backend/controllers/daemonset:DaemonSetController"],
		beego.ControllerComments{
			Method:           "UpdateOrders",
			Router:           `/updateorders`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/wayne/src/backend/controllers/daemonset:DaemonSetTplController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/wayne/src/backend/controllers/daemonset:DaemonSetTplController"],
		beego.ControllerComments{
			Method:           "List",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/wayne/src/backend/controllers/daemonset:DaemonSetTplController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/wayne/src/backend/controllers/daemonset:DaemonSetTplController"],
		beego.ControllerComments{
			Method:           "Create",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/wayne/src/backend/controllers/daemonset:DaemonSetTplController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/wayne/src/backend/controllers/daemonset:DaemonSetTplController"],
		beego.ControllerComments{
			Method:           "Update",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/wayne/src/backend/controllers/daemonset:DaemonSetTplController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/wayne/src/backend/controllers/daemonset:DaemonSetTplController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/wayne/src/backend/controllers/daemonset:DaemonSetTplController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/wayne/src/backend/controllers/daemonset:DaemonSetTplController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/:id([0-9]+)`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

}
