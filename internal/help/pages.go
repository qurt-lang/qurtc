package help

type DocPage string

const (
	docsBaseLink = "https://qurt.tech/docs/"

	QurtTour      DocPage = docsBaseLink + "tour"
	SyntaxPage    DocPage = docsBaseLink + "syntax"
	VarsPage      DocPage = docsBaseLink + "variables"
	StructsPage   DocPage = docsBaseLink + "structs"
	FunctionsPage DocPage = docsBaseLink + "functions"
)
