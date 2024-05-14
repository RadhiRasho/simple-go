package main

type SnakePart struct {
	x, y int
}

type SnakeBody struct {
	Parts []SnakePart
	Xspeed, Yspeed int
}

func (sb *SnakeBody) ChangeDir(vertical, horizontal int) {
	sb.Yspeed = vertical
	sb.Xspeed = horizontal
}

func (sb *SnakeBody) Update(width, height int, longerSnake bool) {
	sb.Parts  = append(sb.Parts, sb.Parts[len(sb.Parts)-1].GetUpdatedPart(sb, width, height))
	if !longerSnake {
		sb.Parts = sb.Parts[1:]
	}
}

func (sp *SnakePart) GetUpdatedPart(sb *SnakeBody, width, height int) SnakePart {
	newPart := *sp;

	newPart.x = (newPart.x + sb.Xspeed) % width
	if newPart.x < 0 {
		newPart.x += width
	}
	newPart.y = (newPart.y + sb.Yspeed) % height
	if newPart.y < 0 {
		newPart.y += height
	}
	return newPart
}

func (sb *SnakeBody) ResetPos(width, height int) {
	snakeParts := []SnakePart{
		{
			x : int(width/2),
			y: int(height/2),
		},
		{
			x : int(width/2) + 1,
			y: int(height/2),
		},
		{
			x : int(width/2) + 2,
			y: int(height/2),
		},
	}

	sb.Parts = snakeParts
	sb.Xspeed = 1
	sb.Yspeed = 0
}