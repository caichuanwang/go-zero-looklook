// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/caichuanwang/go-zero-looklook/common/globalkey"
	"github.com/caichuanwang/go-zero-looklook/common/xerr"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userFieldNames          = builder.RawFieldNames(&User{})
	userRows                = strings.Join(userFieldNames, ",")
	userRowsExpectAutoSet   = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	userRowsWithPlaceHolder = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	userModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *User) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*User, error)
		FindOneByMobile(ctx context.Context, mobile string) (*User, error)
		Update(ctx context.Context, session sqlx.Session, data *User) (sql.Result, error)
		UpdateWithVersion(ctx context.Context, session sqlx.Session, data *User) error
		Delete(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultUserModel struct {
		conn  sqlx.SqlConn
		table string
	}

	User struct {
		Id         int64     `db:"id"`
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
		DeleteTime time.Time `db:"delete_time"`
		DelState   int64     `db:"del_state"`
		Version    int64     `db:"version"` // 版本号
		Mobile     string    `db:"mobile"`
		Password   string    `db:"password"`
		Nickname   string    `db:"nickname"`
		Sex        int64     `db:"sex"` // 性别 0:男 1:女
		Avatar     string    `db:"avatar"`
		Info       string    `db:"info"`
	}
)

func newUserModel(conn sqlx.SqlConn) *defaultUserModel {
	return &defaultUserModel{
		conn:  conn,
		table: "`user`",
	}
}

func (m *defaultUserModel) Insert(ctx context.Context, session sqlx.Session, data *User) (sql.Result, error) {
	data.DeleteTime = time.Unix(0, 0)

	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, userRowsExpectAutoSet)
	if session != nil {
		return session.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.Version, data.Mobile, data.Password, data.Nickname, data.Sex, data.Avatar, data.Info)
	}
	return m.conn.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.Version, data.Mobile, data.Password, data.Nickname, data.Sex, data.Avatar, data.Info)
}

func (m *defaultUserModel) FindOne(ctx context.Context, id int64) (*User, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ?  and del_state = ? limit 1", userRows, m.table)
	var resp User
	err := m.conn.QueryRowCtx(ctx, &resp, query, id, globalkey.DelStateNo)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) FindOneByMobile(ctx context.Context, mobile string) (*User, error) {
	var resp User
	query := fmt.Sprintf("select %s from %s where `mobile` = ?  and del_state = ? limit 1", userRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, mobile, globalkey.DelStateNo)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) Update(ctx context.Context, session sqlx.Session, data *User) (sql.Result, error) {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userRowsWithPlaceHolder)
	if session != nil {
		return session.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.Version, data.Mobile, data.Password, data.Nickname, data.Sex, data.Avatar, data.Info, data.Id)
	}
	return m.conn.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.Version, data.Mobile, data.Password, data.Nickname, data.Sex, data.Avatar, data.Info, data.Id)
}

func (m *defaultUserModel) UpdateWithVersion(ctx context.Context, session sqlx.Session, data *User) error {

	oldVersion := data.Version
	data.Version += 1

	var sqlResult sql.Result
	var err error

	query := fmt.Sprintf("update %s set %s where `id` = ? and version = ? ", m.table, userRowsWithPlaceHolder)
	if session != nil {
		sqlResult, err = session.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.Version, data.Mobile, data.Password, data.Nickname, data.Sex, data.Avatar, data.Info, data.Id, oldVersion)
	} else {
		sqlResult, err = m.conn.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.Version, data.Mobile, data.Password, data.Nickname, data.Sex, data.Avatar, data.Info, data.Id, oldVersion)
	}

	if err != nil {
		return err
	}
	updateCount, err := sqlResult.RowsAffected()
	if err != nil {
		return err
	}
	if updateCount == 0 {
		return xerr.NewErrCode(xerr.DB_UPDATE_AFFECTED_ZERO_ERROR)
	}

	return nil
}

func (m *defaultUserModel) Delete(ctx context.Context, session sqlx.Session, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	if session != nil {
		_, err := session.ExecCtx(ctx, query, id)
		return err
	}
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultUserModel) tableName() string {
	return m.table
}
