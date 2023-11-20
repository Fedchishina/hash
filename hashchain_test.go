package hash

import (
	"reflect"
	"testing"
)

func TestNewHashTableChain(t *testing.T) {
	type args struct {
		size    uint
		indexer Indexer
	}
	tests := []struct {
		name string
		args args
		want *HashTableChain
	}{
		{
			name: "success creating",
			args: args{
				size:    10,
				indexer: NewModuloIndexer(10),
			},
			want: &HashTableChain{
				size:    10,
				items:   make(map[uint]*node, 10),
				indexer: NewModuloIndexer(10),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHashTableChain(tt.args.size, tt.args.indexer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHashTableChain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashTableChain_Insert(t *testing.T) {
	type args struct {
		key   int
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "empty hashTable",
			args:    args{key: 11, value: 11},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ht := NewHashTable(10, NewModuloIndexer(10))
			if err := ht.Insert(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHashTableChain_InsertNext(t *testing.T) {
	ht := NewHashTableChain(10, NewModuloIndexer(10))
	ht.Insert(1, 1)
	first := ht.items[1]
	checkCondition(first.key == 1, "incorrect keys", t)
	checkCondition(first.value == 1, "incorrect values", t)
	checkCondition(first.next == nil, "next is not null", t)

	ht.Insert(11, 11)
	firstAfterInsert := ht.items[1]
	checkCondition(firstAfterInsert.key == 11, "incorrect keys", t)
	checkCondition(firstAfterInsert.value == 11, "incorrect values", t)
	checkCondition(firstAfterInsert.next != nil, "next is null", t)

	checkCondition(firstAfterInsert.next.key == 1, "incorrect keys", t)
	checkCondition(firstAfterInsert.next.value == 1, "incorrect values", t)
	checkCondition(firstAfterInsert.next.next == nil, "next is not null", t)
}

func TestHashTableChain_Search(t *testing.T) {
	tests := []struct {
		name    string
		ht      *HashTableChain
		key     int
		want    interface{}
		wantErr bool
	}{
		{
			name:    "empty hashTable",
			ht:      getHashTable([]int{}),
			key:     1,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "not found",
			ht:      getHashTable([]int{1, 2, 3}),
			key:     8,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "found",
			ht:      getHashTable([]int{1, 2, 3}),
			key:     1,
			want:    1,
			wantErr: false,
		},
		{
			name:    "found in next",
			ht:      getHashTable([]int{1, 2, 3, 11}),
			key:     11,
			want:    11,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ht.Search(tt.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Search() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashTableChain_Delete(t *testing.T) {
	tests := []struct {
		name    string
		ht      *HashTableChain
		key     int
		wantErr bool
	}{
		{
			name:    "empty hashTable",
			ht:      getHashTable([]int{}),
			key:     1,
			wantErr: true,
		},
		{
			name:    "not found",
			ht:      getHashTable([]int{1, 2, 3}),
			key:     8,
			wantErr: true,
		},
		{
			name:    "found",
			ht:      getHashTable([]int{1, 2, 3}),
			key:     1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := tt.ht.Delete(tt.key); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHashTableChain_DeleteNext(t *testing.T) {
	ht := NewHashTableChain(10, NewModuloIndexer(10))
	ht.Insert(1, 1)
	ht.Insert(11, 11)

	ht.Delete(1)

	firstAfterDelete := ht.items[1]
	checkCondition(firstAfterDelete.key == 11, "incorrect keys", t)
	checkCondition(firstAfterDelete.value == 11, "incorrect values", t)
	checkCondition(firstAfterDelete.next == nil, "next is not null", t)
}

func checkCondition(condition bool, message string, t *testing.T) {
	if condition != true {
		t.Errorf(message)
	}
}

func getHashTable(items []int) *HashTableChain {
	ht := NewHashTableChain(10, NewModuloIndexer(10))
	for _, el := range items {
		ht.Insert(el, el)
	}

	return ht
}
