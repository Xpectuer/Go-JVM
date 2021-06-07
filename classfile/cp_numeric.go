package classfile

import "math"

type ConstantIntegerInfo struct {
	val int32
}

func (cii *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	cii.val = int32(bytes)
}

type ConstantFloatInfo struct {
	val float32
}

func (cfi *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	cfi.val = math.Float32frombits(bytes)
}

type ConstantLongInfo struct {
	val int64
}

func (cli *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	cli.val = int64(bytes)
}

type ConstantDoubleInfo struct {
	val float64
}

func (cdi *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	cdi.val = math.Float64frombits(bytes)
}
