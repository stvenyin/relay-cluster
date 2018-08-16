package cache

import (
	"github.com/Loopring/relay-lib/cache"
	"strings"
	"github.com/Loopring/relay-lib/log"
	"github.com/Loopring/relay-cluster/dao"
	"encoding/json"
)

const TxViewSearchPreKey = "TXVIEW_"

func GetCacheTransactions(key string, res *[]dao.TransactionView) (err error, get bool) {
	key = strings.ToUpper(TxViewSearchPreKey + key)
	if txsByte, err := cache.Get(key); err != nil {
		return err, false
	} else if len(txsByte) > 0 {
		json.Unmarshal(txsByte, &res)
		log.Debugf("[GetTransactions Cache] from cache key: %s", key)
		return err, true
	}
	return nil, false
}

func SaveCacheTransactions(key string, res *[]dao.TransactionView, ttl int64) {
	key = strings.ToUpper(TxViewSearchPreKey + key)
	value, _ := json.Marshal(res)
	log.Debugf("[GetTransactions Cache] save cache key: %s", key)
	cache.Set(key, value, ttl)
}

func DelTxViewCacheByOwners(owners []string) {
	keyStrs := make([]string, 0)
	for _, owner := range owners {
		keys, _ := cache.Keys(strings.ToUpper(TxViewSearchPreKey + "OWNER:" + owner + "*"))
		for _, key := range keys {
			log.Debugf("[clear cache] clear key: %s", string(key))
			keyStrs = append(keyStrs, string(key))
		}
	}
	cache.Dels(keyStrs)
}
