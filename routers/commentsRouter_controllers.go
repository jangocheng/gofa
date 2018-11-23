package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["gofa/controllers:GroupController"] = append(beego.GlobalControllerRouter["gofa/controllers:GroupController"],
        beego.ControllerComments{
            Method: "AddFace",
            Router: `/addface`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gofa/controllers:GroupController"] = append(beego.GlobalControllerRouter["gofa/controllers:GroupController"],
        beego.ControllerComments{
            Method: "SearchFace",
            Router: `/searchface`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gofa/controllers:GroupController"] = append(beego.GlobalControllerRouter["gofa/controllers:GroupController"],
        beego.ControllerComments{
            Method: "UserList",
            Router: `/userlist`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
