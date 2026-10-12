# hard / problem1 — Serialize, Deserialize

Theme: **binary trees**. One task with two functions, written in `task.go`.
Run `go test ./problem1` from the `hard/` folder to check.

## Instructions

A tree lives in memory as nodes and pointers. To save it to a file, or to send
it over the network, it has to become a string, and later be turned back into
the very same tree.

Write `Serialize`, which turns a tree into a string, and `Deserialize`, which
reads such a string back into a tree, rejecting any string that is not
correctly formed.

### Format description

The tree is written in **pre-order**: the node, then its left subtree, then its
right subtree. Every node writes its value. Every missing child (`nil`) is
written as `#`. The items are separated by single commas, with nothing else
between them.

```text
    1
   / \
  2   3          ->   "1,2,#,#,3,4,#,#,5,#,#"
     / \
    4   5
```

Reading it item by item: `1` is the root. Its left subtree is `2,#,#` (a 2
with no children). Its right subtree is `3,4,#,#,5,#,#`.

## Allowed packages

`strconv`, `strings`.

## Expected type and functions

```go
package problem1

type Node struct {
    Val         int
    Left, Right *Node
}

func Serialize(root *Node) string
func Deserialize(s string) (*Node, bool)
```

## Rules

### Serialize

1. Output the tree in the format above. A `nil` tree is `"#"`.
2. Values are written in plain decimal, with a `-` in front of negative
   numbers.

### Deserialize

3. A valid string describes exactly one tree in the format above. Return that
   tree and `true`.
4. A value is an optional `-` followed by one or more digits (`0`-`9`), and it
   must fit in an `int`. Leading zeros are allowed (`007` is `7`).
5. Return `nil, false` when the string is not valid. This includes:
   - an empty string, or an empty item (`1,,#`, a trailing comma)
   - an item that is neither a value nor `#` (`+1`, `-`, `1a`, `1.5`, ` 1`)
   - items missing: the tree is not complete when the string ends (`1,#`)
   - items left over after the tree is complete (`1,#,#,#`)

## Examples

```go
Serialize(root)                   // -> "1,2,#,#,3,4,#,#,5,#,#"   (the tree above)
Serialize(nil)                    // -> "#"
Serialize(&Node{Val: -7})         // -> "-7,#,#"

Deserialize("1,2,#,#,3,4,#,#,5,#,#")  // -> the tree above, true
Deserialize("#")                  // -> nil, true     (the empty tree)
Deserialize("007,#,#")            // -> a single node 7, true

Deserialize("")                   // -> nil, false
Deserialize("1,#")                // -> nil, false    (the right child is missing)
Deserialize("1,#,#,#")            // -> nil, false    (one item too many)
Deserialize("+1,#,#")             // -> nil, false    (rule 4: no + sign)
Deserialize("1, #,#")             // -> nil, false    (no spaces)
```

Note: `Deserialize("#")` returns `nil, true`, and an invalid string returns
`nil, false`. The `bool` is what tells them apart.

## Hint

`Deserialize` mirrors `Serialize`: read one item; if it is `#` the subtree is
`nil`, otherwise it is a node, and its left and then its right subtree come
next in the string. Keep track of how far you have read.
