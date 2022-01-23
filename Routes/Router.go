package Routes

import "github.com/AbdulrahmanMasoud/starter/Application"

// ده انا عامله عشان اكسس من  ال Application  في كل الراوتس اللي هعملها عشان  طبعا مينفعش اكسسه ديركت عشان هو مش في نفس الباكيدج
// ف انا هعمل ستراكت جديد هنا واباصيله ال Application  وااكسس منهفي الراوتس

type Router struct {
	*Application.Application
}

func (router Router) Routes() {
	router.visitorsRoute()
}
