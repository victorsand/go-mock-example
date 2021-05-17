package summaries

import (
	"go-mock-example/activities"
	"go-mock-example/mock"
	"testing"
)

func Test_SummaryService_SumsActivityDistances(t *testing.T) {

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

	summariesService := SummariesService{
		activityService: &mockActivityService,
	}

	totalDistance, err := summariesService.TotalDistance()

	if err != nil {
		t.Fatalf(err.Error())
	}

	if totalDistance != 55 {
		t.Fatalf("Incorrect total")
	}

}
