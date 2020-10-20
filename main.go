package main

import (
	"fmt"

	"github.com/vldmkr/merkle-patricia-trie/mpt"
	"github.com/vldmkr/merkle-patricia-trie/storage"
)

func main() {
	store, _ := storage.NewLevelDBAdapter("BASE")
	// store := storage.NewMemoryAdapter()
	trie1 := mpt.New(nil, store)

	err := trie1.Put([]byte("key"), []byte("hello"))
	if err != nil {
		fmt.Printf("%e", err)
	}
	value, err := trie1.Get([]byte("key"))
	if err != nil {
		fmt.Printf("%e", err)
	}
	root1 := trie1.RootHash()
	fmt.Printf("trie1:root1: %x \n", root1)
	fmt.Printf("trie1:value: %s \n", value)
	trie1.Commit()

	err = trie1.Put([]byte("key"), []byte("world"))
	if err != nil {
		fmt.Printf("%e", err)
	}
	value, err = trie1.Get([]byte("key"))
	if err != nil {
		fmt.Printf("%e", err)
	}
	root2 := trie1.RootHash()
	fmt.Printf("trie1:root2: %x \n", root2)
	fmt.Printf("trie1:value: %s \n", value)
	trie1.Commit()

	nodeRoot1 := mpt.HashNode(root1)
	trie2 := mpt.New(&nodeRoot1, store)
	value, err = trie2.Get([]byte("key"))
	if err != nil {
		fmt.Printf("%e", err)
	}
	fmt.Printf("trie2:root1: %x \n", trie2.RootHash())
	fmt.Printf("trie2:value: %s \n", value)

	nodeRoot2 := mpt.HashNode(root2)
	trie3 := mpt.New(&nodeRoot2, store)
	value, err = trie3.Get([]byte("key"))
	if err != nil {
		fmt.Printf("%e", err)
	}
	fmt.Printf("trie3:root2: %x \n", trie3.RootHash())
	fmt.Printf("trie3:value: %s \n", value)
}
