package strings

import (
	"strings"
	"sync"

	"github.com/leonelquinteros/gotext"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Setup a pool of casers to be used for capitalisation
var caserPool = &sync.Pool{
	New: func() any {
		c := cases.Title(language.English)
		return &c
	},
}

// Capitalize capitalizes the first letter of a string.
func Capitalize(str string) string {
	c := caserPool.Get().(*cases.Caser)
	defer caserPool.Put(c)
	return c.String(str)
}

var (
	Archived       = gotext.Get("archived")
	ArchivedUpper  = strings.ToUpper(Archived)
	Key            = gotext.Get("key")
	KeyUpper       = strings.ToUpper(Key)
	Name           = gotext.Get("name")
	NameUpper      = strings.ToUpper(Name)
	No             = gotext.Get("no")
	NoCapitalized  = Capitalize(No)
	NoUpper        = strings.ToUpper(No)
	Path           = gotext.Get("path")
	PathUpper      = strings.ToUpper(Path)
	Yes            = gotext.Get("yes")
	YesCapitalized = Capitalize(Yes)
	YesUpper       = strings.ToUpper(Yes)
)
