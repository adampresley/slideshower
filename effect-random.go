package main

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

// RandomEffect implements a random selection of other effects
type RandomEffect struct {
	currentEffect Effect
	screenWidth   float64
	screenHeight  float64
	duration      int
	effectList    []string
}

func NewRandomEffect(durationInFrames int, screenWidth, screenHeight float64, effectList []string) *RandomEffect {
	result := &RandomEffect{
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
		duration:     durationInFrames,
		effectList:   effectList,
	}

	result.currentEffect = result.getRandomEffect()
	return result
}

func (re *RandomEffect) Update() {
	re.currentEffect.Update()
}

func (re *RandomEffect) Draw(screen *ebiten.Image, current, next *ebiten.Image) {
	if re.currentEffect != nil {
		re.currentEffect.Draw(screen, current, next)
	}
}

func (re *RandomEffect) IsComplete() bool {
	return re.currentEffect.IsComplete()
}

func (re *RandomEffect) Reset() {
	re.currentEffect = re.getRandomEffect()
	re.currentEffect.Reset()
}

func (re *RandomEffect) getRandomEffect() Effect {
	randomIndex := rand.Intn(len(re.effectList))
	return getEffectByName(re.effectList[randomIndex], re.screenWidth, re.screenHeight)
}
