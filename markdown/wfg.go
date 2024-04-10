package markdown

import "regexp"

// regexp for WFG's Phabricator links to diffs, files or pastes
var phabricatorLinkRegexp = regexp.MustCompile(`[Pp]hab:([DFP][[:digit:]]+)`)

// regexp for WFG's Phabricator links to revisions
var phabricatorRevRegexp = regexp.MustCompile(`[Pp]hab:rP([[:digit:]]+)`)

func (converter *DefaultConverter) convertPhabLinks(in string) string {
	out := phabricatorLinkRegexp.ReplaceAllString(in, "[$0](https://code.wildfiregames.com/$1)")

	out = phabricatorRevRegexp.ReplaceAllString(out, "r$1")

	return out
}
