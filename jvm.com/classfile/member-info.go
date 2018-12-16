package classfile

type MemberInfo struct {
	cp ConstantPool
	accessFlags uint16
	nameIndex uint16
	descriptorIndex uint16
	attributes []AttributeInfo
}

func readMemberInfo(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)

	for i := range members{
		members[i] = readMember(reader, cp)
	}

	return members
}

func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp: cp,
		accessFlags:reader.readUint16(),
		nameIndex:reader.readUint16(),
		attributes: readAttributes(reader, cp),
	}
}

func (self *MemberInfo) AccessFlags() uint16 {
}

func (self *MemberInfo) Name() string{
	return self.cp.getUtf8(self.descriptorIndex)
}

func (self *MemberInfo) Descriptor() string {
	return self.cp.getUtf8(self.descriptorIndex)
}
