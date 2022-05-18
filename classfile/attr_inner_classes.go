package classfile

// InnerClassesAttribute
//
/*
InnerClasses_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 number_of_classes;
    {   u2 inner_class_info_index;
        u2 outer_class_info_index;
        u2 inner_name_index;
        u2 inner_class_access_flags;
    } classes[number_of_classes];
}
*/
type InnerClassesAttribute struct {
	classes []*InnerClassInfo
}

type InnerClassInfo struct {
	innerClassInfoIndex   uint16
	outerClassInfoIndex   uint16
	innerNameIndex        uint16
	innerClassAccessFlags uint16
}

func (a *InnerClassesAttribute) readInfo(reader *ClassReader) {
	numberOfClasses := reader.readUint16()

	a.classes = make([]*InnerClassInfo, numberOfClasses)

	for i := range a.classes {
		a.classes[i] = &InnerClassInfo{
			innerClassInfoIndex:   reader.readUint16(),
			outerClassInfoIndex:   reader.readUint16(),
			innerNameIndex:        reader.readUint16(),
			innerClassAccessFlags: reader.readUint16(),
		}
	}
}
