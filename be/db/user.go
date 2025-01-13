package db

import (
	"context"

	"github.com/Armatorix/GoPress/be/db/ent/ent"
	"github.com/Armatorix/GoPress/be/db/ent/ent/user"
)

func (db *DB) UserExists(ctx context.Context, id int) (bool, error) {
	return db.UserClient().Query().
		Where(user.ID(id)).
		Exist(ctx)
}

func (db *DB) UserByEmail(ctx context.Context, email string) (*ent.User, error) {
	return db.UserClient().Query().
		Where(user.Email(email)).
		Only(ctx)
}
