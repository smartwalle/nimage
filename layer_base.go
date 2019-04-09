package image4go

import (
	"image"
	"image/color"
	"image/draw"
)

type BaseLayer struct {
	point   Point
	size    Size
	layers  []Layer
	bgColor color.Color
}

func NewBaseLayer(width, height int) *BaseLayer {
	var l = &BaseLayer{}
	l.point = NewPoint(0, 0)
	l.size = NewSize(width, height)
	l.bgColor = color.Transparent
	return l
}

func (this *BaseLayer) AddLayer(layer Layer) {
	if isNil(layer) {
		return
	}
	this.layers = append(this.layers, layer)
}

func (this *BaseLayer) RemoveLayer(layer Layer) {
	if isNil(layer) {
		return
	}

	var index = -1
	for i, l := range this.layers {
		if l == layer {
			index = i
		}
	}

	if index > -1 {
		this.layers = append(this.layers[:index], this.layers[index+1:]...)
	}
}

func (this *BaseLayer) SetBackgroundColor(bgColor color.Color) {
	this.bgColor = bgColor
	if this.bgColor == nil {
		this.bgColor = color.Transparent
	}
}

func (this *BaseLayer) BackgroundColor() color.Color {
	return this.bgColor
}

func (this *BaseLayer) Render() image.Image {
	var mLayer = image.NewRGBA(image.Rect(0, 0, this.size.Width, this.size.Height))

	// 创建背景层
	if this.bgColor != nil {
		var bgLayer = image.NewUniform(this.bgColor)
		draw.Draw(mLayer, mLayer.Bounds(), bgLayer, image.ZP, draw.Src)
	}

	// 处理子 layer
	for _, layer := range this.layers {
		var img = layer.Render()
		if img != nil {
			draw.Draw(mLayer, layer.Rect(), img, image.ZP, draw.Over)
		}
	}
	return mLayer
}

func (this *BaseLayer) SetPoint(x, y int) {
	this.point = Point{X: x, Y: y}
}

func (this *BaseLayer) Point() Point {
	return this.point
}

func (this *BaseLayer) SetSize(width, height int) {
	this.size = Size{Width: width, Height: height}
}

func (this *BaseLayer) Size() Size {
	return this.size
}

func (this *BaseLayer) Rect() image.Rectangle {
	var r = image.Rectangle{}
	r.Min.X = this.point.X
	r.Min.Y = this.point.Y
	r.Max.X = this.point.X + this.size.Width
	r.Max.Y = this.point.Y + this.size.Height
	return r
}
