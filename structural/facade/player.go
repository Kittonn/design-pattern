package facade

type Player struct {
}

func (p *Player) On() {
	println("Player is ON")
}

func (p *Player) Play(name string) {
	println("Playing music: " + name)
}
