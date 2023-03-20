package twx

import (
	"io"
	"text/tabwriter"
)

func New(w io.Writer) *tabwriter.Writer {
	return tabwriter.NewWriter(w, 0, 0, 3, ' ', 0)
}
