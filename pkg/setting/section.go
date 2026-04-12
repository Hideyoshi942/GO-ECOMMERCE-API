package setting

type Config struct {
	MySql  MySqlSetting  `mapstructure:"mysql"`
	Logger LoggerSetting `mapstructure:"logger"`
	Redis  RedisSetting  `mapstructure:"redis"`
}

type RedisSetting struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

type MySqlSetting struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`

	MaxIdleConns int `mapstructure:"maxIdleConns"`
	MaxOpenConns int `mapstructure:"maxOpenConns"`

	MaxLifetime int `mapstructure:"maxLifetime"`
	MaxIdleTime int `mapstructure:"maxIdleTime"`

	MaxRetries    int `mapstructure:"maxRetries"`
	RetryInterval int `mapstructure:"retryInterval"`

	ReadTimeout  int `mapstructure:"readTimeout"`
	WriteTimeout int `mapstructure:"writeTimeout"`
	Timeout      int `mapstructure:"timeout"`
	IdleTimeout  int `mapstructure:"idleTimeout"`
}

type LoggerSetting struct {
	LogLevel    string `mapstructure:"log_level"`
	FileLogName string `mapstructure:"file_log_name"`

	MaxSize    int  `mapstructure:"max_size"`
	MaxBackups int  `mapstructure:"max_backups"`
	MaxAge     int  `mapstructure:"max_age"`
	Compress   bool `mapstructure:"compress"`
}
