module test_gencert

go 1.16

require (
	github.com/wonderivan/logger v1.0.0
	k8s.io/client-go v0.18.0
)

replace (
	//github.com/fanux/lvscare => github.com/fanux/lvscare v0.0.0-20201224091410-96651f6cbbad
	github.com/wonderivan/logger => ./pkg/logger
)