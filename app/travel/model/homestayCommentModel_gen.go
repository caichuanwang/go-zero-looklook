// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/caichuanwang/go-zero-looklook/deploy/script/mysql/genModel"
	"strings"

	"time"

	"github.com/Masterminds/squirrel"
	"github.com/caichuanwang/go-zero-looklook/common/globalkey"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	homestayCommentFieldNames          = builder.RawFieldNames(&HomestayComment{})
	homestayCommentRows                = strings.Join(homestayCommentFieldNames, ",")
	homestayCommentRowsExpectAutoSet   = strings.Join(stringx.Remove(homestayCommentFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	homestayCommentRowsWithPlaceHolder = strings.Join(stringx.Remove(homestayCommentFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheLooklookTravelHomestayCommentIdPrefix = "cache:looklookTravel:homestayComment:id:"
)

type (
	homestayCommentModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *HomestayComment) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*HomestayComment, error)
		Update(ctx context.Context, session sqlx.Session, data *HomestayComment) (sql.Result, error)
		UpdateWithVersion(ctx context.Context, session sqlx.Session, data *HomestayComment) error
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		SelectBuilder() squirrel.SelectBuilder
		DeleteSoft(ctx context.Context, session sqlx.Session, data *HomestayComment) error
		FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder, field string) (float64, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder, field string) (int64, error)
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*HomestayComment, error)
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*HomestayComment, error)
		FindPageListByPageWithTotal(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*HomestayComment, int64, error)
		FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*HomestayComment, error)
		FindPageListByIdASC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*HomestayComment, error)
		Delete(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultHomestayCommentModel struct {
		sqlc.CachedConn
		table string
	}

	HomestayComment struct {
		Id         int64     `db:"id"`
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
		DeleteTime time.Time `db:"delete_time"`
		DelState   int64     `db:"del_state"`
		HomestayId int64     `db:"homestay_id"` // 民宿id
		UserId     int64     `db:"user_id"`     // 用户id
		Content    string    `db:"content"`     // 评论内容
		Star       string    `db:"star"`        // 星星数,多个维度
		Version    int64     `db:"version"`     // 版本号
	}
)

func newHomestayCommentModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultHomestayCommentModel {
	return &defaultHomestayCommentModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`homestay_comment`",
	}
}

func (m *defaultHomestayCommentModel) Insert(ctx context.Context, session sqlx.Session, data *HomestayComment) (sql.Result, error) {
	data.DeleteTime = time.Unix(0, 0)
	data.DelState = globalkey.DelStateNo
	looklookTravelHomestayCommentIdKey := fmt.Sprintf("%s%v", cacheLooklookTravelHomestayCommentIdPrefix, data.Id)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, homestayCommentRowsExpectAutoSet)
		if session != nil {
			return session.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.HomestayId, data.UserId, data.Content, data.Star, data.Version)
		}
		return conn.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.HomestayId, data.UserId, data.Content, data.Star, data.Version)
	}, looklookTravelHomestayCommentIdKey)
}

func (m *defaultHomestayCommentModel) FindOne(ctx context.Context, id int64) (*HomestayComment, error) {
	looklookTravelHomestayCommentIdKey := fmt.Sprintf("%s%v", cacheLooklookTravelHomestayCommentIdPrefix, id)
	var resp HomestayComment
	err := m.QueryRowCtx(ctx, &resp, looklookTravelHomestayCommentIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? and del_state = ? limit 1", homestayCommentRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id, globalkey.DelStateNo)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, genModel.ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultHomestayCommentModel) Update(ctx context.Context, session sqlx.Session, data *HomestayComment) (sql.Result, error) {
	looklookTravelHomestayCommentIdKey := fmt.Sprintf("%s%v", cacheLooklookTravelHomestayCommentIdPrefix, data.Id)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, homestayCommentRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.HomestayId, data.UserId, data.Content, data.Star, data.Version, data.Id)
		}
		return conn.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.HomestayId, data.UserId, data.Content, data.Star, data.Version, data.Id)
	}, looklookTravelHomestayCommentIdKey)
}

func (m *defaultHomestayCommentModel) UpdateWithVersion(ctx context.Context, session sqlx.Session, data *HomestayComment) error {

	oldVersion := data.Version
	data.Version += 1

	var sqlResult sql.Result
	var err error

	looklookTravelHomestayCommentIdKey := fmt.Sprintf("%s%v", cacheLooklookTravelHomestayCommentIdPrefix, data.Id)
	sqlResult, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ? and version = ? ", m.table, homestayCommentRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.HomestayId, data.UserId, data.Content, data.Star, data.Version, data.Id, oldVersion)
		}
		return conn.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.HomestayId, data.UserId, data.Content, data.Star, data.Version, data.Id, oldVersion)
	}, looklookTravelHomestayCommentIdKey)
	if err != nil {
		return err
	}
	updateCount, err := sqlResult.RowsAffected()
	if err != nil {
		return err
	}
	if updateCount == 0 {
		return genModel.ErrNoRowsUpdate
	}

	return nil
}

func (m *defaultHomestayCommentModel) DeleteSoft(ctx context.Context, session sqlx.Session, data *HomestayComment) error {
	data.DelState = globalkey.DelStateYes
	data.DeleteTime = time.Now()
	if err := m.UpdateWithVersion(ctx, session, data); err != nil {
		return errors.Wrapf(errors.New("delete soft failed "), "HomestayCommentModel delete err : %+v", err)
	}
	return nil
}

func (m *defaultHomestayCommentModel) FindSum(ctx context.Context, builder squirrel.SelectBuilder, field string) (float64, error) {

	if len(field) == 0 {
		return 0, errors.Wrapf(errors.New("FindSum Least One Field"), "FindSum Least One Field")
	}

	builder = builder.Columns("IFNULL(SUM(" + field + "),0)")

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return 0, err
	}

	var resp float64
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *defaultHomestayCommentModel) FindCount(ctx context.Context, builder squirrel.SelectBuilder, field string) (int64, error) {

	if len(field) == 0 {
		return 0, errors.Wrapf(errors.New("FindCount Least One Field"), "FindCount Least One Field")
	}

	builder = builder.Columns("COUNT(" + field + ")")

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return 0, err
	}

	var resp int64
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *defaultHomestayCommentModel) FindAll(ctx context.Context, builder squirrel.SelectBuilder, orderBy string) ([]*HomestayComment, error) {

	builder = builder.Columns(homestayCommentRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*HomestayComment
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultHomestayCommentModel) FindPageListByPage(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*HomestayComment, error) {

	builder = builder.Columns(homestayCommentRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*HomestayComment
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultHomestayCommentModel) FindPageListByPageWithTotal(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*HomestayComment, int64, error) {

	total, err := m.FindCount(ctx, builder, "id")
	if err != nil {
		return nil, 0, err
	}

	builder = builder.Columns(homestayCommentRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, total, err
	}

	var resp []*HomestayComment
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, total, nil
	default:
		return nil, total, err
	}
}

func (m *defaultHomestayCommentModel) FindPageListByIdDESC(ctx context.Context, builder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*HomestayComment, error) {

	builder = builder.Columns(homestayCommentRows)

	if preMinId > 0 {
		builder = builder.Where(" id < ? ", preMinId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id DESC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*HomestayComment
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultHomestayCommentModel) FindPageListByIdASC(ctx context.Context, builder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*HomestayComment, error) {

	builder = builder.Columns(homestayCommentRows)

	if preMaxId > 0 {
		builder = builder.Where(" id > ? ", preMaxId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id ASC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*HomestayComment
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultHomestayCommentModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {

	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})

}

func (m *defaultHomestayCommentModel) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(m.table)
}
func (m *defaultHomestayCommentModel) Delete(ctx context.Context, session sqlx.Session, id int64) error {
	looklookTravelHomestayCommentIdKey := fmt.Sprintf("%s%v", cacheLooklookTravelHomestayCommentIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		if session != nil {
			return session.ExecCtx(ctx, query, id)
		}
		return conn.ExecCtx(ctx, query, id)
	}, looklookTravelHomestayCommentIdKey)
	return err
}
func (m *defaultHomestayCommentModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheLooklookTravelHomestayCommentIdPrefix, primary)
}
func (m *defaultHomestayCommentModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? and del_state = ? limit 1", homestayCommentRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary, globalkey.DelStateNo)
}

func (m *defaultHomestayCommentModel) tableName() string {
	return m.table
}
