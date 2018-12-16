package classfile

type AttributeInfo interface {
	readInfo (reader *ClassReader)
}
