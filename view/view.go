package view

// all views' configuration
//

type BaseView struct {
	Title       string
	CurrentUser string
}

// every view can set its title
func (v *BaseView) SetTitle(title string) {
	v.Title = title
}

func (v *BaseView) SetCurrentUser(currentUser string) {
	v.CurrentUser = currentUser
}
