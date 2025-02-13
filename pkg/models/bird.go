package models

import (
	"time"
)

type Bird struct {
    Pos *Point
    Vel *Vector2D
    Acc *Vector2D
    world World
    lastScaleFactor float64
}

const BIRD_GRAVITY_Y = 40.8;

func CreateBird(world World) *Bird {
    return &Bird {
        Pos: NewPoint2D(0, 0),
        Vel: NewVector2D(0, 0),
        Acc: NewVector2D(0, 0),
        world: world,
        lastScaleFactor: 0,
    }
}

// Note: I am no physicist dammit, I am a scientist
func (b *Bird) Update(t time.Duration) {
    delta := float64(t.Milliseconds()) / 1000.0;
    b.Vel.Y += (BIRD_GRAVITY_Y * delta + b.Acc.Y) * b.world.ScalingYFactor()
    b.Pos.Y += b.Vel.Y * delta

    _, h := b.world.GetBounds()
    b.Pos.Y = min(b.Pos.Y, h)
    b.Acc.Y *= 0.25
}

func (b *Bird) Jump() {
    b.Acc.Y -= 18
    b.Vel.Y = -5
}

func (b *Bird) UpdateScreen() {
    currentScale := b.world.ScalingYFactor()
    if b.lastScaleFactor != 0 {
        // TODO: Adjust position, velocity, accel?
        changed := currentScale / b.lastScaleFactor;
        b.Pos.Y *= changed
        b.Vel.Y *= changed
        b.Acc.Y *= 0.5 * changed
    }

    b.lastScaleFactor = currentScale
}

func min(one float64, two int) float64 {
    two_f := float64(two)
    if two_f > one {
        return one
    }
    return two_f
}

func (b *Bird) Render(renderer Renderer) {
    bird := make([][]byte, 1)
    bird[0] = []byte{'@'}

    renderer.Render(b.Pos, bird)
}

