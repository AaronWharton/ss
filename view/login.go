package view

type LoginViewModel struct {
	BaseViewModel
	Error []string
}

type LVM struct{}

func (LoginViewModel) GetVM() LoginViewModel {
	v := LoginViewModel{}
	v.SetTitle("Login")
	return v
}

func (v *LoginViewModel) Errors(errs ...string)  {
	v.Error = append(v.Error, errs...)
}