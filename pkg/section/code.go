package section

type SectionCode int

const (
	SectionCodeCustom SectionCode = iota
	SectionCodeType
	SectionCodeImport
	SectionCodeFunction
	SectionCodeTable
	SectionCodeMemory
	SectionCodeGlobal
	SectionCodeExport
	SectionCodeStart
	SectionCodeElement
	SectionCodeCode
	SectionCodeData
	SectionCodeDataCount
)

func NewCode(code byte) SectionCode {
	return SectionCode(code)
}
