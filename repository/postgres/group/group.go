package pggroup

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"hangout/entity"
	"hangout/pkg/richerror"
	"time"
)

func (d DB) CheckUserGroup(ctx context.Context, userID string) (bool, error) {
	const op = "GroupRepository.CheckUserGroup"

	row := d.conn.Conn().QueryRowContext(ctx, "select * from users_group where user_id = $1", userID)
	u := entity.User{}
	if err := row.Scan(&u.ID, &u.FirsName, &u.LastName, &u.Password, &u.Username, &u.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return true, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected)
	}

	return true, nil
}

func (d DB) CreateGroup(ctx context.Context, group entity.Group) error {
	const op = "GroupService.CreateGroup"

	_, err := d.conn.Conn().ExecContext(ctx, `insert into groups(id, name, owner_id) values($1, $2, $3)`, group.ID, group.Name, group.Owner)
	if err != nil {
		fmt.Println("fjdslkakd", err)
		return richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected)
	}

	return nil
}

func (d DB) GetAllGroups(ctx context.Context) ([]*entity.Group, error) {
	const op = "GroupRepository.GetAllGroups"

	rows, err := d.conn.Conn().QueryContext(ctx, `select * from groups order by created_at desc`)
	if err != nil {
		return nil, richerror.New(op).WithError(err).WithKind(richerror.KindNotFound)
	}
	groups := make([]*entity.Group, 0)
	defer rows.Close()
	for rows.Next() {
		g := &entity.Group{}
		var updatedAt time.Time
		err := rows.Scan(&g.ID, &g.Name, &g.Owner, &g.CreatedAt, &updatedAt)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, richerror.New(op).WithError(err).WithKind(richerror.KindNotFound)
		}
		groups = append(groups, g)
	}
	if err := rows.Err(); err != nil {
		return nil, richerror.New(op).WithError(err).WithMessage("something went wrong").WithKind(richerror.KindUnexpected)
	}

	return groups, nil
}

func (d DB) GetMyGroup(ctx context.Context, userID string) ([]entity.Member, error) {
	const op = "GroupRepository.GetMyGroup"

	rows, err := d.conn.Conn().QueryContext(ctx,
		`select user_id, group_id, joined_at, role from users_group where group_id = (select group_id from users_group where user_id = $1) order by "joined_at" desc`, userID)
	if err != nil {
		return nil, richerror.New(op).WithError(err).WithKind(richerror.KindNotFound)
	}
	defer rows.Close()
	members := make([]entity.Member, 0)
	defer rows.Close()
	for rows.Next() {
		m := entity.Member{}
		err := rows.Scan(&m.UserID, &m.GroupID, &m.JoinedAt, &m.Role)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, richerror.New(op).WithError(err).WithKind(richerror.KindNotFound)
		}
		members = append(members, m)
	}
	if err := rows.Err(); err != nil {
		return nil, richerror.New(op).WithError(err).WithMessage("something went wrong").WithKind(richerror.KindUnexpected)
	}

	return members, nil
}
