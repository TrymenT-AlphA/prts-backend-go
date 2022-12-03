# prts-backend
prts-backend is a REST style api using golang

## modify config
```
config.prod.json
{
  "port": ":3000",
  "dsn": "<user>:<pwd>@tcp(<host>:<port>)/<database>?charset=utf8mb4&parseTime=True&loc=Local",
  "prefork": true,
  "caseSensitive": false,
  "strictRouting": false,
  "serverHeader": "prts-backend powered by fiber",
  "appName": "prts-backend v1"
}
```

## run
```
./cmd/docker.bat
```
