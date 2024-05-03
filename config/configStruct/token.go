package structures

type TokenConfig struct {
	JTokenUrl      string `mapstructure:"JTOKEN_URL" json:"j_token_url"`
	JTokenUser     string `mapstructure:"JTOKEN_USER" json:"j_token_user"`
	JTokenPassword string `mapstructure:"JTOKEN_PASSWORD" json:"j_token_password"`
	JTokenSubTTL   string `mapstructure:"SUB_TTL_SECONDS" json:"j_token_sub_ttl"`
}
