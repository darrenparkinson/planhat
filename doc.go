// Copyright 2021 The go-planhat AUTHORS. All rights reserved.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

/*
Package planhat provides a client for using the Planhat API.

Usage:

	import "github.com/darrenparkinson/planhat"

Construct a new Planhat client, then use the various services on the client to
access different parts of the Planhat API.  For example:

	client := planhat.NewClient(apikey, cluster, nil)

	// List companies
	companies, err := ph.CompanyService.List(ctx)

Some API methods have optional parameters that can be passed.  For example:

	ph, _ := planhat.NewClient(apikey, cluster, nil)
	opts := &planhat.CompanyListOptions{Limit: planhat.Int(10), Offset: planhat.Int(0)}
	companies, err := ph.CompanyService.List(ctx, opts)

The services of a client divide the API into logical areas and correspond to the structure
of the Planhat API documentation at https://docs.planhat.com/

NOTE: Using the context package, one can easily pass cancelation signals and deadlines
to various services of the client for handling a request. In case there is no context
available, then context.Background() can be used as a starting point.

Authentication

Authentication is provided by an API Key as outlined in the documentation at https://docs.planhat.com/#authentication.
You are able to provide the API Key as part of initialisation using `NewClient`.

You can obtain an API key by navigating to "Service Accounts" and adding a new service account
user that has the permissions you require, followed by clicking the "Generate New Token" button.

If you receive a `planhat: unauthorized request` message, you may be using an incorrect region.

Region

Planhat supports various regions and cluster as outlined in the documentation at https://docs.planhat.com/#base-url).
You must provide the cluster you are using to `NewClient` on initialisation.
Ask your planhat representative or [check the docs](https://docs.planhat.com/#introduction) to see
which one you need to use.  If there is no specified cluster, you may pass an empty string.

Examples:

* "" becomes `api`,
* `eu` becomes `api-eu`,
* `eu2` becomes `api-eu2`,
* `eu3` becomes `api-eu3`,
* `us2` becomes `api-us2`

Pagination

Where pagination is provided, Planhat provides the Offset and Limit query parameters as part of the request
parameters for a given endpoint.  These can be passed via options to the command:

	companies, err := c.CompanyService.List(ctx, &planhat.CompanyListOptions{Limit: planhat.Int(10), Offset: planhat.Int(0)})

By way of an example, you might use the following to work through multiple pages:

	var allCompanies []*planhat.Company
	limit, offset := 10, 0
	for {
		companies, _ := ph.CompanyService.List(ctx, &planhat.CompanyListOptions{Limit: planhat.Int(limit), Offset: planhat.Int(offset)})
		log.Println("Retrieved", len(companies), "companies")
		offset += limit
		if len(companies) == 0 {
			break
		}
		allCompanies = append(allCompanies, companies...)
	}
	log.Println("Found total", len(allCompanies), "companies.")

As you can see, planhat doesn't provide a mechanism to check if there are more values,
so we keep going until there are no results.

Sorting

Where sorting is provided, Planhat provides the Sort query parameter as part of the request parameters
for a given endpoint.  This can also be passed via the options to the command and is used in conjuntion with pagination:

	companies, err := c.CompanyService.List(ctx, &planhat.CompanyListOptions{Limit: planhat.Int(10), Offset: planhat.Int(0), Sort: planhat.String("name")})

Note that the sort string appears to be case sensitive and must currently use the Planhat object name.

*/
package planhat
