//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate ./schema  --idtype string --target ./ent/ --feature namedges --feature sql/upsert --feature sql/execquery --feature sql/modifier --template=./schema/tmpl
package ent
