package database

func (db *Database) Migrate() {
	db.Mariadb.AutoMigrate(&User{}, &Session{})
}
