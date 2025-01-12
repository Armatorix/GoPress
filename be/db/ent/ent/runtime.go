// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/Armatorix/GoPress/be/db/ent/ent/article"
	"github.com/Armatorix/GoPress/be/db/ent/ent/user"
	"github.com/Armatorix/GoPress/be/db/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	articleMixin := schema.Article{}.Mixin()
	articleMixinFields0 := articleMixin[0].Fields()
	_ = articleMixinFields0
	articleFields := schema.Article{}.Fields()
	_ = articleFields
	// articleDescCreatedAt is the schema descriptor for created_at field.
	articleDescCreatedAt := articleMixinFields0[0].Descriptor()
	// article.DefaultCreatedAt holds the default value on creation for the created_at field.
	article.DefaultCreatedAt = articleDescCreatedAt.Default.(time.Time)
	// articleDescUpdatedAt is the schema descriptor for updated_at field.
	articleDescUpdatedAt := articleMixinFields0[1].Descriptor()
	// article.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	article.DefaultUpdatedAt = articleDescUpdatedAt.Default.(time.Time)
	// articleDescBody is the schema descriptor for body field.
	articleDescBody := articleFields[3].Descriptor()
	// article.DefaultBody holds the default value on creation for the body field.
	article.DefaultBody = articleDescBody.Default.(string)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields0[0].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields0[1].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(time.Time)
	// userDescNick is the schema descriptor for nick field.
	userDescNick := userFields[3].Descriptor()
	// user.DefaultNick holds the default value on creation for the nick field.
	user.DefaultNick = userDescNick.Default.(string)
	// userDescAvatarURL is the schema descriptor for avatar_url field.
	userDescAvatarURL := userFields[4].Descriptor()
	// user.DefaultAvatarURL holds the default value on creation for the avatar_url field.
	user.DefaultAvatarURL = userDescAvatarURL.Default.(string)
}
