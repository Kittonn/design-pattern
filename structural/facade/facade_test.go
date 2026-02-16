package facade

import "testing"

func TestHomeTheaterFacade(t *testing.T) {
	player := new(Player)
	speaker := new(Speaker)
	tv := new(TV)

	facade := NewHomeTheaterFacade(player, speaker, tv)
	result := facade.WatchMovie("Inception")

	if result != "Enjoy the movie: Inception" {
		t.Errorf("Expected 'Enjoy the movie: Inception', got '%s'", result)
	}
}
