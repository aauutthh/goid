package main

import (
	"fmt"
	rbt "github.com/aauutthh/go-rbtree"
	"strings"
)

var StrComparetor = func(a, b interface{}) int { return strings.Compare(a.(string), b.(string)) }

func main() {
	tree := rbt.NewTree(StrComparetor)
	kvs := []struct{ k, v string }{
		{"ab", "hello"},
		{"bc", "world"},
	}
	for _, kv := range kvs {
		tree.Put(kv.k, kv.v)
	}
	k := "ab"
	if ok, v := tree.Get(k); ok {
		fmt.Printf("%#v : %#v\n", k, v)
	}
	fmt.Printf("size = %d\n", tree.Size())
	kvs2 := []struct{ k, v string }{
		{"cd", "hello2"},
		{"de", "world2"},
	}
	for _, kv := range kvs2 {
		tree.Put(kv.k, kv.v)
	}
	fmt.Printf("size = %d\n", tree.Size())
}
