package manager

type GameState struct {
	GameStarted       bool
	CurrentGeneration int
	TileSize          int
	TimeTick          int64
	GridSize          int
}
