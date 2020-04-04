package classfile

type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

func readAttributes(reader *ClassReader,cp ConstantPool)[]AttributeInfo  {
	attrCount := reader.readUint16()
	attrs := make([]AttributeInfo,attrCount)
	for i := range attrs  {
		attrs[i] = readAttribute(reader,cp)
	}
	return attrs
}

func readAttribute(reader *ClassReader,cp ConstantPool)AttributeInfo  {
		attrNameIndex := reader.readUint16()
		attrName := cp.getUtf8(attrNameIndex)
		attrLen := reader.readUint32()
		attrInfo := newAttributeInfo(attrName,attrLen,cp)
		attrInfo.readInfo(reader)
		return attrInfo
}


func newAttributeInfo(attrName string,attrlen uint32,cp ConstantPool)AttributeInfo  {
	switch attrName {
	case "CODE":
		break
	case "ConstantValue":
		break
	case "Deprecated":
		break
	case "Excepions":
		break
	case "LineNumberTable":
		break
	case "LocalVariableTable":
		break
	case "SourceFile":
		break
	case "Synthetic":
		break



	}
}
