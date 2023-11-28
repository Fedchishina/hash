package hash

import (
	"reflect"
	"testing"
)

func TestNewHashTable(t *testing.T) {
	tests := []struct {
		name string
		size uint
		want *HashTable[int, int]
	}{
		{
			name: "success creating",
			size: 10,
			want: &HashTable[int, int]{
				size:  10,
				items: make([]*node[int, int], 10),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHashTable[int, int](tt.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHashTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashTable_Insert(t *testing.T) {
	tests := []struct {
		name    string
		ht      *HashTable[int, int]
		key     int
		wantErr bool
	}{
		{
			name:    "empty hashTable",
			ht:      getHashTable([]int{}),
			key:     11,
			wantErr: false,
		},
		{
			name:    "success insert",
			ht:      getHashTable([]int{1}),
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
		ht      *HashTable[int, int]
		key     int
		want    int
		wantErr bool
	}{
		{
			name:    "empty hashTable",
			ht:      getHashTable([]int{}),
			key:     11,
			want:    0,
			wantErr: true,
		},
		{
			name:    "not found",
			ht:      getHashTable([]int{1, 2, 3}),
			key:     9,
			want:    0,
			wantErr: true,
		},
		{
			name:    "found",
			ht:      getHashTable([]int{1, 2, 3}),
			key:     3,
			want:    3,
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
		ht      *HashTable[int, int]
		key     int
		wantErr bool
	}{
		{
			name:    "empty hashTable",
			ht:      getHashTable([]int{}),
			key:     11,
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
		ht      *HashTable[int, int]
		key     int
		wantErr bool
	}{
		{
			name:    "empty table - success insert",
			ht:      getHashTable([]int{}),
			key:     1,
			wantErr: false,
		},
		{
			name:    "success insert",
			ht:      getHashTable([]int{1}),
			key:     2,
			wantErr: false,
		},
		{
			name:    "success insert - neighbour",
			ht:      getHashTable([]int{1}),
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
	ht := NewHashTable[int, int](10)
	ht.Insert(1, 1)
	checkCondition(ht.items[2].key == 1, "incorrect keys", t)

	ht.Insert(11, 11)
	checkCondition(ht.items[8].key == 11, "incorrect keys", t)
}

func getHashTable(items []int) *HashTable[int, int] {
	ht := NewHashTable[int, int](10)
	for _, el := range items {
		ht.Insert(el, el)
	}

	return ht
}
