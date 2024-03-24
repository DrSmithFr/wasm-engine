package list

type Type string

// DOM elements list
const (
	// Main root
	html Type = "html"

	// Document metadata
	base  = "base"
	head  = "head"
	link  = "link"
	meta  = "meta"
	style = "style"
	title = "title"

	// Sectioning root
	body = "body"

	// Content sectioning
	address = "address"
	article = "article"
	aside   = "aside"
	footer  = "footer"
	header  = "header"
	h1      = "h1"
	h2      = "h2"
	h3      = "h3"
	h4      = "h4"
	h5      = "h5"
	h6      = "h6"
	hgroup  = "hgroup"
	main    = "main"
	nav     = "nav"
	section = "section"
	search  = "search"

	// Text content
	blockquote = "blockquote"
	dd         = "dd"
	div        = "div"
	dl         = "dl"
	dt         = "dt"
	figcaption = "figcaption"
	figure     = "figure"
	hr         = "hr"
	li         = "li"
	menu       = "menu"
	ol         = "ol"
	p          = "p"
	pre        = "pre"
	ul         = "ul"

	// Inline text semantics
	a        = "a"
	abbr     = "abbr"
	b        = "b"
	bdi      = "bdi"
	bdo      = "bdo"
	br       = "br"
	cite     = "cite"
	code     = "code"
	data     = "data"
	dfn      = "dfn"
	em       = "em"
	i        = "i"
	kbd      = "kbd"
	mark     = "mark"
	q        = "q"
	rp       = "rp"
	rt       = "rt"
	ruby     = "ruby"
	s        = "s"
	samp     = "samp"
	small    = "small"
	span     = "span"
	strong   = "strong"
	sub      = "sub"
	time     = "time"
	u        = "u"
	variable = "var"
	wbr      = "wbr"

	// Image and multimedia
	area     = "area"
	audio    = "audio"
	img      = "img"
	area_map = "map"
	track    = "track"
	video    = "video"

	// Embedded content
	emded   = "embed"
	iframe  = "iframe"
	object  = "object"
	picture = "picture"
	portal  = "portal"
	source  = "source"

	// SVG and MathML
	svg  = "svg"
	math = "math"

	// Scripting
	canvas   = "canvas"
	noscript = "noscript"
	script   = "script"

	// Demarcating edits
	del = "del"
	ins = "ins"

	// Table content
	caption  = "caption"
	col      = "col"
	colgroup = "colgroup"
	table    = "table"
	tbody    = "tbody"
	td       = "td"
	tfoot    = "tfoot"
	th       = "th"
	thead    = "thead"
	tr       = "tr"

	// Forms
	button     = "button"
	datalist   = "datalist"
	fieldset   = "fieldset"
	form       = "form"
	input      = "input"
	label      = "label"
	legend     = "legend"
	meter      = "meter"
	optgroup   = "optgroup"
	option     = "option"
	output     = "output"
	progress   = "progress"
	select_box = "select"
	textarea   = "textarea"

	// Interactive elements
	details = "details"
	dialog  = "dialog"
	summery = "summery"

	// Web components
	slot     = "slot"
	template = "template"

	// Obsolete and deprecated elements
	acronym   = "acronym"
	big       = "big"
	center    = "center"
	content   = "content"
	dir       = "dir"
	font      = "font"
	frame     = "frame"
	frameset  = "frameset"
	image     = "image"
	marquee   = "marquee"
	menuitem  = "menuitem"
	nobr      = "nobr"
	noembed   = "noembed"
	noframes  = "noframes"
	param     = "param"
	plaintext = "plaintext"
	rb        = "rb"
	stc       = "rtc"
	shadow    = "shadow"
	strike    = "strike"
	tt        = "tt"
	xmp       = "xmp"
)
