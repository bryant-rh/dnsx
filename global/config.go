package global

import "encoding/json"

// Providers DNS 解析供应商
var Providers []string = []string{"aliyun", "qcloud", "dnspod"}

var (
	// CfgPath 指定配置路径
	CfgPath string
	// Profile 指定配置选项
	Profile string
)

// DNSxConfig 配置文件信息
type DNSxConfig struct {
	// Current 当前默认的 Profile 选项
	Current string                    `json:"current"`
	Items   map[string]DNSxConfigItem `json:"items"`
}

// DNSxConfigItem 具体变量信息
type DNSxConfigItem struct {
	AKID     string   `json:"akid,omitempty"`
	AKEY     string   `json:"akey,omitempty"`
	Provider string   `json:"provider,omitempty"`
	Domains  []string `json:"domains,omitempty"`

	// Token 为 DNSPod 的账户信息
	Token string `json:"token,omitempty"`
	// Token 为 DNSPod 的地域信息
	// CN: dnspod.cn
	// Global: dnpod.com
	Region string `json:"region,omitempty"`
}

// Load 加载配置文件
func Load(cfgPath string) (dnsx DNSxConfig) {
	return
}

// Dump 写入配置文件
func (dnsx DNSxConfig) Dump(cfgPath string) {}

// New 新建 配置文件
// 这里应该使用 template 完成
func (dnsx DNSxConfig) New(cfgPath string) {}

// Marshal 格式化配置文件
func (dnsx DNSxConfig) Marshal() (s string) {

	b, err := json.Marshal(dnsx)
	if err != nil {
		panic(err)
	}

	return string(b)
}

// Add 增加 Profile
func (dnsx DNSxConfig) Add(profile string, item DNSxConfigItem) {}

// Delete 删除 Profile
func (dnsx DNSxConfig) Delete(profile string) {}
