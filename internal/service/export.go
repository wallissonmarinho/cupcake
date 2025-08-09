package service

import (
	"log"

	"github.com/jmoiron/sqlx"
	auth "github.com/wallissonmarinho/cupcake/internal/service/auth"
	cart "github.com/wallissonmarinho/cupcake/internal/service/cart"
	dashboard "github.com/wallissonmarinho/cupcake/internal/service/dashboard"
	item "github.com/wallissonmarinho/cupcake/internal/service/item"
	order "github.com/wallissonmarinho/cupcake/internal/service/order"
	product "github.com/wallissonmarinho/cupcake/internal/service/product"
	profile "github.com/wallissonmarinho/cupcake/internal/service/profile"
)

type ServiceFactory interface {
	AuthService() auth.AuthService
	CartService() cart.CartService
	DashboardService() dashboard.DashboardService
	ItemService() item.ItemService
	OrderService() order.OrderService
	ProductService() product.ProductService
	ProfileService() profile.ProfileService
}

type serviceFactory struct {
	authService      auth.AuthService
	cartService      cart.CartService
	dashboardService dashboard.DashboardService
	itemService      item.ItemService
	orderService     order.OrderService
	productService   product.ProductService
	profileService   profile.ProfileService
}

func NewServiceFactory(db *sqlx.DB, logger *log.Logger) ServiceFactory {
	return &serviceFactory{
		authService:      auth.NewAuthService(logger, db),
		cartService:      cart.NewCartService(logger, db),
		dashboardService: dashboard.NewDashboardService(logger, db),
		itemService:      item.NewItemService(logger, db),
		orderService:     order.NewOrderService(logger, db),
		productService:   product.NewProductService(logger, db),
		profileService:   profile.NewProfileService(logger, db),
	}
}

func (s *serviceFactory) AuthService() auth.AuthService {
	return s.authService
}

func (s *serviceFactory) CartService() cart.CartService {
	return s.cartService
}

func (s *serviceFactory) DashboardService() dashboard.DashboardService {
	return s.dashboardService
}

func (s *serviceFactory) ItemService() item.ItemService {
	return s.itemService
}

func (s *serviceFactory) OrderService() order.OrderService {
	return s.orderService
}

func (s *serviceFactory) ProductService() product.ProductService {
	return s.productService
}

func (s *serviceFactory) ProfileService() profile.ProfileService {
	return s.profileService
}
