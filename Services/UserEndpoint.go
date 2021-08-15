package Services

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"
)

type UserRequest struct {
	Uid    int `json:"uid"`
	Method string
}
type UserResponse struct {
	Reslut string `json :"reslust"`
}

func GetUserEndpoint(UserService IUserService) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		r := request.(UserRequest)
		result := "nothing do"
		//fmt.Println(r.Method)
		if r.Method == "GET" {
			result = UserService.GetName(r.Uid)

		} else if r.Method == "DELETE" {
			err := UserService.DeleteName(r.Uid)
			if err != nil {
				result = err.Error()
			} else {
				result = fmt.Sprintf("userid 为%d的用户删除成功", r.Uid)
			}
		}

		return UserResponse{
			Reslut: result,
		}, nil
	}
}
