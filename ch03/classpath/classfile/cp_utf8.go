package classfile

/**
CONSTANT_Utf8 {
	u1 tag
    u2 length
	u1 bytes[length]
}
字段名，字段描述
 */
type ConstantUtf8Info struct {
	str string
}

func (self * ConstantUtf8Info)readInfo(reader * ClassReader)  {
	len := uint32(reader.readUint16());
	bytes := reader.readBytes(len)
	self.str = decodeMUTF8(bytes);
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

