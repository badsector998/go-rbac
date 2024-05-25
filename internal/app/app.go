package app

import (
	"log"

	"github.com/badsector998/go-rbac/internal/adapter"
	"github.com/badsector998/go-rbac/internal/app/command"
	"github.com/badsector998/go-rbac/internal/app/query"
	"github.com/badsector998/go-rbac/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	Queries  Queries
	Commands Commands
}

type Queries struct {
	BurjoGroupReadByID *query.BurjoGroupReadByIDHandler
	BurjoGroupReadAll  *query.BurjoGroupReadAllHandler

	BurjoReadByID *query.BurjoReadByIDHandler
	BurjoReadAll  *query.BurjoReadAllHandler

	EmployeeReadByID *query.EmployeeReadByIDHandler
	EmployeeReadAll  *query.EmployeeReadAllHandler

	UserReadByEmail *query.UserReadByEmailHandler
	UserReadByID    *query.UserGetByIDHandler
	UserReadAll     *query.UserReadAllHandler
}

type Commands struct {
	BurjoGroupCreate *command.BurjoGroupCreateHandler
	BurjoGroupUpdate *command.BurjoGroupUpdateHandler
	BurjoGroupDelete *command.BurjoGroupDeleteHandler

	BurjoCreate *command.BurjoCreateHandler
	BurjoUpdate *command.BurjoUpdateHandler
	BurjoDelete *command.BurjoDeleteHandler

	EmployeeCreate *command.EmployeeCreateHandler
	EmployeeUpdate *command.EmployeeUpdateHanlder
	EmployeeDelete *command.EmployeeDeleteHandler

	UserCreate *command.UserCreateHandler
	UserUpdate *command.UserUpdateHandler
	UserDelete *command.UserDeleteHandler
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
		&domain.User{},
		&domain.BurjoGroup{},
		&domain.Burjo{},
		&domain.Employee{},
	)

	burjoGroupRepo := adapter.NewBurjoGroupRepository(db)
	burjoRepo := adapter.NewBurjoRepository(db)
	employeeRepo := adapter.NewEmployeeRepository(db)
	userRepo := adapter.NewUserRepository(db)

	return &App{
		Queries: Queries{
			BurjoGroupReadByID: query.NewBurjoGroupReadByID(burjoGroupRepo),
			BurjoGroupReadAll:  query.NewBurjoGroupReadAll(burjoGroupRepo),

			BurjoReadByID: query.NewBurjoReadByID(burjoRepo),
			BurjoReadAll:  query.NewBurjoReadAll(burjoRepo),

			EmployeeReadByID: query.NewEmployeeReadByID(employeeRepo),
			EmployeeReadAll:  query.NewEmployeeReadAll(employeeRepo),

			UserReadByEmail: query.NewUserReadByEmail(userRepo),
			UserReadByID:    query.NewUserGetByID(userRepo),
			UserReadAll:     query.NewUserReadAll(userRepo),
		},
		Commands: Commands{
			BurjoGroupCreate: command.NewBurjoGroupCreate(burjoGroupRepo),
			BurjoGroupUpdate: command.NewBurjoGroupUpdate(burjoGroupRepo),
			BurjoGroupDelete: command.NewBurjoGroupDelete(burjoGroupRepo),

			BurjoCreate: command.NewBurjoCreate(burjoRepo),
			BurjoUpdate: command.NewBurjoUpdate(burjoRepo),
			BurjoDelete: command.NewBurjoDelete(burjoRepo),

			EmployeeCreate: command.NewEmployeeCreate(employeeRepo),
			EmployeeUpdate: command.NewEmployeeUpdate(employeeRepo),
			EmployeeDelete: command.NewEmployeeDelete(employeeRepo),

			UserCreate: command.NewUserCreate(userRepo),
			UserUpdate: command.NewUserUpdate(userRepo),
			UserDelete: command.NewUserDelete(userRepo),
		},
	}
}
