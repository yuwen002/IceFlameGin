// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                          db,
		ArticleCategory:             newArticleCategory(db, opts...),
		ArticleChannel:              newArticleChannel(db, opts...),
		ArticleTag:                  newArticleTag(db, opts...),
		ArticleTagOrm:               newArticleTagOrm(db, opts...),
		SinglePage:                  newSinglePage(db, opts...),
		UcAccount:                   newUcAccount(db, opts...),
		UcSystemMaster:              newUcSystemMaster(db, opts...),
		UcSystemMasterRole:          newUcSystemMasterRole(db, opts...),
		UcSystemMasterRoleRelation:  newUcSystemMasterRoleRelation(db, opts...),
		UcSystemMasterVisitCategory: newUcSystemMasterVisitCategory(db, opts...),
		UcSystemMasterVisitorLog:    newUcSystemMasterVisitorLog(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	ArticleCategory             articleCategory
	ArticleChannel              articleChannel
	ArticleTag                  articleTag
	ArticleTagOrm               articleTagOrm
	SinglePage                  singlePage
	UcAccount                   ucAccount
	UcSystemMaster              ucSystemMaster
	UcSystemMasterRole          ucSystemMasterRole
	UcSystemMasterRoleRelation  ucSystemMasterRoleRelation
	UcSystemMasterVisitCategory ucSystemMasterVisitCategory
	UcSystemMasterVisitorLog    ucSystemMasterVisitorLog
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                          db,
		ArticleCategory:             q.ArticleCategory.clone(db),
		ArticleChannel:              q.ArticleChannel.clone(db),
		ArticleTag:                  q.ArticleTag.clone(db),
		ArticleTagOrm:               q.ArticleTagOrm.clone(db),
		SinglePage:                  q.SinglePage.clone(db),
		UcAccount:                   q.UcAccount.clone(db),
		UcSystemMaster:              q.UcSystemMaster.clone(db),
		UcSystemMasterRole:          q.UcSystemMasterRole.clone(db),
		UcSystemMasterRoleRelation:  q.UcSystemMasterRoleRelation.clone(db),
		UcSystemMasterVisitCategory: q.UcSystemMasterVisitCategory.clone(db),
		UcSystemMasterVisitorLog:    q.UcSystemMasterVisitorLog.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:                          db,
		ArticleCategory:             q.ArticleCategory.replaceDB(db),
		ArticleChannel:              q.ArticleChannel.replaceDB(db),
		ArticleTag:                  q.ArticleTag.replaceDB(db),
		ArticleTagOrm:               q.ArticleTagOrm.replaceDB(db),
		SinglePage:                  q.SinglePage.replaceDB(db),
		UcAccount:                   q.UcAccount.replaceDB(db),
		UcSystemMaster:              q.UcSystemMaster.replaceDB(db),
		UcSystemMasterRole:          q.UcSystemMasterRole.replaceDB(db),
		UcSystemMasterRoleRelation:  q.UcSystemMasterRoleRelation.replaceDB(db),
		UcSystemMasterVisitCategory: q.UcSystemMasterVisitCategory.replaceDB(db),
		UcSystemMasterVisitorLog:    q.UcSystemMasterVisitorLog.replaceDB(db),
	}
}

type queryCtx struct {
	ArticleCategory             *articleCategoryDo
	ArticleChannel              *articleChannelDo
	ArticleTag                  *articleTagDo
	ArticleTagOrm               *articleTagOrmDo
	SinglePage                  *singlePageDo
	UcAccount                   *ucAccountDo
	UcSystemMaster              *ucSystemMasterDo
	UcSystemMasterRole          *ucSystemMasterRoleDo
	UcSystemMasterRoleRelation  *ucSystemMasterRoleRelationDo
	UcSystemMasterVisitCategory *ucSystemMasterVisitCategoryDo
	UcSystemMasterVisitorLog    *ucSystemMasterVisitorLogDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		ArticleCategory:             q.ArticleCategory.WithContext(ctx),
		ArticleChannel:              q.ArticleChannel.WithContext(ctx),
		ArticleTag:                  q.ArticleTag.WithContext(ctx),
		ArticleTagOrm:               q.ArticleTagOrm.WithContext(ctx),
		SinglePage:                  q.SinglePage.WithContext(ctx),
		UcAccount:                   q.UcAccount.WithContext(ctx),
		UcSystemMaster:              q.UcSystemMaster.WithContext(ctx),
		UcSystemMasterRole:          q.UcSystemMasterRole.WithContext(ctx),
		UcSystemMasterRoleRelation:  q.UcSystemMasterRoleRelation.WithContext(ctx),
		UcSystemMasterVisitCategory: q.UcSystemMasterVisitCategory.WithContext(ctx),
		UcSystemMasterVisitorLog:    q.UcSystemMasterVisitorLog.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
