package gortc_service_env

type Environment struct {
	MYSQL struct {
		HOST     string `env:"MYSQL_HOST,default=127.0.0.1"`
		PORT     int    `env:"MYSQL_PORT,default=3306"`
		USER     string `env:"MYSQL_USER,default=root"`
		PASSWORD string `env:"MYSQL_PASSWORD,default=root"`
		DATABASE string `env:"MYSQL_DATABASE,default=gortc_dev"`
	}
	APP struct {
		PEPPER string `env:"APP_PEPPER,default=qwerty1234"`
	}
}
