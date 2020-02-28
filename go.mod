module gitlab.com/krayohu/logtest

go 1.13

require (
	github.com/golang/protobuf v1.3.3
	github.com/micro/go-micro/v2 v2.1.2
	github.com/micro/go-plugins/logger/zerolog/v2 v2.0.3
)

replace github.com/micro/go-plugins/logger/zerolog/v2 => ../../../github.com/micro/go-plugins/logger/zerolog
