package classfile

const (
	CONSTANT_Class              = 7
	CONSTANT_String             = 8
	CONSTANT_Fieldred           = 9
	CONSTANT_Methodrep          = 10
	CONSTANT_InterfaceMethodref = 11
	CONSTANT_Interger           = 2
	CONSTANT_Float              = 4
	CONSTANT_Long               = 5
	CONSTANT_Double             = 6
	CONSTANT_NameAndType        = 12
	CONSTANT_Utf8               = 1
	CONSTANT_MethodHandler      = 15
	CONSTANT_MethodType         = 16
	CONSTANT_InvokeDynamic      = 18
)

type ConstantInfo interface {
	readInfo(read *ClassReader)
}

func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag := reader.readUint8()
	c := newConstantInfo(tag, cp)
	c.readInfo(reader)
	return c
}

func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANT_Interger:
		return &ConstantIntegerInfo{}
	case CONSTANT_Float:
		return &ConstantFloatInfo{}
	case CONSTANT_Long:
		return &ConstantLongInfo{}
	case CONSTANT_Double:
		return &ConstantDoubleInfo{}
	case CONSTANT_NameAndType:
		return &ConstantNameAndTypeInfo{}
	case CONSTANT_Utf8:
		return &ConstantUtf8Info{}
	case CONSTANT_Class:
		return &ConstantClassInfo{}
	case CONSTANT_String:
		return &ConstantStringInfo{}
	case CONSTANT_Fieldred:
		return &ConstantFieldrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_Methodrep:
		return &ConstantMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_InterfaceMethodref:
		return &ConstantInterfaceMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_MethodHandler:
	case CONSTANT_MethodType:
	case CONSTANT_InvokeDynamic:
	default:
		panic("java.lang.classFormatError: constant pool tag")

	}
}
