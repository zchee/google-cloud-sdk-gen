// Copyright 2018 The google-cloud-sdk-gen Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/blang/semver"
)

const (
	releaseNoteURI = "https://cloud.google.com/sdk/docs/release-notes"
)

var (
	versionThreshold = semver.MustParse("0.9.83")
)

func main() {
	doc, err := goquery.NewDocument(releaseNoteURI)
	if err != nil {
		log.Fatal(err)
	}

	versions := []semver.Version{}
	versionm := make(map[string]string)
	doc.Find("h2").Each(func(i int, s *goquery.Selection) {
		ss := strings.Split(s.Text(), " ")
		version, err := semver.Parse(ss[0])
		if err != nil { // handle secondary patch number such as 0.9.13.1
			return
		}
		if version.LT(versionThreshold) { // handle version => 0.9.83
			return
		}

		// fmt.Println(ss[0])
		versions = append(versions, version)
		versionm[version.String()] = s.Nodes[0].Attr[0].Val
	})
	semver.Sort(versions)

	b := new(strings.Builder)
	for _, v := range versions {
		b.WriteString(v.String() + "\n")
	}

	fmt.Println(strings.TrimSpace(b.String()))
}
