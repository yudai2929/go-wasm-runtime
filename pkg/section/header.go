package section

type Header struct {
	Code SectionCode
	Size uint32
}

func NewHeader(code SectionCode, size uint32) *Header {
	return &Header{
		Code: code,
		Size: size,
	}
}
