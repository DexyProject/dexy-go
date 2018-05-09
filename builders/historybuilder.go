package builders

import "github.com/DexyProject/dexy-go/types"

type HistoryBuilder struct {

}

func (hb *HistoryBuilder) build(transactions []types.Transaction) ([]types.Tick, error) {
	ticks := make([]types.Tick, 0)

	return ticks, nil
}
