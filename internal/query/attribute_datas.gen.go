// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/Thingsly/backend/internal/model"
)

func newAttributeData(db *gorm.DB, opts ...gen.DOOption) attributeData {
	_attributeData := attributeData{}

	_attributeData.attributeDataDo.UseDB(db, opts...)
	_attributeData.attributeDataDo.UseModel(&model.AttributeData{})

	tableName := _attributeData.attributeDataDo.TableName()
	_attributeData.ALL = field.NewAsterisk(tableName)
	_attributeData.ID = field.NewString(tableName, "id")
	_attributeData.DeviceID = field.NewString(tableName, "device_id")
	_attributeData.Key = field.NewString(tableName, "key")
	_attributeData.T = field.NewTime(tableName, "ts")
	_attributeData.BoolV = field.NewBool(tableName, "bool_v")
	_attributeData.NumberV = field.NewFloat64(tableName, "number_v")
	_attributeData.StringV = field.NewString(tableName, "string_v")
	_attributeData.TenantID = field.NewString(tableName, "tenant_id")

	_attributeData.fillFieldMap()

	return _attributeData
}

type attributeData struct {
	attributeDataDo

	ALL      field.Asterisk
	ID       field.String
	DeviceID field.String 
	Key      field.String 
	T        field.Time   
	BoolV    field.Bool
	NumberV  field.Float64
	StringV  field.String
	TenantID field.String

	fieldMap map[string]field.Expr
}

func (a attributeData) Table(newTableName string) *attributeData {
	a.attributeDataDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a attributeData) As(alias string) *attributeData {
	a.attributeDataDo.DO = *(a.attributeDataDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *attributeData) updateTableName(table string) *attributeData {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewString(table, "id")
	a.DeviceID = field.NewString(table, "device_id")
	a.Key = field.NewString(table, "key")
	a.T = field.NewTime(table, "ts")
	a.BoolV = field.NewBool(table, "bool_v")
	a.NumberV = field.NewFloat64(table, "number_v")
	a.StringV = field.NewString(table, "string_v")
	a.TenantID = field.NewString(table, "tenant_id")

	a.fillFieldMap()

	return a
}

func (a *attributeData) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *attributeData) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 8)
	a.fieldMap["id"] = a.ID
	a.fieldMap["device_id"] = a.DeviceID
	a.fieldMap["key"] = a.Key
	a.fieldMap["ts"] = a.T
	a.fieldMap["bool_v"] = a.BoolV
	a.fieldMap["number_v"] = a.NumberV
	a.fieldMap["string_v"] = a.StringV
	a.fieldMap["tenant_id"] = a.TenantID
}

func (a attributeData) clone(db *gorm.DB) attributeData {
	a.attributeDataDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a attributeData) replaceDB(db *gorm.DB) attributeData {
	a.attributeDataDo.ReplaceDB(db)
	return a
}

type attributeDataDo struct{ gen.DO }

type IAttributeDataDo interface {
	gen.SubQuery
	Debug() IAttributeDataDo
	WithContext(ctx context.Context) IAttributeDataDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAttributeDataDo
	WriteDB() IAttributeDataDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAttributeDataDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAttributeDataDo
	Not(conds ...gen.Condition) IAttributeDataDo
	Or(conds ...gen.Condition) IAttributeDataDo
	Select(conds ...field.Expr) IAttributeDataDo
	Where(conds ...gen.Condition) IAttributeDataDo
	Order(conds ...field.Expr) IAttributeDataDo
	Distinct(cols ...field.Expr) IAttributeDataDo
	Omit(cols ...field.Expr) IAttributeDataDo
	Join(table schema.Tabler, on ...field.Expr) IAttributeDataDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAttributeDataDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAttributeDataDo
	Group(cols ...field.Expr) IAttributeDataDo
	Having(conds ...gen.Condition) IAttributeDataDo
	Limit(limit int) IAttributeDataDo
	Offset(offset int) IAttributeDataDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAttributeDataDo
	Unscoped() IAttributeDataDo
	Create(values ...*model.AttributeData) error
	CreateInBatches(values []*model.AttributeData, batchSize int) error
	Save(values ...*model.AttributeData) error
	First() (*model.AttributeData, error)
	Take() (*model.AttributeData, error)
	Last() (*model.AttributeData, error)
	Find() ([]*model.AttributeData, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AttributeData, err error)
	FindInBatches(result *[]*model.AttributeData, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.AttributeData) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAttributeDataDo
	Assign(attrs ...field.AssignExpr) IAttributeDataDo
	Joins(fields ...field.RelationField) IAttributeDataDo
	Preload(fields ...field.RelationField) IAttributeDataDo
	FirstOrInit() (*model.AttributeData, error)
	FirstOrCreate() (*model.AttributeData, error)
	FindByPage(offset int, limit int) (result []*model.AttributeData, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAttributeDataDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a attributeDataDo) Debug() IAttributeDataDo {
	return a.withDO(a.DO.Debug())
}

func (a attributeDataDo) WithContext(ctx context.Context) IAttributeDataDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a attributeDataDo) ReadDB() IAttributeDataDo {
	return a.Clauses(dbresolver.Read)
}

func (a attributeDataDo) WriteDB() IAttributeDataDo {
	return a.Clauses(dbresolver.Write)
}

func (a attributeDataDo) Session(config *gorm.Session) IAttributeDataDo {
	return a.withDO(a.DO.Session(config))
}

func (a attributeDataDo) Clauses(conds ...clause.Expression) IAttributeDataDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a attributeDataDo) Returning(value interface{}, columns ...string) IAttributeDataDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a attributeDataDo) Not(conds ...gen.Condition) IAttributeDataDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a attributeDataDo) Or(conds ...gen.Condition) IAttributeDataDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a attributeDataDo) Select(conds ...field.Expr) IAttributeDataDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a attributeDataDo) Where(conds ...gen.Condition) IAttributeDataDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a attributeDataDo) Order(conds ...field.Expr) IAttributeDataDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a attributeDataDo) Distinct(cols ...field.Expr) IAttributeDataDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a attributeDataDo) Omit(cols ...field.Expr) IAttributeDataDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a attributeDataDo) Join(table schema.Tabler, on ...field.Expr) IAttributeDataDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a attributeDataDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAttributeDataDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a attributeDataDo) RightJoin(table schema.Tabler, on ...field.Expr) IAttributeDataDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a attributeDataDo) Group(cols ...field.Expr) IAttributeDataDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a attributeDataDo) Having(conds ...gen.Condition) IAttributeDataDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a attributeDataDo) Limit(limit int) IAttributeDataDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a attributeDataDo) Offset(offset int) IAttributeDataDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a attributeDataDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAttributeDataDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a attributeDataDo) Unscoped() IAttributeDataDo {
	return a.withDO(a.DO.Unscoped())
}

func (a attributeDataDo) Create(values ...*model.AttributeData) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a attributeDataDo) CreateInBatches(values []*model.AttributeData, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a attributeDataDo) Save(values ...*model.AttributeData) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a attributeDataDo) First() (*model.AttributeData, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AttributeData), nil
	}
}

func (a attributeDataDo) Take() (*model.AttributeData, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AttributeData), nil
	}
}

func (a attributeDataDo) Last() (*model.AttributeData, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AttributeData), nil
	}
}

func (a attributeDataDo) Find() ([]*model.AttributeData, error) {
	result, err := a.DO.Find()
	return result.([]*model.AttributeData), err
}

func (a attributeDataDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AttributeData, err error) {
	buf := make([]*model.AttributeData, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a attributeDataDo) FindInBatches(result *[]*model.AttributeData, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a attributeDataDo) Attrs(attrs ...field.AssignExpr) IAttributeDataDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a attributeDataDo) Assign(attrs ...field.AssignExpr) IAttributeDataDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a attributeDataDo) Joins(fields ...field.RelationField) IAttributeDataDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a attributeDataDo) Preload(fields ...field.RelationField) IAttributeDataDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a attributeDataDo) FirstOrInit() (*model.AttributeData, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AttributeData), nil
	}
}

func (a attributeDataDo) FirstOrCreate() (*model.AttributeData, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AttributeData), nil
	}
}

func (a attributeDataDo) FindByPage(offset int, limit int) (result []*model.AttributeData, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a attributeDataDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a attributeDataDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a attributeDataDo) Delete(models ...*model.AttributeData) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *attributeDataDo) withDO(do gen.Dao) *attributeDataDo {
	a.DO = *do.(*gen.DO)
	return a
}
