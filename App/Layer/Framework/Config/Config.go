package config

type config struct {
	Database struct {
		User     string
		Password string
		Addr     string
		DBName   string
	}
	Server struct {
		Address string
	}
}

var C config

func ReadConfig() {
	Config := &C

	Config.Database.User = "dsdd"
	Config.Database.Password = "12345"
	Config.Database.Addr = "10.130.70.1"
	Config.Database.DBName = "HBDSD"

	Config.Server.Address = "8081"
	// fmt.Println(*&Config.Database)
}
