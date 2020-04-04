package classfile

import (
	"fmt"
	"go/ast"
)

type ClassFile struct {
	maginc uint32
	minorVersion uint16
	majorVersion uint16

	constantPool ConstantPool
	accessFlag uint16
	thisClass uint16
	superClass uint16
	interfaces []uint16
	fields []*MemberInfo
	method []*MemberInfo
	attributes []AttributeInfo

}

func (self* ClassFile)Parse(classData []byte)(cf *ClassFile,err error)  {

}

func (self* ClassFile)read(read *ClassReader)()  {
	self.readAndCheckMaging(read)
	self.readAndCheckVersion(read)
	// add read pool, read access flag
	// read className, read Superclass
	// read interfers, read field
	// read method  read attributed
	self.constantPool = read
}

func (self* ClassFile)readAndCheckMaging(read *ClassReader)  {
		magic := read.readUint32()
		if magic != 0xCAFEBABE{
			fmt.Errorf("read magic error \n")
			panic("java.class.ClasssFormateError: maigc")
		}
}

func (self* ClassFile)readAndCheckVersion(read *ClassReader)  {
	self.minorVersion = read.readUint16()
	self.majorVersion = read.readUint16()

	switch self.MagicVersion() {
	case 45:
		// java 1.2 has minorVersion
		return
	case 46,47,48,49,50,51,52:
		// after jdk 1.2, minorVersion is 0
		if self.minorVersion == 0{
			return
		}

	panic("java.lang.UnsupportedClassVersionError!")
	}
}

func (self* ClassFile)MinorVersion() uint16  {
		return self.minorVersion
}

func (self* ClassFile)MagicVersion()uint16  {

}

func (self* ClassFile)ConstantPool()ConstantPool  {

}

func (self* ClassFile)AccessFlags()uint16  {

}

func (self* ClassFile)Fields()[]*MemberInfo  {

}

func (self* ClassFile)ClassName()[]*MemberInfo  {
	return self.constantPool.getClassName(self.thisClass)
}

func (self* ClassFile)SuperClassName()string  {
	if(self.superClass > 0){
		return self.constantPool.getClassName(self.superClass)
	}
	return ""
}

func (self* ClassFile)InterfaceNames()[]string  {
	interfaceName := make([]string,len(self.interfaces))
	for i,cpIndex := range self.interfaces{
		interfaceName[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceName
}

func Parse(classData []byte)(cf *ClassFile,err error)  {
	defer func() {
		if r := recover();r != nil{
			var ok bool
			err,ok = r.(error)
			if !ok{
				err = fmt.Errorf("%v",r)
			}
		}
	}()

	cr := &ClassReader{data:classData}

	cf = &ClassFile{};
	cf.read(cr)
	return
}
