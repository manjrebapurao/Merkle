package main

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
)
type Node struct {
    Left  *Node
    Right *Node
    Hash  string
}

type MerkleTree struct {
    Root *Node
}

func NewNode(left, right *Node, hash string) *Node {
    return &Node{
        Left:  left,
        Right: right,
        Hash:  hash,
    }
}

func NewLeafNode(data string) *Node {
    hash := sha256.Sum256([]byte(data))
    return NewNode(nil, nil, hex.EncodeToString(hash[:]))
}

func NewInternalNode(left, right *Node) *Node {
    hash := sha256.Sum256([]byte(left.Hash + right.Hash))
    return NewNode(left, right, hex.EncodeToString(hash[:]))
}
func BuildMerkleTree(data []string) *MerkleTree {
    var nodes []*Node

    for _, datum := range data {
        nodes = append(nodes, NewLeafNode(datum))
    }

    for len(nodes) > 1 {
        var newLevel []*Node

        for i := 0; i < len(nodes); i += 2 {
            if i+1 == len(nodes) {
                // If the number of nodes is odd, duplicate the last node
                newLevel = append(newLevel, nodes[i])
            } else {
                newLevel = append(newLevel, NewInternalNode(nodes[i], nodes[i+1]))
            }
        }

        nodes = newLevel
    }

    return &MerkleTree{Root: nodes[0]}
}

func PrintTree(node *Node, level int) {
    if node == nil {
        return
    }
    fmt.Printf("%s%s\n", string(level*'\t'), node.Hash)
    PrintTree(node.Left, level+1)
    PrintTree(node.Right, level+1)
}

func main() {
    data := []string{"a", "b", "c", "d"}
    merkleTree := BuildMerkleTree(data)
    PrintTree(merkleTree.Root, 0)
}
