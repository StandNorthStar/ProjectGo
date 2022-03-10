package main


type ConfigV1 struct {
	File   string        `json:"file"`
	Parsed []DirectiveV1   `json:"parsed"`
}

type DirectiveV1 struct {
	Directive string       `json:"directive"`
	Args      []string     `json:"args"`
	Block     *[]DirectiveV1 `json:"block,omitempty"`  // 指向自身的指针变量
}

func main() {
	var config ConfigV1
	var directive DirectiveV1

	config.File = "/usr/local/openresty/nginx/conf/nginx.conf"

	directive.Directive = "location"
	directive.Args = []string{"^/v1/api/test"}

}

