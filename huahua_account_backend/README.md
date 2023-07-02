
```shell
go install github.com/zeromicro/go-zero/tools/goctl@latest

goctl model mysql datasource -url="root:123456@tcp(192.168.163.121:3306)/huahua_account_backend" -table="*"  -dir="./account/internal/model" --style go_zero

```

```shell

goctl api go -api account/account.api -dir account --style go_zero

```
