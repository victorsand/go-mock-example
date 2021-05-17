package activities

type Walk struct {
	Distance int
}

type BikeRide struct {
	Distance int
}

type ActivityServiceInterface interface {
	Walks() ([]Walk, error)
	BikeRides() ([]BikeRide, error)
}

type ActivityService struct {
}

func (s *ActivityService) Walks() ([]Walk, error) {
	// Go off to some external API to fetch data
	panic("Not implemented")
}

func (s *ActivityService) BikeRides() ([]BikeRide, error) {
	// Go off to some external API to fetch data
	panic("Not implemented")
}
