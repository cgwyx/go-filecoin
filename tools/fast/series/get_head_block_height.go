package series

import (
	"context"

	"github.com/filecoin-project/specs-actors/actors/abi"

	"github.com/filecoin-project/go-filecoin/tools/fast"
)

// GetHeadBlockHeight will inspect the chain head and return the height
func GetHeadBlockHeight(ctx context.Context, client *fast.Filecoin) (abi.ChainEpoch, error) {
	tipset, err := client.ChainHead(ctx)
	if err != nil {
		return 0, err
	}

	block, err := client.ShowHeader(ctx, tipset[0])
	if err != nil {
		return 0, err
	}

	return block.Height, nil
}
