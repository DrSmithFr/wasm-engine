package node

type Types int

const (
	ELEMENT_NODE                Types = 1
	ATTRIBUTE_NODE              Types = 2
	TEXT_NODE                   Types = 3
	CDATA_SECTION_NODE          Types = 4
	PROCESSING_INSTRUCTION_NODE Types = 7
	COMMENT_NODE                Types = 8
	DOCUMENT_NODE               Types = 9
	DOCUMENT_TYPE_NODE          Types = 10
	DOCUMENT_FRAGMENT_NODE      Types = 11
)
