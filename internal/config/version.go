package config

import "github.com/lancer-kit/uwe/v2"

// The variables are set when compiling the binary and used to make sure which version of the backend is running.
// Example: go build -ldflags "-X ports_api/internal/config.version=$VERSION\
// -X ports_api/internal/config.build=$COMMIT \
// -X ports_api/internal/config.tag=$TAG" .

// nolint:gochecknoglobals
var (
	version     = "n/a"
	build       = "n/a"
	tag         = "n/a"
	ServiceName = "ports_client_api"
)

func AppInfo() uwe.AppInfo {
	return uwe.AppInfo{
		Name:    ServiceName,
		Version: version,
		Build:   build,
		Tag:     tag,
	}
}
