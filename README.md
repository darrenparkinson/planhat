# Go Planhat - A simple Planhat Go Library

[![Status](https://img.shields.io/badge/status-wip-yellow)](https://github.com/darrenparkinson/planhat) ![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/darrenparkinson/planhat) ![GitHub](https://img.shields.io/github/license/darrenparkinson/planhat?color=brightgreen) [![GoDoc](https://pkg.go.dev/badge/darrenparkinson/planhat)](https://pkg.go.dev/github.com/darrenparkinson/planhat) [![Go Report Card](https://goreportcard.com/badge/github.com/darrenparkinson/planhat)](https://goreportcard.com/report/github.com/darrenparkinson/planhat)

This repository is intended as a simple to use library for the Go language to interact with the [Planhat API](https://docs.planhat.com/).

In order to use this library you will need a [planhat subscription](https://www.planhat.com/).  

## Installing

You can install the library in the usual way as follows:

```sh
$ go get github.com/darrenparkinson/planhat
```

## Usage

In your code, import the library:

```go
import "github.com/darrenparkinson/planhat"
```

You can then construct a new Planhat client, and use the various services on the client to access different parts of the API.  For example:

```go
ph, _ := planhat.NewClient(apikey, cluster, nil)
companies, err := ph.CompanyService.List(context.Background())
```

Some API methods have optional parameters that can be passed, for example:

```go
ph, _ := planhat.NewClient(apikey, cluster, nil)
opts := &planhat.CompanyListOptions{Limit: planhat.Int(10), Offset: planhat.Int(0)}
companies, err := ph.CompanyService.List(context.Background(), opts)
```

The services of a client divide the API into logical areas and correspond to the structure of the Planhat API documentation at https://docs.planhat.com/

NOTE: Using the context package, one can easily pass cancelation signals and deadlines to various services of the client for handling a request. In case there is no context available, then context.Background() can be used as a starting point.

## Authentication

Authentication is provided by an API Key as outlined [in the documentation](https://docs.planhat.com/#authentication).  You are able to provide the API Key as part of initialisation using `NewClient`.  

You can obtain an API key by navigating to "Service Accounts" and adding a new service account user that has the permissions you require, followed by clicking the "Generate New Token" button.

If you receive a `planhat: unauthorized request` message, you may be using an incorrect region.

## Region

Planhat supports various regions and cluster as outlined [in the documentation](https://docs.planhat.com/#base-url).  You must provide the cluster you are using to `NewClient` on initialisation.  Ask your planhat representative or [check the docs](https://docs.planhat.com/#introduction) to see which one you need to use.  If there is no specified cluster, you may pass an empty string.

Examples:

* "" becomes `api`, 
* `eu` becomes `api-eu`, 
* `eu2` becomes `api-eu2`, 
* `eu3` becomes `api-eu3`, 
* `us2` becomes `api-us2`

## Helper Functions

Most structs for resources use pointer values.  This allows distinguishing between unset fields and those set to a zero value.  Some helper functions have been provided to easily create these pointers for string, bool and int values as you saw above and here, for example:

```go
opts := &planhat.CompanyListOptions{
    Limit:  planhat.Int(10),
    Offset: planhat.Int(0),
	Sort:   planhat.String("name"),
}
```

This can cause challenges when receiving results since you may encounter a panic if you access a nil pointer, e.g:

```
company, _ := ph.CompanyService.GetCompany(ctx, "123123123abcabcabc")
log.Println(*company.ExternalID)
```

In this case, if the external id has no value, you would receive a `panic: runtime error: invalid memory address or nil pointer dereference` error.  Clearly this isn't a very nice user experience, so where appropriate, "getter" accessor functions are generated automatically for structs with pointer fields to enable you to safely retrieve values:

```
company, _ := ph.CompanyService.GetCompany(ctx, "123123123abcabcabc")
log.Println(company.GetExternalID())
```

## Pagination

Where pagination is provided, Planhat provides the Offset and Limit query parameters as part of the request parameters for a given endpoint.  These can be passed via options to the command:

```go
companies, err := c.CompanyService.List(ctx, &planhat.CompanyListOptions{Limit: planhat.Int(10), Offset: planhat.Int(0)})
```

By way of an example, you might use the following to work through multiple pages:

```go
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
```

As you can see, planhat doesn't provide a mechanism to check if there are more values, so we keep going until there are no results.

## Sorting

Where sorting is provided, Planhat provides the Sort query parameter as part of the request parameters for a given endpoint.  This can also be passed via the options to the command and is used in conjuntion with pagination:

```go
companies, err := c.CompanyService.List(ctx, &planhat.CompanyListOptions{Limit: planhat.Int(10), Offset: planhat.Int(0), Sort: planhat.String("name")})
```

Note that the sort string appears to be case sensitive and must currently use the Planhat object name.

## Errors

In the [documentation](https://docs.planhat.com/), Planhat identifies the following returned errors.  These are provided as constants so that you may check against them:

| Code | Error                | Constant           |
|------|----------------------|--------------------|
| 400  | Bad Request          | `ErrBadRequest`    |
| 401  | Unauthorized Request | `ErrUnauthorized`  |
| 403  | Forbidden            | `ErrForbidden`     |
| 500  | Internal Error       | `ErrInternalError` |

All other errors are returned as `ErrUnknown`

As an example:

```go
companies, err := ph.CompanyService.List(ctx)
if errors.Is(err, planhat.ErrUnauthorized) {
	log.Fatal("Sorry, you're not allowed to do that.")
}
```

# Services

The following outlines the planhat models and their implementation status:

| Model        | Service             | Implementation Status |
|--------------|---------------------|-----------------------|
| Asset        | AssetService        | Not Implemented       |
| Churn        | ChurnService        | Not Implemented       |
| Company      | CompanyService      | Complete              |
| Conversation | ConversationService | Not Implemented       |
| Custom Field | CustomFieldService  | Not Implemented       |
| Enduser      | EnduserService      | Not Implemented       |
| Invoice      | InvoiceService      | Not Implemented       |
| Issue        | IssueService        | Not Implemented       |
| License      | LicenseService      | Not Implemented       |
| Note         | NoteService         | Not Implemented       |
| NPS          | NPSService          | Not Implemented       |
| Opportunity  | OpportunityService  | Not Implemented       |
| Project      | ProjectService      | Not Implemented       |
| Sale         | SaleService         | Not Implemented       |
| Task         | TaskService         | Not Implemented       |
| Ticket       | TicketService       | Not Implemented       |
| User         | UserService         | Partial               |

In addition to the Planhat Models, there are some additional endpoints in the documentation as outlined below:

| Section         | Service             | Implementation Status |
|-----------------|---------------------|-----------------------|
| User Activities | UserActivityService | Not Implemented       |
| Metrics         | MetricsService      | Complete              |

# Contributing

Since all endpoints would ideally be covered, contributions are always welcome.  Adding new methods should be relatively straightforward.

# Versioning

In general this planhat library follows [semver](https://semver.org/) for tagging releases of the package.  As yet, it is still in development and has not had a tag added, but will in due course.  Since it is still in development, you may expect some changes.

# License

This library is distributed under the MIT license found in the [LICENSE](LICENSE) file.