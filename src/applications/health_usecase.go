package applications

import (
	"restapi/src/shared"
	"time"
)

// HealthUsecase ...
type HealthUsecase struct {
	StartupTime time.Time
}

// Health ...
func (u *HealthUsecase) Health() *shared.HTTPResponse {
	var response = make(map[string]interface{})

	response["startup"] = u.StartupTime
	response["now"] = time.Now()

	return shared.HTTPSuccess(response)
}
