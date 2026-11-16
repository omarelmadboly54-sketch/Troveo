package models


type Configuration struct{
	DBDriver				string				`mapstructure:"DB_DRIVER"`
	DBSource				string				`mapstructure:"DB_SOURCE"`
	ServerAddress			string				`mapstructure:"SERVER_ADDRESS"`
	FrontendUrl				string				`mapstructure:"FRONTEND_URL"`
	CloudinaryUrl			string				`mapstructure:"CLOUDINARY_URL"`			
	SecretKey				string				`mapstructure:"SECRET_KEY"`
	AccessTokenMinutes		int					`mapstructure:"ACCESS_TOKEN_MINUTES"`
}

