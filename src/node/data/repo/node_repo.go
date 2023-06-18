package repo

import "github.com/apm-dev/eth_getBalance-proxy/src/domain"

type nodeRepository struct {
	data map[string][]domain.Node
}

func NewNodeRepository() domain.NodeRepository {
	data := map[string][]domain.Node{
		"eth": {
			{
				ID:         "51a8c2b2-ac70-4ea7-bc6b-5d634538253a",
				Url:        "https://rpc.ankr.com/eth",
				Blockchain: "eth",
				Weight:     100,
			},
			{
				ID:         "a2544b72-66f9-4075-a8b0-2aea88a27ebc",
				Url:        "https://mainnet.infura.io/v3/b3bd9456e8d44150b963248668023317",
				Blockchain: "eth",
				Weight:     95,
			},
		},
	}
	return &nodeRepository{data}
}

func (r *nodeRepository) GetNodesByBlockchain(blockchain string) ([]domain.Node, error) {
	return r.data[blockchain], nil
}
