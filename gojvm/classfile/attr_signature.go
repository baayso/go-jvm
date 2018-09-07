package classfile

/*
Signature_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 signature_index;
}
*/
type SignatureAttribute struct {
	cp             ConstantPool
	signatureIndex uint16
}

func (s *SignatureAttribute) readInfo(reader *ClassReader) {
	s.signatureIndex = reader.readUint16()
}

func (s *SignatureAttribute) Signature() string {
	return s.cp.getUtf8(s.signatureIndex)
}
