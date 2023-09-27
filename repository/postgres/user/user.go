package pguser

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"hangout/entity"
	"hangout/pkg/errmsg"
	"hangout/pkg/richerror"
)

func (d *DB) Register(ctx context.Context, user *entity.User) error {
	const op = "UserRepository.Register"

	_, err := d.conn.Conn().ExecContext(ctx,
		`insert into users("id", "username", "password", "first_name", "last_name") values($1, $2, $3, $4, $5)`,
		user.ID, user.Username, user.Password, user.FirsName, user.LastName)
	if err != nil {
		return richerror.New(op).WithKind(richerror.KindUnexpected).WithError(err).WithMessage("something went wrong")
	}

	return nil
}

func (d *DB) IsUserExists(ctx context.Context, username string) (bool, error) {
	const op = "UserRepository.IsUserExists"

	res := d.conn.Conn().QueryRowContext(ctx, "select * from users where username = $1", username)
	u := &entity.User{}

	if err := res.Scan(&u.ID, &u.FirsName, &u.LastName, &u.Password, &u.Username, &u.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return true, richerror.New(op).WithKind(richerror.KindUnexpected).WithError(err).WithMessage("something went wrong")
	}

	return true, nil
}

func (d *DB) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	const op = "UserRepository.GetUserByUsername"

	res := d.conn.Conn().QueryRowContext(ctx, "select * from users where username = $1", username)
	u := &entity.User{}

	if err := res.Scan(&u.ID, &u.FirsName, &u.LastName, &u.Password, &u.Username, &u.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, richerror.New(op).WithKind(richerror.KindNotFound).WithError(err).WithMessage("user not found")
		}
		return nil, richerror.New(op).WithKind(richerror.KindUnexpected).WithError(err).WithMessage("something went wrong")
	}

	return u, nil
}

func (d *DB) GetUserByID(ctx context.Context, userID string) (*entity.User, error) {
	const op = "UserRepository.GetUserById"

	res := d.conn.Conn().QueryRowContext(ctx, `select * from users where "id" = $1`, userID)
	u := &entity.User{}

	if err := res.Scan(&u.ID, &u.FirsName, &u.LastName, &u.Password, &u.Username, &u.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, richerror.New(op).WithKind(richerror.KindNotFound).WithError(err).WithMessage("user not found")
		}
		return nil, richerror.New(op).WithKind(richerror.KindUnexpected).WithError(err).WithMessage("something went wrong")
	}

	return u, nil
}

func (d *DB) GetLastInsertedProfileImageOrder(ctx context.Context, userID string) (sql.NullInt32, error) {
	const op = "UserRepository.GetLastInsertedProfileImageOrder"

	row := d.conn.Conn().QueryRowContext(ctx, `select MAX("order") from "account_images" where "account" = $1`, userID)
	var order sql.NullInt32
	if err := row.Scan(&order); err != nil {
		fmt.Println(err)
		if d.conn.IsEmptyRowError(err) {
			return order, nil
		}
		return order, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}

	return order, nil
}

func (d *DB) SaveProfileImageInfo(ctx context.Context, imageUrl, useID string) error {
	const op = "UserRepository.SaveProfileImageInfo"

	order, err := d.GetLastInsertedProfileImageOrder(ctx, useID)
	if err != nil {
		return richerror.New(op).WithError(err)
	}

	isPrimary := false
	order.Int32 += 1
	if order.Int32 == 1 {
		isPrimary = true
	}
	_, err = d.conn.Conn().ExecContext(ctx, `insert into "account_images" values ($1, $2, $3, $4, $5)`, uuid.NewString(), useID, imageUrl, isPrimary, order.Int32)
	if err != nil {
		return richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}

	return nil
}
