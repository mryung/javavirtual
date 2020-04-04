package classfile

import (
	"fmt"
	"encoding/binary"
)

type ClassReader struct {
	data[] byte
}

func (self * ClassReader)readUint8()uint8  {
	if len(self.data) < 1{
		fmt.Println("not data left 1")
		panic("not data left")
	}
	val := self.data[0]
	self.data = self.data[1:]
	return val;
}

func (self * ClassReader)readUint16()uint16  {
	if len(self.data) < 2{
		fmt.Println("not data left 2,now:",len(self.data))
		panic("not data left")
	}
	val := binary.BigEndian.Uint16(self.data);
	self.data = self.data[2:]
	return val
}

func (self * ClassReader)readUint32()uint32  {
	if len(self.data) < 4{
		fmt.Println("not data left 4,now:",len(self.data))
		panic("not data left")
	}
	val := binary.BigEndian.Uint32(self.data);
	self.data = self.data[4:]
	return val
}

func (self * ClassReader)readUint64()uint64  {
	if len(self.data) < 8{
		fmt.Println("not data left 8,now:",len(self.data))
		panic("not data left")
	}
	val := binary.BigEndian.Uint64(self.data);
	self.data = self.data[8:]
	return val
}

func (self * ClassReader)readUint16s()[]uint16{
	if len(self.data) < 2{
		fmt.Println("not data readUint16s left 2,now:",len(self.data))
		panic("not data left")
	}
	n := self.readUint16()
	val := make([]uint16,n)
	for _,i:= range val  {
		val[i] =self.readUint16()
	}
	return val
}

func (self * ClassReader)readBytes(n uint32) []byte  {
	fmt.Println("data left ",n,",now:",len(self.data))

	bytes := self.data[:n]
	self.data = self.data[n:]
	return bytes
}
