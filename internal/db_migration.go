package internal

//Migrator migrates db tables
type Migrator interface {
	Migrate() error
}

//Migrations runs all migrations
type Migrations struct {
	Migrators []Migrator
}

//NewDBMigration ...
func NewDBMigration(migrators ...Migrator) Migrations{
	return Migrations{
		Migrators: migrators,
	}
}

//AutoMigrate to initiate the migration process
func (m *Migrations) AutoMigrate() error {
	for _, dm := range m.Migrators {
		err := dm.Migrate()
		if err != nil {
			return err
		}
	}

	return nil
}