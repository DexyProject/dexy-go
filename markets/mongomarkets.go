package markets

import "github.com/DexyProject/dexy-go/types"

type MongoMarkets struct {
	
}

func (markets *MongoMarkets) InsertMarkets([]types.Market) error {
	panic("implement me")
}

func (markets *MongoMarkets) GetMarkets([]types.Address) ([]types.Market, error) {
	panic("implement me")
}
