package routes

func Routes() map[string]string {
	routes := make(map[string]string)

	routes["/"] = "/"
	routes["/contact"] = "/contact"

	return routes
}
