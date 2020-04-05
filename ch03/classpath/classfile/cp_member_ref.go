package classfile

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

