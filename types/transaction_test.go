package types

import (
	"testing"
	"github.com/ethereum/go-ethereum/common"
	"github.com/DexyProject/dexy-go/types"
)

func TestTick_Aggregate(t *testing.T) {

	trans1 := Transaction{
		TransactionID: "0x87012a0d870d47c3c93526c05c4a2f494054c3f4dd8584e94af7d8dd90a535f8",
		OrderHash: "0xEEAD6DBFC7340A56CAEDC044696A168870549A6A7F6F56961E84A54BD9970B8A",
		BlockNumber: 4862998,
		Timestamp: "1515233752",
		Taker: HexToAddress("0x997919a608788621dd48b3896f78dcda682fe91d"),
		Maker: HexToAddress("0x9f612fcb422d1971c1be7416c37e3ebc77c0de19"),
		Give: Trade{Token: HexToAddress("0x0000000000000000000000000000000000000000"), Amount: "30"},
		Get:  Trade{Token: HexToAddress("0xd26114cd6EE289AccF82350c8d8487fedB8A0C07"), Amount: "300"},
	}


}