/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2018 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * AUTHOR: Bob van Luijt (bob@kub.design)
 * See www.creativesoftwarefdn.org for details
 * Contact: @CreativeSofwFdn / bob@kub.design
 */
package local

import (
	"fmt"

	"github.com/creativesoftwarefdn/weaviate/database/schema"
	"github.com/creativesoftwarefdn/weaviate/graphqlapi/descriptions"
	local_aggregate "github.com/creativesoftwarefdn/weaviate/graphqlapi/local/aggregate"
	local_get "github.com/creativesoftwarefdn/weaviate/graphqlapi/local/get"
	local_get_meta "github.com/creativesoftwarefdn/weaviate/graphqlapi/local/get_meta"
	"github.com/graphql-go/graphql"
)

// Build the local queries from the database schema.
func Build(dbSchema *schema.Schema) (*graphql.Field, error) {
	getField, err := local_get.Build(dbSchema)
	if err != nil {
		return nil, err
	}
	getMetaField, err := local_get_meta.Build(dbSchema)
	if err != nil {
		return nil, err
	}
	getAggregateField, err := local_aggregate.Build(dbSchema)
	if err != nil {
		return nil, err
	}

	localFields := graphql.Fields{
		"Get":       getField,
		"GetMeta":   getMetaField,
		"Aggregate": getAggregateField,
	}

	localObject := graphql.NewObject(graphql.ObjectConfig{
		Name:        "WeaviateLocalObj",
		Fields:      localFields,
		Description: descriptions.LocalObjDesc,
	})

	localField := graphql.Field{
		Type:        localObject,
		Description: descriptions.WeaviateLocalDesc,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			fmt.Printf("- localGetAndMetaObjectResolver (pass on source; the resolver)\n")
			// This step does nothing; all ways allow the resolver to continue
			return p.Source, nil
		},
	}

	return &localField, nil
}