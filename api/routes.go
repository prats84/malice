package api

import "net/http"

// Route api route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes route array
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"APIFileScan",
		"POST",
		"/file/scan",
		APIFileScan,
	},
	Route{
		"APIIntelLookup",
		"POST",
		"/intel/lookup/{hash}",
		APIIntelLookup,
	},
	Route{
		"APIPluginList",
		"POST",
		"/plugin/list",
		APIPluginList,
	},
	Route{
		"APIPluginInstall",
		"POST",
		"/plugin/install",
		APIPluginInstall,
	},
	Route{
		"APIPluginRemove",
		"POST",
		"/plugin/remove",
		APIPluginRemove,
	},
	Route{
		"APIPluginUpdate",
		"POST",
		"/plugin/update",
		APIPluginUpdate,
	},
}
