package Services

import "errors"

type IUserService interface {
	GetName(userid int) string
	DeleteName(userid int) error
}

type UserService struct{}

func (this UserService) GetName(userid int) string {

	if userid == 101 {
		return "shenyi"
	}
	return "guest"
}

func (this UserService) DeleteName(userid int) error {

	if userid == 101 {
		return errors.New("wu quanxian")
	}
	return nil
}
