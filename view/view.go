package view

// all views' configuration
//

type BaseView struct {
	Title string
}

// every view can set its title
func (v *BaseView) SetTitle(title string) {
	v.Title = title
}
