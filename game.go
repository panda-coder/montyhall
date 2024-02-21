package main

type Game struct {
	host   *Host
	player *Player
}

func NewGame(maintain bool) *Game {
	host := NewHost()
	host.SetupGate()

	return &Game{
		host:   host,
		player: NewPlayer(maintain),
	}
}

func (g *Game) TestGame() bool {
	hGate := g.host.GetGate()
	pGate := g.player.SelectGate()

	if g.player.GetMaintain() {
		return hGate == pGate
	}

	return hGate != pGate
}
