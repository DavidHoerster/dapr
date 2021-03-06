// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package http

import routing "github.com/qiangxue/fasthttp-routing"

// Endpoint is a collection of route information for an Dapr API
type Endpoint struct {
	Methods []string
	Route   string
	Version string
	Handler func(c *routing.Context) error
}
