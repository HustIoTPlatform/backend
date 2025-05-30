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

func newAlarmInfo(db *gorm.DB, opts ...gen.DOOption) alarmInfo {
	_alarmInfo := alarmInfo{}

	_alarmInfo.alarmInfoDo.UseDB(db, opts...)
	_alarmInfo.alarmInfoDo.UseModel(&model.AlarmInfo{})

	tableName := _alarmInfo.alarmInfoDo.TableName()
	_alarmInfo.ALL = field.NewAsterisk(tableName)
	_alarmInfo.ID = field.NewString(tableName, "id")
	_alarmInfo.AlarmConfigID = field.NewString(tableName, "alarm_config_id")
	_alarmInfo.Name = field.NewString(tableName, "name")
	_alarmInfo.AlarmTime = field.NewTime(tableName, "alarm_time")
	_alarmInfo.Description = field.NewString(tableName, "description")
	_alarmInfo.Content = field.NewString(tableName, "content")
	_alarmInfo.Processor = field.NewString(tableName, "processor")
	_alarmInfo.ProcessingResult = field.NewString(tableName, "processing_result")
	_alarmInfo.TenantID = field.NewString(tableName, "tenant_id")
	_alarmInfo.Remark = field.NewString(tableName, "remark")
	_alarmInfo.AlarmLevel = field.NewString(tableName, "alarm_level")

	_alarmInfo.fillFieldMap()

	return _alarmInfo
}

type alarmInfo struct {
	alarmInfoDo

	ALL              field.Asterisk
	ID               field.String
	AlarmConfigID    field.String 
	Name             field.String 
	AlarmTime        field.Time   
	Description      field.String 
	Content          field.String 
	Processor        field.String 
	ProcessingResult field.String 
	TenantID         field.String 
	Remark           field.String
	AlarmLevel       field.String 

	fieldMap map[string]field.Expr
}

func (a alarmInfo) Table(newTableName string) *alarmInfo {
	a.alarmInfoDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a alarmInfo) As(alias string) *alarmInfo {
	a.alarmInfoDo.DO = *(a.alarmInfoDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *alarmInfo) updateTableName(table string) *alarmInfo {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewString(table, "id")
	a.AlarmConfigID = field.NewString(table, "alarm_config_id")
	a.Name = field.NewString(table, "name")
	a.AlarmTime = field.NewTime(table, "alarm_time")
	a.Description = field.NewString(table, "description")
	a.Content = field.NewString(table, "content")
	a.Processor = field.NewString(table, "processor")
	a.ProcessingResult = field.NewString(table, "processing_result")
	a.TenantID = field.NewString(table, "tenant_id")
	a.Remark = field.NewString(table, "remark")
	a.AlarmLevel = field.NewString(table, "alarm_level")

	a.fillFieldMap()

	return a
}

func (a *alarmInfo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *alarmInfo) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 11)
	a.fieldMap["id"] = a.ID
	a.fieldMap["alarm_config_id"] = a.AlarmConfigID
	a.fieldMap["name"] = a.Name
	a.fieldMap["alarm_time"] = a.AlarmTime
	a.fieldMap["description"] = a.Description
	a.fieldMap["content"] = a.Content
	a.fieldMap["processor"] = a.Processor
	a.fieldMap["processing_result"] = a.ProcessingResult
	a.fieldMap["tenant_id"] = a.TenantID
	a.fieldMap["remark"] = a.Remark
	a.fieldMap["alarm_level"] = a.AlarmLevel
}

func (a alarmInfo) clone(db *gorm.DB) alarmInfo {
	a.alarmInfoDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a alarmInfo) replaceDB(db *gorm.DB) alarmInfo {
	a.alarmInfoDo.ReplaceDB(db)
	return a
}

type alarmInfoDo struct{ gen.DO }

type IAlarmInfoDo interface {
	gen.SubQuery
	Debug() IAlarmInfoDo
	WithContext(ctx context.Context) IAlarmInfoDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAlarmInfoDo
	WriteDB() IAlarmInfoDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAlarmInfoDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAlarmInfoDo
	Not(conds ...gen.Condition) IAlarmInfoDo
	Or(conds ...gen.Condition) IAlarmInfoDo
	Select(conds ...field.Expr) IAlarmInfoDo
	Where(conds ...gen.Condition) IAlarmInfoDo
	Order(conds ...field.Expr) IAlarmInfoDo
	Distinct(cols ...field.Expr) IAlarmInfoDo
	Omit(cols ...field.Expr) IAlarmInfoDo
	Join(table schema.Tabler, on ...field.Expr) IAlarmInfoDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAlarmInfoDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAlarmInfoDo
	Group(cols ...field.Expr) IAlarmInfoDo
	Having(conds ...gen.Condition) IAlarmInfoDo
	Limit(limit int) IAlarmInfoDo
	Offset(offset int) IAlarmInfoDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAlarmInfoDo
	Unscoped() IAlarmInfoDo
	Create(values ...*model.AlarmInfo) error
	CreateInBatches(values []*model.AlarmInfo, batchSize int) error
	Save(values ...*model.AlarmInfo) error
	First() (*model.AlarmInfo, error)
	Take() (*model.AlarmInfo, error)
	Last() (*model.AlarmInfo, error)
	Find() ([]*model.AlarmInfo, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AlarmInfo, err error)
	FindInBatches(result *[]*model.AlarmInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.AlarmInfo) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAlarmInfoDo
	Assign(attrs ...field.AssignExpr) IAlarmInfoDo
	Joins(fields ...field.RelationField) IAlarmInfoDo
	Preload(fields ...field.RelationField) IAlarmInfoDo
	FirstOrInit() (*model.AlarmInfo, error)
	FirstOrCreate() (*model.AlarmInfo, error)
	FindByPage(offset int, limit int) (result []*model.AlarmInfo, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAlarmInfoDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a alarmInfoDo) Debug() IAlarmInfoDo {
	return a.withDO(a.DO.Debug())
}

func (a alarmInfoDo) WithContext(ctx context.Context) IAlarmInfoDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a alarmInfoDo) ReadDB() IAlarmInfoDo {
	return a.Clauses(dbresolver.Read)
}

func (a alarmInfoDo) WriteDB() IAlarmInfoDo {
	return a.Clauses(dbresolver.Write)
}

func (a alarmInfoDo) Session(config *gorm.Session) IAlarmInfoDo {
	return a.withDO(a.DO.Session(config))
}

func (a alarmInfoDo) Clauses(conds ...clause.Expression) IAlarmInfoDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a alarmInfoDo) Returning(value interface{}, columns ...string) IAlarmInfoDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a alarmInfoDo) Not(conds ...gen.Condition) IAlarmInfoDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a alarmInfoDo) Or(conds ...gen.Condition) IAlarmInfoDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a alarmInfoDo) Select(conds ...field.Expr) IAlarmInfoDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a alarmInfoDo) Where(conds ...gen.Condition) IAlarmInfoDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a alarmInfoDo) Order(conds ...field.Expr) IAlarmInfoDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a alarmInfoDo) Distinct(cols ...field.Expr) IAlarmInfoDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a alarmInfoDo) Omit(cols ...field.Expr) IAlarmInfoDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a alarmInfoDo) Join(table schema.Tabler, on ...field.Expr) IAlarmInfoDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a alarmInfoDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAlarmInfoDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a alarmInfoDo) RightJoin(table schema.Tabler, on ...field.Expr) IAlarmInfoDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a alarmInfoDo) Group(cols ...field.Expr) IAlarmInfoDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a alarmInfoDo) Having(conds ...gen.Condition) IAlarmInfoDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a alarmInfoDo) Limit(limit int) IAlarmInfoDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a alarmInfoDo) Offset(offset int) IAlarmInfoDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a alarmInfoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAlarmInfoDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a alarmInfoDo) Unscoped() IAlarmInfoDo {
	return a.withDO(a.DO.Unscoped())
}

func (a alarmInfoDo) Create(values ...*model.AlarmInfo) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a alarmInfoDo) CreateInBatches(values []*model.AlarmInfo, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a alarmInfoDo) Save(values ...*model.AlarmInfo) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a alarmInfoDo) First() (*model.AlarmInfo, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AlarmInfo), nil
	}
}

func (a alarmInfoDo) Take() (*model.AlarmInfo, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AlarmInfo), nil
	}
}

func (a alarmInfoDo) Last() (*model.AlarmInfo, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AlarmInfo), nil
	}
}

func (a alarmInfoDo) Find() ([]*model.AlarmInfo, error) {
	result, err := a.DO.Find()
	return result.([]*model.AlarmInfo), err
}

func (a alarmInfoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AlarmInfo, err error) {
	buf := make([]*model.AlarmInfo, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a alarmInfoDo) FindInBatches(result *[]*model.AlarmInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a alarmInfoDo) Attrs(attrs ...field.AssignExpr) IAlarmInfoDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a alarmInfoDo) Assign(attrs ...field.AssignExpr) IAlarmInfoDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a alarmInfoDo) Joins(fields ...field.RelationField) IAlarmInfoDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a alarmInfoDo) Preload(fields ...field.RelationField) IAlarmInfoDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a alarmInfoDo) FirstOrInit() (*model.AlarmInfo, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AlarmInfo), nil
	}
}

func (a alarmInfoDo) FirstOrCreate() (*model.AlarmInfo, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AlarmInfo), nil
	}
}

func (a alarmInfoDo) FindByPage(offset int, limit int) (result []*model.AlarmInfo, count int64, err error) {
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

func (a alarmInfoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a alarmInfoDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a alarmInfoDo) Delete(models ...*model.AlarmInfo) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *alarmInfoDo) withDO(do gen.Dao) *alarmInfoDo {
	a.DO = *do.(*gen.DO)
	return a
}
