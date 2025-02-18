package flags

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
	"travel-server/global"

	"github.com/sirupsen/logrus"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// 初始化models
func GenModels(Name string) {
	g := gen.NewGenerator(gen.Config{
		/* 相对执行go run的路径，自动创建目录 */
		OutPath:      "gen/query",
		ModelPkgPath: "gen/model",
		// WithDefaultQuery 生成默认查询结构体(作为全局变量使用), 即`Q`结构体和其字段(各表模型)
		// WithoutContext 生成没有context调用限制的代码供查询
		// WithQueryInterface 生成interface形式的查询代码(可导出), 如`Where()`方法返回的就是一个可导出的接口类型
		Mode: gen.WithoutContext,

		// 表字段可为 null 值时, 对应结体字段使用指针类型
		FieldNullable: false,

		//如果数据库中字段有默认值，则生成指针类型的字段，以避免零值（zero-value）问题
		FieldCoverable: false,

		// 模型结构体字段的数字类型的符号表示是否与表字段的一致, `false`指示都用有符号类型
		FieldSignable: false, // detect integer field's unsigned type, adjust generated data type
		// 生成 gorm 标签的字段索引属性
		FieldWithIndexTag: false, // generate with gorm index tag
		// 生成 gorm 标签的字段类型属性
		FieldWithTypeTag: true, // generate with gorm column type tag
	})

	g.UseDB(global.DB)

	// 自定义字段的数据类型
	// 统一数字类型为int,兼容protobuf和thrift
	dataMap := map[string]func(detailType gorm.ColumnType) (dataType string){
		"tinyint":   func(detailType gorm.ColumnType) (dataType string) { return "int" },
		"smallint":  func(detailType gorm.ColumnType) (dataType string) { return "int" },
		"mediumint": func(detailType gorm.ColumnType) (dataType string) { return "int" },
		"bigint":    func(detailType gorm.ColumnType) (dataType string) { return "int" },
		"int":       func(detailType gorm.ColumnType) (dataType string) { return "int" },
		"decimal":   func(detailType gorm.ColumnType) (dataType string) { return "Decimal" }, // 金额类型全部转换为第三方库,github.com/shopspring/decimal
	}
	g.WithDataTypeMap(dataMap)
	if Name != "" {
		g.ApplyBasic(g.GenerateModel(Name))
	} else {
		g.ApplyBasic(g.GenerateAllTable()...)
	}
	g.Execute()

	// 替换model名称，将.gen.go 变为 .go
	filepath.Walk("gen/model", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".gen.go") {
			// 获取新的文件名
			baseName := strings.TrimSuffix(info.Name(), ".gen.go") + ".go"
			newPath := filepath.Join("model/", baseName)

			// 移动文件
			if err := os.Rename(path, newPath); err != nil {
				return err
			}
		}

		return nil
	})

	// 删除query目录文件
	err := os.RemoveAll("gen")
	if err != nil {
		panic(err)
	}
}

// 导入数据库文件
func Load(sqlPath string) {
	byteData, err := os.ReadFile(sqlPath)
	if err != nil {
		logrus.Fatalf("%s err: %s", sqlPath, err.Error())
	}

	// 一定要按照\r\n分割
	sqlList := strings.Split(string(byteData), ";\r\n")
	for _, sql := range sqlList {
		sql = strings.TrimSpace(sql)
		if sql == "" {
			continue
		}
		err = global.DB.Exec(sql).Error
		if err != nil {
			logrus.Errorf("%s err:%s", sql, err.Error())
			continue
		}
	}

	logrus.Infof("%s sql导入成功", sqlPath)

}

// 导出数据库文件
func Dump() {
	mysql := global.Config.Mysql

	timer := time.Now().Format("20060102")

	sqlPath := fmt.Sprintf("sql/%s_%s.sql", mysql.DB, timer)

	// 调用系统命令， 执行mysqldump进行数据库导出
	cmder := fmt.Sprintf("mysqldump -u%s -p%s %s > %s", mysql.User, mysql.Password, mysql.DB, sqlPath)
	cmd := exec.Command("sh", "-c", cmder)

	var out, stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		logrus.Errorln(err.Error(), stderr.String())
		return
	}
	logrus.Infof("sql文件 %s 导出成功", sqlPath)
}
