package problem4

import "testing"

func eq(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// sample is the Index from the README.
func sample() *Index {
	ix := NewIndex()
	ix.Put("mango", 3)
	ix.Put("apple", 5)
	ix.Put("peach", 1)
	ix.Put("cherry", 8)
	ix.Put("kiwi", 2)
	ix.Put("apple", 7)
	return ix
}

func TestEmpty(t *testing.T) {
	ix := NewIndex()
	if got := ix.Len(); got != 0 {
		t.Errorf("empty: Len() = %d, want 0", got)
	}
	if got := ix.Keys(); len(got) != 0 {
		t.Errorf("empty: Keys() = %q, want []", got)
	}
	if v, ok := ix.Get("a"); v != 0 || ok {
		t.Errorf("empty: Get(\"a\") = %d, %v, want 0, false", v, ok)
	}
	if k, ok := ix.Floor("a"); k != "" || ok {
		t.Errorf("empty: Floor(\"a\") = %q, %v, want \"\", false", k, ok)
	}
	if k, ok := ix.Ceiling("a"); k != "" || ok {
		t.Errorf("empty: Ceiling(\"a\") = %q, %v, want \"\", false", k, ok)
	}
}

func TestPutGetLen(t *testing.T) {
	ix := sample()
	if got := ix.Len(); got != 5 {
		t.Errorf("Len() = %d, want 5 (apple was put twice)", got)
	}
	cases := []struct {
		key    string
		want   int
		wantOK bool
	}{
		{"apple", 7, true},
		{"cherry", 8, true},
		{"kiwi", 2, true},
		{"mango", 3, true},
		{"peach", 1, true},
		{"banana", 0, false},
		{"Apple", 0, false},
		{"", 0, false},
		{"appl", 0, false},
		{"apples", 0, false},
	}
	for _, c := range cases {
		v, ok := ix.Get(c.key)
		if v != c.want || ok != c.wantOK {
			t.Errorf("Get(%q) = %d, %v, want %d, %v", c.key, v, ok, c.want, c.wantOK)
		}
	}
}

func TestZeroValueAndEmptyKey(t *testing.T) {
	ix := NewIndex()
	ix.Put("", 0)
	ix.Put("z", 0)
	if v, ok := ix.Get(""); v != 0 || !ok {
		t.Errorf("Get(\"\") = %d, %v, want 0, true", v, ok)
	}
	if v, ok := ix.Get("z"); v != 0 || !ok {
		t.Errorf("Get(\"z\") = %d, %v, want 0, true", v, ok)
	}
	if got := ix.Len(); got != 2 {
		t.Errorf("Len() = %d, want 2", got)
	}
	if got, want := ix.Keys(), []string{"", "z"}; !eq(got, want) {
		t.Errorf("Keys() = %q, want %q", got, want)
	}
}

func TestKeys(t *testing.T) {
	ix := sample()
	want := []string{"apple", "cherry", "kiwi", "mango", "peach"}
	got := ix.Keys()
	if !eq(got, want) {
		t.Fatalf("Keys() = %q, want %q", got, want)
	}
	got[0] = "zzz"
	if again := ix.Keys(); !eq(again, want) {
		t.Errorf("changing the slice returned by Keys changed the Index: Keys() = %q", again)
	}
	if v, ok := ix.Get("apple"); v != 7 || !ok {
		t.Errorf("after changing the Keys slice, Get(\"apple\") = %d, %v, want 7, true", v, ok)
	}
}

func TestKeysSortedOrder(t *testing.T) {
	ix := NewIndex()
	for i, k := range []string{"b", "B", "ba", "a", "Z", "aa", "", "b0", "é", "z"} {
		ix.Put(k, i)
	}
	want := []string{"", "B", "Z", "a", "aa", "b", "b0", "ba", "z", "é"}
	if got := ix.Keys(); !eq(got, want) {
		t.Errorf("Keys() = %q, want %q", got, want)
	}
}

func TestFloorCeiling(t *testing.T) {
	ix := sample()
	cases := []struct {
		key       string
		floor     string
		floorOK   bool
		ceiling   string
		ceilingOK bool
	}{
		{"kiwi", "kiwi", true, "kiwi", true},
		{"banana", "apple", true, "cherry", true},
		{"lemon", "kiwi", true, "mango", true},
		{"aardvark", "", false, "apple", true},
		{"", "", false, "apple", true},
		{"zebra", "peach", true, "", false},
		{"apple", "apple", true, "apple", true},
		{"peach", "peach", true, "peach", true},
		{"peaches", "peach", true, "", false},
		{"cherr", "apple", true, "cherry", true},
		{"cherryx", "cherry", true, "kiwi", true},
		{"Mango", "", false, "apple", true},
	}
	for _, c := range cases {
		k, ok := ix.Floor(c.key)
		if k != c.floor || ok != c.floorOK {
			t.Errorf("Floor(%q) = %q, %v, want %q, %v", c.key, k, ok, c.floor, c.floorOK)
		}
		k, ok = ix.Ceiling(c.key)
		if k != c.ceiling || ok != c.ceilingOK {
			t.Errorf("Ceiling(%q) = %q, %v, want %q, %v", c.key, k, ok, c.ceiling, c.ceilingOK)
		}
	}
}

func TestSeparateIndexes(t *testing.T) {
	a := NewIndex()
	b := NewIndex()
	a.Put("x", 1)
	if b.Len() != 0 {
		t.Errorf("putting into one Index changed another: Len() = %d, want 0", b.Len())
	}
	if _, ok := b.Get("x"); ok {
		t.Errorf("putting into one Index changed another: Get(\"x\") found a value")
	}
}

func TestManyKeys(t *testing.T) {
	ix := NewIndex()
	words := []string{"m", "f", "t", "c", "i", "p", "w", "a", "d", "g", "k", "n", "r", "u", "y"}
	for i, w := range words {
		ix.Put(w, i)
	}
	for i, w := range words {
		ix.Put(w, i*10)
	}
	if ix.Len() != len(words) {
		t.Errorf("Len() = %d, want %d", ix.Len(), len(words))
	}
	for i, w := range words {
		if v, ok := ix.Get(w); v != i*10 || !ok {
			t.Errorf("Get(%q) = %d, %v, want %d, true", w, v, ok, i*10)
		}
	}
	want := []string{"a", "c", "d", "f", "g", "i", "k", "m", "n", "p", "r", "t", "u", "w", "y"}
	if got := ix.Keys(); !eq(got, want) {
		t.Errorf("Keys() = %q, want %q", got, want)
	}
	if k, ok := ix.Floor("j"); k != "i" || !ok {
		t.Errorf("Floor(\"j\") = %q, %v, want \"i\", true", k, ok)
	}
	if k, ok := ix.Ceiling("j"); k != "k" || !ok {
		t.Errorf("Ceiling(\"j\") = %q, %v, want \"k\", true", k, ok)
	}
	if k, ok := ix.Floor("b"); k != "a" || !ok {
		t.Errorf("Floor(\"b\") = %q, %v, want \"a\", true", k, ok)
	}
	if k, ok := ix.Ceiling("x"); k != "y" || !ok {
		t.Errorf("Ceiling(\"x\") = %q, %v, want \"y\", true", k, ok)
	}
}
