package blockchain

import (
	"DataCertPlatform/utils"
	"bytes"
	"math/big"
)
/**
 *工作量证明算法结构体
 */
type ProofOfWork struct {
	Target big.Int//系统的目标值
	Block Block //要找的nonce值对应的区块
}

/**
 *实例化一个pow算法的实例
 */
func NewPoW(block Block) ProofOfWork{
	t := big.NewInt(1)
	t = t.Lsh(t,255)
	pow := ProofOfWork{
		Target: t,
		Block: block ,
	}
	return pow
}

/**
 *run方法用于寻找合适的nonce值
 */

func (p ProofOfWork) Run() int64{
	var nonce int64
	nonce = 0
	for { //不知道什么时候结束，使用无限循环
		block := p.Block
		heightBytes, _ := utils.Int64TpByte(Block.Height)
		timeBytes, _ := utils.Int64ToByte(p.Block.TimeStamp)
		versionBytes, _ := utils.StringToBytes(Block{}.Version)

		nonceBytes, _ := utils.Int64ToByte(nonce)
		//已有区块信息和尝试的nonce值的拼接信息
		blockBytes := bytes.Join([][]byte{
			heightBytes,
			timeBytes,
			block.PrevHash,
			block.Date,
			versionBytes,
			nonceBytes,
		}, []byte{})
		//区块和尝试的monce值拼接后得到的hash值
		blockHash:= utils.SHA256HashBlock(blockBytes)

		target := p.Target     //目标值
		var hashBig *big.Int   //声明和定义
		hashBig = new(big.Int) //分配内存空间，为变量分配地址
		//xx : invalid memory or nil pointer dereference 空指针错误
		hashBig = hashBig.SetBytes(blockHash)
		if  hashBig.Cmp(target)== -1{
			//停止寻找
			break
		}
		nonce++//自增，继续寻找
	}
	//将找到符合规则的nonce返回
	return BlockHash ,nonce
}