CREATE TABLE users (
                       id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                       created_at datetime(3) NOT NULL,
                       updated_at datetime(3) NOT NULL,
                       deleted_at datetime(3),
                       PRIMARY KEY (id),
                       KEY idx_users_created_at (created_at),
                       KEY idx_users_updated_at (updated_at),
                       KEY deleted_at (deleted_at),
                       nome varchar(255) NOT NULL,
                       token varchar(255) NOT NULL,
                       conta_vinculada varchar(255) NOT NULL
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;