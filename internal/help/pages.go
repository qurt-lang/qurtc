package help

type DocPage string

const (
	docsBaseLink = "https://qurt.tech/docs"

	QurtTour      DocPage = docsBaseLink + "/tour"
	SyntaxPage    DocPage = docsBaseLink + "/syntax"
	FunctionsPage DocPage = docsBaseLink + "/functions"
	StructsPage   DocPage = docsBaseLink + "/structs"
	VarsPage      DocPage = docsBaseLink + "/variables"
)
