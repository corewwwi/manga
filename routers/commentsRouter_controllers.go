package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:GroupsScanlationController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:GroupsScanlationController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:GroupsScanlationController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:GroupsScanlationController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:GroupsScanlationController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:GroupsScanlationController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:name`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:GroupsScanlationController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:GroupsScanlationController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:name`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:GroupsScanlationController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:GroupsScanlationController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:name`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ImportController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ImportController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ImportController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ImportController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ImportController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ImportController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ImportController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ImportController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ImportController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ImportController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:LanguagesController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:LanguagesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:LanguagesController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:LanguagesController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:LanguagesController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:LanguagesController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:LoginController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:LoginController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:LogoutController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:LogoutController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ReaderController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ReaderController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:hash`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:RegisterController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:RegisterController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesChaptersController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesChaptersController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesChaptersController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesChaptersController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesChaptersController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesChaptersController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesChaptersController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesChaptersController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesChaptersController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesChaptersController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesStatusController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesStatusController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesStatusController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesStatusController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesStatusController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesStatusController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesStatusController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesStatusController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesStatusController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesStatusController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:Series_user_ratingsController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:Series_user_ratingsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:Series_user_ratingsController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:Series_user_ratingsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:StatusesController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:StatusesController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:TagsController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:TagsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:UsersController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:UsersController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:UsersController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:UsersController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
