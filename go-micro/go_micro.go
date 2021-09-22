package go_micro

// go-micro v2版本和v1不太一样，v2移除了对consul的支持。https://learnku.com/docs/go-micro/2.x/getting-started.html/8460
import (
	"github.com/micro/go-micro/v2"
)

var (
	service micro.Service
)

func init() {
	service = micro.NewService()
	service.Init()
}
