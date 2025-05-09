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

func newDeviceModelCommand(db *gorm.DB, opts ...gen.DOOption) deviceModelCommand {
	_deviceModelCommand := deviceModelCommand{}

	_deviceModelCommand.deviceModelCommandDo.UseDB(db, opts...)
	_deviceModelCommand.deviceModelCommandDo.UseModel(&model.DeviceModelCommand{})

	tableName := _deviceModelCommand.deviceModelCommandDo.TableName()
	_deviceModelCommand.ALL = field.NewAsterisk(tableName)
	_deviceModelCommand.ID = field.NewString(tableName, "id")
	_deviceModelCommand.DeviceTemplateID = field.NewString(tableName, "device_template_id")
	_deviceModelCommand.DataName = field.NewString(tableName, "data_name")
	_deviceModelCommand.DataIdentifier = field.NewString(tableName, "data_identifier")
	_deviceModelCommand.Param = field.NewString(tableName, "params")
	_deviceModelCommand.Description = field.NewString(tableName, "description")
	_deviceModelCommand.AdditionalInfo = field.NewString(tableName, "additional_info")
	_deviceModelCommand.CreatedAt = field.NewTime(tableName, "created_at")
	_deviceModelCommand.UpdatedAt = field.NewTime(tableName, "updated_at")
	_deviceModelCommand.Remark = field.NewString(tableName, "remark")
	_deviceModelCommand.TenantID = field.NewString(tableName, "tenant_id")

	_deviceModelCommand.fillFieldMap()

	return _deviceModelCommand
}

type deviceModelCommand struct {
	deviceModelCommandDo

	ALL              field.Asterisk
	ID               field.String 
	DeviceTemplateID field.String 
	DataName         field.String 
	DataIdentifier   field.String 
	Param            field.String 
	Description      field.String 
	AdditionalInfo   field.String 
	CreatedAt        field.Time   
	UpdatedAt        field.Time   
	Remark           field.String 
	TenantID         field.String

	fieldMap map[string]field.Expr
}

func (d deviceModelCommand) Table(newTableName string) *deviceModelCommand {
	d.deviceModelCommandDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d deviceModelCommand) As(alias string) *deviceModelCommand {
	d.deviceModelCommandDo.DO = *(d.deviceModelCommandDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *deviceModelCommand) updateTableName(table string) *deviceModelCommand {
	d.ALL = field.NewAsterisk(table)
	d.ID = field.NewString(table, "id")
	d.DeviceTemplateID = field.NewString(table, "device_template_id")
	d.DataName = field.NewString(table, "data_name")
	d.DataIdentifier = field.NewString(table, "data_identifier")
	d.Param = field.NewString(table, "params")
	d.Description = field.NewString(table, "description")
	d.AdditionalInfo = field.NewString(table, "additional_info")
	d.CreatedAt = field.NewTime(table, "created_at")
	d.UpdatedAt = field.NewTime(table, "updated_at")
	d.Remark = field.NewString(table, "remark")
	d.TenantID = field.NewString(table, "tenant_id")

	d.fillFieldMap()

	return d
}

func (d *deviceModelCommand) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *deviceModelCommand) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 11)
	d.fieldMap["id"] = d.ID
	d.fieldMap["device_template_id"] = d.DeviceTemplateID
	d.fieldMap["data_name"] = d.DataName
	d.fieldMap["data_identifier"] = d.DataIdentifier
	d.fieldMap["params"] = d.Param
	d.fieldMap["description"] = d.Description
	d.fieldMap["additional_info"] = d.AdditionalInfo
	d.fieldMap["created_at"] = d.CreatedAt
	d.fieldMap["updated_at"] = d.UpdatedAt
	d.fieldMap["remark"] = d.Remark
	d.fieldMap["tenant_id"] = d.TenantID
}

func (d deviceModelCommand) clone(db *gorm.DB) deviceModelCommand {
	d.deviceModelCommandDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d deviceModelCommand) replaceDB(db *gorm.DB) deviceModelCommand {
	d.deviceModelCommandDo.ReplaceDB(db)
	return d
}

type deviceModelCommandDo struct{ gen.DO }

type IDeviceModelCommandDo interface {
	gen.SubQuery
	Debug() IDeviceModelCommandDo
	WithContext(ctx context.Context) IDeviceModelCommandDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDeviceModelCommandDo
	WriteDB() IDeviceModelCommandDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDeviceModelCommandDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDeviceModelCommandDo
	Not(conds ...gen.Condition) IDeviceModelCommandDo
	Or(conds ...gen.Condition) IDeviceModelCommandDo
	Select(conds ...field.Expr) IDeviceModelCommandDo
	Where(conds ...gen.Condition) IDeviceModelCommandDo
	Order(conds ...field.Expr) IDeviceModelCommandDo
	Distinct(cols ...field.Expr) IDeviceModelCommandDo
	Omit(cols ...field.Expr) IDeviceModelCommandDo
	Join(table schema.Tabler, on ...field.Expr) IDeviceModelCommandDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDeviceModelCommandDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDeviceModelCommandDo
	Group(cols ...field.Expr) IDeviceModelCommandDo
	Having(conds ...gen.Condition) IDeviceModelCommandDo
	Limit(limit int) IDeviceModelCommandDo
	Offset(offset int) IDeviceModelCommandDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDeviceModelCommandDo
	Unscoped() IDeviceModelCommandDo
	Create(values ...*model.DeviceModelCommand) error
	CreateInBatches(values []*model.DeviceModelCommand, batchSize int) error
	Save(values ...*model.DeviceModelCommand) error
	First() (*model.DeviceModelCommand, error)
	Take() (*model.DeviceModelCommand, error)
	Last() (*model.DeviceModelCommand, error)
	Find() ([]*model.DeviceModelCommand, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DeviceModelCommand, err error)
	FindInBatches(result *[]*model.DeviceModelCommand, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.DeviceModelCommand) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDeviceModelCommandDo
	Assign(attrs ...field.AssignExpr) IDeviceModelCommandDo
	Joins(fields ...field.RelationField) IDeviceModelCommandDo
	Preload(fields ...field.RelationField) IDeviceModelCommandDo
	FirstOrInit() (*model.DeviceModelCommand, error)
	FirstOrCreate() (*model.DeviceModelCommand, error)
	FindByPage(offset int, limit int) (result []*model.DeviceModelCommand, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDeviceModelCommandDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d deviceModelCommandDo) Debug() IDeviceModelCommandDo {
	return d.withDO(d.DO.Debug())
}

func (d deviceModelCommandDo) WithContext(ctx context.Context) IDeviceModelCommandDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d deviceModelCommandDo) ReadDB() IDeviceModelCommandDo {
	return d.Clauses(dbresolver.Read)
}

func (d deviceModelCommandDo) WriteDB() IDeviceModelCommandDo {
	return d.Clauses(dbresolver.Write)
}

func (d deviceModelCommandDo) Session(config *gorm.Session) IDeviceModelCommandDo {
	return d.withDO(d.DO.Session(config))
}

func (d deviceModelCommandDo) Clauses(conds ...clause.Expression) IDeviceModelCommandDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d deviceModelCommandDo) Returning(value interface{}, columns ...string) IDeviceModelCommandDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d deviceModelCommandDo) Not(conds ...gen.Condition) IDeviceModelCommandDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d deviceModelCommandDo) Or(conds ...gen.Condition) IDeviceModelCommandDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d deviceModelCommandDo) Select(conds ...field.Expr) IDeviceModelCommandDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d deviceModelCommandDo) Where(conds ...gen.Condition) IDeviceModelCommandDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d deviceModelCommandDo) Order(conds ...field.Expr) IDeviceModelCommandDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d deviceModelCommandDo) Distinct(cols ...field.Expr) IDeviceModelCommandDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d deviceModelCommandDo) Omit(cols ...field.Expr) IDeviceModelCommandDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d deviceModelCommandDo) Join(table schema.Tabler, on ...field.Expr) IDeviceModelCommandDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d deviceModelCommandDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDeviceModelCommandDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d deviceModelCommandDo) RightJoin(table schema.Tabler, on ...field.Expr) IDeviceModelCommandDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d deviceModelCommandDo) Group(cols ...field.Expr) IDeviceModelCommandDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d deviceModelCommandDo) Having(conds ...gen.Condition) IDeviceModelCommandDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d deviceModelCommandDo) Limit(limit int) IDeviceModelCommandDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d deviceModelCommandDo) Offset(offset int) IDeviceModelCommandDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d deviceModelCommandDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDeviceModelCommandDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d deviceModelCommandDo) Unscoped() IDeviceModelCommandDo {
	return d.withDO(d.DO.Unscoped())
}

func (d deviceModelCommandDo) Create(values ...*model.DeviceModelCommand) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d deviceModelCommandDo) CreateInBatches(values []*model.DeviceModelCommand, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d deviceModelCommandDo) Save(values ...*model.DeviceModelCommand) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d deviceModelCommandDo) First() (*model.DeviceModelCommand, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeviceModelCommand), nil
	}
}

func (d deviceModelCommandDo) Take() (*model.DeviceModelCommand, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeviceModelCommand), nil
	}
}

func (d deviceModelCommandDo) Last() (*model.DeviceModelCommand, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeviceModelCommand), nil
	}
}

func (d deviceModelCommandDo) Find() ([]*model.DeviceModelCommand, error) {
	result, err := d.DO.Find()
	return result.([]*model.DeviceModelCommand), err
}

func (d deviceModelCommandDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DeviceModelCommand, err error) {
	buf := make([]*model.DeviceModelCommand, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d deviceModelCommandDo) FindInBatches(result *[]*model.DeviceModelCommand, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d deviceModelCommandDo) Attrs(attrs ...field.AssignExpr) IDeviceModelCommandDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d deviceModelCommandDo) Assign(attrs ...field.AssignExpr) IDeviceModelCommandDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d deviceModelCommandDo) Joins(fields ...field.RelationField) IDeviceModelCommandDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d deviceModelCommandDo) Preload(fields ...field.RelationField) IDeviceModelCommandDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d deviceModelCommandDo) FirstOrInit() (*model.DeviceModelCommand, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeviceModelCommand), nil
	}
}

func (d deviceModelCommandDo) FirstOrCreate() (*model.DeviceModelCommand, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeviceModelCommand), nil
	}
}

func (d deviceModelCommandDo) FindByPage(offset int, limit int) (result []*model.DeviceModelCommand, count int64, err error) {
	result, err = d.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = d.Offset(-1).Limit(-1).Count()
	return
}

func (d deviceModelCommandDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d deviceModelCommandDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d deviceModelCommandDo) Delete(models ...*model.DeviceModelCommand) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *deviceModelCommandDo) withDO(do gen.Dao) *deviceModelCommandDo {
	d.DO = *do.(*gen.DO)
	return d
}
