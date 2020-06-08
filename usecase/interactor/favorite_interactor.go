package interactor

import (
	"virtual-travel/domain/repository"
	"virtual-travel/interfaces/presenter"
	"virtual-travel/usecase/dto/favoritedto"
)

// FavoriteInteractor お気に入りインタラクタ
type FavoriteInteractor struct {
	userRepository     repository.IUserRepository
	favoriteRepository repository.IFavoriteRepository
}

// NewFavoriteInteractor コンストラクタ
func NewFavoriteInteractor(
	userRepository repository.IUserRepository,
	favoriteRepository repository.IFavoriteRepository) *FavoriteInteractor {

	return &FavoriteInteractor{userRepository: userRepository, favoriteRepository: favoriteRepository}
}

// Add お気に入りを追加する
func (interactor *FavoriteInteractor) Add(in favoritedto.FavoriteAddInput) {
	userID := interactor.userRepository.Save(in.LineUserID)

	isSuccess := true
	if userID == 0 {
		isSuccess = false
	}

	isAlreadyAdded := interactor.favoriteRepository.Save(userID, in.PlaceID)

	out := favoritedto.FavoriteAddOutput{
		Bot:            in.Bot,
		ReplyToken:     in.ReplyToken,
		IsSuccess:      isSuccess,
		IsAlreadyAdded: isAlreadyAdded,
	}
	presenter.Add(out)
}

// Get お気に入り全件を取得する
func (interactor *FavoriteInteractor) Get(in favoritedto.FavoriteGetInput) favoritedto.FavoriteGetOutput {
	LineUserID := in.LineUserID
	PlaceIDs := interactor.favoriteRepository.FindAll(LineUserID)

	return favoritedto.FavoriteGetOutput{PlaceIDs: PlaceIDs}
}
