package model

import (
	"github.com/sirupsen/logrus"
	"sync"
	"sync/atomic"
)

var SN *MySyncMap

type MySyncMap struct {
	SyncMap sync.Map
	Counter uint64
}

func NewSyncMap() *MySyncMap {
	return &MySyncMap{SyncMap: sync.Map{}, Counter: 0}
}

func (s *MySyncMap) Store(key any, value any) {
	s.SyncMap.Store(key, value)
	atomic.AddUint64(&s.Counter, 1)
}
func (s *MySyncMap) Load(key any) (result string) {
	res, ok := s.SyncMap.Load(key)
	if !ok {
		logrus.Error("There is not such key")
	}
	return res.(string)
}

func (s *MySyncMap) GetCounter() (result uint64) {
	return atomic.LoadUint64(&s.Counter)
}
