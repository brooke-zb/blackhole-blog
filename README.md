# BlackHole Blog

A simple, security blog system based on [Gin](https://github.com/gin-gonic/gin)

⚠ Still under development ⚠

## TODO

- [x] custom setting with config file based on [viper](https://github.com/spf13/viper)
- [x] RESTful api
- [x] cache service data with [CCache v3](https://github.com/karlseguin/ccache)
- [x] cache article read count with Redis
- [x] jwt authentication
- [x] regenerate jwt while about to expire
- [x] RBAC authorization
- [x] csrf protection
- [x] sensitive words filter base on [sensitive](https://github.com/importcjj/sensitive)
- [ ] mail notification base on [go-mail](https://github.com/wneessen/go-mail)
- [x] cron task base on [gocron](https://github.com/go-co-op/gocron)
- [x] static file storage base on Aliyun OSS

## Development Convention

- DAO return result and error
- Service return result, and panic service.Error while error occurred
- Router should not handle error, handle it in recovery middleware
- JSON field name should be camel case
- struct field name should be upper camel case