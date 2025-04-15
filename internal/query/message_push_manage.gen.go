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

func newMessagePushManage(db *gorm.DB, opts ...gen.DOOption) messagePushManage {
	_messagePushManage := messagePushManage{}

	_messagePushManage.messagePushManageDo.UseDB(db, opts...)
	_messagePushManage.messagePushManageDo.UseModel(&model.MessagePushManage{})

	tableName := _messagePushManage.messagePushManageDo.TableName()
	_messagePushManage.ALL = field.NewAsterisk(tableName)
	_messagePushManage.ID = field.NewString(tableName, "id")
	_messagePushManage.UserID = field.NewString(tableName, "user_id")
	_messagePushManage.PushID = field.NewString(tableName, "push_id")
	_messagePushManage.DeviceType = field.NewString(tableName, "device_type")
	_messagePushManage.Status = field.NewInt16(tableName, "status")
	_messagePushManage.CreateTime = field.NewTime(tableName, "create_time")
	_messagePushManage.UpdateTime = field.NewTime(tableName, "update_time")
	_messagePushManage.DeleteTime = field.NewTime(tableName, "delete_time")
	_messagePushManage.LastPushTime = field.NewTime(tableName, "last_push_time")
	_messagePushManage.ErrCount = field.NewInt32(tableName, "err_count")
	_messagePushManage.InactiveTime = field.NewTime(tableName, "inactive_time")

	_messagePushManage.fillFieldMap()

	return _messagePushManage
}

type messagePushManage struct {
	messagePushManageDo

	ALL          field.Asterisk
	ID           field.String
	UserID       field.String 
	PushID       field.String 
	DeviceType   field.String 
	Status       field.Int16  
	CreateTime   field.Time   
	UpdateTime   field.Time   
	DeleteTime   field.Time   
	LastPushTime field.Time   
	ErrCount     field.Int32  
	InactiveTime field.Time   
	fieldMap map[string]field.Expr
}

func (m messagePushManage) Table(newTableName string) *messagePushManage {
	m.messagePushManageDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m messagePushManage) As(alias string) *messagePushManage {
	m.messagePushManageDo.DO = *(m.messagePushManageDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *messagePushManage) updateTableName(table string) *messagePushManage {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewString(table, "id")
	m.UserID = field.NewString(table, "user_id")
	m.PushID = field.NewString(table, "push_id")
	m.DeviceType = field.NewString(table, "device_type")
	m.Status = field.NewInt16(table, "status")
	m.CreateTime = field.NewTime(table, "create_time")
	m.UpdateTime = field.NewTime(table, "update_time")
	m.DeleteTime = field.NewTime(table, "delete_time")
	m.LastPushTime = field.NewTime(table, "last_push_time")
	m.ErrCount = field.NewInt32(table, "err_count")
	m.InactiveTime = field.NewTime(table, "inactive_time")

	m.fillFieldMap()

	return m
}

func (m *messagePushManage) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *messagePushManage) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 10)
	m.fieldMap["id"] = m.ID
	m.fieldMap["user_id"] = m.UserID
	m.fieldMap["push_id"] = m.PushID
	m.fieldMap["device_type"] = m.DeviceType
	m.fieldMap["status"] = m.Status
	m.fieldMap["create_time"] = m.CreateTime
	m.fieldMap["update_time"] = m.UpdateTime
	m.fieldMap["delete_time"] = m.DeleteTime
	m.fieldMap["last_push_time"] = m.LastPushTime
	m.fieldMap["err_count"] = m.ErrCount
	m.fieldMap["InactiveTime"] = m.InactiveTime
	//m.InactiveTime = field.NewTime(table, "inactive_time")
}

func (m messagePushManage) clone(db *gorm.DB) messagePushManage {
	m.messagePushManageDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m messagePushManage) replaceDB(db *gorm.DB) messagePushManage {
	m.messagePushManageDo.ReplaceDB(db)
	return m
}

type messagePushManageDo struct{ gen.DO }

type IMessagePushManageDo interface {
	gen.SubQuery
	Debug() IMessagePushManageDo
	WithContext(ctx context.Context) IMessagePushManageDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IMessagePushManageDo
	WriteDB() IMessagePushManageDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IMessagePushManageDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IMessagePushManageDo
	Not(conds ...gen.Condition) IMessagePushManageDo
	Or(conds ...gen.Condition) IMessagePushManageDo
	Select(conds ...field.Expr) IMessagePushManageDo
	Where(conds ...gen.Condition) IMessagePushManageDo
	Order(conds ...field.Expr) IMessagePushManageDo
	Distinct(cols ...field.Expr) IMessagePushManageDo
	Omit(cols ...field.Expr) IMessagePushManageDo
	Join(table schema.Tabler, on ...field.Expr) IMessagePushManageDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IMessagePushManageDo
	RightJoin(table schema.Tabler, on ...field.Expr) IMessagePushManageDo
	Group(cols ...field.Expr) IMessagePushManageDo
	Having(conds ...gen.Condition) IMessagePushManageDo
	Limit(limit int) IMessagePushManageDo
	Offset(offset int) IMessagePushManageDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IMessagePushManageDo
	Unscoped() IMessagePushManageDo
	Create(values ...*model.MessagePushManage) error
	CreateInBatches(values []*model.MessagePushManage, batchSize int) error
	Save(values ...*model.MessagePushManage) error
	First() (*model.MessagePushManage, error)
	Take() (*model.MessagePushManage, error)
	Last() (*model.MessagePushManage, error)
	Find() ([]*model.MessagePushManage, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MessagePushManage, err error)
	FindInBatches(result *[]*model.MessagePushManage, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.MessagePushManage) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IMessagePushManageDo
	Assign(attrs ...field.AssignExpr) IMessagePushManageDo
	Joins(fields ...field.RelationField) IMessagePushManageDo
	Preload(fields ...field.RelationField) IMessagePushManageDo
	FirstOrInit() (*model.MessagePushManage, error)
	FirstOrCreate() (*model.MessagePushManage, error)
	FindByPage(offset int, limit int) (result []*model.MessagePushManage, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IMessagePushManageDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (m messagePushManageDo) Debug() IMessagePushManageDo {
	return m.withDO(m.DO.Debug())
}

func (m messagePushManageDo) WithContext(ctx context.Context) IMessagePushManageDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m messagePushManageDo) ReadDB() IMessagePushManageDo {
	return m.Clauses(dbresolver.Read)
}

func (m messagePushManageDo) WriteDB() IMessagePushManageDo {
	return m.Clauses(dbresolver.Write)
}

func (m messagePushManageDo) Session(config *gorm.Session) IMessagePushManageDo {
	return m.withDO(m.DO.Session(config))
}

func (m messagePushManageDo) Clauses(conds ...clause.Expression) IMessagePushManageDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m messagePushManageDo) Returning(value interface{}, columns ...string) IMessagePushManageDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m messagePushManageDo) Not(conds ...gen.Condition) IMessagePushManageDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m messagePushManageDo) Or(conds ...gen.Condition) IMessagePushManageDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m messagePushManageDo) Select(conds ...field.Expr) IMessagePushManageDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m messagePushManageDo) Where(conds ...gen.Condition) IMessagePushManageDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m messagePushManageDo) Order(conds ...field.Expr) IMessagePushManageDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m messagePushManageDo) Distinct(cols ...field.Expr) IMessagePushManageDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m messagePushManageDo) Omit(cols ...field.Expr) IMessagePushManageDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m messagePushManageDo) Join(table schema.Tabler, on ...field.Expr) IMessagePushManageDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m messagePushManageDo) LeftJoin(table schema.Tabler, on ...field.Expr) IMessagePushManageDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m messagePushManageDo) RightJoin(table schema.Tabler, on ...field.Expr) IMessagePushManageDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m messagePushManageDo) Group(cols ...field.Expr) IMessagePushManageDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m messagePushManageDo) Having(conds ...gen.Condition) IMessagePushManageDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m messagePushManageDo) Limit(limit int) IMessagePushManageDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m messagePushManageDo) Offset(offset int) IMessagePushManageDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m messagePushManageDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IMessagePushManageDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m messagePushManageDo) Unscoped() IMessagePushManageDo {
	return m.withDO(m.DO.Unscoped())
}

func (m messagePushManageDo) Create(values ...*model.MessagePushManage) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m messagePushManageDo) CreateInBatches(values []*model.MessagePushManage, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m messagePushManageDo) Save(values ...*model.MessagePushManage) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m messagePushManageDo) First() (*model.MessagePushManage, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.MessagePushManage), nil
	}
}

func (m messagePushManageDo) Take() (*model.MessagePushManage, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.MessagePushManage), nil
	}
}

func (m messagePushManageDo) Last() (*model.MessagePushManage, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.MessagePushManage), nil
	}
}

func (m messagePushManageDo) Find() ([]*model.MessagePushManage, error) {
	result, err := m.DO.Find()
	return result.([]*model.MessagePushManage), err
}

func (m messagePushManageDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MessagePushManage, err error) {
	buf := make([]*model.MessagePushManage, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m messagePushManageDo) FindInBatches(result *[]*model.MessagePushManage, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m messagePushManageDo) Attrs(attrs ...field.AssignExpr) IMessagePushManageDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m messagePushManageDo) Assign(attrs ...field.AssignExpr) IMessagePushManageDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m messagePushManageDo) Joins(fields ...field.RelationField) IMessagePushManageDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m messagePushManageDo) Preload(fields ...field.RelationField) IMessagePushManageDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m messagePushManageDo) FirstOrInit() (*model.MessagePushManage, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.MessagePushManage), nil
	}
}

func (m messagePushManageDo) FirstOrCreate() (*model.MessagePushManage, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.MessagePushManage), nil
	}
}

func (m messagePushManageDo) FindByPage(offset int, limit int) (result []*model.MessagePushManage, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m messagePushManageDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m messagePushManageDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m messagePushManageDo) Delete(models ...*model.MessagePushManage) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *messagePushManageDo) withDO(do gen.Dao) *messagePushManageDo {
	m.DO = *do.(*gen.DO)
	return m
}
