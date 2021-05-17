# A simple pattern for mocks in Go unit testing

## Purpose
Illustrate a simple approach for mocking Go interfaces when testing internal dependencies.

## When is it useful?
When you have several internal and interconnected services in your application and want to test one at a time, providing mock data or functionality for the dependencies.

## Prerequisites
The services need to talk to each other using `interfaces` in order to make mock implementations possible. If the services are interconnected using concrete `structs`, you need to refactor a bit first.

## Example app

Two services:
- `ActivitityService` which fetches `Walk` and `BikeRide` data from some imagined external API. This is the service we want to provide mock implementations for.
- `SummaryService` which internally uses an `ActivityService` instance, uses it to get some data and then perform some simple calculations on it. This is the service we want to test.

## Run the test
`cd summaries`
`go test`


## Setup

`ActivityService` implements `ActivityServiceInterface`. So does `MockActivityService`. The idea is to, for each interface method, provide the ability for a unit test to provide its own implementation using `func`s like so (in [mock_activity_service.go](mock/mock_activity_service.go)):

```go
package mock

import  "go-mock-example/activities"

type  MockActivityService  struct {
    WalksFunc func() ([]activities.Walk, error)
    BikeRidesFunc func() ([]activities.BikeRide, error)
}

func (s *MockActivityService) Walks() ([]activities.Walk, error) {
    return s.WalksFunc()
}

func (s *MockActivityService) BikeRides() ([]activities.BikeRide, error) {
    return s.BikeRidesFunc()
}
```
In the test ( [summary_service_test.go](summaries/summary_service_test.go)), we provide the data we need in the setup stage:

```go
mockActivityService := mock.MockActivityService{
	WalksFunc: func() ([]activities.Walk, error) {
		return []activities.Walk{
			{Distance: 10},
			{Distance: 25},
		}, nil
	},
	BikeRidesFunc: func() ([]activities.BikeRide, error) {
		return []activities.BikeRide{
			{Distance: 20},
		}, nil
	},
}
```
We then instantiated the service want to test and provide the mock as the dependency:
```go
summariesService := SummariesService{
	activityService: &mockActivityService,
}
```

## Benefits
- The mock class does not contain any business logic on its own, and doesn't have to change unless the interface changes.
-  Each unit test is 100% responsible for providing exactly the data needed. Default implementations are of course possible and is a matter of choice. As is, the test crash if one tries to call a method without a provided implementation.
