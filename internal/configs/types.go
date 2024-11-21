package configs

type (
	Config struct {
		Service  Service  `mapstructure:"service"`
		Database Database `mapstructure:"database"`
	}

	Service struct {
		Port      string `mapstructure:"port"`
		SecretJWT string `mapstructure:"secret_jwt"`
	}

	Database struct {
		DatabaseSourceName string `mapstructure:"db_source_name"`
	}
)
