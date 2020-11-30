package json

import "github.com/kyle-pollock/go-clean/pkg/usecases/user"

type UsersViewModel struct {
	Users []*UserViewModel `json:"users"`
}

type UserViewModel struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Presenter struct {
	viewModel *UsersViewModel
}

func (p *Presenter) ViewModel() *UsersViewModel {
	return p.viewModel
}

func (p *Presenter) Present(model []*user.UserResponseModel) {
	p.viewModel = new(UsersViewModel)
	for _, value := range model {
		p.viewModel.Users = append(p.viewModel.Users, &UserViewModel{
			ID:   value.ID,
			Name: value.Name,
		})
	}
}
