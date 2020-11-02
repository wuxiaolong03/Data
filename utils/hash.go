package utils

/**
 *对区块数据进行SHA256哈希计算
 */
func SHA256HashBlock(block blockchain.Block) []byte{
	//1.将block结构体数据转换为[]byte类型
	heighBytes,_ := Int64Byte(block.Height)

	//2.将转换后的[]byte字节切片输入Write方法

	sha256Hash := sha256.New()
	sha256Hasha.Write()
	hash := sha256Hasha.Sun(nil)
	return hash
}
