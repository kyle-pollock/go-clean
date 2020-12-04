package presenters

import "github.com/kyle-pollock/go-clean/pkg/usecases/user"

type ViewModel struct {
	Users []*User `json:"users"`
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UsersPresenter struct {
	viewModel ViewModel
}

func (p *UsersPresenter) ViewModel() ViewModel {
	return p.viewModel
}

func (p *UsersPresenter) Present(model []*user.UserResponseModel) {
	for _, user := range model {
		p.viewModel.Users = append(p.viewModel.Users, &User{
			ID:   user.ID,
			Name: user.Name,
		})
	}
}
