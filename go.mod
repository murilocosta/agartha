module github.com/murilocosta/agartha

go 1.16

// +heroku goVersion 1.16
// +heroku install ./cmd/...

require (
	github.com/appleboy/gin-jwt/v2 v2.7.0
	github.com/fastly/go-utils v0.0.0-20180712184237-d95a45783239 // indirect
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.7.7
	github.com/go-playground/validator/v10 v10.9.0
	github.com/go-redis/redis/v8 v8.11.4
	github.com/jehiah/go-strftime v0.0.0-20171201141054-1d33003b3869 // indirect
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/lestrrat/go-envload v0.0.0-20180220120943-6ed08b54a570 // indirect
	github.com/lestrrat/go-file-rotatelogs v0.0.0-20180223000712-d3151e2a480f
	github.com/lestrrat/go-strftime v0.0.0-20180220042222-ba3bf9c1d042 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/sirupsen/logrus v1.8.1
	github.com/tebeka/strftime v0.1.5 // indirect
	go.etcd.io/etcd/client/v3 v3.5.1
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	go.uber.org/zap v1.19.1 // indirect
	golang.org/x/crypto v0.0.0-20211202192323-5770296d904e
	golang.org/x/sys v0.0.0-20210917161153-d61c044b1678 // indirect
	google.golang.org/genproto v0.0.0-20210917145530-b395a37504d4 // indirect
	gopkg.in/yaml.v2 v2.4.0
	gorm.io/driver/postgres v1.2.3
	gorm.io/gorm v1.22.4
)
