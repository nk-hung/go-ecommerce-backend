package setting

type Config struct {
	Redis  RedisSetting  `mapstructure:"redis"`
	Mysql  MysqlSetting  `mapstructure:"mysql"`
	Logger LoggerSetting `mapstructure:"logger"`
	Server ServerSetting `mapstructure:"mode"`
}

type ServerSetting struct {
	Mode string `mapstructure:"mode"`
	Port int    `mapstructure:"port"`
}

type MysqlSetting struct {
	Host            string `mapstructure:"host"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	Dbname          string `mapstructure:"dbname"`
	Port            int    `mapstructure:"port"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime"`
}

type LoggerSetting struct {
	Log_level     string `mapstructure:"log_level"`
	File_log_name string `mapstructure:"file_log_name"`
	Max_backups   int    `mapstructure:"max_backups"`
	Max_age       int    `mapstructure:"max_age"`
	Max_size      int    `mapstructure:"max_size"`
	Compress      bool   `mapstructure:"compress"`
}

type RedisSetting struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Database int    `mapstructure:"database"`
	Port     int    `mapstructure:"port"`
}
