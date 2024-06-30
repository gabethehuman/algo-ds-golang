## Hash Map with linear probing

### Features
* Collision resolution: Utilizes linear probing to resolve collisions.
* Dynamic resizing: Automatically doubles the hash map size and rehashes the existing entries when the load factor exceeds 50%.
* Hashmap operations:
    - `Insert`: Add key-value pairs, automatically resizing (doubling in capacity) and rehashing if load factor > 0.5.
    - `Delete`: Remove key-value pairs by marking them as deleted.
    - `Get`: Retrieve values using keys.

### Limitations
* Does not shrink in size after deletions, potentially wasting space.
* Works only for `string` keys and values.
* Error handling is not the best, but it's not the point of the exercise.

### Usage
```golang
hashmap := NewHashMap(10)

hashmap.Insert("key1", "value1")
hashmap.Insert("key2", "value2")

if value, ok := hashmap.Get("key1"); ok {
    fmt.Println("Found key1:", value)
}

hashmap.Delete("key1")
```
