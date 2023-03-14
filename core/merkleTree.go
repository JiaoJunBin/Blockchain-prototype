package core

import (
	"crypto/sha256"
	"github.com/pkg/errors"
)

// a binary merkle tree
type MerkleNode struct {
	Value  hash
	Parent *MerkleNode
	Lchild *MerkleNode
	Rchind *MerkleNode
}

// build a merkle tree from given transaction
func NewMerkleTree(txs []*Transaction) (root *MerkleNode, err error) {
	if len(txs) == 0 {
		return nil, errors.New("NewMerkleTree() error: doesn't cotain any transactions")
	}
	hList := make([]hash, 0)
	for _, tx := range txs {
		hList = append(hList, tx.ID)
	}
	return buildTree(hList), nil
}

// arg: hash of transactions return: root
func buildTree(hList []hash) (root *MerkleNode) {
	length := len(hList)
	root = &MerkleNode{}
	// if the number of txs is odd, add the last tx to the (root->Rchild)
	if len(hList)&1 == 1 {
		root.Rchind = &MerkleNode{
			Value:  hList[length-1],
			Parent: root,
			Lchild: nil,
			Rchind: nil,
		}
		hList = hList[:len(hList)-1]
	}

	oldNodeList := make([]*MerkleNode, 0)
	length /= 2
	newNodeList := make([]*MerkleNode, 0)

	// leaves nodes
	for _, h := range hList {
		node := &MerkleNode{
			Value:  h,
			Lchild: nil,
			Rchind: nil,
		}
		oldNodeList = append(oldNodeList, node)
	}

	tuple := make([]*MerkleNode, 2)
	// intermediate nodes, root | (root->left)
	for length != 0 {
		for i, oldNode := range oldNodeList {
			// create parent node for a pair
			if i&1 == 1 {
				tuple[1] = oldNode
				newNode := &MerkleNode{
					Value:  ComputeParentHash(tuple[0], tuple[1]),
					Lchild: tuple[0],
					Rchind: tuple[1],
				}
				newNodeList = append(newNodeList, newNode)
				tuple[0].Parent, tuple[1].Parent = newNode, newNode
			} else {
				tuple[0] = oldNode
			}
		}
		oldNodeList = newNodeList
		length /= 2
		newNodeList = make([]*MerkleNode, 0)
	}

	if len(hList)&1 == 1 {
		oldNodeList[0].Parent = root
		root.Lchild = oldNodeList[0]
		root.Value = ComputeParentHash(root.Lchild, root.Rchind)
	} else {
		root = oldNodeList[0]
	}
	return
}

func ComputeParentHash(m, n *MerkleNode) hash {
	mb := m.Serialize()
	nb := n.Serialize()
	return sha256.Sum256(append(mb, nb...))
}
