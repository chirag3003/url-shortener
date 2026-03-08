package idgen

import (
	"sync"

	"github.com/chirag3003/hyperflake-go"
)

// Generator defines ID generation behavior.
type Generator interface {
	NewID() (int64, error)
}

var (
	defaultGenerator Generator
	once             sync.Once
)

// SetDefault sets a global generator instance.
func SetDefault(g Generator) {
	defaultGenerator = g
}

// Init initializes the default hyperflake generator.
func Init(datacenterID int, machineID int, epochMS int64) {
	once.Do(func() {
		SetDefault(NewHyperflakeGenerator(datacenterID, machineID, epochMS))
	})
}

// NewID generates a new ID using the default generator.
func NewID() (int64, error) {
	if defaultGenerator == nil {
		// Fallback for cases where Init wasn't called (e.g., tests)
		// but ideally Init should be called explicitly.
		Init(0, 0, 0)
	}
	return defaultGenerator.NewID()
}

type hyperflakeGenerator struct {
	cfg *hyperflake.Config
}

// NewHyperflakeGenerator creates a hyperflake-backed ID generator.
func NewHyperflakeGenerator(datacenterID int, machineID int, epochMS int64) Generator {
	if epochMS > 0 {
		return &hyperflakeGenerator{cfg: hyperflake.NewHyperflakeConfigWithEpoch(datacenterID, machineID, epochMS)}
	}
	return &hyperflakeGenerator{cfg: hyperflake.NewHyperflakeConfig(datacenterID, machineID)}
}

func (g *hyperflakeGenerator) NewID() (int64, error) {
	return g.cfg.GenerateHyperflakeID()
}
