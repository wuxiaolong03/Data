package utils

/**
 *将一个int64转换为[]byte字节切片
 */
func Int64Tobyte(num int64){
	//Buffer:缓冲区。增益
	buff := new(bytes.Buffer)//通过new实例化一个缓冲区
	//bytes.Wuffer 通过一些列的Write方法向缓冲区写入数据
	//buff.Bytes() 通过Bytes方法从缓冲区中获得数据

	/**
	 *两种排列方式：
	 *			大端位序排列：binary.BigEndian
	 *			小端位序排列：binary.LitteEndian
	 */

	err := binary.Write(buff,binary.BigEndian,num)
	if err != nil{
		return nil,err
	}
	//从缓冲区读取数据
	return buff.Bytes(),nil
}
