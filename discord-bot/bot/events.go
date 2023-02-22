package bot

type event struct {
	id              string
	creator         string
	title           string
	minParticipants int
	active          bool
}
