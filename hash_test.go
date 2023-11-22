package hash

import (
	"reflect"
	"testing"
)

func TestNewHashTable(t *testing.T) {
	type args struct {
		s       uint
		indexer Indexer
	}
	tests := []struct {
		name string
		args args
		want *HashTable
	}{
		{
			name: "success creating",
			args: args{
				s:       10,
				indexer: NewModuloIndexer(10),
			},
			want: &HashTable{
				size:    10,
				items:   make([]any, 10),
				indexer: NewModuloIndexer(10),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHashTable(tt.args.s, tt.args.indexer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHashTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashTable_Insert(t *testing.T) {
	type fields struct {
		size    uint
		items   []any
		indexer Indexer
	}
	type args struct {
		key   int
		value any
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "empty hashTable",
			fields:  fields{size: 0, items: make([]any, 0), indexer: NewModuloIndexer(0)},
			args:    args{key: 11, value: 11},
			wantErr: true,
		},
		{
			name:    "success insert",
			fields:  fields{size: 10, items: make([]any, 10), indexer: NewModuloIndexer(10)},
			args:    args{key: 1, value: 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HashTable{
				size:    tt.fields.size,
				items:   tt.fields.items,
				indexer: tt.fields.indexer,
			}
			if err := h.Insert(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHashTable_Search(t *testing.T) {
	type fields struct {
		size    uint
		items   []any
		indexer Indexer
	}
	type args struct {
		key int
	}
	items := make([]any, 10)
	items[5] = "success"

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    any
		wantErr bool
	}{
		{
			name:    "empty hashTable",
			fields:  fields{size: 0, items: make([]any, 0), indexer: NewModuloIndexer(0)},
			args:    args{key: 11},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "not found",
			fields:  fields{size: 10, items: make([]any, 10), indexer: NewModuloIndexer(10)},
			args:    args{key: 1},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "found",
			fields:  fields{size: 10, items: items, indexer: NewModuloIndexer(10)},
			args:    args{key: 5},
			want:    "success",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HashTable{
				size:    tt.fields.size,
				items:   tt.fields.items,
				indexer: tt.fields.indexer,
			}
			got, err := h.Search(tt.args.key)
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

func TestHashTable_Delete(t *testing.T) {
	type fields struct {
		size    uint
		items   []any
		indexer Indexer
	}
	type args struct {
		key int
	}
	items := make([]any, 10)
	items[5] = "success"

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "empty hashTable",
			fields:  fields{size: 0, items: make([]any, 0), indexer: NewModuloIndexer(0)},
			args:    args{key: 11},
			wantErr: true,
		},
		{
			name:    "not found",
			fields:  fields{size: 10, items: make([]any, 10), indexer: NewModuloIndexer(10)},
			args:    args{key: 1},
			wantErr: true,
		},
		{
			name:    "found",
			fields:  fields{size: 10, items: items, indexer: NewModuloIndexer(10)},
			args:    args{key: 5},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HashTable{
				size:    tt.fields.size,
				items:   tt.fields.items,
				indexer: tt.fields.indexer,
			}
			if err := h.Delete(tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHashTable_InsertLinearProbing(t *testing.T) {
	type fields struct {
		size          uint
		items         []any
		indexer       Indexer
		countElements uint
	}
	type args struct {
		key   int
		value any
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "empty hashTable",
			fields:  fields{size: 0, items: make([]any, 0), indexer: NewModuloIndexer(0)},
			args:    args{key: 11, value: 11},
			wantErr: true,
		},
		{
			name:    "success insert",
			fields:  fields{size: 10, items: make([]any, 10), indexer: NewModuloIndexer(10)},
			args:    args{key: 1, value: 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HashTable{
				size:          tt.fields.size,
				items:         tt.fields.items,
				indexer:       tt.fields.indexer,
				countElements: tt.fields.countElements,
			}
			if err := h.InsertLinearProbing(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("InsertLinearProbing() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHashTable_InsertLinearProbing_details(t *testing.T) {
	ht := NewHashTable(10, NewModuloIndexer(10))
	ht.InsertLinearProbing(1, 1)
	checkCondition(ht.items[1] == 1, "incorrect keys", t)

	ht.InsertLinearProbing(11, 11)
	checkCondition(ht.items[2] == 11, "incorrect keys", t)
}
