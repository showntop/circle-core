package migrations

import (
	"github.com/go-gorp/gorp"
	"github.com/showntop/circle-core/models"
)

func createUsers(db *gorp.DbMap) {
	table := db.AddTableWithName(models.User{}, "users").SetKeys(false, "Id")
	table.ColMap("Id").Rename("id").SetMaxSize(26)
	table.ColMap("LoginName").Rename("login_name").SetMaxSize(64).SetUnique(true)
	table.ColMap("Password").Rename("password").SetMaxSize(128)
	table.ColMap("CreatedAt").Rename("created_at").SetMaxSize(128)
	table.ColMap("UpdatedAt").Rename("updated_at").SetMaxSize(128)
	table.ColMap("DeletedAt").Rename("deleted_at").SetMaxSize(128)

}
