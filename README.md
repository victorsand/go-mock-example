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

`ActivityService` implements `ActivityServiceInterface` ([activity_service.go](activities/activity_service.go)) with its two data-returning methods.

```go
type ActivityServiceInterface interface {
    Walks() ([]Walk, error)
	  BikeRides() ([]BikeRide, error)
}
```


`MockActivityService` implements the same interface.

The idea is to, for each interface method, provide the ability for a unit test to provide its own implementation using corresponding `func`s like so (in [mock_activity_service.go](mock/mock_activity_service.go)):

```go
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
We the instantiated the service want to test and provide the mock as the dependency:
```go
summariesService := SummariesService{
	activityService: &mockActivityService,
}
Then we make the test assertions we need.
```

## Benefits
- No additional libraries or tools needed.
- The mock classes do not contain any logic on their own, and don't have to change unless the interfaces change.
-  Each unit test is 100% responsible for providing exactly the data needed (and only that data). Default implementations are of course possible and is a matter of choice. As is, the tests crash if one tries to call a methods without a provided implementations.
