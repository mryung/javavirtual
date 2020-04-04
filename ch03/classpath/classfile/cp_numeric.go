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

/**
CONSTANT_Utf8 {
	u1 tag
    u2 length
	u1 bytes[length]
}
字段名，字段描述
 */
type ConstantUtf8Info struct {
	val string
}

func (self * ConstantUtf8Info)readInfo(reader * ClassReader)  {
	len := uint32(reader.readUint32())
	bytes :=reader.readBytes(len)
	self.val = decodeMUTF8(bytes)
}

func decodeMUTF8(bytes []byte) string {
	return string(bytes)
}

/**
CONSTANT_String {
	u1 tag
    u2 string_index
}
指向常量池中Utf8类型
string 字面量
 */
type ConstantStringInfo struct {
	cp ConstantPool
	stringIndex uint16
}

func (self * ConstantStringInfo)readInfo(reader * ClassReader)  {
	self.stringIndex = reader.readUint16()
}

func (self *ConstantStringInfo)String() string  {
	return self.cp.getUtf8(self.stringIndex)
}

/**
CONSTANT_Class_info {
	u1 tag
    u2 name_index
}
指向常量池中Utf8类型
string 字面量
 */
type ConstantClassInfo struct {
	cp ConstantPool
	name_index uint16
}

func (self * ConstantClassInfo)readInfo(reader * ClassReader)  {
	self.name_index = reader.readUint16()
}

func (self * ConstantClassInfo)Name()string  {
	return self.cp.getUtf8(self.name_index)
}

/**
	CONSTANT_Fieldred = 9
	CONSTANT_Methodrep = 10
	CONSTANT_InterfaceMethodref = 11
	接口相同

 	CONSTANT_Fieldred_info {
		u1 tag
		u2 class_index
		u2 name_and_type_index
	}
 */
type ConstantMemberrefInfo struct {
	cp ConstantPool
	classIndex uint16
	nameAndTypeIndex uint16
}

func (self * ConstantMemberrefInfo)readInfo(reader * ClassReader) {
	self.classIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}

func (self * ConstantMemberrefInfo)ClassName()string  {
	return self.cp.getUtf8(self.classIndex)
}

func (self * ConstantMemberrefInfo)NameAndDescriptor()string  {
	return self.cp.getUtf8(self.nameAndTypeIndex)
}


type ConstantFieldrefInfo struct {
	ConstantMemberrefInfo
}

type ConstantMethodrefInfo struct {
	ConstantMemberrefInfo
}

type ConstantInterfaceMethodrefInfo struct {
	ConstantMemberrefInfo
}

