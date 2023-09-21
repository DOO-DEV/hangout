package pggroup

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"hangout/entity"
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
		if isDuplicateKeyError(err) {
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
		if isEmptyRowError(err) {
			return nil, richerror.New(op).WithError(err).WithKind(richerror.KindForbidden).WithMessage(errmsg.ErrorMsgUserNotAllowed)
		}
		return nil, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgCantScanQueryResult)
	}

	return g, nil
}

func (d DB) ListAllJoinRequestToMyGroup(ctx context.Context, groupID string) ([]entity.PendingList, error) {
	const op = "GroupRepository.ListAllJoinRequestToMyGroup"

	fmt.Println(groupID)
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
			if isEmptyRowError(err) {
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
		if isEmptyRowError(err) {
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

func isDuplicateKeyError(err error) bool {
	var pgErr *pq.Error
	if errors.As(err, &pgErr) && pgErr.Code == "23505" {
		return true
	}

	return false
}

func isEmptyRowError(err error) bool {
	if errors.Is(err, sql.ErrNoRows) {
		return true
	}

	return false
}
