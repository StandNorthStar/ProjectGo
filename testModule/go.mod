module TestModules

go 1.16

require (
	github.com/aluttik/go-crossplane v0.0.0-20210526174032-f987c53bd056
	github.com/go-git/go-git/v5 v5.4.2
	github.com/google/go-cmp v0.4.0 // indirect
	github.com/tufanbarisyildirim/gonginx v0.0.0-20210817111223-7fdce97d53d6
)

replace github.com/caddyserver/nginx-adapter => ./pkg/nginx-adapter
