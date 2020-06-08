//+build wireinject

package main

import (
	"virtual-travel/domain/repository"
	"virtual-travel/infrastructure"
	"virtual-travel/infrastructure/database"
	"virtual-travel/interfaces/controllers"
	"virtual-travel/interfaces/gateway"
	"virtual-travel/usecases/igateway"
	"virtual-travel/usecases/interactor"
	"virtual-travel/usecases/interactor/usecase"

	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"googlemaps.github.io/maps"
)

var superSet = wire.NewSet(
	database.NewFavoriteRepository,
	wire.Bind(new(repository.IFavoriteRepository), new(*database.FavoriteRepository)),
	database.NewUserRepository,
	wire.Bind(new(repository.IUserRepository), new(*database.UserRepository)),

	gateway.NewGoogleMapGateway,
	wire.Bind(new(igateway.IGoogleMapGateway), new(*gateway.GoogleMapGateway)),

	interactor.NewFavoriteInteractor,
	wire.Bind(new(usecase.IFavoriteUseCase), new(*interactor.FavoriteInteractor)),

	controllers.NewLinebotController,

	infrastructure.NewRouter,
)

// Initialize DI
func Initialize(e *echo.Echo, db *gorm.DB, gmc *maps.Client) *infrastructure.Router {
	wire.Build(superSet)
	return &infrastructure.Router{}
}
