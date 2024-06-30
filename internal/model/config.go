package model

// configuration model
type Config struct {
	Port         int    `mapstructure:"PORT"`
	DBPath       string `mapstructure:"DB_PATH"`
	MapAssetPath string `mapstructure:"MAP_ASSET_PATH"`
	GridSize     string `mapstructure:"GRID_SIZE"`
	CanvasHeight string `mapstructure:"CANVAS_HEIGHT"`
	CanvasWidth  string `mapstructure:"CANVAS_WIDTH"`
}
