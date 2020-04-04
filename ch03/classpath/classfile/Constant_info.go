package classfile

const (
	CONSTANT_Class = 7
	CONSTANT_String = 8
	CONSTANT_Fieldred = 9
	CONSTANT_Methodrep = 10
	CONSTANT_InterfaceMethodref = 11
	CONSTANT_Interger = 2
	CONSTANT_Float = 4
	CONSTANT_Long = 5
	CONSTANT_Double = 6
	CONSTANT_NameAndType = 12
	CONSTANT_Utf8 = 1
	CONSTANT_MethodHandler = 15
	CONSTANT_MethodType = 16
	CONSTANT_InvokeDynamic = 18
)

type ConstantInfo interface {
	readINfo(read * ClassReader)
}

func reaConstantInfo(reader *ClassReader,cp ConstantPool) ConstantInfo  {
	 tag := reader.readUint8();
	 c := newConstanInfo(tag,cp)
	 return c
}



func newConstantInfo(tag uint8,cp ConstantPool)  {
	switch tag {
	case CONSTANT_Interger:
		
	}
}


