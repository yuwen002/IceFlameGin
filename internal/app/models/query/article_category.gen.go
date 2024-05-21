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

func newArticleCategory(db *gorm.DB, opts ...gen.DOOption) articleCategory {
	_articleCategory := articleCategory{}

	_articleCategory.articleCategoryDo.UseDB(db, opts...)
	_articleCategory.articleCategoryDo.UseModel(&model.ArticleCategory{})

	tableName := _articleCategory.articleCategoryDo.TableName()
	_articleCategory.ALL = field.NewAsterisk(tableName)
	_articleCategory.ID = field.NewUint32(tableName, "id")
	_articleCategory.Fid = field.NewUint32(tableName, "fid")
	_articleCategory.Name = field.NewString(tableName, "name")
	_articleCategory.Remark = field.NewString(tableName, "remark")
	_articleCategory.Sort = field.NewUint32(tableName, "sort")
	_articleCategory.Status = field.NewUint32(tableName, "status")
	_articleCategory.CreatedAt = field.NewTime(tableName, "created_at")
	_articleCategory.UpdatedAt = field.NewTime(tableName, "updated_at")

	_articleCategory.fillFieldMap()

	return _articleCategory
}

// articleCategory 文章分类表
type articleCategory struct {
	articleCategoryDo articleCategoryDo

	ALL       field.Asterisk
	ID        field.Uint32
	Fid       field.Uint32 // 父级ID
	Name      field.String // 分类名称
	Remark    field.String // 备注信息
	Sort      field.Uint32 // 排序顺序
	Status    field.Uint32 // 显示状态（1=显示，2=隐藏）
	CreatedAt field.Time
	UpdatedAt field.Time

	fieldMap map[string]field.Expr
}

func (a articleCategory) Table(newTableName string) *articleCategory {
	a.articleCategoryDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a articleCategory) As(alias string) *articleCategory {
	a.articleCategoryDo.DO = *(a.articleCategoryDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *articleCategory) updateTableName(table string) *articleCategory {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewUint32(table, "id")
	a.Fid = field.NewUint32(table, "fid")
	a.Name = field.NewString(table, "name")
	a.Remark = field.NewString(table, "remark")
	a.Sort = field.NewUint32(table, "sort")
	a.Status = field.NewUint32(table, "status")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")

	a.fillFieldMap()

	return a
}

func (a *articleCategory) WithContext(ctx context.Context) *articleCategoryDo {
	return a.articleCategoryDo.WithContext(ctx)
}

func (a articleCategory) TableName() string { return a.articleCategoryDo.TableName() }

func (a articleCategory) Alias() string { return a.articleCategoryDo.Alias() }

func (a articleCategory) Columns(cols ...field.Expr) gen.Columns {
	return a.articleCategoryDo.Columns(cols...)
}

func (a *articleCategory) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *articleCategory) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 8)
	a.fieldMap["id"] = a.ID
	a.fieldMap["fid"] = a.Fid
	a.fieldMap["name"] = a.Name
	a.fieldMap["remark"] = a.Remark
	a.fieldMap["sort"] = a.Sort
	a.fieldMap["status"] = a.Status
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
}

func (a articleCategory) clone(db *gorm.DB) articleCategory {
	a.articleCategoryDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a articleCategory) replaceDB(db *gorm.DB) articleCategory {
	a.articleCategoryDo.ReplaceDB(db)
	return a
}

type articleCategoryDo struct{ gen.DO }

func (a articleCategoryDo) Debug() *articleCategoryDo {
	return a.withDO(a.DO.Debug())
}

func (a articleCategoryDo) WithContext(ctx context.Context) *articleCategoryDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a articleCategoryDo) ReadDB() *articleCategoryDo {
	return a.Clauses(dbresolver.Read)
}

func (a articleCategoryDo) WriteDB() *articleCategoryDo {
	return a.Clauses(dbresolver.Write)
}

func (a articleCategoryDo) Session(config *gorm.Session) *articleCategoryDo {
	return a.withDO(a.DO.Session(config))
}

func (a articleCategoryDo) Clauses(conds ...clause.Expression) *articleCategoryDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a articleCategoryDo) Returning(value interface{}, columns ...string) *articleCategoryDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a articleCategoryDo) Not(conds ...gen.Condition) *articleCategoryDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a articleCategoryDo) Or(conds ...gen.Condition) *articleCategoryDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a articleCategoryDo) Select(conds ...field.Expr) *articleCategoryDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a articleCategoryDo) Where(conds ...gen.Condition) *articleCategoryDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a articleCategoryDo) Order(conds ...field.Expr) *articleCategoryDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a articleCategoryDo) Distinct(cols ...field.Expr) *articleCategoryDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a articleCategoryDo) Omit(cols ...field.Expr) *articleCategoryDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a articleCategoryDo) Join(table schema.Tabler, on ...field.Expr) *articleCategoryDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a articleCategoryDo) LeftJoin(table schema.Tabler, on ...field.Expr) *articleCategoryDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a articleCategoryDo) RightJoin(table schema.Tabler, on ...field.Expr) *articleCategoryDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a articleCategoryDo) Group(cols ...field.Expr) *articleCategoryDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a articleCategoryDo) Having(conds ...gen.Condition) *articleCategoryDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a articleCategoryDo) Limit(limit int) *articleCategoryDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a articleCategoryDo) Offset(offset int) *articleCategoryDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a articleCategoryDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *articleCategoryDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a articleCategoryDo) Unscoped() *articleCategoryDo {
	return a.withDO(a.DO.Unscoped())
}

func (a articleCategoryDo) Create(values ...*model.ArticleCategory) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a articleCategoryDo) CreateInBatches(values []*model.ArticleCategory, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a articleCategoryDo) Save(values ...*model.ArticleCategory) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a articleCategoryDo) First() (*model.ArticleCategory, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArticleCategory), nil
	}
}

func (a articleCategoryDo) Take() (*model.ArticleCategory, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArticleCategory), nil
	}
}

func (a articleCategoryDo) Last() (*model.ArticleCategory, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArticleCategory), nil
	}
}

func (a articleCategoryDo) Find() ([]*model.ArticleCategory, error) {
	result, err := a.DO.Find()
	return result.([]*model.ArticleCategory), err
}

func (a articleCategoryDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ArticleCategory, err error) {
	buf := make([]*model.ArticleCategory, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a articleCategoryDo) FindInBatches(result *[]*model.ArticleCategory, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a articleCategoryDo) Attrs(attrs ...field.AssignExpr) *articleCategoryDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a articleCategoryDo) Assign(attrs ...field.AssignExpr) *articleCategoryDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a articleCategoryDo) Joins(fields ...field.RelationField) *articleCategoryDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a articleCategoryDo) Preload(fields ...field.RelationField) *articleCategoryDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a articleCategoryDo) FirstOrInit() (*model.ArticleCategory, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArticleCategory), nil
	}
}

func (a articleCategoryDo) FirstOrCreate() (*model.ArticleCategory, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArticleCategory), nil
	}
}

func (a articleCategoryDo) FindByPage(offset int, limit int) (result []*model.ArticleCategory, count int64, err error) {
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

func (a articleCategoryDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a articleCategoryDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a articleCategoryDo) Delete(models ...*model.ArticleCategory) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *articleCategoryDo) withDO(do gen.Dao) *articleCategoryDo {
	a.DO = *do.(*gen.DO)
	return a
}
