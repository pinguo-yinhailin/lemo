package cmd

func (c *Cmd) registerCommands() {
	cmds := []Commander{
		NewArticle(c),
		NewUser(c),
	}

	for _, cmd := range cmds {
		c.app.Commands = append(c.app.Commands, cmd.Command())
	}
}
