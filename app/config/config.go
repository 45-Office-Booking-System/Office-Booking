package config

import (
	_bo "Office-Booking/domain/booking"
	_ge "Office-Booking/domain/gedung"
	_je "Office-Booking/domain/jenisgedung"
	_ne "Office-Booking/domain/nearby"
	_re "Office-Booking/domain/review"
	_us "Office-Booking/domain/users"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() *gorm.DB {
	config := Config{
		DB_Username: "doadmin",
		DB_Password: "AVNS_hi6OU7gGH3lcMdZZ0zP",
		DB_Port:     "25060",
		DB_Host:     "officebooking-do-user-11917513-0.b.db.ondigitalocean.com",
		DB_Name:     "defaultdb",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",

		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	InitialMigration()

	return DB
}

func InitialMigration() {
	DB.AutoMigrate(
		&_us.User{},
		&_ge.Gedung{},
		&_je.Jenisgedung{},
		&_ne.Nearby{},
		&_re.Review{},
		&_bo.Booking{},
	)
}
