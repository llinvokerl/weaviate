//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2024 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package parameters

import (
	"github.com/tailor-inc/graphql/language/ast"
	"github.com/weaviate/weaviate/usecases/modulecomponents/gqlparser"
)

type Params struct {
	ApiEndpoint string
	Model       string
	Temperature *float64
	Images      []string
}

func extract(field *ast.ObjectField) interface{} {
	out := Params{}
	fields, ok := field.Value.GetValue().([]*ast.ObjectField)
	if ok {
		for _, f := range fields {
			switch f.Name.Value {
			case "apiEndpoint":
				out.ApiEndpoint = gqlparser.GetValueAsStringOrEmpty(f)
			case "model":
				out.Model = gqlparser.GetValueAsStringOrEmpty(f)
			case "temperature":
				out.Temperature = gqlparser.GetValueAsFloat64(f)
			case "images":
				out.Images = gqlparser.GetValueAsStringArray(f)
			default:
				// do nothing
			}
		}
	}
	return out
}
