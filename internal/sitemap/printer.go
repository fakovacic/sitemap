package sitemap

import "context"

//go:generate moq -out ./mocks/printer.go -pkg mocks  . Printer
type Printer interface {
	Print(context.Context, []Page, string) error
}
