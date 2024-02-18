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

	"ice_flame_gin/internal/app/models/model"
)

func newUcSystemMaster(db *gorm.DB, opts ...gen.DOOption) ucSystemMaster {
	_ucSystemMaster := ucSystemMaster{}

	_ucSystemMaster.ucSystemMasterDo.UseDB(db, opts...)
	_ucSystemMaster.ucSystemMasterDo.UseModel(&model.UcSystemMaster{})

	tableName := _ucSystemMaster.ucSystemMasterDo.TableName()
	_ucSystemMaster.ALL = field.NewAsterisk(tableName)
	_ucSystemMaster.ID = field.NewUint32(tableName, "id")
	_ucSystemMaster.AccountID = field.NewUint32(tableName, "account_id")
	_ucSystemMaster.Name = field.NewString(tableName, "name")
	_ucSystemMaster.Email = field.NewString(tableName, "email")
	_ucSystemMaster.Tel = field.NewString(tableName, "tel")
	_ucSystemMaster.SuperMaster = field.NewBool(tableName, "super_master")
	_ucSystemMaster.CreatedAt = field.NewTime(tableName, "created_at")
	_ucSystemMaster.UpdatedAt = field.NewTime(tableName, "updated_at")

	_ucSystemMaster.fillFieldMap()

	return _ucSystemMaster
}

// ucSystemMaster 管理员用户表
type ucSystemMaster struct {
	ucSystemMasterDo ucSystemMasterDo

	ALL         field.Asterisk
	ID          field.Uint32 // 自增ID
	AccountID   field.Uint32 // 关联account id
	Name        field.String // 用户姓名
	Email       field.String // 用户邮箱
	Tel         field.String // 用户电话
	SuperMaster field.Bool   // 超级管理员(=1为超级管理员)
	CreatedAt   field.Time
	UpdatedAt   field.Time

	fieldMap map[string]field.Expr
}

func (u ucSystemMaster) Table(newTableName string) *ucSystemMaster {
	u.ucSystemMasterDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u ucSystemMaster) As(alias string) *ucSystemMaster {
	u.ucSystemMasterDo.DO = *(u.ucSystemMasterDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *ucSystemMaster) updateTableName(table string) *ucSystemMaster {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewUint32(table, "id")
	u.AccountID = field.NewUint32(table, "account_id")
	u.Name = field.NewString(table, "name")
	u.Email = field.NewString(table, "email")
	u.Tel = field.NewString(table, "tel")
	u.SuperMaster = field.NewBool(table, "super_master")
	u.CreatedAt = field.NewTime(table, "created_at")
	u.UpdatedAt = field.NewTime(table, "updated_at")

	u.fillFieldMap()

	return u
}

func (u *ucSystemMaster) WithContext(ctx context.Context) *ucSystemMasterDo {
	return u.ucSystemMasterDo.WithContext(ctx)
}

func (u ucSystemMaster) TableName() string { return u.ucSystemMasterDo.TableName() }

func (u ucSystemMaster) Alias() string { return u.ucSystemMasterDo.Alias() }

func (u ucSystemMaster) Columns(cols ...field.Expr) gen.Columns {
	return u.ucSystemMasterDo.Columns(cols...)
}

func (u *ucSystemMaster) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *ucSystemMaster) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 8)
	u.fieldMap["id"] = u.ID
	u.fieldMap["account_id"] = u.AccountID
	u.fieldMap["name"] = u.Name
	u.fieldMap["email"] = u.Email
	u.fieldMap["tel"] = u.Tel
	u.fieldMap["super_master"] = u.SuperMaster
	u.fieldMap["created_at"] = u.CreatedAt
	u.fieldMap["updated_at"] = u.UpdatedAt
}

func (u ucSystemMaster) clone(db *gorm.DB) ucSystemMaster {
	u.ucSystemMasterDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u ucSystemMaster) replaceDB(db *gorm.DB) ucSystemMaster {
	u.ucSystemMasterDo.ReplaceDB(db)
	return u
}

type ucSystemMasterDo struct{ gen.DO }

func (u ucSystemMasterDo) Debug() *ucSystemMasterDo {
	return u.withDO(u.DO.Debug())
}

func (u ucSystemMasterDo) WithContext(ctx context.Context) *ucSystemMasterDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u ucSystemMasterDo) ReadDB() *ucSystemMasterDo {
	return u.Clauses(dbresolver.Read)
}

func (u ucSystemMasterDo) WriteDB() *ucSystemMasterDo {
	return u.Clauses(dbresolver.Write)
}

func (u ucSystemMasterDo) Session(config *gorm.Session) *ucSystemMasterDo {
	return u.withDO(u.DO.Session(config))
}

func (u ucSystemMasterDo) Clauses(conds ...clause.Expression) *ucSystemMasterDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u ucSystemMasterDo) Returning(value interface{}, columns ...string) *ucSystemMasterDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u ucSystemMasterDo) Not(conds ...gen.Condition) *ucSystemMasterDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u ucSystemMasterDo) Or(conds ...gen.Condition) *ucSystemMasterDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u ucSystemMasterDo) Select(conds ...field.Expr) *ucSystemMasterDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u ucSystemMasterDo) Where(conds ...gen.Condition) *ucSystemMasterDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u ucSystemMasterDo) Order(conds ...field.Expr) *ucSystemMasterDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u ucSystemMasterDo) Distinct(cols ...field.Expr) *ucSystemMasterDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u ucSystemMasterDo) Omit(cols ...field.Expr) *ucSystemMasterDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u ucSystemMasterDo) Join(table schema.Tabler, on ...field.Expr) *ucSystemMasterDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u ucSystemMasterDo) LeftJoin(table schema.Tabler, on ...field.Expr) *ucSystemMasterDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u ucSystemMasterDo) RightJoin(table schema.Tabler, on ...field.Expr) *ucSystemMasterDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u ucSystemMasterDo) Group(cols ...field.Expr) *ucSystemMasterDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u ucSystemMasterDo) Having(conds ...gen.Condition) *ucSystemMasterDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u ucSystemMasterDo) Limit(limit int) *ucSystemMasterDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u ucSystemMasterDo) Offset(offset int) *ucSystemMasterDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u ucSystemMasterDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *ucSystemMasterDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u ucSystemMasterDo) Unscoped() *ucSystemMasterDo {
	return u.withDO(u.DO.Unscoped())
}

func (u ucSystemMasterDo) Create(values ...*model.UcSystemMaster) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u ucSystemMasterDo) CreateInBatches(values []*model.UcSystemMaster, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u ucSystemMasterDo) Save(values ...*model.UcSystemMaster) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u ucSystemMasterDo) First() (*model.UcSystemMaster, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UcSystemMaster), nil
	}
}

func (u ucSystemMasterDo) Take() (*model.UcSystemMaster, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UcSystemMaster), nil
	}
}

func (u ucSystemMasterDo) Last() (*model.UcSystemMaster, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UcSystemMaster), nil
	}
}

func (u ucSystemMasterDo) Find() ([]*model.UcSystemMaster, error) {
	result, err := u.DO.Find()
	return result.([]*model.UcSystemMaster), err
}

func (u ucSystemMasterDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UcSystemMaster, err error) {
	buf := make([]*model.UcSystemMaster, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u ucSystemMasterDo) FindInBatches(result *[]*model.UcSystemMaster, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u ucSystemMasterDo) Attrs(attrs ...field.AssignExpr) *ucSystemMasterDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u ucSystemMasterDo) Assign(attrs ...field.AssignExpr) *ucSystemMasterDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u ucSystemMasterDo) Joins(fields ...field.RelationField) *ucSystemMasterDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u ucSystemMasterDo) Preload(fields ...field.RelationField) *ucSystemMasterDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u ucSystemMasterDo) FirstOrInit() (*model.UcSystemMaster, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UcSystemMaster), nil
	}
}

func (u ucSystemMasterDo) FirstOrCreate() (*model.UcSystemMaster, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UcSystemMaster), nil
	}
}

func (u ucSystemMasterDo) FindByPage(offset int, limit int) (result []*model.UcSystemMaster, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u ucSystemMasterDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u ucSystemMasterDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u ucSystemMasterDo) Delete(models ...*model.UcSystemMaster) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *ucSystemMasterDo) withDO(do gen.Dao) *ucSystemMasterDo {
	u.DO = *do.(*gen.DO)
	return u
}
