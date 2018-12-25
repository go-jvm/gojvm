package classfile

type BootstrapMethod struct {
	bootstrapMethodRef uint16
	bootstrapArguments []uint16
}

type BootstrapMethodsAttribute struct {
	bootstrapMethods []*BootstrapMethod
}

func (self *BootstrapMethodsAttribute) readInfo(reader *ClassReader) {
	numBootstrapMethods := reader.readUint16()
	self.bootstrapMethods = make([]*BootstrapMethod, numBootstrapMethods)
	for i := range self.bootstrapMethods {
		self.bootstrapMethods[i] = &BootstrapMethod{
			bootstrapMethodRef: reader.readUint16(),
			bootstrapArguments: reader.readUint16s(),
		}
	}
}
