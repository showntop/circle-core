package main

import (
	"github.com/showntop/circle-core/migrations"
)

func main() {
	///可以根据参数确定要做的动作，create 、migrate 、drop 、reset， rollback -n 、redo -n
	migrations.Migrate()
}
