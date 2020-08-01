package config

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //import driver
	"github.com/spf13/viper"
)

var (
	// DB for Database
	DB *gorm.DB
	// Redis
	Redis *redis.Client
	// BridBD         *gorm.DB
	// Log Severity
	ERRORLOG   = "ERROR"
	INFOLOG    = "INFO"
	WARNINGLOG = "WARNING"
	// Dashboard URL
	DashboardURL string
	// AES Key
	AESKey string
	// Sadmin Secret
	SadminSecret string
	// JWT Secret
	JWTSecret string
	// SMTP sender email
	SMTPEmail string
)

// MaxAgeSessionTTL this is for store maxage session
type MaxAgeSessionTTL struct {
	Days   int
	Months int
	Years  int
}

// Database this is for store DB
type Database struct {
	Driver            string
	Host              string
	User              string
	Password          string
	DBName            string
	DBNumber          int
	Port              int
	APIUrl            string
	ReconnectRetry    int
	ReconnectInterval int64
	DebugMode         bool
	MaxAge            MaxAgeSessionTTL
}

// WorkerInstance for worker
type WorkerInstance struct {
	Server struct {
		Host string
		Port int
	}
	MachineName string
	IPAddress   string
	BindPort    int
}

// LoadDBConfig load database configuration
func LoadDBConfig(name string) Database {
	db := viper.Sub("database." + name)
	conf := Database{
		Driver:            db.GetString("driver"),
		Host:              db.GetString("host"),
		User:              db.GetString("user"),
		Password:          db.GetString("password"),
		DBName:            db.GetString("db_name"),
		DBNumber:          db.GetInt("db_number"),
		Port:              db.GetInt("port"),
		APIUrl:            db.GetString("api_url"),
		ReconnectRetry:    db.GetInt("reconnect_retry"),
		ReconnectInterval: db.GetInt64("reconnect_interval"),
		DebugMode:         db.GetBool("debug"),
		MaxAge: MaxAgeSessionTTL{
			Days:   db.GetInt("max_age.days"),
			Months: db.GetInt("max_age.months"),
			Years:  db.GetInt("max_age.years"),
		},
	}

	// get Global Variable from ENV File
	dashboardURL := viper.GetString("dashboard_url")
	aesKey := viper.GetString("aes_key")
	// sadminSecret := viper.GetString("sadmin_secret")
	jwtSecret := viper.GetString("jwt_secret")

	smtpConf := new(SMTPConfig)
	smtpConf.Host = viper.GetString("smtp.host")
	smtpConf.Password = viper.GetString("smtp.password")
	smtpConf.Username = viper.GetString("smtp.username")
	smtpConf.Port = viper.GetInt("smtp.port")
	smtpEmail := viper.GetString("smtp.email")

	hedwigConf := new(HedwigConfig)
	hedwigConf.URL = viper.GetString("hedwig.url")
	hedwigConf.ClientID = viper.GetString("hedwig.client_id")
	hedwigConf.ClientSecret = viper.GetString("hedwig.client_secret")

	administratorConf := new(AdministratorConfig)
	administratorConf.Username = viper.GetString("administrator.username")
	administratorConf.Password = viper.GetString("administrator.password")
	administratorConf.Name = viper.GetString("administrator.name")
	administratorConf.Email = viper.GetString("administrator.email")
	administratorConf.Phone = viper.GetString("administrator.phone")

	// declare variable from env to global variable
	DashboardURL = dashboardURL
	AESKey = aesKey
	// SadminSecret = sadminSecret
	JWTSecret = jwtSecret
	SMTPEmail = smtpEmail
	SMTP = smtpConf
	Hedwig = hedwigConf
	Administrator = administratorConf

	return conf
}

// MysqlConnect connect to mysql using config name. return *gorm.DB incstance
func MysqlConnect(configName string) *gorm.DB {
	mysql := LoadDBConfig(configName)
	connectionString := mysql.User + ":" + mysql.Password + "@tcp(" + mysql.Host + ":" + strconv.Itoa(mysql.Port) + ")/" + mysql.DBName + "?charset=utf8&parseTime=True&loc=Local"
	connection, err := gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection Established")

	if mysql.DebugMode {
		return connection.Debug()
	}

	return connection
}

// OpenDbPool for open DB Pool
func OpenDbPool() {
	DB = MysqlConnect("mysql")
	pool := viper.Sub("database.mysql.pool")
	DB.DB().SetMaxOpenConns(pool.GetInt("maxOpenConns"))
	DB.DB().SetMaxIdleConns(pool.GetInt("maxIdleConns"))
	DB.DB().SetConnMaxLifetime(pool.GetDuration("maxLifetime") * time.Second)
	DB.SingularTable(true)
	DB.Callback().Create().Remove("gorm:update_time_stamp")
}

// RedisConnect this is for connect to redis
func RedisConnect() {
	conf := LoadDBConfig("redis")
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Host + ":" + strconv.Itoa(conf.Port),
		Password: conf.Password,
		DB:       conf.DBNumber,
	})

	Redis = client
}

func getLocalIPAddress() string {
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			return ipv4.String()
		}
	}
	return ""
}
