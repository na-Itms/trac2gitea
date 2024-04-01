// Copyright 2020 Steve Jefferson. All rights reserved.
// Use of this source code is governed by a GPL-style
// license that can be found in the LICENSE file.

package importer

import (
	"github.com/stevejefferson/trac2gitea/accessor/gitea"
	"github.com/stevejefferson/trac2gitea/accessor/trac"
)

// importCommentIssueComment imports a Trac ticket comment into Gitea, returns id of created Gitea issue comment or NullID if cannot create comment
func (importer *Importer) importCommentIssueComment(issueID int64, change *trac.TicketChange, userMap, revisionMap map[string]string) (int64, error) {
	issueComment, err := importer.createIssueComment(issueID, change, userMap)
	if err != nil {
		return gitea.NullID, err
	}

	issueComment.CommentType = gitea.CommentIssueCommentType
	issueComment.Text = importer.markdownConverter.TicketConvert(change.TicketID, change.NewValue)
	issueComment.Text = MapRevisions(issueComment.Text, revisionMap)

	issueCommentID, err := importer.giteaAccessor.AddIssueComment(issueID, issueComment)
	if err != nil {
		return gitea.NullID, err
	}

	return issueCommentID, nil
}

// --- WFG
// importPatchChangeIssueComment imports a Trac ticket patch change into Gitea, as a formatted comment
// it returns id of created Gitea issue comment or NullID if cannot create comment
func (importer *Importer) importPatchChangeIssueComment(issueID int64, change *trac.TicketChange, userMap map[string]string) (int64, error) {
	issueComment, err := importer.createIssueComment(issueID, change, userMap)
	if err != nil {
		return gitea.NullID, err
	}

	issueComment.CommentType = gitea.CommentIssueCommentType

	if change.NewValue != "" {
		issueComment.Text = "**Patch set to:** " + importer.markdownConverter.TicketConvert(change.TicketID, change.NewValue)
	} else {
		issueComment.Text = "**Patch removed.**"
	}

	issueCommentID, err := importer.giteaAccessor.AddIssueComment(issueID, issueComment)
	if err != nil {
		return gitea.NullID, err
	}

	return issueCommentID, nil
}
