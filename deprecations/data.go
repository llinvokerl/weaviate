//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2020 SeMI Holding B.V. (registered @ Dutch Chamber of Commerce no 75221632). All rights reserved.
//  LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
//  LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
//  CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

// Code generated by go generate; DO NOT EDIT.
// This file was generated by go generate ./deprecations at 2020-06-16
package deprecations

import (
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/semi-technologies/weaviate/entities/models"
)

func timeMust(t time.Time, err error) strfmt.DateTime {
	if err != nil {
		panic(err)
	}

	return strfmt.DateTime(t)
}

func timeMustPtr(t time.Time, err error) *strfmt.DateTime {
	if err != nil {
		panic(err)
	}

	parsed := strfmt.DateTime(t)
	return &parsed
}

func ptString(in string) *string {
	return &in
}

var ByID = map[string]models.Deprecation{
	"rest-meta-prop": models.Deprecation{
		ID:           "rest-meta-prop",
		Status:       "deprecated",
		APIType:      "REST",
		Mitigation:   "Use ?include=<propName>, e.g. ?include=_classification for classification meta or ?include=_vector to show the vector position or ?include=_classification,_vector for both. When consuming the response use the underscore fields such as _vector, as the meta object in the reponse, such as meta.vector will be removed.",
		Msg:          "use of deprecated property ?meta=true/false",
		SinceVersion: "0.22.8",
		SinceTime:    timeMust(time.Parse(time.RFC3339, "2020-06-15T16:18:06.000Z")),
	},
}
