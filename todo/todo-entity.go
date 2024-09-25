package todo

type TodoEntity struct {
	Name    string
	Checked bool
	State   string `enum:"created,inReview,done"`
}
