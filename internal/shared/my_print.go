package shared

import (
	"github.com/gookit/color"
)

func PrintLnCyan(s interface{}) {
	color.Cyan.Printf("%v\n", s)
}
