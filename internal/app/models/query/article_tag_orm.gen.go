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

func newArticleTagOrm(db *gorm.DB, opts ...gen.DOOption) articleTagOrm {
	_articleTagOrm := articleTagOrm{}

	_articleTagOrm.articleTagOrmDo.UseDB(db, opts...)
	_articleTagOrm.articleTagOrmDo.UseModel(&model.ArticleTagOrm{})

	tableName := _articleTagOrm.articleTagOrmDo.TableName()
	_articleTagOrm.ALL = field.NewAsterisk(tableName)
	_articleTagOrm.ID = field.NewUint32(tableName, "id")
	_articleTagOrm.TagID = field.NewInt32(tableName, "tag_id")
	_articleTagOrm.ArticleID = field.NewInt32(tableName, "article_id")
	_articleTagOrm.CreatedAt = field.NewTime(tableName, "created_at")
	_articleTagOrm.UpdatedAt = field.NewTime(tableName, "updated_at")

	_articleTagOrm.fillFieldMap()

	return _articleTagOrm
}

// articleTagOrm 文章tags
type articleTagOrm struct {
	articleTagOrmDo articleTagOrmDo

	ALL       field.Asterisk
	ID        field.Uint32
	TagID     field.Int32 // tag ID
	ArticleID field.Int32 // 备注
	CreatedAt field.Time
	UpdatedAt field.Time

	fieldMap map[string]field.Expr
}

func (a articleTagOrm) Table(newTableName string) *articleTagOrm {
	a.articleTagOrmDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a articleTagOrm) As(alias string) *articleTagOrm {
	a.articleTagOrmDo.DO = *(a.articleTagOrmDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *articleTagOrm) updateTableName(table string) *articleTagOrm {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewUint32(table, "id")
	a.TagID = field.NewInt32(table, "tag_id")
	a.ArticleID = field.NewInt32(table, "article_id")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")

	a.fillFieldMap()

	return a
}

func (a *articleTagOrm) WithContext(ctx context.Context) *articleTagOrmDo {
	return a.articleTagOrmDo.WithContext(ctx)
}

func (a articleTagOrm) TableName() string { return a.articleTagOrmDo.TableName() }

func (a articleTagOrm) Alias() string { return a.articleTagOrmDo.Alias() }

func (a articleTagOrm) Columns(cols ...field.Expr) gen.Columns {
	return a.articleTagOrmDo.Columns(cols...)
}

func (a *articleTagOrm) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *articleTagOrm) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 5)
	a.fieldMap["id"] = a.ID
	a.fieldMap["tag_id"] = a.TagID
	a.fieldMap["article_id"] = a.ArticleID
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
}

func (a articleTagOrm) clone(db *gorm.DB) articleTagOrm {
	a.articleTagOrmDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a articleTagOrm) replaceDB(db *gorm.DB) articleTagOrm {
	a.articleTagOrmDo.ReplaceDB(db)
	return a
}

type articleTagOrmDo struct{ gen.DO }

func (a articleTagOrmDo) Debug() *articleTagOrmDo {
	return a.withDO(a.DO.Debug())
}

func (a articleTagOrmDo) WithContext(ctx context.Context) *articleTagOrmDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a articleTagOrmDo) ReadDB() *articleTagOrmDo {
	return a.Clauses(dbresolver.Read)
}

func (a articleTagOrmDo) WriteDB() *articleTagOrmDo {
	return a.Clauses(dbresolver.Write)
}

func (a articleTagOrmDo) Session(config *gorm.Session) *articleTagOrmDo {
	return a.withDO(a.DO.Session(config))
}

func (a articleTagOrmDo) Clauses(conds ...clause.Expression) *articleTagOrmDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a articleTagOrmDo) Returning(value interface{}, columns ...string) *articleTagOrmDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a articleTagOrmDo) Not(conds ...gen.Condition) *articleTagOrmDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a articleTagOrmDo) Or(conds ...gen.Condition) *articleTagOrmDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a articleTagOrmDo) Select(conds ...field.Expr) *articleTagOrmDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a articleTagOrmDo) Where(conds ...gen.Condition) *articleTagOrmDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a articleTagOrmDo) Order(conds ...field.Expr) *articleTagOrmDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a articleTagOrmDo) Distinct(cols ...field.Expr) *articleTagOrmDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a articleTagOrmDo) Omit(cols ...field.Expr) *articleTagOrmDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a articleTagOrmDo) Join(table schema.Tabler, on ...field.Expr) *articleTagOrmDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a articleTagOrmDo) LeftJoin(table schema.Tabler, on ...field.Expr) *articleTagOrmDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a articleTagOrmDo) RightJoin(table schema.Tabler, on ...field.Expr) *articleTagOrmDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a articleTagOrmDo) Group(cols ...field.Expr) *articleTagOrmDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a articleTagOrmDo) Having(conds ...gen.Condition) *articleTagOrmDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a articleTagOrmDo) Limit(limit int) *articleTagOrmDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a articleTagOrmDo) Offset(offset int) *articleTagOrmDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a articleTagOrmDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *articleTagOrmDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a articleTagOrmDo) Unscoped() *articleTagOrmDo {
	return a.withDO(a.DO.Unscoped())
}

func (a articleTagOrmDo) Create(values ...*model.ArticleTagOrm) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a articleTagOrmDo) CreateInBatches(values []*model.ArticleTagOrm, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a articleTagOrmDo) Save(values ...*model.ArticleTagOrm) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a articleTagOrmDo) First() (*model.ArticleTagOrm, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArticleTagOrm), nil
	}
}

func (a articleTagOrmDo) Take() (*model.ArticleTagOrm, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArticleTagOrm), nil
	}
}

func (a articleTagOrmDo) Last() (*model.ArticleTagOrm, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArticleTagOrm), nil
	}
}

func (a articleTagOrmDo) Find() ([]*model.ArticleTagOrm, error) {
	result, err := a.DO.Find()
	return result.([]*model.ArticleTagOrm), err
}

func (a articleTagOrmDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ArticleTagOrm, err error) {
	buf := make([]*model.ArticleTagOrm, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a articleTagOrmDo) FindInBatches(result *[]*model.ArticleTagOrm, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a articleTagOrmDo) Attrs(attrs ...field.AssignExpr) *articleTagOrmDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a articleTagOrmDo) Assign(attrs ...field.AssignExpr) *articleTagOrmDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a articleTagOrmDo) Joins(fields ...field.RelationField) *articleTagOrmDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a articleTagOrmDo) Preload(fields ...field.RelationField) *articleTagOrmDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a articleTagOrmDo) FirstOrInit() (*model.ArticleTagOrm, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArticleTagOrm), nil
	}
}

func (a articleTagOrmDo) FirstOrCreate() (*model.ArticleTagOrm, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArticleTagOrm), nil
	}
}

func (a articleTagOrmDo) FindByPage(offset int, limit int) (result []*model.ArticleTagOrm, count int64, err error) {
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

func (a articleTagOrmDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a articleTagOrmDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a articleTagOrmDo) Delete(models ...*model.ArticleTagOrm) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *articleTagOrmDo) withDO(do gen.Dao) *articleTagOrmDo {
	a.DO = *do.(*gen.DO)
	return a
}
