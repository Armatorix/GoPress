package db

import (
	"context"

	"github.com/Armatorix/GoPress/be/db/ent/ent"
	"github.com/Armatorix/GoPress/be/db/ent/ent/user"
	"github.com/Armatorix/GoPress/be/db/ext"
)

func (db *DB) UserExists(ctx context.Context, id int) (bool, error) {
	return db.UserClient().Query().
		Where(user.ID(id)).
		Exist(ctx)
}

func (db *DB) UserByEmailOrNick(ctx context.Context, email string) (*ent.User, error) {
	return db.UserClient().Query().
		Where(
			user.Or(
				user.Email(email),
				user.Nick(email),
			),
		).
		Only(ctx)
}

func (db *DB) InitAdminUser(ctx context.Context, passwd string) error {
	totalUsers, err := db.UserClient().Query().Count(ctx)
	if err != nil {
		return err
	}

	if totalUsers > 0 {
		return nil
	}

	return db.UserClient().Create().
		SetEmail("").
		SetNick("admin").
		SetPassword(ext.Password(passwd)).
		Exec(ctx)
}
