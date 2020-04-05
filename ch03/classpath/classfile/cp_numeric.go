package classfile

import "math"

/**
CONSTANT_Float{
u1 tag
u4 bytes
}
*/
type ConstantIntegerInfo struct {
	val int32
}

func (self *ConstantIntegerInfo) readInfo(reader *ClassReader)  {
	bytes := reader.readUint32()
	self.val = int32(bytes)
}

/**
CONSTANT_Float{
u1 tag
u4 bytes
}
*/
type ConstantFloatInfo struct {
	val float32
}

func (self * ConstantFloatInfo)readInfo(reader * ClassReader)  {
	bytes := reader.readUint32()
	self.val = math.Float32frombits(bytes)
}

/**
CONSTANT_Long{
u1 tag
u4 high_bytes
u4 low_bytes
}
 */
type ConstantLongInfo struct {
	val int64
}

func (self * ConstantLongInfo)readInfo(reader * ClassReader)  {
	bytes := reader.readUint64()
	self.val = int64(bytes)
}


/**
CONSTANT_Double{
u1 tag
u4 high_bytes
u4 low_bytes
}
 */
type ConstantDoubleInfo struct {
	val float64
}

func (self * ConstantDoubleInfo)readInfo(reader * ClassReader)  {
	bytes := reader.readUint64()
	self.val = math.Float64frombits(bytes)
}





