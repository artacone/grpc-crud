package cache

import (
	"gitlab.ozon.dev/artacone/workshop-1/pkg/generator"
	"gitlab.ozon.dev/artacone/workshop-1/pkg/object"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sync"
	"time"
)

type ObjectRepository interface {
	Create(string) (uint64, error)
	Get(uint64) (*object.Object, error)
	Edit(uint64, string) error
	Delete(uint64) error
}

type Cache struct {
	cache       map[uint64]*object.Object
	mu          sync.RWMutex
	idGenerator generator.Generator
}

func New() ObjectRepository {
	m := make(map[uint64]*object.Object)
	idGen := generator.New(0, 1)
	return &Cache{cache: m, idGenerator: idGen}
}

func (c *Cache) Create(name string) (uint64, error) {
	newId := c.idGenerator.Next()
	newObject := object.New(newId, time.Now().Unix(), name)
	c.mu.Lock()
	c.cache[newId] = newObject
	c.mu.Unlock()
	return newId, nil
}

func (c *Cache) Get(id uint64) (*object.Object, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if obj, ok := c.cache[id]; ok {
		return obj, nil
	}
	return nil, status.Error(codes.NotFound, "no object")
}

func (c *Cache) Edit(id uint64, newName string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if obj, ok := c.cache[id]; ok {
		obj.Data.Ts = time.Now().Unix()
		obj.Data.Name = newName
		return nil
	}
	return status.Error(codes.NotFound, "no object")
}

func (c *Cache) Delete(id uint64) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.cache[id]; ok {
		delete(c.cache, id)
		return nil
	}
	return status.Error(codes.NotFound, "no object")
}
