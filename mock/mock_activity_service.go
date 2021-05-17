package mock

import "go-mock-example/activities"

type MockActivityService struct {
	WalksFunc     func() ([]activities.Walk, error)
	BikeRidesFunc func() ([]activities.BikeRide, error)
}

func (s *MockActivityService) Walks() ([]activities.Walk, error) {
	return s.WalksFunc()
}

func (s *MockActivityService) BikeRides() ([]activities.BikeRide, error) {
	return s.BikeRidesFunc()
}
