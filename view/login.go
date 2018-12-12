package view

type LoginView struct {
	BaseView
	Errors []string			// errors information
}

type LVM struct{}

func (LVM) GetView() LoginView {
	v := LoginView{}
	v.SetTitle("Login")
	return v
}

// add login errors information
func (v *LoginView) AddError(err ...string) {
	v.Errors = append(v.Errors, err...)
}
