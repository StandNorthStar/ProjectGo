package main

import (
	"bytes"
	"fmt"
	"github.com/tufanbarisyildirim/gonginx"
	"github.com/tufanbarisyildirim/gonginx/parser"
	"sort"
	"strings"
)

// nginx总配置
type NginxConfig struct {
	File   		string        	  `json:"file"`
	Directives []NginxDirective   `json:"parsed"`
}


/*
// nginx的单条指令数据结构
NginxDirective：
Key: 指令名称
Args: 参数列表
Block： 指向子参数集的指针。比如 location指令，此执行的Block值是location内部所有指令集的指针。
 */
type NginxDirective struct {
	Key 	  string       `json:"Key"`
	Args      []string     `json:"args"`
	Block     *[]NginxDirective `json:"block,omitempty"`
}

/*
GetDirective: 处理单条指令
参数：单条指令，类型：gonginx.IDirective
返回：新结构体
工作流程：
1. 根据传递的指令IDirective，获取指令的Key、Args值，并赋值给NginxDirective；
2. 如果传递的指令IDirective Block不为空，那么获取Block所有的指令列表集；
3. 调用 GetBlockV1() 获取Block指令集的指针。
 */
func GetDirective(d gonginx.IDirective) NginxDirective {
	var directive NginxDirective

	directive.Key = d.GetName()
	directive.Args = d.GetParameters()

	if d.GetBlock() != nil {
		directivelist := GetBlockV1(d.GetBlock())
		directive.Block = directivelist
	}
	return directive
}

/*
GetBlockV1:
参数：nginx配置的实例化对象IBlock
返回：实例化新结构体[]NginxDirective，返回所有指令列表的指针。
工作：
1. 读取nginx的配置，实例化对象IBlock
2. 循环指令列表，调用GetDirective方法，返回一个NginxDirective对象。
3. 把返回新的NginxDirective对象 append到 []NginxDirective 中。
 */
func GetBlockV1(block gonginx.IBlock) *[]NginxDirective {
	/*
		args : //  []IDirective
	*/
	var ngdirectives []NginxDirective
	directives := block.GetDirectives()
	for _, directive := range directives {
		direct := GetDirective(directive)
		ngdirectives = append(ngdirectives, direct)
		//fmt.Println("block", r)
	}

	return &ngdirectives
}
/*
核心思想：
思维不要定式，nginx配置里面的所有项都是指令，唯一有区别的是有普通指令和有子集的指令。如果这样理解就容易处理nginx配置了。
思路：
1. 便利获取指令
2. 处理单个指令
	2.1 查看指令是否包含Block快，如果存在就查询这个块下所有的指令并返回指向这个列表的指针。
	2.2
 */

func main() {
	n := `
upstream testserver{
 server 172.16.1.22:8781;
 server 172.16.1.21:8781;
}
server {
    listen 80;
    server_name test03-json.class.com;
    client_max_body_size 60m;
    location / {
        access_log logs/test-access-loc.log reverseRealIpFormat1;
        proxy_pass https://localhost:8080;
    }
}
server {
    listen 8200;
    proxy_http_version 1.1;
    proxy_buffering off;
    proxy_request_buffering off;
    location / {
    }
    # internal redirect
    location @oss {
        #// endpoint eg: oss.aliyuncs.com
        proxy_pass http://test.com;
    }
}
`
	r := parser.NewStringParser(n)
	rv1 := r.Parse()
	fmt.Println(rv1.FilePath)
	result := GetBlockV1(rv1.Block)
	fmt.Println(result)

	fmt.Println("-------------------------")
	fmt.Println(DumpNginxBlock(result, gonginx.IndentedStyle))

	/*
	示例一: 读取nginx配置，并格式化输出
	 */
	//r := parser.NewStringParser(n)
	//rv1 := r.Parse()
	//fmt.Println(gonginx.DumpConfig(rv1, gonginx.IndentedStyle))
	//fmt.Println(gonginx.DumpDirective(rv1.Directives[1], gonginx.IndentedStyle))
	//ib := rv1.GetDirectives()[1].GetBlock()
	//fmt.Println(gonginx.DumpBlock(ib, gonginx.IndentedStyle))

}


func DumpNginxDirective(d NginxDirective, style *gonginx.Style) string {
	var buf bytes.Buffer
	if style.SpaceBeforeBlocks && d.Block != nil {
		buf.WriteString("\n")
	}

	buf.WriteString(fmt.Sprintf("%s%s", strings.Repeat(" ", style.StartIndent), d.Key))
	if len(d.Args) > 0 {
		buf.WriteString(fmt.Sprintf(" %s", strings.Join(d.Args, " ")))
	}
	if d.Block == nil {
		buf.WriteRune(';')
	} else {
		buf.WriteString(" {\n")
		buf.WriteString(DumpNginxBlock(d.Block, style.Iterate()))
		buf.WriteString(fmt.Sprintf("\n%s}", strings.Repeat(" ", style.StartIndent)))
	}
	return buf.String()
}

func DumpNginxBlock(block *[]NginxDirective, style *gonginx.Style) string {
	var buf bytes.Buffer

	directives := *block

	if style.SortDirectives {
		sort.SliceStable(directives, func(i, j int) bool {
			return directives[i].Key < directives[j].Key
		})
	}

	for i, directive := range directives {
		buf.WriteString(DumpNginxDirective(directive, style))
		if i != len(directives) - 1 {
			buf.WriteString("\n")
		}
	}
	return buf.String()
}

