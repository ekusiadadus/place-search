package interactor

import (
	"virtual-travel/domain/repository"
	"virtual-travel/usecase/dto/favoritedto"
)

// FavoriteInteractor お気に入りインタラクタ
type FavoriteInteractor struct {
	repository repository.IFavoriteRepository
}

// NewFavoriteInteractor コンストラクタ
func NewFavoriteInteractor(repository repository.IFavoriteRepository) *FavoriteInteractor {
	return &FavoriteInteractor{repository: repository}
}

// Add お気に入りを追加する
func (interactor *FavoriteInteractor) Add(in favoritedto.FavoriteAddInput) favoritedto.FavoriteAddOutput {
	UserID := in.UserID
	PlaceID := in.PlaceID
	isAlreadyAdded := interactor.repository.Save(UserID, PlaceID)

	return favoritedto.FavoriteAddOutput{IsAlreadyAdded: isAlreadyAdded}
}

// Get お気に入り全件を取得する
func (interactor *FavoriteInteractor) Get(in favoritedto.FavoriteGetInput) favoritedto.FavoriteGetOutput {
	LineUserID := in.LineUserID
	PlaceIDs := interactor.repository.FindAll(LineUserID)

	return favoritedto.FavoriteGetOutput{PlaceIDs: PlaceIDs}
}
