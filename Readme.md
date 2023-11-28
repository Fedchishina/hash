hashtable
=======================
Library for work with hash tables.

You can create a hash table and use a list of functions to work with it.

## Hash functions
- [Empty hash table creation example](#empty-hash-table-creation-example)
- [Insert key into hash table](#insert-key-into-hash-table)
- [Search element by key](#search-element-by-key)
- [Delete element by key](#delete-element-by-key)

### Empty hash table creation example
```
htInt := hash.NewHashTable[int, int](3) // simple hash table with int keys and values
htString := hash.NewHashTable[string, string](3) // simple hash table with string keys and values
htChain := hash.NewHashTableChain[int, int](3) // create hash table with int keys and values and chain realisation
```

### Insert key into hash table
```
ht := hash.NewHashTable[int, int](3) // simple hash table
ht.Insert(1, "value 1")
ht.Insert(2, "value 2")
```

### Search element by key
```
ht := hash.NewHashTable[int, string](3) // simple hash table
ht.Insert(1, "value 1")
ht.Insert(2, "value 2")

result, err := ht.Search(1) // return "value 1"
resultNil, err := ht.Search(3) // return "not found" error
```

### Delete element by key
```
ht := hash.NewHashTable[int, string](3) // simple hash table
ht.Insert(1, "value 1")
ht.Insert(2, "value 2")

err := ht.Delete(1) // deleting without error
err := ht.Delete(3) // return "not found" error
```
