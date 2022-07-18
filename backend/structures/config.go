package structures

type Config struct {
	Database struct {
		User string `mapstructure:"POSTGRES_USER"`
		Name string `mapstructure:"POSTGRES_DB"`
		Pass string `mapstructure:"POSTGRES_PASSWORD"`
	} `mapstructure:"database"`
	System struct {
		Save_Dir    string `mapstructure:"save_dir"`
		Max_Storage int64  `mapstructure:"max_storage"`
		Expiry      int64  `mapStructure:"expiry"`
	} `mapstructure:"system"`
}
