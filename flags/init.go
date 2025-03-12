package flags

import (
	"flag"
)

type Option struct {
	User   string // -u admin  -u user
	DB     bool   // 根据models生成表结构，不好用不建议用
	Load   string // 导入数据库文件
	Dump   bool   // 导出数据库
	Model  string // 初始化models
	Es     bool   // 创建索引
	ESDump bool   // 导出es索引
	ESLoad string // 导入es索引
	API    bool   // 初始化API到数据库
}

func Parse() (option *Option) {
	option = new(Option)
	flag.StringVar(&option.Load, "load", "", "导入数据库文件")
	flag.BoolVar(&option.Dump, "dump", false, "导出数据库")
	flag.StringVar(&option.Model, "model", "", "初始化models")
	flag.Parse()
	return option
}

func (option Option) Run() bool {
	if option.Load != "" {
		Load(option.Load)
		return true
	}
	if option.Dump {
		Dump()
		return true
	}
	if option.Model != "" {
		if option.Model == "all" {
			GenModels("")
		} else {
			GenModels(option.Model)
		}
		return true
	}
	return false
}
