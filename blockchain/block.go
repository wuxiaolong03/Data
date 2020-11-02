package blockchain

/**
 *定义区块结构体，用于表示区块
 */
type Block struct {
	Height    int64  //区块的高度，第几个区块
	TimeStamp int64  //区块产生的时间戳
	PrevHash  []byte //前一个区块的hash
	Date      []byte //数据字段
	Hash      []byte //当前区块的hash值
	Version   string //版本号
}

/**
 *创建一个新区快
 */

func NewBlock() {
	block := Block{
		Height:    height + 1,
		TimeStamp: time.Now().Unix(),
		PrebHash:  prevHasH,
		Data:      data,
		Version:   "0x01"
	}
	block.Hash = SHA256
	return block
}

/**
 *创建创世区块
 */
func CreateGenesisBlock() Block {
	genesisBlock := NewBlock(0,
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		nil)
	return genesisBlock
}
