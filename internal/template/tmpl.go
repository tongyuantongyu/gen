package template

const HeaderTmpl = `
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package {{.}}

import(
	"gorm.io/gorm"

	"gorm.io/gen/field"
	"gorm.io/gen/helper"
)
`

const FuncTmpl = `
/*
{{.Doc}}*/
func ({{.S}} {{.MethodStruct}}){{.MethodName}}({{range $index,$params:=.Params}}{{if ne $index 0}},{{end}}{{$params.Name}} {{if ne $params.Package ""}}{{$params.Package}}.{{end}}{{$params.Type}}{{end}})({{range $index,$params:=.Result}}{{if ne $index 0}},{{end}}{{$params.Name}} {{if $params.IsArray}}[]{{end}}{{if $params.IsPointer}}*{{end}}{{if ne $params.Package ""}}{{$params.Package}}.{{end}}{{$params.Type}}{{end}}){
	{{if .HasSqlData}}params := map[string]interface{}{ {{range $index,$data:=.SqlData}}
		"{{$data}}":{{$data}},{{end}}
	}
	{{end}}
	{{if eq .Table "_"}}table:={{.S}}.UnderlyingDB().Statement.Table{{end}}

	var generateSQL string
	{{range $line:=.SqlTmplList}}{{$line}}
	{{end}}

	{{if .HasNeedNewResult}}result =new({{if ne .ResultData.Package ""}}{{.ResultData.Package}}.{{end}}{{.ResultData.Type}}){{end}}
	{{.ExecuteResult}} = {{.S}}.UnderlyingDB().{{.GormOption}}(generateSQL{{if .HasSqlData}},params{{end}}){{if not .ResultData.IsNull}}.Find({{if .HasGotPoint}}&{{end}}{{.ResultData.Name}}){{end}}.Error
	return
}

`

const BaseStruct = `
type {{.NewStructName}} struct {
	gen.DO

	{{range $p :=.Members}}{{$p.Name}}  field.{{$p.NewType}}
	{{end}}
}

func New{{.StructName}}(db *gorm.DB) *{{.NewStructName}} {
	_{{.NewStructName}} := new({{.NewStructName}})

	_{{.NewStructName}}.UseDB(db)
	_{{.NewStructName}}.UseModel({{.StructInfo.Package}}.{{.StructInfo.Type}}{})

	{{if .HasMember}}tableName := _{{.NewStructName}}.TableName(){{end}}
	{{range $p :=.Members}} _{{$.NewStructName}}.{{$p.Name}} = field.New{{$p.NewType}}(tableName, "{{$p.ColumnName}}")
	{{end}}
	
	return _{{.NewStructName}}
}

`

const BaseGormFunc = `
func ({{.S}} {{.NewStructName}}) Debug() *{{.NewStructName}} {
	{{.S}}.DO = *{{.S}}.DO.Debug().(*gen.DO)
	return &{{.S}}
}

func ({{.S}} {{.NewStructName}}) Not(conds ...gen.Condition) *{{.NewStructName}} {
	{{.S}}.DO = *{{.S}}.DO.Not(conds...).(*gen.DO)
	return &{{.S}}
}

func ({{.S}} {{.NewStructName}}) Or(conds ...gen.Condition) *{{.NewStructName}} {
	{{.S}}.DO = *{{.S}}.DO.Or(conds...).(*gen.DO)
	return &{{.S}}
}

func ({{.S}} {{.NewStructName}}) Select(conds ...field.Expr) *{{.NewStructName}} {
	{{.S}}.DO = *{{.S}}.DO.Select(conds...).(*gen.DO)
	return &{{.S}}
}

func ({{.S}} {{.NewStructName}}) Where(conds ...gen.Condition) *{{.NewStructName}} {
	{{.S}}.DO = *{{.S}}.DO.Where(conds...).(*gen.DO)
	return &{{.S}}
}

func ({{.S}} {{.NewStructName}}) Order(conds ...field.Expr) *{{.NewStructName}} {
	{{.S}}.DO = *{{.S}}.DO.Order(conds...).(*gen.DO)
	return &{{.S}}
}

func ({{.S}} {{.NewStructName}}) Distinct(conds ...field.Expr) *{{.NewStructName}} {
	{{.S}}.DO = *{{.S}}.DO.Distinct(conds...).(*gen.DO)
	return &{{.S}}
}

func ({{.S}} {{.NewStructName}}) Join(table schema.Tabler, on ...gen.Condition) *{{.NewStructName}} {
	{{.S}}.DO = *{{.S}}.DO.Join(table, on...).(*gen.DO)
	return &{{.S}}
}

func ({{.S}} {{.NewStructName}}) LeftJoin(table schema.Tabler, on ...gen.Condition) *{{.NewStructName}} {
	{{.S}}.DO = *{{.S}}.DO.Join(table, on...).(*gen.DO)
	return &{{.S}}
}

func ({{.S}} {{.NewStructName}}) RightJoin(table schema.Tabler, on ...gen.Condition) *{{.NewStructName}} {
	{{.S}}.DO = *{{.S}}.DO.Join(table, on...).(*gen.DO)
	return &{{.S}}
}

func ({{.S}} {{.NewStructName}}) Take() (*{{.StructInfo.Package}}.{{.StructInfo.Type}}, error) {
	result := new({{.StructInfo.Package}}.{{.StructInfo.Type}})
	if err := {{.S}}.DO.Take(result); err != nil {
		return nil, err
	}
	return result, nil
}

func ({{.S}} {{.NewStructName}}) First() (*{{.StructInfo.Package}}.{{.StructInfo.Type}}, error) {
	result := new({{.StructInfo.Package}}.{{.StructInfo.Type}})
	if err := {{.S}}.DO.First(result); err != nil {
		return nil, err
	}
	return result, nil
}

func ({{.S}} {{.NewStructName}}) Last() (*{{.StructInfo.Package}}.{{.StructInfo.Type}}, error) {
	result := new({{.StructInfo.Package}}.{{.StructInfo.Type}})
	if err := {{.S}}.DO.Last(result); err != nil {
		return nil, err
	}
	return result, nil
}

func ({{.S}} {{.NewStructName}}) Find() (result []*{{.StructInfo.Package}}.{{.StructInfo.Type}},err error) {
	return result, {{.S}}.DO.Find(&result)
}

func ({{.S}} {{.NewStructName}}) Create(info *{{.StructInfo.Package}}.{{.StructInfo.Type}}) error {
	return {{.S}}.DO.Create(info)
}

func ({{.S}} {{.NewStructName}}) BathCreate(infos []*{{.StructInfo.Package}}.{{.StructInfo.Type}}) error {
	return {{.S}}.DO.CreateInBatches(infos, len(infos))
}

func ({{.S}} {{.NewStructName}}) Update(updates map[string]interface{}) error {
	return {{.S}}.DO.Updates(updates)
}

func ({{.S}} {{.NewStructName}}) FindByPage(offset int, limit int) (result []*{{.StructInfo.Package}}.{{.StructInfo.Type}}, count int64, err error) {
	err = {{.S}}.DO.Count(&count)
	if err != nil {
		return
	}
	err = {{.S}}.DO.Offset(offset).Limit(limit).Find(&result)
	return
}

func ({{.S}} {{.NewStructName}}) Delete(conds ...field.Expr) error {
	result := new({{.StructInfo.Package}}.{{.StructInfo.Type}})
	return {{.S}}.DO.Delete(result, conds...)
}
`

const UseTmpl = `
type DB struct{
	db *gorm.DB

	{{range $name,$d :=.Data}}{{$d.StructName}} *{{$d.NewStructName}}
	{{end}}
}

func (d *DB) Transaction(fc func(db *DB) error, opts ...*sql.TxOptions) error {
	return d.db.Transaction(func(tx *gorm.DB) error { return fc(d.withTx(tx)) }, opts...)
}

func (d *DB) Begin(opts ...*sql.TxOptions) *DB {
	d.db = d.db.Begin(opts...)
	return d
}

func (d *DB) Commit() *DB {
	d.db = d.db.Commit()
	return d
}

func (d *DB) Rollback() *DB {
	d.db = d.db.Rollback()
	return d
}

func (d *DB) SavePoint(name string) *DB {
	d.db = d.db.SavePoint(name)
	return d
}

func (d *DB) RollbackTo(name string) *DB {
	d.db = d.db.RollbackTo(name)
	return d
}

func (d *DB) withTx(tx *gorm.DB) *DB {
	newDB := *d
	newDB.db = tx
	return &newDB
}

func Use(db *gorm.DB) *DB {
	return &DB{
		db: db,
		{{range $name,$d :=.Data}}{{$d.StructName}}: New{{$d.StructName}}(db),
		{{end}}
	}
}
`

// ModelTemplate used as a variable because it cannot load template file after packed, params still can pass file
const ModelTemplate = `
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
package {{.StructInfo.Package}}

import "time"

const TableName{{.StructName}} = "{{.TableName}}"

// {{.TableName}}
type {{.StructName}} struct {
    {{range .Members}}
    {{.Name}} {{.ModelType}} ` + "`json:\"{{.ColumnName}}\" gorm:\"column:{{.ColumnName}}\"` // {{.ColumnComment}}" +
	`{{end}}
}

// TableName .
func (*{{.StructName}}) TableName() string {
    return TableName{{.StructName}}
}
`
