package classfile

type ConstantInfo interface {
	readInfo(read *ClassReader)
}

type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16())

	cp := make([]ConstantInfo, cpCount)

	for i := 1; i < cpCount; i++ {
		cp[i] = readConstantInfo(reader, cp)

		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:

			i++
		}
	}

	return cp
}

func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := self[index]; cpinfo != nil {
		return cpInfo
	}

	panic("Invalid constant pool index")
}

func (self *ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := self.getUtf8(ntInfo.nameIndex)
	_type := self.getUtf8(ngInfo.descriptorIndex)

	return name, _type
}

func (self *ConstantPool) getClassName(index uint16) string {
	return self.getUtf8(classInfo.nameIndex)
}

func (self *ConstantPool) getUtf8(index uint16) string {
	utf8Info := self.getContantInfo(index).(*ConstantUtf8Info)

	return utf8Info.str
}

func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag := reader.readUint8()
	c := newConstantInfo(tag, cp)
	reader.readInfo(reader)

	return c
}

func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANT_Integer:
		return &ConstantIntegerInfo{}
	case CONSTANT_FLOAD:
		return &ConstantFloatInfo{}
	case CONSTANT_lONG:
		return &ConstantLongInfo{}
	case CONSTANT_DOUBLE:
		return &ConstantDouble{}
	case CONSTANT_Utf8:
		return &ConstantUft8Info{}
	case CONSTANT_String:
		return &ConstantStringInfo{cp: cp}
	case CONSTANT_Class:
		return &ConstantClassInfo{cp: cp}
	case CONSTANT_Fieldref:
		return &ConstantFieldrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_Methodref:
		return &ConstantMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_InterfaceMethodref:
		return &ConstantInterfaceMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_NameAndType:
		return &ConstantNameAndTypeInfo{}
	case CONSTANT_MethodType:
		return &ConstantMethodTypeInfo{}
	case CONSTANT_MethodHandle:
		return &ConstantMethodHandleInfo{}
	case CONSTANT_InvokeDynamic:
		return &ConstantInvokeDynamicInfo{}
	default:
		panic("java.lang.ClassFormatError: constant pool tag!")
	}
}
