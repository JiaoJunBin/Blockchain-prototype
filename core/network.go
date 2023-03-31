package core

import (
	"encoding/gob"
	"fmt"
	"net"
)

/*
总体思路：
1.每个cli都开了一个端口，端口之间通信
2.使用net在不同cli之间发送信息
3.进行监听后，使用channel在进程里进行通信
*/

// 监听来自其他节点的信息，需要输入想获取的node的nodeID
func listen(block chan []*Block, transactions chan []*Transaction, nodeID string, informationType string) {
	node_Address := fmt.Sprintf("localhost:%s", nodeID)
	conn, err := net.Dial("tcp", node_Address)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	//判断接受到的信息是block还是Transaction
	if informationType == "block" {
		go func() {
			for {
				// 接收来自其他节点的区块信息
				var receivedBlocks []*Block
				decoder_block := gob.NewDecoder(conn)
				err := decoder_block.Decode(&receivedBlocks)
				if err != nil {
					panic(err)
				}

				block <- receivedBlocks
			}
		}()
	}

	if informationType == "transaction" {
		go func() {
			for {
				var receivedTransactions []*Transaction
				decoder_tran := gob.NewDecoder(conn)
				err := decoder_tran.Decode(&receivedTransactions)
				if err != nil {
					panic(err)
				}

				transactions <- receivedTransactions
			}
		}()
	}

}

// 向其他节点发送区块链信息
func sendblock(bc []Block, nodeID string) {
	node_Address := fmt.Sprintf("localhost:%s", nodeID)
	conn, err := net.Dial("tcp", node_Address)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	//io.WriteString(conn, "block\n")
	encoder := gob.NewEncoder(conn)
	err = encoder.Encode(bc)
	if err != nil {
		panic(err)
	}
}

// 向其他节点发送交易信息
func sendTransaction(tx []Transaction, nodeID string) {
	node_Address := fmt.Sprintf("localhost:%s", nodeID)
	conn, err := net.Dial("tcp", node_Address)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	//io.WriteString(conn, "transaction\n")
	encoder := gob.NewEncoder(conn)
	err = encoder.Encode(tx)
	if err != nil {
		panic(err)
	}
}
