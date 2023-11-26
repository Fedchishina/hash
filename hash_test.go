package hash

import (
	"reflect"
	"testing"
)

func TestNewHashTable(t *testing.T) {
	tests := []struct {
		name string
		size uint
		want *HashTable[IntKey]
	}{
		{
			name: "success creating",
			size: 10,
			want: &HashTable[IntKey]{
				size:  10,
				items: make(map[uint]*node[IntKey], 10),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHashTable[IntKey](tt.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHashTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashTable_Insert(t *testing.T) {
	tests := []struct {
		name    string
		ht      *HashTable[IntKey]
		key     IntKey
		wantErr bool
	}{
		{
			name:    "empty hashTable",
			ht:      getHashTable([]IntKey{}),
			key:     11,
			wantErr: false,
		},
		{
			name:    "success insert",
			ht:      getHashTable([]IntKey{1}),
			key:     2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.ht.Insert(tt.key, tt.key); (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHashTable_Search(t *testing.T) {

	tests := []struct {
		name    string
		ht      *HashTable[IntKey]
		key     IntKey
		want    any
		wantErr bool
	}{
		{
			name:    "empty hashTable",
			ht:      getHashTable([]IntKey{}),
			key:     11,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "not found",
			ht:      getHashTable([]IntKey{1, 2, 3}),
			key:     9,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "found",
			ht:      getHashTable([]IntKey{1, 2, 3}),
			key:     3,
			want:    IntKey(3),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ht.Search(tt.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else {
				if got != tt.want {
					t.Errorf("Search() got = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestHashTable_Delete(t *testing.T) {
	tests := []struct {
		name    string
		ht      *HashTable[IntKey]
		key     IntKey
		wantErr bool
	}{
		{
			name:    "empty hashTable",
			ht:      getHashTable([]IntKey{}),
			key:     11,
			wantErr: true,
		},
		{
			name:    "not found",
			ht:      getHashTable([]IntKey{1, 2, 3}),
			key:     8,
			wantErr: true,
		},
		{
			name:    "found",
			ht:      getHashTable([]IntKey{1, 2, 3}),
			key:     3,
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

func TestHashTable_InsertLinearProbing(t *testing.T) {
	tests := []struct {
		name    string
		ht      *HashTable[IntKey]
		key     IntKey
		wantErr bool
	}{
		{
			name:    "empty table - success insert",
			ht:      getHashTable([]IntKey{}),
			key:     1,
			wantErr: false,
		},
		{
			name:    "success insert",
			ht:      getHashTable([]IntKey{1}),
			key:     2,
			wantErr: false,
		},
		{
			name:    "success insert - neighbour",
			ht:      getHashTable([]IntKey{1}),
			key:     11,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.ht.Insert(tt.key, tt.key); (err != nil) != tt.wantErr {
				t.Errorf("InsertLinearProbing() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHashTable_InsertLinearProbing_details(t *testing.T) {
	ht := NewHashTable[IntKey](10)
	ht.Insert(1, 1)
	checkCondition(ht.items[1].key == 1, "incorrect keys", t)

	ht.Insert(11, 11)
	checkCondition(ht.items[2].key == 11, "incorrect keys", t)
}

func getHashTable(items []IntKey) *HashTable[IntKey] {
	ht := NewHashTable[IntKey](10)
	for _, el := range items {
		ht.Insert(el, el)
	}

	return ht
}
