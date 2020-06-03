package usecase

// IUserCreateUseCase ユーザ登録ユースケース
type IUserCreateUseCase interface {
	Handle(UserCreateInput)
}