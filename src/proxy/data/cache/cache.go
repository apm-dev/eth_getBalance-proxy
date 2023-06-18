package cache

import (
	"fmt"
	"time"

	"github.com/ReneKroon/ttlcache"
	"github.com/apm-dev/eth_getBalance-proxy/src/domain"
)

type rpcProxyCache struct {
	cache *ttlcache.Cache
}

func NewRpcProxyCache(cache *ttlcache.Cache) domain.RpcProxyCache {
	return &rpcProxyCache{
		cache: cache,
	}
}

func (c *rpcProxyCache) GetCachedResponse(blockchain string, req *domain.JsonRpcRequest) (*domain.JsonRpcResponse, bool) {
	key := buildCacheKey(blockchain, req)
	item, ok := c.cache.Get(key)
	if ok {
		return item.(*domain.JsonRpcResponse), ok
	}
	return nil, ok
}

func (c *rpcProxyCache) CacheResponse(blockchain string, req *domain.JsonRpcRequest, rsp *domain.JsonRpcResponse, ttl time.Duration) {
	key := buildCacheKey(blockchain, req)
	c.cache.SetWithTTL(key, rsp, ttl)
}

func buildCacheKey(blockchain string, req *domain.JsonRpcRequest) string {
	return fmt.Sprintf("rpc-proxy/%s/%s/%v", blockchain, req.Method, req.Params)
}
