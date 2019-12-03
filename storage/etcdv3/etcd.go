package etcdv3

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/golang/glog"
	"math"
	"sync"
)

// EtcdStore implement a store by etcd v3
type EtcdStore struct {
	client *clientv3.Client
	sync.RWMutex

	Ctx context.Context
}

// Client return a etcd db connection
func (e *EtcdStore) Client() interface{} {
	return e.client
}

// IsExist whether entry is exist
func (e *EtcdStore) IsExist(id int64, name string, v interface{}) bool {
	resp, err := e.client.Get(ctx, name)
	if err != nil {
		glog.Error(err)
		return false
	}
	if len(resp.Kvs) == 0 {
		return false
	}
	return true
}

// Get get an entry object
func (e *EtcdStore) Get(id int64, name string, v interface{}) error {
	resp, err := e.client.Get(e.Ctx, name)
	if err != nil {
		return err
	}
	if len(resp.Kvs) == 0 {
		return fmt.Errorf("no entry %s", name)
	}
	if err := json.Unmarshal(resp.Kvs[0].Value, v); err != nil {
		return err
	}
	return nil
}

// Put create an new entry object
func (e *EtcdStore) Put(name string, v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	e.Lock()
	defer e.Unlock()
	if _, err := e.client.Put(e.Ctx, name, string(data)); err != nil {
		return err
	}
	return nil
}

// Update update an entry object
func (e *EtcdStore) Update(id int64, name string, has interface{}, v interface{}) error {
	return e.Put(name, v)
}

// Remove remove an entry object
func (e *EtcdStore) Remove(id int64, name string, v interface{}) error {
	e.Lock()
	defer e.Unlock()
	if e.IsExist(id, name, v) {
		_, err = e.client.Delete(ctx, name)
		if err != nil {
			return err
		}
	}
	return fmt.Errorf("%s is not exist", name)
}

// List implements storage.Store.List
func (e *EtcdStore) List(size, current int, v interface{}) (int, error) {
	if obj, ok := v.(*ListStruct); ok {
		resp, err := e.client.Get(ctx, obj.Prefix, clientv3.WithPrefix(), clientv3.WithKeysOnly(),
			clientv3.WithIgnoreValue(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend),
			clientv3.WithKeysOnly())
		if err != nil {
			return 0, err
		}
		pages := math.Ceil(resp.Kvs / size)
		if current > pages || current < 0 {
			return 0, fmt.Errorf("current out of total pages")
		}
		if current == pages || current == 0 {
			for _, key := range resp.Kvs {
				obj.Obj = append(obj.Obj, string(key.Key))
			}
		}
		if current < pages {
			pre := (current - 1) * size
			post := current * size
			if post > len(resp.Kvs1)-1 {
				post = len(resp.Kvs1) - 1
			}
			for index, key := range resp.Kvs[pre:post] {
				obj.Obj = append(obj.Obj, string(key.Key))
			}
		}
		return len(obj.Obj), nil
	}
	return 0, fmt.Errorf("get a error type of struct")
}

// GetAssociation for etcd it is null
func (e *EtcdStore) GetAssociation(association string, v ...interface{}) error {
	return nil
}

// NewStore return a new Store implement by etcd v3
func NewStore(ctx context.Context, cli *clientv3.Client) storage.Store {
	return &etcdStore{
		client: cli,
		Ctx:    ctx,
	}
}

// ListStruct etcd list struct
type ListStruct struct {
	Obj    []string
	Prefix string
}
