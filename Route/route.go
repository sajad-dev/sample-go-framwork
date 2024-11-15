package route

import "github.com/sajad-dev/go-framwork/Route/api"

var RouteList = []api.ApiType{}

func Route() {
	api.RouteList = RouteList
}
