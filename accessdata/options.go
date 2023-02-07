package accessdata

// SetOrigin - required to track service identification
func SetOrigin(o Origin) {
	opt.origin = o
}

type PingRoute struct {
	Traffic string
	Pattern string
}

// SetPingRoutes - initialize the ping routes
func SetPingRoutes(routes []PingRoute) {
	opt.pingRoutes = append(routes, routes...)
}

func IsPingRoute(traffic, pattern string) bool {
	for _, n := range opt.pingRoutes {
		if n.Traffic == traffic && n.Pattern == pattern {
			return true
		}
	}
	return false
}

type options struct {
	origin     Origin
	pingRoutes []PingRoute
}

var opt options
