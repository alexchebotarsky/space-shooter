package game

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/goodleby/space-shooter/assets"
	"github.com/goodleby/space-shooter/asteroid"
	"github.com/goodleby/space-shooter/player"
	"github.com/goodleby/space-shooter/timer"
	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	assets *assets.Assets

	player *player.Player

	asteroidSpawnTimer *timer.Timer
	asteroids          []*asteroid.Asteroid
}

func New() (*Game, error) {
	var g Game
	var err error

	g.assets, err = assets.LoadAssets()
	if err != nil {
		return nil, fmt.Errorf("error loading assets: %v", err)
	}

	g.player = player.New(&g.assets.Player)

	g.asteroidSpawnTimer = timer.New(5 * time.Second)
	g.asteroids = []*asteroid.Asteroid{}

	return &g, nil
}

func (g *Game) Update() error {
	g.player.Update()

	if g.asteroidSpawnTimer.IsReady() {
		g.asteroidSpawnTimer.Reset()

		asteroidImg := g.assets.Asteroids[rand.Intn(len(g.assets.Asteroids))]

		g.asteroids = append(g.asteroids, asteroid.New(asteroidImg))
	}
	g.asteroidSpawnTimer.Update()

	for i, asteroid := range g.asteroids {
		if asteroid.IsOutOfBounds() {
			g.asteroids = append(g.asteroids[:i], g.asteroids[i+1:]...)
		}
		asteroid.Update()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)

	for _, asteroid := range g.asteroids {
		asteroid.Draw(screen)
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("%d", len(g.asteroids)))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
