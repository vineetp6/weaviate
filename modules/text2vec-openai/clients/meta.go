//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

package clients

func (v *vectorizer) MetaInfo() (map[string]interface{}, error) {
	return map[string]interface{}{
		"documentation_href": "https://beta.openai.com/docs/guides/embeddings/what-are-embeddings",
	}, nil
}
