# hard / problem4 — Index

Theme: **binary search trees**. One task, written in `task.go`.
Run `go test ./problem4` from the `hard/` folder to check.

## Instructions

A dictionary app stores words with a number for each (how many times the word
was looked up). Besides finding a word, it must list all words in alphabetical
order, and, when a word is not there, suggest the closest words before and
after it.

A Go map can't do the last two without sorting everything each time. A binary
search tree keeps its keys in order all the time.

Write an `Index` type backed by a binary search tree you build yourself, with
the methods below.

## Allowed packages

None.

## Expected types and functions

```go
package problem4

type Index struct {
    // your fields
}

func NewIndex() *Index

func (ix *Index) Put(key string, val int)
func (ix *Index) Get(key string) (int, bool)
func (ix *Index) Len() int
func (ix *Index) Keys() []string
func (ix *Index) Floor(key string) (string, bool)
func (ix *Index) Ceiling(key string) (string, bool)
```

The node type of the tree is yours to define. The tests only ever use an
`*Index` returned by `NewIndex`.

## Rules

1. Keys are compared with Go's `<` on strings (byte by byte), so `"B" < "a"`
   and `"a" < "aa" < "b"`. Any string is a valid key, including `""`.
2. `Put` stores `val` under `key`. If `key` is already there, its value is
   replaced; the key is not added twice.
3. `Get` returns the value under `key` and `true`, or `0, false` if `key` is
   not there.
4. `Len` returns how many different keys are stored.
5. `Keys` returns all keys in sorted order, as a new slice: changing that slice
   must not change the `Index`. An empty `Index` returns an empty slice (`nil`
   is fine).
6. `Floor(key)` returns the largest stored key that is `<= key`, and `true`.
   If there is none, it returns `"", false`.
7. `Ceiling(key)` returns the smallest stored key that is `>= key`, and `true`.
   If there is none, it returns `"", false`.
8. Two `Index` values from two `NewIndex()` calls are independent.
9. Do not store the keys in a map or a sorted slice: the point is the tree.

## Examples

```go
ix := NewIndex()
ix.Put("mango", 3)
ix.Put("apple", 5)
ix.Put("peach", 1)
ix.Put("cherry", 8)
ix.Put("kiwi", 2)
ix.Put("apple", 7)      // replaces 5

ix.Len()                // -> 5
ix.Get("apple")         // -> 7, true
ix.Get("banana")        // -> 0, false
ix.Keys()               // -> ["apple" "cherry" "kiwi" "mango" "peach"]

ix.Floor("lemon")       // -> "kiwi", true     (kiwi < lemon < mango)
ix.Ceiling("lemon")     // -> "mango", true
ix.Floor("kiwi")        // -> "kiwi", true     (the key itself counts)
ix.Ceiling("kiwi")      // -> "kiwi", true
ix.Floor("aardvark")    // -> "", false        (nothing comes before it)
ix.Ceiling("zebra")     // -> "", false        (nothing comes after it)
```

The tree after the puts above (each node holds a key and its value):

```text
             mango:3
            /       \
       apple:7     peach:1
            \
          cherry:8
               \
              kiwi:2
```

## Hint

`Floor`: walk down from the root. When the node's key is bigger than `key`, the
answer is to the left. When it is smaller, this node is a candidate, but there
may be a better one to the right.
