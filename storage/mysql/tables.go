package mysql

var cUsersTable = `
    CREATE TABLE if not exists %s (
		uuid          VARCHAR(128) NOT NULL UNIQUE,
		user_name     VARCHAR(128) NOT NULL UNIQUE,
		pass_hash     VARCHAR(128) NOT NULL,
		created_at    DATETIME     DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY (uuid)
	) ENGINE=InnoDB COMMENT='user info';`
