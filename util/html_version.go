package util

import "strings"

var DocType = make(map[string]string)

func init() {
	DocType["HTML 4.01 Strict"] = `"-//W3C//DTD HTML 4.01//EN"`
	DocType["HTML 4.01 Transitional"] = `"-//W3C//DTD HTML 4.01 Transitional//EN"`
	DocType["HTML 4.01 Frameset"] = `"-//W3C//DTD HTML 4.01 Frameset//EN"`
	DocType["XHTML 1.0 Strict"] = `"-//W3C//DTD XHTML 1.0 Strict//EN"`
	DocType["XHTML 1.0 Transitional"] = `"-//W3C//DTD XHTML 1.0 Transitional//EN"`
	DocType["XHTML 1.0 Frameset"] = `"-//W3C//DTD XHTML 1.0 Frameset//EN"`
	DocType["XHTML 1.1"] = `"-//W3C//DTD XHTML 1.1//EN"`
	DocType["HTML 5"] = `<!DOCTYPE html>`
}

/*
GetDocVersion This function returns the version html page based on doctype
*/
func GetDocVersion(input string) (docVersion string) {

	docVersion = "UNKNOWN"

	for doctype, version := range DocType {

		if strings.Contains(input, version) {
			docVersion = doctype
			return
		}
	}

	return
}
