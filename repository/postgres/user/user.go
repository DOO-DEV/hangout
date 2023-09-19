package pguser

import (
	"context"
	"database/sql"
	"errors"
	"hangout/entity"
	customerr "hangout/pkg/error"
)

func (d *DB) Register(ctx context.Context, user *entity.User) error {
	_, err := d.conn.Conn().ExecContext(ctx,
		`insert into users("id", "username", "password", "first_name", "last_name") values($1, $2, $3, $4, $5)`,
		user.ID, user.Username, user.Password, user.FirsName, user.LastName)
	if err != nil {
		return err
	}

	return nil
}

func (d *DB) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	res := d.conn.Conn().QueryRowContext(ctx, "select * from users where username = $1", username)
	u := &entity.User{}

	if err := res.Scan(&u.ID, &u.FirsName, &u.LastName, &u.Password, &u.Username, &u.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, customerr.RecordNotFoundErr
		}
		return nil, err
	}

	return u, nil
}
