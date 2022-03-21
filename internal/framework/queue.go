package framework

type (
	Song struct {
		Media    string
		Title    string
		Duration *string
		Id       string
	}
	SongQueue struct {
		list    []Song
		current *Song
		Running bool
	}
)

func newSongQueue() *SongQueue {
	queue := new(SongQueue)
	queue.list = make([]Song, 0)
	return queue
}
