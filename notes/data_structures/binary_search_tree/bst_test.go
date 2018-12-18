// Тест для пакета "Двоичное дерево поиска"
package main

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	tree := initTree()
	fmt.Println("Inital tree:\n", tree)
	os.Exit(m.Run())
}

func TestInsert(t *testing.T) {
	tree := initTree()
	var failed bool
	if tree.root.Value != 10 {
		t.Error("Root should be 10, but it is:", tree.root.Value)
		failed = true
	}
	if tree.root.left.Value != 5 {
		t.Error("Root left should be 6, but it is:", tree.root.left.Value)
		failed = true
	}
	if tree.root.right.Value != 20 {
		t.Error("Root right should be 20, but it is:", tree.root.right.Value)
		failed = true
	}

	if !failed {
		t.Log("Insert Test passed - OK")
	}

}

func TestSearch(t *testing.T) {
	tree := initTree()
	var failed bool
	var tt = []struct {
		val int
		res bool
	}{
		{
			val: 10,
			res: true,
		},
		{
			val: 20,
			res: true,
		},
		{
			val: 12,
			res: false,
		},
		{
			val: 1,
			res: true,
		},
	}

	for _, tc := range tt {
		if res := tree.Search(tc.val); res != tc.res {
			t.Errorf("Searching for %d should return %t, but we got %t", tc.val, tc.res, res)
			failed = true
		}
	}

	if !failed {
		t.Log("Search Test passed - OK")
	}

}

//	Output:
//	Inital tree:
//					35
//				30
//			25
//		20
//			15
//	10
//			6
//		5
//				2
//			1
//
//	=== RUN   TestInsert
//	--- PASS: TestInsert (0.00s)
//	    c:\Users\dtsp\YandexDisk\asubk\src\gonotes\notes\data_structures\binary_search_tree\bst_test.go:32: Insert Test passed - OK
//	=== RUN   TestSearch
//	--- PASS: TestSearch (0.00s)
//	    c:\Users\dtsp\YandexDisk\asubk\src\gonotes\notes\data_structures\binary_search_tree\bst_test.go:70: Search Test passed - OK
//	PASS
//	coverage: 95.7% of statements
//	ok  	gonotes/notes/data_structures/binary_search_tree	0.203s	coverage: 95.7% of statements
//	Success: Tests passed.
