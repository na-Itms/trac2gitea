package markdown

import "regexp"

// regexp for WFG's Phabricator links
var phabricatorLinkRegexp = regexp.MustCompile(`Phab:(D[[:digit:]]+)`)

func (converter *DefaultConverter) convertPhabLinks(in string) string {
	return phabricatorLinkRegexp.ReplaceAllString(in, "[$0](https://code.wildfiregames.com/$1)")
}
