package memdb

import (
	"container/list"
	"sync"
)

type EntityKey struct {
	/*
		Use a uuid or ulid
	*/
	Id string
	/*
		Functions the same as a person's surname, in an ideal world relatives (with same surname, at least one) would be aware of each others existence, in our case hold references.
		Ideally related records would be either clustered closely or hold reference or links to each other.
	*/
	Surnames []string
}

type Entity[T any] struct {
	/*
		Unique key
	*/
	Key EntityKey
	/*
		Value of the record, typed at a higher, "instance" level
	*/
	Value T
}

type Storage[T any] struct {
	store sync.Map
}

func (s *Storage[T]) DebugInit() {
	s.store.Clear()
}

/*
Stores reference of this record at each surname that this Entity has.

This record is not copied, if holding a reference and modifying the value of the record then that value is changed when fetched elsewhere as well.

Set itself is thread safe as it uses sync.Map but your struct type T should have a mutex if you are changing the value after the data is stored in this mem store.

In that case T itself should have setters and getters on itself to abtract locking and unlocking and prevent race conditions.
*/
func (s *Storage[T]) Set(e *Entity[T]) {
	for _, prefix := range e.Key.Surnames {
		cont := list.List{}
		cont.Init()
		v, ok := s.store.Load(prefix)
		if ok {
			cont = v.(list.List)
		}
		cont.PushBack(e)
		s.store.Store(prefix, cont)
	}
	s.store.Store(e.Key.Id, e)
}

func (s *Storage[T]) Get(id string) [](*Entity[T]) {
	var ls [](*Entity[T])
	if en, ok := s.store.Load(id); ok {
		if t, ok := en.(*Entity[T]); ok {
			return [](*Entity[T]){t}
		}
		if t, ok := en.(list.List); ok {
			for e := t.Front(); e != nil; e = e.Next() {
				if v, ok := e.Value.(*Entity[T]); ok {
					ls = append(ls, v)
				}
			}
			return ls
		}
	}
	return [](*Entity[T]){}
}
