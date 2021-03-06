// Copyright (c) 2015 RxnWeaver
//
// Part of the RxnWeaver suite of projects.  See README.md and LICENSE
// for more details.

package tokenizer

import (
	"fmt"
	"strconv"
	"strings"
)

// Annotation represents a curated annotation of a logical word in a
// text.
//
// Each annotated word belongs to exactly one input document, and
// exactly one identified section within that (title, abstract, etc.).
// The annotation also holds information about a particular property
// of the word.  Annotations are used for training the tools.
type Annotation struct {
	DocumentID string
	Section    string
	Begin      int
	End        int
	Entity     string
	Property   string
}

// NewAnnotation creates and initialises a new annotation for the
// given input word.
//
// It expects its input to be in six columns that are tab-separated.
// The order of the fields is:
//   - document identifier,
//   - section,
//   - beginning index of the word in the input text,
//   - corresponding ending index,
//   - word itself and
//   - entity type.
func NewAnnotation(in string) (*Annotation, error) {
	fields := strings.Split(in, "\t")
	if len(fields) != 6 {
		return nil, fmt.Errorf("Input does not have 6 columns : %s\n", in)
	}

	a := &Annotation{}
	a.DocumentID = fields[0]
	a.Section = fields[1]

	if n, err := strconv.Atoi(fields[2]); err == nil {
		a.Begin = n
	} else {
		return nil, err
	}
	if n, err := strconv.Atoi(fields[3]); err == nil {
		a.End = n
	} else {
		return nil, err
	}
	a.Entity = fields[4]

	a.Property = fields[5]

	return a, nil
}
