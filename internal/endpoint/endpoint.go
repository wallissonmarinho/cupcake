package endpoint

import (
	"github.com/go-kit/log"
	"github.com/wallissonmarinho/cupcake/internal/service"
)

type Endpoints struct {
}

func MakeEndpoints(s service.ServiceFactory, logger log.Logger) Endpoints {
	return Endpoints{}
}
