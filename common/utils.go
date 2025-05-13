package common

type Interactable interface {
	GetBounds() (float32, float32, float32, float32)
}

func Collide(x, y int, button Interactable) bool {
	left, top, right, bottom := button.GetBounds()
	X, Y := float32(x), float32(y)
	return X > left && X < right && Y > top && Y < bottom
}
