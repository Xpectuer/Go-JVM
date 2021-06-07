package classfile

import "encoding/binary"

type ClassReader struct {
	data []byte
}

// readUint8 读取一位Uint8数字
func (cr *ClassReader) readUint8() uint8 {
	val := cr.data[0]
	// 跳过1个byte
	cr.data = cr.data[1:]
	return val
}

// readUint16 读取一位Uint16数字
func (cr *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(cr.data)
	cr.data = cr.data[2:]
	return val
}

// readUint32 读取一位Uint32数字
func (cr *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(cr.data)
	cr.data = cr.data[4:]
	return val
}

// readUint64 读取一位Uint64数字
func (cr *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(cr.data)
	cr.data = cr.data[8:]
	return val
}

// readUint16s 读取Uint16表格，表的大小由表头uint16数据给出
func (cr *ClassReader) readUint16s() []uint16 {
	n := cr.readUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = cr.readUint16()
	}
	return s

}

// readBytes 读取指定数量的字节
func (cr *ClassReader) readBytes(n uint32) []byte {
	bytes := cr.data[:n]
	cr.data = cr.data[n:]
	return bytes
}
