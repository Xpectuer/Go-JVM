package classfile

import "fmt"

type ClassFile struct {
	// magic number 很多文件格式以魔数开头，来标定特定的文件类型
	// 如class文件就是0xCAFEBABE
	// magic 	uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

// parse the class info from byte stream
func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

func (cf *ClassFile) read(reader *ClassReader) {
	cf.readAndCheckMagic(reader)
	cf.readAndCheckVersion(reader)
	cf.constantPool = readConstantPool(reader)
	cf.accessFlags = reader.readUint16()
	cf.thisClass = reader.readUint16()
	// mono extension supported
	cf.superClass = reader.readUint16()
	// multi impl supported
	cf.interfaces = reader.readUint16s()
	cf.fields = readMembers(reader, cf.constantPool)
	cf.methods = readMembers(reader, cf.constantPool)
	cf.attributes = readAttributes(reader, cf.constantPool)

}

// read 按照规范顺序读取各个class文件的各个字段
func (cf *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

// version id: class文件的版本号在magic number之后
// 版本号包括：主版本号 和 次版本号 M.m
func (cf *ClassFile) readAndCheckVersion(reader *ClassReader) {
	cf.minorVersion = reader.readUint16()
	cf.majorVersion = reader.readUint16()
	switch cf.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if cf.minorVersion == 0 {
			return
		}

	}
	panic("java.lang.UnsupportedClassVersionError!")
}

// --- getters ---
func (cf *ClassFile) MinorVersion() uint16 {
	return cf.minorVersion
}

func (cf *ClassFile) MajorVersion() uint16 {
	return cf.majorVersion
}

func (cf *ClassFile) ConstantPool() ConstantPool {
	return cf.constantPool
}

func (cf *ClassFile) AccessFlags() uint16 {
	return 0
}

func (cf *ClassFile) Fields() []*MemberInfo {
	return cf.fields
}

func (cf *ClassFile) Methods() []*MemberInfo {
	return cf.methods
}

func (cf *ClassFile) ClassName() string {
	return cf.constantPool.getClassName(cf.thisClass)
}

func (cf *ClassFile) SuperClassName() string {

	if cf.superClass > 0 {
		return cf.constantPool.getClassName(cf.superClass)
	}
	// 没有父类
	return ""
}

func (cf *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(cf.interfaces))
	for i, cpIndex := range cf.interfaces {
		interfaceNames[i] = cf.constantPool.getClassName(cpIndex)

	}
	return interfaceNames
}

// --- end of getters ---
