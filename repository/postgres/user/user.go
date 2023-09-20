package pguser

import (
	"context"
	"database/sql"
	"errors"
	"hangout/entity"
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
