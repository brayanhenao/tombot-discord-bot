package framework

type (
	Command       func(ctx Context)
	CommandStruct struct {
		Command
		Help string
	}
	Commands map[string]CommandStruct
	//Handler made to manage the pointer to commands created
	Handler struct {
		Commands
	}
)

func NewCommandHandler() *Handler {
	return &Handler{map[string]CommandStruct{}}
}

func (handler Handler) GetCommand(name string) (*CommandStruct, bool) {
	cmd, found := handler.Commands[name]
	return &cmd, found
}
