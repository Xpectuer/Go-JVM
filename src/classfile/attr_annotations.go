package classfile


type DeprecatedAttribute struct {AnnotationAttribute}
type SyntheticAttribute struct {AnnotationAttribute}

type AnnotationAttribute struct{}

func (aa *AnnotationAttribute) readInfo(reader *ClassReader) {
	// pass
}