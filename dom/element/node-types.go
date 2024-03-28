package element

// Node https://developer.mozilla.org/en-US/docs/Web/API/Node
// NodeType https://developer.mozilla.org/en-US/docs/Web/API/Node/nodeType

type NodeType int

const (
	NODE_ELEMENT                NodeType = 1
	NODE_ATTRIBUTE              NodeType = 2
	NODE_TEXT                   NodeType = 3
	NODE_CDATA_SECTION          NodeType = 4
	NODE_PROCESSING_INSTRUCTION NodeType = 7
	NODE_COMMENT                NodeType = 8
	NODE_DOCUMENT               NodeType = 9
	NODE_DOCUMENT_TYPE          NodeType = 10
	NODE_DOCUMENT_FRAGMENT      NodeType = 11
)
