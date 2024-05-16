package app

import (
	"log"

	"github.com/badsector998/go-rbac/internal/adapter"
	"github.com/badsector998/go-rbac/internal/app/command"
	"github.com/badsector998/go-rbac/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	Queries  Queries
	Commands Commands
}

type Queries struct {
}

type Commands struct {
	BurjoGroupCreate *command.BurjoGroupCreateHandler

	BurjoCreate *command.BurjoCreateHandler

	EmployeeCreate *command.EmployeeCreateHandler
}

func NewApp() *App {

	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		log.Panic("error opening db")
	}

	db.AutoMigrate(
		&domain.BurjoGroup{},
		&domain.Burjo{},
		&domain.Employee{},
	)

	burjoGroupRepo := adapter.NewBurjoGroupRepository(db)
	burjoRepo := adapter.NewBurjoRepository(db)
	employeeRepo := adapter.NewEmployeeRepository(db)

	return &App{
		Queries: Queries{},
		Commands: Commands{
			BurjoGroupCreate: command.NewBurjoGroupCreate(burjoGroupRepo),

			BurjoCreate: command.NewBurjoCreate(burjoRepo),

			EmployeeCreate: command.NewEmployeeCreate(employeeRepo),
		},
	}
}
