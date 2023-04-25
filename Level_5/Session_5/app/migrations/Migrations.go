package migrations

import (
	"idstar.com/session12/app/configs"
	"idstar.com/session12/app/migrations/dummy"
	"idstar.com/session12/app/models"
)

func Apply() error {
	db := configs.GetDB()

	if err := db.AutoMigrate(&models.RoleModel{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&models.ModuleModel{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&models.RoleModuleModel{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&models.UserModel{}); err != nil {
		return err
	}

	role := dummy.CreateDummyRole{}
	if err := role.Apply(db); err != nil {
		return err
	}

	module := dummy.CreateDummyModule{}
	if err := module.Apply(db); err != nil {
		return err
	}

	roleModule := dummy.CreateDummyRoleModule{}
	if err := roleModule.Apply(db); err != nil {
		return err
	}

	user := dummy.CreateDummyUser{}
	if err := user.Apply(db); err != nil {
		return err
	}

	return nil
}
