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

func newArticleTag(db *gorm.DB, opts ...gen.DOOption) articleTag {
	_articleTag := articleTag{}

	_articleTag.articleTagDo.UseDB(db, opts...)
	_articleTag.articleTagDo.UseModel(&model.ArticleTag{})

	tableName := _articleTag.articleTagDo.TableName()
	_articleTag.ALL = field.NewAsterisk(tableName)
	_articleTag.ID = field.NewUint32(tableName, "id")
	_articleTag.Name = field.NewString(tableName, "name")
	_articleTag.Remark = field.NewString(tableName, "remark")
	_articleTag.Sort = field.NewUint32(tableName, "sort")
	_articleTag.Status = field.NewUint32(tableName, "status")
	_articleTag.CreatedAt = field.NewTime(tableName, "created_at")
	_articleTag.UpdatedAt = field.NewTime(tableName, "updated_at")

	_articleTag.fillFieldMap()

	return _articleTag
}

// articleTag 文章tags
type articleTag struct {
	articleTagDo articleTagDo

	ALL       field.Asterisk
	ID        field.Uint32
	Name      field.String // tag内容
	Remark    field.String // 备注
	Sort      field.Uint32 // 排序
	Status    field.Uint32 // 显示状态（0=显示，1=隐藏）
	CreatedAt field.Time
	UpdatedAt field.Time

	fieldMap map[string]field.Expr
}

func (a articleTag) Table(newTableName string) *articleTag {
	a.articleTagDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a articleTag) As(alias string) *articleTag {
	a.articleTagDo.DO = *(a.articleTagDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *articleTag) updateTableName(table string) *articleTag {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewUint32(table, "id")
	a.Name = field.NewString(table, "name")
	a.Remark = field.NewString(table, "remark")
	a.Sort = field.NewUint32(table, "sort")
	a.Status = field.NewUint32(table, "status")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")

	a.fillFieldMap()

	return a
}

func (a *articleTag) WithContext(ctx context.Context) *articleTagDo {
	return a.articleTagDo.WithContext(ctx)
}

func (a articleTag) TableName() string { return a.articleTagDo.TableName() }

func (a articleTag) Alias() string { return a.articleTagDo.Alias() }

func (a articleTag) Columns(cols ...field.Expr) gen.Columns { return a.articleTagDo.Columns(cols...) }

func (a *articleTag) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *articleTag) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 7)
	a.fieldMap["id"] = a.ID
	a.fieldMap["name"] = a.Name
	a.fieldMap["remark"] = a.Remark
	a.fieldMap["sort"] = a.Sort
	a.fieldMap["status"] = a.Status
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
}

func (a articleTag) clone(db *gorm.DB) articleTag {
	a.articleTagDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a articleTag) replaceDB(db *gorm.DB) articleTag {
	a.articleTagDo.ReplaceDB(db)
	return a
}

type articleTagDo struct{ gen.DO }

func (a articleTagDo) Debug() *articleTagDo {
	return a.withDO(a.DO.Debug())
}

func (a articleTagDo) WithContext(ctx context.Context) *articleTagDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a articleTagDo) ReadDB() *articleTagDo {
	return a.Clauses(dbresolver.Read)
}

func (a articleTagDo) WriteDB() *articleTagDo {
	return a.Clauses(dbresolver.Write)
}

func (a articleTagDo) Session(config *gorm.Session) *articleTagDo {
	return a.withDO(a.DO.Session(config))
}

func (a articleTagDo) Clauses(conds ...clause.Expression) *articleTagDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a articleTagDo) Returning(value interface{}, columns ...string) *articleTagDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a articleTagDo) Not(conds ...gen.Condition) *articleTagDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a articleTagDo) Or(conds ...gen.Condition) *articleTagDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a articleTagDo) Select(conds ...field.Expr) *articleTagDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a articleTagDo) Where(conds ...gen.Condition) *articleTagDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a articleTagDo) Order(conds ...field.Expr) *articleTagDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a articleTagDo) Distinct(cols ...field.Expr) *articleTagDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a articleTagDo) Omit(cols ...field.Expr) *articleTagDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a articleTagDo) Join(table schema.Tabler, on ...field.Expr) *articleTagDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a articleTagDo) LeftJoin(table schema.Tabler, on ...field.Expr) *articleTagDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a articleTagDo) RightJoin(table schema.Tabler, on ...field.Expr) *articleTagDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a articleTagDo) Group(cols ...field.Expr) *articleTagDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a articleTagDo) Having(conds ...gen.Condition) *articleTagDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a articleTagDo) Limit(limit int) *articleTagDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a articleTagDo) Offset(offset int) *articleTagDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a articleTagDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *articleTagDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a articleTagDo) Unscoped() *articleTagDo {
	return a.withDO(a.DO.Unscoped())
}

func (a articleTagDo) Create(values ...*model.ArticleTag) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a articleTagDo) CreateInBatches(values []*model.ArticleTag, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a articleTagDo) Save(values ...*model.ArticleTag) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a articleTagDo) First() (*model.ArticleTag, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArticleTag), nil
	}
}

func (a articleTagDo) Take() (*model.ArticleTag, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArticleTag), nil
	}
}

func (a articleTagDo) Last() (*model.ArticleTag, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArticleTag), nil
	}
}

func (a articleTagDo) Find() ([]*model.ArticleTag, error) {
	result, err := a.DO.Find()
	return result.([]*model.ArticleTag), err
}

func (a articleTagDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ArticleTag, err error) {
	buf := make([]*model.ArticleTag, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a articleTagDo) FindInBatches(result *[]*model.ArticleTag, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a articleTagDo) Attrs(attrs ...field.AssignExpr) *articleTagDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a articleTagDo) Assign(attrs ...field.AssignExpr) *articleTagDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a articleTagDo) Joins(fields ...field.RelationField) *articleTagDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a articleTagDo) Preload(fields ...field.RelationField) *articleTagDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a articleTagDo) FirstOrInit() (*model.ArticleTag, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArticleTag), nil
	}
}

func (a articleTagDo) FirstOrCreate() (*model.ArticleTag, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArticleTag), nil
	}
}

func (a articleTagDo) FindByPage(offset int, limit int) (result []*model.ArticleTag, count int64, err error) {
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

func (a articleTagDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a articleTagDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a articleTagDo) Delete(models ...*model.ArticleTag) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *articleTagDo) withDO(do gen.Dao) *articleTagDo {
	a.DO = *do.(*gen.DO)
	return a
}