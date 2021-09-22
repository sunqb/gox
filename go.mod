module gox

go 1.16

require (
	github.com/axgle/mahonia v0.0.0-20180208002826-3358181d7394
	github.com/hashicorp/go-uuid v1.0.2 // indirect
	github.com/jponge/vertx-go-tcp-eventbus-bridge v0.0.0-20180315155916-c9060137d6c9
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/consul v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/micro v1.18.0 // indirect
	github.com/micro/micro/v2 v2.9.3 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible
	go.etcd.io/etcd/client/v3 v3.5.0
	golang.org/x/sys v0.0.0-20210603081109-ebe580a85c40
	golang.org/x/text v0.3.5
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
