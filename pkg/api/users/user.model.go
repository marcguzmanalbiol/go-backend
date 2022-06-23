package users

type todo struct {
	id      int
	userId  int
	name    string
	checked bool
}

func (t *todo) changeTodoState() {
	t.checked = !t.checked
}

func (t *todo) changeTodoName(newName string) {
	if newName != "" {
		t.name = newName
	}
}

func (t *todo) createTodo(id int, userId int, name string, checked bool) {
	// TODO: Handle errors when creating a to do.
}
