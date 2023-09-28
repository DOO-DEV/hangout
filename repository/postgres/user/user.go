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

func (d *DB) GetPrimaryProfileImage(ctx context.Context, userID string) (string, error) {
	const op = "UserRepository.GetPrimaryProfileImage"

	row := d.conn.Conn().QueryRowContext(ctx, `select "url" from "account_images" where "account" = $1 and "is_primary" = 'true'`, userID)
	var imageUrl string
	if err := row.Scan(&imageUrl); err != nil {
		fmt.Println(err)
		if d.conn.IsEmptyRowError(err) {
			return imageUrl, richerror.New(op).WithError(err).WithKind(richerror.KindNotFound).WithMessage(errmsg.ErrorMsgNotFound)
		}
		return imageUrl, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}

	return imageUrl, nil
}

func (d *DB) GetAllProfileImages(ctx context.Context, userID string) ([]string, error) {
	const op = "UserRepository.GetAllProfileImages"

	rows, err := d.conn.Conn().QueryContext(ctx, `select "url" from "account_images" where "account" = $1`, userID)
	if err != nil {
		return nil, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}

	urls := make([]string, 0)
	for rows.Next() {
		var u string
		if err := rows.Scan(&u); err != nil {
			return nil, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
		}
		urls = append(urls, u)
	}

	return urls, nil
}

func (d *DB) DeleteProfileImage(ctx context.Context, userID, imgID string) (string, error) {
	const op = "UserRepository.DeleteProfileImage"

	tx, err := d.conn.Conn().BeginTx(ctx, nil)
	if err != nil {
		fmt.Println("tx", err)
		return "", richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	row := tx.QueryRowContext(ctx, `delete from "account_images" where "account" = $1 and "id" = $2 RETURNING "url", "is_primary", "order"`, userID, imgID)
	var url string
	var order int
	var isPrimary bool
	if err = row.Scan(&url, &isPrimary, &order); err != nil {
		if d.conn.IsEmptyRowError(err) {
			return url, richerror.New(op).WithError(err).WithKind(richerror.KindNotFound).WithMessage(errmsg.ErrorMsgNotFound)
		}
		return url, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}

	// update the order of profile images
	_, err = tx.ExecContext(ctx, `update "account_images" set "order" = "order" - 1 where "order" > $1 and "account" = $1`, order, userID)
	if err != nil {
		return url, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}

	// update primary
	if isPrimary {
		_, err = tx.ExecContext(ctx, `update "account_images" set "is_primary" = 'true' where "order" = 1 and "account" = $1`, userID)
		if err != nil {
			return url, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
		}
	}

	return url, nil
}

func (d *DB) SetImageAsPrimary(ctx context.Context, userID, imgID string) error {
	const op = "UserRepository.SetImageAsPrimary"

	tx, err := d.conn.Conn().BeginTx(ctx, nil)
	if err != nil {
		return richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	// first update the new primary image to 'true' and set the past image to 'false'
	_, err = tx.ExecContext(ctx, `update "account_images" set "is_primary" = 'false' where "account" = $1 and "is_primary" = 'true'`, userID)
	if err != nil {
		return richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}
	row := tx.QueryRowContext(ctx, `update "account_images" set "is_primary" = 'true' where "id" = $1 RETURNING "order"`, imgID)
	if err != nil {
		return richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}
	var order int
	if err = row.Scan(&order); err != nil {
		if d.conn.IsEmptyRowError(err) {
			return richerror.New(op).WithError(err).WithKind(richerror.KindNotFound).WithMessage(errmsg.ErrorMsgNotFound)
		}
		return richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}
	// reorder the images
	_, err = tx.ExecContext(ctx, `update "account_images" set "order" = "order" + 1 where "account" = $1 and "order" < $2`, userID, order)
	if err != nil {
		return richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}
	_, err = tx.ExecContext(ctx, `update "account_images" set "order" = 1 where "account" = $1 and "is_primary" = 'true'`, userID)
	if err != nil {
		return richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}

	return nil
}
