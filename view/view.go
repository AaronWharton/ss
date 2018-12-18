package view

// all views' configuration
//

type BaseView struct {
	Title       string
	CurrentUser string
}

type BasePageView struct {
	PrevPage    int
	CurrentPage int
	NextPage    int
	TotalPage   int
	Limit       int // Limit is the max number of posts that displayed in one page
}

// every view can set its title
func (v *BaseView) SetTitle(title string) {
	v.Title = title
}

func (v *BaseView) SetCurrentUser(currentUser string) {
	v.CurrentUser = currentUser
}

func (v *BasePageView) SetPrevAndNextPage() {
	if v.CurrentPage > 1 {
		v.PrevPage = v.CurrentPage - 1
	}

	// TODO: What does it mean?
	if (v.TotalPage-1)/v.Limit >= v.CurrentPage {
		v.NextPage = v.CurrentPage + 1
	}
}

func (v *BasePageView) SetBasePageView(total, current, limit int) {
	v.TotalPage = total
	v.CurrentPage = current
	v.Limit = limit
}
