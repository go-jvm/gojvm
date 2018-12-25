package classfile

type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

// read attributes
func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	attributesCount := reader.readUint16()
	attributes := make([]AttributeInfo, attributesCount)

	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}

	return attributes
}

// read attr

func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {

	attrNameIndex := reader.readUint16()

	attrName := cp.getUtf8(attrNameIndex)

	attrLen := reader.readUint32()

	attrInfo := newAttributeInfo(attrName, attrLen, cp)
	attrInfo.readInfo(reader)

	return attrInfo
}


// attr Info
func newAttributeInfo(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {
	switch attrName {
	case "Code":
		return &CodeAttribute{cp: cp}
	case "ConstantValue":
		return &ConstantValueAttribute{}
	case "Deprecated":
		return &DeprecatedAttribute{}
	case "Exceptions":
		return &ExceptionsAttribute{}
	case "LineNumberTable":
		return &LineNumberTableAttribute{}
	case "LocalVariableTable":
		return &LocalVariableTableAttribute{}
	case "SourceFile":
		return &SourceFileAttribute{cp: cp}
	case "Synthetic":
		return &SyntheticAttribute{}
	default:
		return &UnparsedAttribute{attrName, attrLen, nil}
	}
}