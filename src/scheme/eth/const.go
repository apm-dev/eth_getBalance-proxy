package eth

import "time"

var (
	cachingMethods = map[string]time.Duration{
		"eth_blockNumber": 2 * time.Second,
		"eth_getBalance":  time.Minute,
	}

	blacklistedMethods = map[string]bool{
		"eth_hashrate": true,
		"eth_mining":   true,
		"eth_getWork":  true,
	}

	blacklistedPrefixes = map[string]bool{
		"miner":  true,
		"engine": true,
		"parity": true,
	}
)
