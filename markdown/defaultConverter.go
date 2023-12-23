// Copyright 2020 Steve Jefferson. All rights reserved.
// Use of this source code is governed by a GPL-style
// license that can be found in the LICENSE file.

package markdown

import (
	"github.com/stevejefferson/trac2gitea/accessor/gitea"
	"github.com/stevejefferson/trac2gitea/accessor/trac"
)

// CreateDefaultConverter creates a default implementation of the markdown converter
func CreateDefaultConverter(tracAccessor trac.Accessor, giteaAccessor gitea.Accessor) *DefaultConverter {
	converter := DefaultConverter{tracAccessor: tracAccessor, giteaAccessor: giteaAccessor}
	return &converter
}

// DefaultConverter is the default implementation of the Trac markdown to Gitea markdown converter.
// This is used in two circumstances:
// 1. for ticket comments - in which case ticketID != NullID and wikiAccessor == nil
// 2. for wiki imports - in which case ticketID == NullID and wikiAccessor != nil
type DefaultConverter struct {
	tracAccessor  trac.Accessor
	giteaAccessor gitea.Accessor
}

func (converter *DefaultConverter) convertNonCodeBlockText(ticketID int64, wikiPage string, in string) string {
	out := in

	// remove TOC special tags
	out = converter.removeTOC(out)

	// do simple one-line constructs first
	out = converter.convertLinks(ticketID, wikiPage, out)
	out = converter.convertAnchors(out)
	out = converter.convertEscapes(out)
	out = converter.convertLists(out)
	out = converter.convertDefinitionLists(out)
	out = converter.convertHeadings(out)

	// font styles require links to be disguised
	out = converter.disguiseLinks(out)
	out = converter.convertFontStyles(out)
	out = converter.undisguiseLinks(out)

	// now do potentially more complex constructs
	out = converter.convertBlockQuotes(out)
	out = converter.convertTables(out)

	// do paragraphs last because this results in the insertion of newlines
	// - this can upset other regexps which assume everything on a single line
	out = converter.convertParagraphs(out)

	// custom WFG changes
	out = converter.convertPhabLinks(out)

	return out
}

func (converter *DefaultConverter) convert(ticketID int64, wikiPage string, in string) string {
	out := in

	// ensure we have Unix EOLs
	out = converter.convertEOL(out)

	// convert Trac "code blocks"
	// - some of them are converted to HTML tags
	// - others are actual code blocks, inside which the Trac syntax must not be converted - mark
	// those with custom brackets
	out = converter.convertCodeBlocks(out)

	// perform conversions on text not in a code block using the ticket-specific link conversion
	out = converter.convertNonCodeBlocks(out, func(in string) string {
		return converter.convertNonCodeBlockText(ticketID, wikiPage, in)
	})

	// finally, replace custom brackets with markdown code block boundaries
	out = converter.undisguiseCodeBlocks(out)

	return out
}

// TicketConvert converts a comment/description string associated with a Trac ticket to Gitea markdown
func (converter *DefaultConverter) TicketConvert(ticketID int64, in string) string {
	return converter.convert(ticketID, "", in)
}

// WikiConvert converts a comment/description string associated with a Trac wiki page to Gitea markdown
func (converter *DefaultConverter) WikiConvert(wikiPage string, in string) string {
	return converter.convert(trac.NullID, wikiPage, in)
}
