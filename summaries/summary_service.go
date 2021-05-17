package summaries

import "go-mock-example/activities"

type SummariesServiceInterface interface {
	TotalDistance() (int, error)
}

type SummariesService struct {
	activityService activities.ActivityServiceInterface
}

func (s *SummariesService) TotalDistance() (int, error) {
	walks, err := s.activityService.Walks()
	if err != nil {
		return 0, err
	}
	bikeRides, err := s.activityService.BikeRides()
	if err != nil {
		return 0, err
	}

	totalDistance := 0
	for _, walk := range walks {
		totalDistance += walk.Distance
	}
	for _, bikeRide := range bikeRides {
		totalDistance += bikeRide.Distance
	}
	return totalDistance, nil
}
