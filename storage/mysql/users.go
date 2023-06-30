package mysql

import (
	"context"
	"fmt"

	"vps-provider/types"
)

// SaveUserInfo save user info
func SaveUserInfo(ctx context.Context, user *types.User) error {
	_, err := DB.NamedExecContext(ctx, fmt.Sprintf(
		`INSERT INTO %s (uuid, user_name, pass_hash)
			VALUES (:uuid, :user_name, :pass_hash);`, tableNameUser,
	), user)
	return err
}

// ResetPassword reset password of user
func ResetPassword(ctx context.Context, passHash, username string) error {
	_, err := DB.ExecContext(ctx, fmt.Sprintf(
		`UPDATE %s SET pass_hash = '%s', WHERE user_name = '%s'`, tableNameUser, passHash, username))
	return err
}

// GetUserByUserName get user info by user name
func GetUserByUserName(ctx context.Context, username string) (*types.User, error) {
	var out types.User
	if err := DB.QueryRowxContext(ctx, fmt.Sprintf(
		`SELECT uuid,user_name,pass_hash FROM %s WHERE user_name = ?`, tableNameUser), username,
	).StructScan(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetUserByUserUUID get user info by uuid
func GetUserByUserUUID(ctx context.Context, UUID string) (*types.User, error) {
	var out types.User
	if err := DB.QueryRowxContext(ctx, fmt.Sprintf(
		`SELECT * FROM %s WHERE uuid = ?`, tableNameUser), UUID,
	).StructScan(&out); err != nil {
		return nil, err
	}
	return &out, nil
}
