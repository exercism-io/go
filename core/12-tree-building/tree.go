// Package tree includes a solution for the "Tree Building" problem in the Go track on https://exercism.io.
package tree

import (
	"errors"
	"sort"
)

// Record contains an ID number and a parent ID number.
type Record struct {
	ID     int
	Parent int
}

// Node is used to build a Tree.
type Node struct {
	ID       int
	Children []*Node
}

// Build builds a tree from an array of input records.
func Build(input []Record) (*Node, error) {
	if len(input) == 0 {
		return nil, nil
	}

	sort.Slice(input, func(i, j int) bool {
		return input[i].ID < input[j].ID
	})

	nodes := make([]*Node, len(input))
	for i, r := range input {
		if i != r.ID {
			return nil, errors.New("invalid input")
		}
		if r.ID == 0 && input[i].Parent != 0 {
			return nil, errors.New("Parent ID must be 0 for root record")
		}
		if r.ID < r.Parent || (r.ID == r.Parent && r.ID != 0) {
			return nil, errors.New("Parent ID must be smaller than the current record ID")
		}

		node := Node{
			ID: i,
		}
		nodes[i] = &node

		if i == 0 {
			continue
		}

		nodes[r.Parent].Children = append(nodes[r.Parent].Children, &node)
	}

	return nodes[0], nil
}
