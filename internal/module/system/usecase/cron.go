package usecase

type CronUsecase struct {
}

func NewCronUsecase() *CronUsecase {
	return &CronUsecase{}
}

func (u *CronUsecase) InitIfNeeded() error {
	return nil
}
