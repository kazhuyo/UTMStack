package cache

import (
	"sync"
	"time"

	"github.com/AtlasInsideCorp/UTMStackCorrelationEngine/rules"
	"github.com/AtlasInsideCorp/UTMStackCorrelationEngine/utils"
	"github.com/quantfall/holmes"
	"github.com/tidwall/gjson"
)

var cacheStorageMutex = &sync.RWMutex{}

var cacheStorage []string

var h = holmes.New(utils.GetConfig().ErrorLevel, "CACHE")

func Status() {
	for {
		h.Info("Logs in cache: %v", len(cacheStorage))
		if len(cacheStorage) != 0 {
			est := gjson.Get(cacheStorage[0], "@timestamp").String()
			h.Info("Old document in cache: %s", est)
		}
		time.Sleep(60 * time.Second)
	}

}

func Search(allOf []rules.AllOf, oneOf []rules.OneOf, seconds int) []string {
	var elements []string
	cacheStorageMutex.RLock()
	start := time.Now()
	cToBreak := 0
	ait := time.Now().UTC().Unix() - int64(seconds)
	for i := len(cacheStorage) - 1; i >= 0; i-- {
		est := gjson.Get(cacheStorage[i], "@timestamp").String()
		eit, err := time.Parse(time.RFC3339Nano, est)
		if err != nil {
			h.Error("Could not parse @timestamp: %v", err)
			continue
		}
		if eit.Unix() < ait {
			if cToBreak <= 5 {
				cToBreak++
			} else {
				break
			}
		} else {
			cToBreak = 0
			var allCatch bool
			var oneCatch bool
			for _, of := range oneOf {
				oneCatch = evalElement(cacheStorage[i], of.Field, of.Operator, of.Value)
				if oneCatch {
					break
				}
			}
			for _, af := range allOf {
				allCatch = evalElement(cacheStorage[i], af.Field, af.Operator, af.Value)
				if !allCatch {
					break
				}
			}
			if (len(allOf) == 0 || allCatch) && (len(oneOf) == 0 || oneCatch) {
				elements = append(elements, cacheStorage[i])
			}
		}
	}
	cacheStorageMutex.RUnlock()
	h.Debug("Cache storage unlocked by search. Search took: %s", time.Since(start))
	h.Debug("Returned %v elements", len(elements))
	return elements
}

var logs = make(chan string, 100)

func AddToCache(l string) {
	logs <- l
}

func ProccessQueue() {
	for {
		l := <-logs
		cacheStorageMutex.Lock()
		cacheStorage = append(cacheStorage, l)
		cacheStorageMutex.Unlock()
	}

}

func Clean() {
	for {
		if utils.AssignedMemory >= 50 && len(cacheStorage) > 500 {
			cacheStorageMutex.Lock()
			cacheStorage = cacheStorage[1:]
			cacheStorageMutex.Unlock()
		} else {
			time.Sleep(5 * time.Second)
		}
	}
}