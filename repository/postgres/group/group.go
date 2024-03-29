package pggroup

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"hangout/entity"
	dbparam "hangout/param/pgdb"
	"hangout/pkg/errmsg"
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

func (d DB) AddToPendingList(ctx context.Context, p entity.PendingList) error {
	const op = "GroupRepository.AddToPendingList"

	if _, err := d.conn.Conn().ExecContext(ctx, `insert into "pending_list"("user_id", "group_id") values ($1, $2)`, p.UserID, p.GroupId); err != nil {
		if d.conn.IsDuplicateKeyError(err) {
			return richerror.New(op).WithError(err).WithKind(richerror.KindInvalid).WithMessage(errmsg.ErrorMsgYouAlreadySendRequest)
		}
		return richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected)
	}

	return nil
}

func (d DB) ListJoinRequest(ctx context.Context, userID string) ([]entity.PendingList, error) {
	const op = "GroupRepository.ListJoinRequest"

	rows, err := d.conn.Conn().QueryContext(ctx, `select * from "pending_list" where "user_id" = $1 order by "sent_at" desc`, userID)
	defer rows.Close()
	if err != nil {
		return nil, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected)
	}

	list := make([]entity.PendingList, 0)
	for rows.Next() {
		p := entity.PendingList{}
		err := rows.Scan(&p.UserID, &p.GroupId, &p.SentAt)
		if err != nil {
			return nil, richerror.New(op).WithError(err).WithKind(richerror.KindNotFound).WithMessage(errmsg.ErrorMsgNotFound)
		}
		list = append(list, p)
	}

	return list, nil
}

func (d DB) GetOwnedGroup(ctx context.Context, userID string) (*entity.Group, error) {
	const op = "GroupRepository.CheckGroupOwner"
	g := &entity.Group{}
	row := d.conn.Conn().QueryRowContext(ctx, `select "id" from "groups" where "owner_id" = $1`, userID)
	if err := row.Scan(&g.ID); err != nil {
		if d.conn.IsEmptyRowError(err) {
			fmt.Println(err, userID)
			return nil, richerror.New(op).WithError(err).WithKind(richerror.KindForbidden).WithMessage(errmsg.ErrorMsgUserNotAllowed)
		}
		return nil, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgCantScanQueryResult)
	}

	return g, nil
}

func (d DB) ListAllJoinRequestToMyGroup(ctx context.Context, groupID string) ([]entity.PendingList, error) {
	const op = "GroupRepository.ListAllJoinRequestToMyGroup"

	rows, err := d.conn.Conn().QueryContext(ctx, `select * from "pending_list" where "group_id" = $1 and "active"='true' order by "sent_at" desc`, groupID)
	defer rows.Close()
	if err != nil {
		return nil, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected)
	}

	list := make([]entity.PendingList, 0)
	for rows.Next() {
		p := entity.PendingList{}
		err := rows.Scan(&p.UserID, &p.GroupId, &p.SentAt, &p.Active)
		if err != nil {
			if d.conn.IsEmptyRowError(err) {
				return nil, richerror.New(op).WithError(err).WithKind(richerror.KindNotFound).WithMessage(errmsg.ErrorMsgNotFound)
			}
			return nil, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
		}

		list = append(list, p)
	}

	return list, nil
}

func (d DB) MoveFromPendingListToGroup(ctx context.Context, groupID string, userID string) error {
	const op = "GroupRepository.MoveFromPendingListToGroup"

	fmt.Println("user: ", userID, "group: ", groupID)
	p := entity.PendingList{}
	row := d.conn.Conn().QueryRowContext(ctx, `select * from "pending_list" where "user_id" = $1 and "group_id" = $2`, userID, groupID)
	if err := row.Scan(&p.UserID, &p.GroupId, &p.SentAt, &p.Active); err != nil {
		if d.conn.IsEmptyRowError(err) {
			return richerror.New(op).WithError(err).WithKind(richerror.KindNotFound).WithMessage(errmsg.ErrorMsgNotFound)
		}
		return richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgCantScanQueryResult)
	}

	// remove this pending request and change the active status of rest of them to false
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
	// remove this request
	_, err = tx.ExecContext(ctx, `delete from "pending_list" where "user_id" = $1 and "group_id" = $2`, userID, groupID)
	if err != nil {
		return richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}
	_, err = tx.ExecContext(ctx, `update "pending_list" set "active" = 'false' where "user_id" = $1`, userID)
	if err != nil {
		return richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}
	// move this user to "users_group"
	_, err = tx.ExecContext(ctx, `insert into "users_group"("user_id", "group_id", "role") values ($1, $2,$3)`, userID, groupID, "normal")
	if err != nil {
		return richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}

	return nil
}

func (d DB) ConnectGroups(ctx context.Context, g1, g2 string) error {
	const op = "GroupRepository.ConnectGroup"

	_, err := d.conn.Conn().ExecContext(ctx, `insert into "groups_connections"("from", "to") values ($1, $2)`, g1, g2)
	if err != nil {
		if d.conn.IsForeignKeyError(err) {
			return richerror.New(op).WithError(err).WithKind(richerror.KindInvalid).WithMessage(errmsg.ErrorMsgGroupNotFound)
		}
		if d.conn.IsDuplicateKeyError(err) {
			return richerror.New(op).WithError(err).WithKind(richerror.KindInvalid).WithMessage(errmsg.ErrorMsgByDirectional)
		}
		return richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}

	return nil
}

func (d DB) ListMyGroupConnections(ctx context.Context, groupID string) ([]dbparam.GroupConnection, error) {
	const op = "GroupRepository.ListMyGroupConnections"

	rows, err := d.conn.Conn().QueryContext(ctx, `select "from", "created_at" from "groups_connections" where "to" = $1 and "accept" = 'false' order by "created_at" desc`, groupID)
	if err != nil {
		return nil, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}
	defer rows.Close()
	list := make([]dbparam.GroupConnection, 0)
	for rows.Next() {
		g := dbparam.GroupConnection{}
		if err := rows.Scan(&g.From, &g.CreatedAt); err != nil {
			if d.conn.IsEmptyRowError(err) {
				return nil, richerror.New(op).WithError(err).WithKind(richerror.KindNotFound).WithMessage(errmsg.ErrorMsgNotFound)
			}
			return nil, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
		}
		list = append(list, g)
	}
	if rows.Err() != nil {
		return nil, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}

	return list, nil
}

func (d DB) AcceptGroupConnection(ctx context.Context, from, to string) error {
	const op = "GroupRepository.AcceptGroupConnection"

	row := d.conn.Conn().QueryRowContext(ctx, `update "groups_connections" set "accept" = 'true' where "from" = $1 and "to" = $2 returning "accept"`, from, to)
	var accept bool
	if err := row.Scan(&accept); err != nil {
		if d.conn.IsEmptyRowError(err) {
			return richerror.New(op).WithError(err).WithKind(richerror.KindNotFound).WithMessage(errmsg.ErrorMsgGroupNotFound)
		}
		return richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}

	return nil
}

func (d DB) CheckUserGroupConnection(ctx context.Context, u1, u2 string) (bool, error) {
	const op = "GroupRepository.CheckUserGroupConnection"

	row := d.conn.Conn().QueryRowContext(ctx, `select "group_id" from "users_group" where "user_id" = $1`, u1)
	g1 := entity.Group{}
	if err := row.Scan(&g1.ID); err != nil {
		if d.conn.IsEmptyRowError(err) {
			return false, richerror.New(op).WithError(err).WithKind(richerror.KindNotFound).WithMessage(errmsg.ErrorMsgUserNotFound)
		}
		return false, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}

	row = d.conn.Conn().QueryRowContext(ctx, `select "group_id" from "users_group" where "user_id" = $1`, u2)
	g2 := entity.Group{}
	if err := row.Scan(&g2.ID); err != nil {
		if d.conn.IsEmptyRowError(err) {
			return false, richerror.New(op).WithError(err).WithKind(richerror.KindNotFound).WithMessage(errmsg.ErrorMsgUserNotFound)
		}
		return false, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}

	row = d.conn.Conn().QueryRowContext(ctx, `select "accept" from "groups_connections" where (("from"= $1 and "to"= $2) or ("from" = $3 and "to" = $4) and "accept" = 'true')`, g1.ID, g2.ID, g2.ID, g1.ID)
	var accept bool
	if err := row.Scan(&accept); err != nil {
		if d.conn.IsEmptyRowError(err) {
			return false, richerror.New(op).WithError(err).WithKind(richerror.KindNotFound).WithMessage(errmsg.ErrorMsgUsersAreNotConnected)
		}
		return false, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}

	return accept, nil
}
