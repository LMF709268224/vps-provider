package storage

import (
	"context"
	"fmt"

	"vps-provider/types"
)

func CreateUser(ctx context.Context, user *types.User) error {
	_, err := DB.NamedExecContext(ctx, fmt.Sprintf(
		`INSERT INTO %s (uuid, user_name, pass_hash)
			VALUES (:uuid, :user_name, :pass_hash);`, tableNameUser,
	), user)
	return err
}

func ResetPassword(ctx context.Context, passHash, username string) error {
	_, err := DB.ExecContext(ctx, fmt.Sprintf(
		`UPDATE %s SET pass_hash = '%s', WHERE user_name = '%s'`, tableNameUser, passHash, username))
	return err
}

func GetUserByUsername(ctx context.Context, username string) (*types.User, error) {
	var out types.User
	if err := DB.QueryRowxContext(ctx, fmt.Sprintf(
		`SELECT * FROM %s WHERE user_name = ?`, tableNameUser), username,
	).StructScan(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetUserByUserUUID(ctx context.Context, UUID string) (*types.User, error) {
	var out types.User
	if err := DB.QueryRowxContext(ctx, fmt.Sprintf(
		`SELECT * FROM %s WHERE uuid = ?`, tableNameUser), UUID,
	).StructScan(&out); err != nil {
		return nil, err
	}
	return &out, nil
}
