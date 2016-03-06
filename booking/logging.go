package booking

import (
	"time"

	"github.com/go-kit/kit/log"

	"github.com/marcusolsson/goddd/cargo"
	"github.com/marcusolsson/goddd/location"
)

type loggingService struct {
	logger log.Logger
	Service
}

// NewLoggingService ...
func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func (s *loggingService) BookNewCargo(origin location.UNLocode, destination location.UNLocode, arrivalDeadline time.Time) (id cargo.TrackingID, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "book",
			"origin", origin,
			"destination", destination,
			"arrival_deadline", arrivalDeadline,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.BookNewCargo(origin, destination, arrivalDeadline)
}

func (s *loggingService) LoadCargo(id cargo.TrackingID) (c *Cargo, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "load",
			"tracking_id", id,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.LoadCargo(id)
}

func (s *loggingService) RequestPossibleRoutesForCargo(id cargo.TrackingID) []cargo.Itinerary {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "request_routes",
			"tracking_id", id,
			"took", time.Since(begin),
		)
	}(time.Now())
	return s.Service.RequestPossibleRoutesForCargo(id)
}

func (s *loggingService) AssignCargoToRoute(id cargo.TrackingID, itinerary cargo.Itinerary) (err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "assign_to_route",
			"tracking_id", id,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.AssignCargoToRoute(id, itinerary)
}

func (s *loggingService) ChangeDestination(id cargo.TrackingID, l location.UNLocode) (err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "change_destination",
			"tracking_id", id,
			"destination", l,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.ChangeDestination(id, l)
}

func (s *loggingService) Cargos() []Cargo {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "list_cargos",
			"took", time.Since(begin),
		)
	}(time.Now())
	return s.Service.Cargos()
}

func (s *loggingService) Locations() []Location {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "list_locations",
			"took", time.Since(begin),
		)
	}(time.Now())
	return s.Service.Locations()
}