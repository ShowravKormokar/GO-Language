package migrations

import "gorm.io/gorm"

func CreateUserSearchIndexes(db *gorm.DB) error {

	return db.Exec(`
		CREATE EXTENSION IF NOT EXISTS pg_trgm;

		CREATE INDEX IF NOT EXISTS idx_users_name_trgm
		ON users
		USING gin(name gin_trgm_ops);

		CREATE INDEX IF NOT EXISTS idx_users_email_trgm
		ON users
		USING gin(email gin_trgm_ops);

		CREATE INDEX IF NOT EXISTS idx_users_created_at
		ON users(created_at);

		CREATE INDEX IF NOT EXISTS idx_users_active
		ON users(is_active);

		CREATE INDEX IF NOT EXISTS idx_users_role
		ON users(role_id);
	`).Error
}
