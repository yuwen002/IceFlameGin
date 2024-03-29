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

func newUcSystemMasterVisitorLog(db *gorm.DB, opts ...gen.DOOption) ucSystemMasterVisitorLog {
	_ucSystemMasterVisitorLog := ucSystemMasterVisitorLog{}

	_ucSystemMasterVisitorLog.ucSystemMasterVisitorLogDo.UseDB(db, opts...)
	_ucSystemMasterVisitorLog.ucSystemMasterVisitorLogDo.UseModel(&model.UcSystemMasterVisitorLog{})

	tableName := _ucSystemMasterVisitorLog.ucSystemMasterVisitorLogDo.TableName()
	_ucSystemMasterVisitorLog.ALL = field.NewAsterisk(tableName)
	_ucSystemMasterVisitorLog.ID = field.NewUint32(tableName, "id")
	_ucSystemMasterVisitorLog.AccountID = field.NewUint32(tableName, "account_id")
	_ucSystemMasterVisitorLog.OsCategory = field.NewUint32(tableName, "os_category")
	_ucSystemMasterVisitorLog.VisitCategory = field.NewUint32(tableName, "visit_category")
	_ucSystemMasterVisitorLog.UnionID = field.NewUint32(tableName, "union_id")
	_ucSystemMasterVisitorLog.Description = field.NewString(tableName, "description")
	_ucSystemMasterVisitorLog.IPLong = field.NewString(tableName, "ip_long")
	_ucSystemMasterVisitorLog.IP = field.NewString(tableName, "ip")
	_ucSystemMasterVisitorLog.CreatedAt = field.NewTime(tableName, "created_at")
	_ucSystemMasterVisitorLog.UpdatedAt = field.NewTime(tableName, "updated_at")

	_ucSystemMasterVisitorLog.fillFieldMap()

	return _ucSystemMasterVisitorLog
}

// ucSystemMasterVisitorLog 管理员访问记录
type ucSystemMasterVisitorLog struct {
	ucSystemMasterVisitorLogDo ucSystemMasterVisitorLogDo

	ALL           field.Asterisk
	ID            field.Uint32 // 自增ID
	AccountID     field.Uint32 // 关联用户ID
	OsCategory    field.Uint32 // 系统访问类型（1=web端，2=android端，3=IOS端）
	VisitCategory field.Uint32 // 访问类型
	UnionID       field.Uint32 // 关联ID
	Description   field.String // 访问信息描述
	IPLong        field.String
	IP            field.String
	CreatedAt     field.Time
	UpdatedAt     field.Time

	fieldMap map[string]field.Expr
}

func (u ucSystemMasterVisitorLog) Table(newTableName string) *ucSystemMasterVisitorLog {
	u.ucSystemMasterVisitorLogDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u ucSystemMasterVisitorLog) As(alias string) *ucSystemMasterVisitorLog {
	u.ucSystemMasterVisitorLogDo.DO = *(u.ucSystemMasterVisitorLogDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *ucSystemMasterVisitorLog) updateTableName(table string) *ucSystemMasterVisitorLog {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewUint32(table, "id")
	u.AccountID = field.NewUint32(table, "account_id")
	u.OsCategory = field.NewUint32(table, "os_category")
	u.VisitCategory = field.NewUint32(table, "visit_category")
	u.UnionID = field.NewUint32(table, "union_id")
	u.Description = field.NewString(table, "description")
	u.IPLong = field.NewString(table, "ip_long")
	u.IP = field.NewString(table, "ip")
	u.CreatedAt = field.NewTime(table, "created_at")
	u.UpdatedAt = field.NewTime(table, "updated_at")

	u.fillFieldMap()

	return u
}

func (u *ucSystemMasterVisitorLog) WithContext(ctx context.Context) *ucSystemMasterVisitorLogDo {
	return u.ucSystemMasterVisitorLogDo.WithContext(ctx)
}

func (u ucSystemMasterVisitorLog) TableName() string { return u.ucSystemMasterVisitorLogDo.TableName() }

func (u ucSystemMasterVisitorLog) Alias() string { return u.ucSystemMasterVisitorLogDo.Alias() }

func (u ucSystemMasterVisitorLog) Columns(cols ...field.Expr) gen.Columns {
	return u.ucSystemMasterVisitorLogDo.Columns(cols...)
}

func (u *ucSystemMasterVisitorLog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *ucSystemMasterVisitorLog) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 10)
	u.fieldMap["id"] = u.ID
	u.fieldMap["account_id"] = u.AccountID
	u.fieldMap["os_category"] = u.OsCategory
	u.fieldMap["visit_category"] = u.VisitCategory
	u.fieldMap["union_id"] = u.UnionID
	u.fieldMap["description"] = u.Description
	u.fieldMap["ip_long"] = u.IPLong
	u.fieldMap["ip"] = u.IP
	u.fieldMap["created_at"] = u.CreatedAt
	u.fieldMap["updated_at"] = u.UpdatedAt
}

func (u ucSystemMasterVisitorLog) clone(db *gorm.DB) ucSystemMasterVisitorLog {
	u.ucSystemMasterVisitorLogDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u ucSystemMasterVisitorLog) replaceDB(db *gorm.DB) ucSystemMasterVisitorLog {
	u.ucSystemMasterVisitorLogDo.ReplaceDB(db)
	return u
}

type ucSystemMasterVisitorLogDo struct{ gen.DO }

func (u ucSystemMasterVisitorLogDo) Debug() *ucSystemMasterVisitorLogDo {
	return u.withDO(u.DO.Debug())
}

func (u ucSystemMasterVisitorLogDo) WithContext(ctx context.Context) *ucSystemMasterVisitorLogDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u ucSystemMasterVisitorLogDo) ReadDB() *ucSystemMasterVisitorLogDo {
	return u.Clauses(dbresolver.Read)
}

func (u ucSystemMasterVisitorLogDo) WriteDB() *ucSystemMasterVisitorLogDo {
	return u.Clauses(dbresolver.Write)
}

func (u ucSystemMasterVisitorLogDo) Session(config *gorm.Session) *ucSystemMasterVisitorLogDo {
	return u.withDO(u.DO.Session(config))
}

func (u ucSystemMasterVisitorLogDo) Clauses(conds ...clause.Expression) *ucSystemMasterVisitorLogDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u ucSystemMasterVisitorLogDo) Returning(value interface{}, columns ...string) *ucSystemMasterVisitorLogDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u ucSystemMasterVisitorLogDo) Not(conds ...gen.Condition) *ucSystemMasterVisitorLogDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u ucSystemMasterVisitorLogDo) Or(conds ...gen.Condition) *ucSystemMasterVisitorLogDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u ucSystemMasterVisitorLogDo) Select(conds ...field.Expr) *ucSystemMasterVisitorLogDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u ucSystemMasterVisitorLogDo) Where(conds ...gen.Condition) *ucSystemMasterVisitorLogDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u ucSystemMasterVisitorLogDo) Order(conds ...field.Expr) *ucSystemMasterVisitorLogDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u ucSystemMasterVisitorLogDo) Distinct(cols ...field.Expr) *ucSystemMasterVisitorLogDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u ucSystemMasterVisitorLogDo) Omit(cols ...field.Expr) *ucSystemMasterVisitorLogDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u ucSystemMasterVisitorLogDo) Join(table schema.Tabler, on ...field.Expr) *ucSystemMasterVisitorLogDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u ucSystemMasterVisitorLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) *ucSystemMasterVisitorLogDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u ucSystemMasterVisitorLogDo) RightJoin(table schema.Tabler, on ...field.Expr) *ucSystemMasterVisitorLogDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u ucSystemMasterVisitorLogDo) Group(cols ...field.Expr) *ucSystemMasterVisitorLogDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u ucSystemMasterVisitorLogDo) Having(conds ...gen.Condition) *ucSystemMasterVisitorLogDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u ucSystemMasterVisitorLogDo) Limit(limit int) *ucSystemMasterVisitorLogDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u ucSystemMasterVisitorLogDo) Offset(offset int) *ucSystemMasterVisitorLogDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u ucSystemMasterVisitorLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *ucSystemMasterVisitorLogDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u ucSystemMasterVisitorLogDo) Unscoped() *ucSystemMasterVisitorLogDo {
	return u.withDO(u.DO.Unscoped())
}

func (u ucSystemMasterVisitorLogDo) Create(values ...*model.UcSystemMasterVisitorLog) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u ucSystemMasterVisitorLogDo) CreateInBatches(values []*model.UcSystemMasterVisitorLog, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u ucSystemMasterVisitorLogDo) Save(values ...*model.UcSystemMasterVisitorLog) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u ucSystemMasterVisitorLogDo) First() (*model.UcSystemMasterVisitorLog, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UcSystemMasterVisitorLog), nil
	}
}

func (u ucSystemMasterVisitorLogDo) Take() (*model.UcSystemMasterVisitorLog, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UcSystemMasterVisitorLog), nil
	}
}

func (u ucSystemMasterVisitorLogDo) Last() (*model.UcSystemMasterVisitorLog, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UcSystemMasterVisitorLog), nil
	}
}

func (u ucSystemMasterVisitorLogDo) Find() ([]*model.UcSystemMasterVisitorLog, error) {
	result, err := u.DO.Find()
	return result.([]*model.UcSystemMasterVisitorLog), err
}

func (u ucSystemMasterVisitorLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UcSystemMasterVisitorLog, err error) {
	buf := make([]*model.UcSystemMasterVisitorLog, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u ucSystemMasterVisitorLogDo) FindInBatches(result *[]*model.UcSystemMasterVisitorLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u ucSystemMasterVisitorLogDo) Attrs(attrs ...field.AssignExpr) *ucSystemMasterVisitorLogDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u ucSystemMasterVisitorLogDo) Assign(attrs ...field.AssignExpr) *ucSystemMasterVisitorLogDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u ucSystemMasterVisitorLogDo) Joins(fields ...field.RelationField) *ucSystemMasterVisitorLogDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u ucSystemMasterVisitorLogDo) Preload(fields ...field.RelationField) *ucSystemMasterVisitorLogDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u ucSystemMasterVisitorLogDo) FirstOrInit() (*model.UcSystemMasterVisitorLog, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UcSystemMasterVisitorLog), nil
	}
}

func (u ucSystemMasterVisitorLogDo) FirstOrCreate() (*model.UcSystemMasterVisitorLog, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UcSystemMasterVisitorLog), nil
	}
}

func (u ucSystemMasterVisitorLogDo) FindByPage(offset int, limit int) (result []*model.UcSystemMasterVisitorLog, count int64, err error) {
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

func (u ucSystemMasterVisitorLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u ucSystemMasterVisitorLogDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u ucSystemMasterVisitorLogDo) Delete(models ...*model.UcSystemMasterVisitorLog) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *ucSystemMasterVisitorLogDo) withDO(do gen.Dao) *ucSystemMasterVisitorLogDo {
	u.DO = *do.(*gen.DO)
	return u
}
