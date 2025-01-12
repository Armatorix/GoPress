package db

import (
	"context"

	"github.com/Armatorix/GoPress/be/db/ent/ent/user"
)

func (db *DB) UserExists(ctx context.Context, id int) (bool, error) {
	return db.UserClient().Query().
		Where(user.ID(id)).
		Exist(ctx)
}
