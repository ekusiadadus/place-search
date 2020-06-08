package usecase

import "virtual-travel/usecase/dto/favoritedto"

// IFavoriteUseCase お気に入りユースケース
type IFavoriteUseCase interface {
	Add(favoritedto.AddInput)
	Get(favoritedto.GetInput)
}
