package objects

import (
	"github.com/wieku/gosu-pp/beatmap/audio"
	difficulty2 "github.com/wieku/gosu-pp/beatmap/difficulty"
	"github.com/wieku/gosu-pp/beatmap/timing"
	"github.com/wieku/gosu-pp/math/vector"
)

type IHitObject interface {
	Update(time float64) bool
	SetTiming(timings *timing.Timings)
	SetDifficulty(difficulty *difficulty2.Difficulty)

	GetStartTime() float64
	GetEndTime() float64
	GetDuration() float64

	GetPositionAt(float64) vector.Vector2f
	GetStackedPositionAt(float64) vector.Vector2f
	GetStackedPositionAtMod(time float64, modifier difficulty2.Modifier) vector.Vector2f

	GetStartPosition() vector.Vector2f
	GetStackedStartPosition() vector.Vector2f
	GetStackedStartPositionMod(modifier difficulty2.Modifier) vector.Vector2f

	GetEndPosition() vector.Vector2f
	GetStackedEndPosition() vector.Vector2f
	GetStackedEndPositionMod(modifier difficulty2.Modifier) vector.Vector2f

	GetID() int
	SetID(int)
	SetComboNumber(cn int)
	GetComboSet() int
	SetComboSet(set int)
	GetComboSetHax() int
	SetComboSetHax(set int)

	GetStackIndex(modifier difficulty2.Modifier) int
	SetStackIndex(index int, modifier difficulty2.Modifier)
	SetStackOffset(offset float32, modifier difficulty2.Modifier)

	GetSounds() []audio.HitSound

	GetColorOffset() int
	IsNewCombo() bool
	SetNewCombo(b bool)

	GetType() Type
}

type HitObject struct {
	StartPosRaw vector.Vector2f
	EndPosRaw   vector.Vector2f

	StartTime float64
	EndTime   float64

	StackOffset   vector.Vector2f
	StackOffsetEZ vector.Vector2f
	StackOffsetHR vector.Vector2f

	PositionDelegate func(time float64) vector.Vector2f

	StackIndex   int
	StackIndexEZ int
	StackIndexHR int

	HitObjectID int

	sounds []audio.HitSound

	NewCombo    bool
	ComboNumber int
	ComboSet    int
	ComboSetHax int
	ColorOffset int
}

func (hitObject *HitObject) Update(_ float64) bool { return true }

func (hitObject *HitObject) SetTiming(_ *timing.Timings) {}

func (hitObject *HitObject) UpdateStacking() {}

func (hitObject *HitObject) SetDifficulty(_ *difficulty2.Difficulty) {}

func (hitObject *HitObject) GetStartTime() float64 {
	return hitObject.StartTime
}

func (hitObject *HitObject) GetEndTime() float64 {
	return hitObject.EndTime
}

func (hitObject *HitObject) GetDuration() float64 {
	return hitObject.EndTime - hitObject.StartTime
}

func (hitObject *HitObject) GetPositionAt(time float64) vector.Vector2f {
	if hitObject.PositionDelegate != nil {
		return hitObject.PositionDelegate(time)
	}

	return hitObject.StartPosRaw
}

func (hitObject *HitObject) GetStackedPositionAt(time float64) vector.Vector2f {
	return hitObject.GetPositionAt(time).Add(hitObject.StackOffset)
}

func (hitObject *HitObject) GetStackedPositionAtMod(time float64, modifier difficulty2.Modifier) vector.Vector2f {
	return hitObject.modifyPosition(hitObject.GetPositionAt(time), modifier)
}

func (hitObject *HitObject) GetStartPosition() vector.Vector2f {
	return hitObject.StartPosRaw
}

func (hitObject *HitObject) GetStackedStartPosition() vector.Vector2f {
	return hitObject.GetStartPosition().Add(hitObject.StackOffset)
}

func (hitObject *HitObject) GetStackedStartPositionMod(modifier difficulty2.Modifier) vector.Vector2f {
	return hitObject.modifyPosition(hitObject.GetStartPosition(), modifier)
}

func (hitObject *HitObject) GetEndPosition() vector.Vector2f {
	return hitObject.EndPosRaw
}

func (hitObject *HitObject) GetStackedEndPosition() vector.Vector2f {
	return hitObject.GetEndPosition().Add(hitObject.StackOffset)
}

func (hitObject *HitObject) GetStackedEndPositionMod(modifier difficulty2.Modifier) vector.Vector2f {
	return hitObject.modifyPosition(hitObject.GetEndPosition(), modifier)
}

func (hitObject *HitObject) GetID() int {
	return hitObject.HitObjectID
}

func (hitObject *HitObject) SetID(id int) {
	hitObject.HitObjectID = id
}

func (hitObject *HitObject) SetComboNumber(cn int) {
	hitObject.ComboNumber = cn
}

func (hitObject *HitObject) GetComboSet() int {
	return hitObject.ComboSet
}

func (hitObject *HitObject) SetComboSet(set int) {
	hitObject.ComboSet = set
}

func (hitObject *HitObject) GetComboSetHax() int {
	return hitObject.ComboSetHax
}

func (hitObject *HitObject) SetComboSetHax(set int) {
	hitObject.ComboSetHax = set
}

func (hitObject *HitObject) GetStackIndex(modifier difficulty2.Modifier) int {
	switch {
	case modifier&difficulty2.HardRock > 0:
		return hitObject.StackIndexHR
	case modifier&difficulty2.Easy > 0:
		return hitObject.StackIndexEZ
	default:
		return hitObject.StackIndex
	}
}

func (hitObject *HitObject) SetStackIndex(index int, modifier difficulty2.Modifier) {
	switch {
	case modifier&difficulty2.HardRock > 0:
		hitObject.StackIndexHR = index
	case modifier&difficulty2.Easy > 0:
		hitObject.StackIndexEZ = index
	default:
		hitObject.StackIndex = index
	}
}

func (hitObject *HitObject) SetStackOffset(offset float32, modifier difficulty2.Modifier) {
	switch {
	case modifier&difficulty2.HardRock > 0:
		hitObject.StackOffsetHR = vector.NewVec2f(1, 1).Scl(offset)
	case modifier&difficulty2.Easy > 0:
		hitObject.StackOffsetEZ = vector.NewVec2f(1, 1).Scl(offset)
	default:
		hitObject.StackOffset = vector.NewVec2f(1, 1).Scl(offset)
	}
}

func (hitObject *HitObject) GetSounds() []audio.HitSound {
	return hitObject.sounds
}

func (hitObject *HitObject) GetColorOffset() int {
	return hitObject.ColorOffset
}

func (hitObject *HitObject) IsNewCombo() bool {
	return hitObject.NewCombo
}

func (hitObject *HitObject) SetNewCombo(b bool) {
	hitObject.NewCombo = b
}

func (hitObject *HitObject) modifyPosition(basePosition vector.Vector2f, modifier difficulty2.Modifier) vector.Vector2f {
	switch {
	case modifier&difficulty2.HardRock > 0:
		basePosition.Y = 384 - basePosition.Y
		return basePosition.Add(hitObject.StackOffsetHR)
	case modifier&difficulty2.Easy > 0:
		return basePosition.Add(hitObject.StackOffsetEZ)
	}

	return basePosition.Add(hitObject.StackOffset)
}
