package facade

type HomeTheaterFacade struct {
	player  *Player
	speaker *Speaker
	tv      *TV
}

func NewHomeTheaterFacade(player *Player, speaker *Speaker, tv *TV) *HomeTheaterFacade {
	return &HomeTheaterFacade{
		player:  player,
		speaker: speaker,
		tv:      tv,
	}
}

func (h *HomeTheaterFacade) WatchMovie(name string) string {
	h.tv.On()
	h.speaker.On()
	h.player.On()
	h.player.Play(name)

	return "Enjoy the movie: " + name
}
