// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package wand provides API definitions for
//accessing CORE_RL_wand_.dll.
//Based on GraphicsMagick v1.3.8.
package wand

import (
	. "github.com/tHinqa/outside"
	G "github.com/tHinqa/outside-graphicsmagick"
	. "github.com/tHinqa/outside/types"
)

// Opaque types
type (
	DrawingWand struct{}
	MagickWand  struct{}
	PixelWand   struct{}
)

func init() {
	AddDllApis(DLL, false, allApis)
	allApis = nil
}

// Drawing Wand

var CloneDrawingWand func(d *DrawingWand) *DrawingWand

func (d *DrawingWand) Clone() *DrawingWand { return CloneDrawingWand(d) }

var DestroyDrawingWand func(d *DrawingWand) *DrawingWand

func (d *DrawingWand) Destroy() *DrawingWand { return DestroyDrawingWand(d) }

var DrawAnnotation func(d *DrawingWand, x, y float64, text *uint8)

func (d *DrawingWand) Annotation(x, y float64, text *uint8) { DrawAnnotation(d, x, y, text) }

var DrawAffine func(d *DrawingWand, affine *AffineMatrix)

func (d *DrawingWand) Affine(affine *AffineMatrix) { DrawAffine(d, affine) }

type AffineMatrix struct{ Sx, Rx, Ry, Sy, Tx, Ty float64 }

var DrawAllocateWand func(*DrawInfo, *G.Image) *DrawingWand

type DrawInfo struct {
	Primitive *VString
	Geometry  *VString
	// Viewbox          G.RectangleInfo
	Affine           AffineMatrix
	Gravity          G.GravityType
	Fill             G.PixelPacket
	Stroke           G.PixelPacket
	StrokeWidth      float64
	Gradient         G.GradientInfo
	FillPattern      *G.Image
	Tile             *G.Image
	StrokePattern    *G.Image
	StrokeAntialias  G.MagickBooleanType
	TextAntialias    G.MagickBooleanType
	FillRule         G.FillRule
	Linecap          G.LineCap
	Linejoin         G.LineJoin
	Miterlimit       G.Size
	DashOffset       float64
	Decorate         G.DecorationType
	Compose          G.CompositeOperator
	Text             *VString
	Font             *VString
	Family           *VString
	Style            G.StyleType
	Stretch          G.StretchType
	Weight           G.Size
	Encoding         *VString
	Pointsize        float64
	Density          *VString
	Align            G.AlignType
	Undercolor       G.PixelPacket
	BorderColor      G.PixelPacket
	ServerName       *VString
	DashPattern      *float64
	ClipPath         *VString
	Bounds           G.SegmentInfo
	ClipUnits        G.ClipPathUnits
	Opacity          G.Quantum
	Render           G.MagickBooleanType
	Debug            G.MagickBooleanType
	ElementReference G.ElementReference
	Signature        G.Size
}

var DrawArc func(d *DrawingWand, sx, sy, ex, ey, sd, ed float64)

func (d *DrawingWand) Arc(sx, sy, ex, ey, sd, ed float64) { DrawArc(d, sx, sy, ex, ey, sd, ed) }

var DrawBezier func(d *DrawingWand, numberCoordinates uint32, coordinates *G.PointInfo)

func (d *DrawingWand) Bezier(numberCoordinates uint32, coordinates *G.PointInfo) {
	DrawBezier(d, numberCoordinates, coordinates)
}

var DrawCircle func(d *DrawingWand, ox, oy, px, py float64)

func (d *DrawingWand) Circle(ox, oy, px, py float64) { DrawCircle(d, ox, oy, px, py) }

var DrawClearException func(d *DrawingWand) bool

func (d *DrawingWand) ClearException() bool { return DrawClearException(d) }

var DrawGetClipPath func(d *DrawingWand) string

func (d *DrawingWand) ClipPath() string { return DrawGetClipPath(d) }

var DrawSetClipPath func(d *DrawingWand, clipMask string) bool

func (d *DrawingWand) SetClipPath(clipMask string) bool { return DrawSetClipPath(d, clipMask) }

var DrawGetClipRule func(d *DrawingWand) G.FillRule

func (d *DrawingWand) ClipRule() G.FillRule { return DrawGetClipRule(d) }

var DrawSetClipRule func(d *DrawingWand, fillRule G.FillRule)

func (d *DrawingWand) SetClipRule(fillRule G.FillRule) { DrawSetClipRule(d, fillRule) }

var DrawGetClipUnits func(d *DrawingWand) G.ClipPathUnits

func (d *DrawingWand) ClipUnits() G.ClipPathUnits { return DrawGetClipUnits(d) }

var DrawGetException func(d *DrawingWand, severity *G.ExceptionType) string

func (d *DrawingWand) Exception(severity *G.ExceptionType) string {
	return DrawGetException(d, severity)
}

var DrawSetClipUnits func(d *DrawingWand, clipUnits G.ClipPathUnits)

func (d *DrawingWand) SetClipUnits(clipUnits G.ClipPathUnits) { DrawSetClipUnits(d, clipUnits) }

var DrawColor func(d *DrawingWand, x, y float64, paintMethod G.PaintMethod)

func (d *DrawingWand) Color(x, y float64, paintMethod G.PaintMethod) {
	DrawColor(d, x, y, paintMethod)
}

var DrawComment func(d *DrawingWand, comment string)

func (d *DrawingWand) Comment(comment string) { DrawComment(d, comment) }

var DrawEllipse func(d *DrawingWand, ox, oy, rx, ry, start, end float64)

func (d *DrawingWand) Ellipse(ox, oy, rx, ry, start, end float64) {
	DrawEllipse(d, ox, oy, rx, ry, start, end)
}

var DrawGetFillColor func(d *DrawingWand, fillColor *PixelWand)

func (d *DrawingWand) FillColor(fillColor *PixelWand) { DrawGetFillColor(d, fillColor) }

var DrawSetFillColor func(d *DrawingWand, fillWand *PixelWand)

func (d *DrawingWand) SetFillColor(fillWand *PixelWand) { DrawSetFillColor(d, fillWand) }

var DrawSetFillPatternURL func(d *DrawingWand, fillUrl string) bool

func (d *DrawingWand) SetFillPatternURL(fillUrl string) bool {
	return DrawSetFillPatternURL(d, fillUrl)
}

var DrawGetFillOpacity func(d *DrawingWand) float64

func (d *DrawingWand) FillOpacity() float64 { return DrawGetFillOpacity(d) }

var DrawSetFillOpacity func(d *DrawingWand, fillOpacity float64)

func (d *DrawingWand) SetFillOpacity(fillOpacity float64) { DrawSetFillOpacity(d, fillOpacity) }

var DrawGetFillRule func(d *DrawingWand) G.FillRule

func (d *DrawingWand) FillRule() G.FillRule { return DrawGetFillRule(d) }

var DrawSetFillRule func(d *DrawingWand, fillRule G.FillRule)

func (d *DrawingWand) SetFillRule(fillRule G.FillRule) { DrawSetFillRule(d, fillRule) }

var DrawGetFont func(d *DrawingWand) string

func (d *DrawingWand) Font() string { return DrawGetFont(d) }

var DrawSetFont func(d *DrawingWand, fontName string) bool

func (d *DrawingWand) SetFont(fontName string) bool { return DrawSetFont(d, fontName) }

var DrawGetFontFamily func(d *DrawingWand) string

func (d *DrawingWand) FontFamily() string { return DrawGetFontFamily(d) }

var DrawSetFontFamily func(d *DrawingWand, fontFamily string) bool

func (d *DrawingWand) SetFontFamily(fontFamily string) bool {
	return DrawSetFontFamily(d, fontFamily)
}

var DrawGetFontSize func(d *DrawingWand) float64

func (d *DrawingWand) FontSize() float64 { return DrawGetFontSize(d) }

var DrawSetFontSize func(d *DrawingWand, pointsize float64)

func (d *DrawingWand) SetFontSize(pointsize float64) { DrawSetFontSize(d, pointsize) }

var DrawGetFontStretch func(d *DrawingWand) G.StretchType

func (d *DrawingWand) FontStretch() G.StretchType { return DrawGetFontStretch(d) }

var DrawSetFontStretch func(d *DrawingWand, fontStretch G.StretchType)

func (d *DrawingWand) SetFontStretch(fontStretch G.StretchType) { DrawSetFontStretch(d, fontStretch) }

var DrawGetFontStyle func(d *DrawingWand) G.StyleType

func (d *DrawingWand) FontStyle() G.StyleType { return DrawGetFontStyle(d) }

var DrawSetFontStyle func(d *DrawingWand, style G.StyleType)

func (d *DrawingWand) SetFontStyle(style G.StyleType) { DrawSetFontStyle(d, style) }

var DrawGetFontWeight func(d *DrawingWand) uint32

func (d *DrawingWand) FontWeight() uint32 { return DrawGetFontWeight(d) }

var DrawSetFontWeight func(d *DrawingWand, fontWeight uint32)

func (d *DrawingWand) SetFontWeight(fontWeight uint32) { DrawSetFontWeight(d, fontWeight) }

var DrawGetGravity func(d *DrawingWand) G.GravityType

func (d *DrawingWand) Gravity() G.GravityType { return DrawGetGravity(d) }

var DrawSetGravity func(d *DrawingWand, gravity G.GravityType)

func (d *DrawingWand) SetGravity(gravity G.GravityType) { DrawSetGravity(d, gravity) }

var DrawComposite func(d *DrawingWand, compose G.CompositeOperator, x, y, width, height float64, magickWand *MagickWand) bool

func (d *DrawingWand) Composite(compose G.CompositeOperator, x, y, width, height float64, magickWand *MagickWand) bool {
	return DrawComposite(d, compose, x, y, width, height, magickWand)
}

var DrawLine func(d *DrawingWand, sx, sy, ex, ey float64)

func (d *DrawingWand) Line(sx, sy, ex, ey float64) { DrawLine(d, sx, sy, ex, ey) }

var DrawMatte func(d *DrawingWand, x, y float64, paintMethod G.PaintMethod)

func (d *DrawingWand) Matte(x, y float64, paintMethod G.PaintMethod) {
	DrawMatte(d, x, y, paintMethod)
}

var DrawPathClose func(d *DrawingWand)

func (d *DrawingWand) PathClose() { DrawPathClose(d) }

var DrawPathCurveToAbsolute func(d *DrawingWand, x1, y1, x2, y2, x, y float64)

func (d *DrawingWand) PathCurveToAbsolute(x1, y1, x2, y2, x, y float64) {
	DrawPathCurveToAbsolute(d, x1, y1, x2, y2, x, y)
}

var DrawPathCurveToRelative func(d *DrawingWand, x1, y1, x2, y2, x, y float64)

func (d *DrawingWand) PathCurveToRelative(x1, y1, x2, y2, x, y float64) {
	DrawPathCurveToRelative(d, x1, y1, x2, y2, x, y)
}

var DrawPathCurveToQuadraticBezierAbsolute func(d *DrawingWand, x1, y1, x, y float64)

func (d *DrawingWand) PathCurveToQuadraticBezierAbsolute(x1, y1, x, y float64) {
	DrawPathCurveToQuadraticBezierAbsolute(d, x1, y1, x, y)
}

var DrawPathCurveToQuadraticBezierRelative func(d *DrawingWand, x1, y1, x, y float64)

func (d *DrawingWand) PathCurveToQuadraticBezierRelative(x1, y1, x, y float64) {
	DrawPathCurveToQuadraticBezierRelative(d, x1, y1, x, y)
}

var DrawPathCurveToQuadraticBezierSmoothAbsolute func(d *DrawingWand, x, y float64)

func (d *DrawingWand) PathCurveToQuadraticBezierSmoothAbsolute(x, y float64) {
	DrawPathCurveToQuadraticBezierSmoothAbsolute(d, x, y)
}

var DrawPathCurveToQuadraticBezierSmoothRelative func(d *DrawingWand, x, y float64)

func (d *DrawingWand) PathCurveToQuadraticBezierSmoothRelative(x, y float64) {
	DrawPathCurveToQuadraticBezierSmoothRelative(d, x, y)
}

var DrawPathCurveToSmoothAbsolute func(d *DrawingWand, x2, y2, x, y float64)

func (d *DrawingWand) PathCurveToSmoothAbsolute(x2, y2, x, y float64) {
	DrawPathCurveToSmoothAbsolute(d, x2, y2, x, y)
}

var DrawPathCurveToSmoothRelative func(d *DrawingWand, x2, y2, x, y float64)

func (d *DrawingWand) PathCurveToSmoothRelative(x2, y2, x, y float64) {
	DrawPathCurveToSmoothRelative(d, x2, y2, x, y)
}

var DrawPathEllipticArcAbsolute func(d *DrawingWand, rx, ry, xAxisRotation float64, largeArcFlag, sweepFlag bool, x, y float64)

func (d *DrawingWand) PathEllipticArcAbsolute(rx, ry, xAxisRotation float64, largeArcFlag, sweepFlag bool, x, y float64) {
	DrawPathEllipticArcAbsolute(d, rx, ry, xAxisRotation, largeArcFlag, sweepFlag, x, y)
}

var DrawPathEllipticArcRelative func(d *DrawingWand, rx, ry, xAxisRotation float64, largeArcFlag, sweepFlag bool, x, y float64)

func (d *DrawingWand) PathEllipticArcRelative(rx, ry, xAxisRotation float64, largeArcFlag, sweepFlag bool, x, y float64) {
	DrawPathEllipticArcRelative(d, rx, ry, xAxisRotation, largeArcFlag, sweepFlag, x, y)
}

var DrawPathFinish func(d *DrawingWand)

func (d *DrawingWand) PathFinish() { DrawPathFinish(d) }

var DrawPathLineToAbsolute func(d *DrawingWand, x, y float64)

func (d *DrawingWand) PathLineToAbsolute(x, y float64) { DrawPathLineToAbsolute(d, x, y) }

var DrawPathLineToRelative func(d *DrawingWand, x, y float64)

func (d *DrawingWand) PathLineToRelative(x, y float64) { DrawPathLineToRelative(d, x, y) }

var DrawPathLineToHorizontalAbsolute func(d *DrawingWand, x float64)

func (d *DrawingWand) PathLineToHorizontalAbsolute(x float64) {
	DrawPathLineToHorizontalAbsolute(d, x)
}

var DrawPathLineToHorizontalRelative func(d *DrawingWand, x float64)

func (d *DrawingWand) PathLineToHorizontalRelative(x float64) {
	DrawPathLineToHorizontalRelative(d, x)
}

var DrawPathLineToVerticalAbsolute func(d *DrawingWand, y float64)

func (d *DrawingWand) PathLineToVerticalAbsolute(y float64) { DrawPathLineToVerticalAbsolute(d, y) }

var DrawPathLineToVerticalRelative func(d *DrawingWand, y float64)

func (d *DrawingWand) PathLineToVerticalRelative(y float64) { DrawPathLineToVerticalRelative(d, y) }

var DrawPathMoveToAbsolute func(d *DrawingWand, x, y float64)

func (d *DrawingWand) PathMoveToAbsolute(x, y float64) { DrawPathMoveToAbsolute(d, x, y) }

var DrawPathMoveToRelative func(d *DrawingWand, x, y float64)

func (d *DrawingWand) PathMoveToRelative(x, y float64) { DrawPathMoveToRelative(d, x, y) }

var DrawPathStart func(d *DrawingWand)

func (d *DrawingWand) PathStart() { DrawPathStart(d) }

var DrawPeekGraphicContext func(d *DrawingWand) *DrawInfo
var DrawPoint func(d *DrawingWand, x, y float64)

func (d *DrawingWand) Point(x, y float64) { DrawPoint(d, x, y) }

var DrawPolygon func(d *DrawingWand, numberCoordinates uint32, coordinates *G.PointInfo)

func (d *DrawingWand) Polygon(numberCoordinates uint32, coordinates *G.PointInfo) {
	DrawPolygon(d, numberCoordinates, coordinates)
}

var DrawPolyline func(d *DrawingWand, numberCoordinates uint32, coordinates *G.PointInfo)

func (d *DrawingWand) Polyline(numberCoordinates uint32, coordinates *G.PointInfo) {
	DrawPolyline(d, numberCoordinates, coordinates)
}

var DrawPopClipPath func(d *DrawingWand)

func (d *DrawingWand) PopClipPath() { DrawPopClipPath(d) }

var DrawPopDefs func(d *DrawingWand)

func (d *DrawingWand) PopDefs() { DrawPopDefs(d) }

var DrawPopGraphicContext func(d *DrawingWand) bool

func (d *DrawingWand) PopGraphicContext() bool { return DrawPopGraphicContext(d) }

var DrawPopPattern func(d *DrawingWand) bool

func (d *DrawingWand) PopPattern() bool { return DrawPopPattern(d) }

var DrawPushClipPath func(d *DrawingWand, clipMaskId string)

func (d *DrawingWand) PushClipPath(clipMaskId string) { DrawPushClipPath(d, clipMaskId) }

var DrawPushDefs func(d *DrawingWand)

func (d *DrawingWand) PushDefs() { DrawPushDefs(d) }

var DrawPushGraphicContext func(d *DrawingWand) bool

func (d *DrawingWand) PushGraphicContext() bool { return DrawPushGraphicContext(d) }

var DrawPushPattern func(d *DrawingWand, patternId string, x, y, width, height float64) bool

func (d *DrawingWand) PushPattern(patternId string, x, y, width, height float64) bool {
	return DrawPushPattern(d, patternId, x, y, width, height)
}

var DrawRectangle func(d *DrawingWand, x1, y1, x2, y2 float64)

func (d *DrawingWand) Rectangle(x1, y1, x2, y2 float64) { DrawRectangle(d, x1, y1, x2, y2) }

var DrawRender func(d *DrawingWand) bool

func (d *DrawingWand) Render() bool { return DrawRender(d) }

var DrawRotate func(d *DrawingWand, degrees float64)

func (d *DrawingWand) Rotate(degrees float64) { DrawRotate(d, degrees) }

var DrawRoundRectangle func(d *DrawingWand, x1, y1, x2, y2, rx, ry float64)

func (d *DrawingWand) RoundRectangle(x1, y1, x2, y2, rx, ry float64) {
	DrawRoundRectangle(d, x1, y1, x2, y2, rx, ry)
}

var DrawScale func(d *DrawingWand, x, y float64)

func (d *DrawingWand) Scale(x, y float64) { DrawScale(d, x, y) }

var DrawSkewX func(d *DrawingWand, degrees float64)

func (d *DrawingWand) SkewX(degrees float64) { DrawSkewX(d, degrees) }

var DrawSkewY func(d *DrawingWand, degrees float64)

func (d *DrawingWand) SkewY(degrees float64) { DrawSkewY(d, degrees) }

// DrawSetStopColor documented but unimplemented (#if 0)
var DrawGetStrokeColor func(d *DrawingWand, strokeColor *PixelWand)

func (d *DrawingWand) StrokeColor(strokeColor *PixelWand) { DrawGetStrokeColor(d, strokeColor) }

var DrawSetStrokeColor func(d *DrawingWand, strokeWand *PixelWand)

func (d *DrawingWand) SetStrokeColor(strokeWand *PixelWand) { DrawSetStrokeColor(d, strokeWand) }

var DrawSetStrokePatternURL func(d *DrawingWand, strokeUrl string) bool

func (d *DrawingWand) SetStrokePatternURL(strokeUrl string) bool {
	return DrawSetStrokePatternURL(d, strokeUrl)
}

var DrawGetStrokeAntialias func(d *DrawingWand) bool

func (d *DrawingWand) StrokeAntialias() bool { return DrawGetStrokeAntialias(d) }

var DrawSetStrokeAntialias func(d *DrawingWand, strokeAntialias bool)

func (d *DrawingWand) SetStrokeAntialias(strokeAntialias bool) {
	DrawSetStrokeAntialias(d, strokeAntialias)
}

var DrawSetStrokeDashArray func(d *DrawingWand, numberElements uint32, dashArray *float64) bool

func (d *DrawingWand) SetStrokeDashArray(numberElements uint32, dashArray *float64) bool {
	return DrawSetStrokeDashArray(d, numberElements, dashArray)
}

var DrawGetStrokeDashArray func(d *DrawingWand, numberElements *uint32) *float64

func (d *DrawingWand) StrokeDashArray(numberElements *uint32) *float64 {
	return DrawGetStrokeDashArray(d, numberElements)
}

var DrawGetStrokeDashOffset func(d *DrawingWand) float64

func (d *DrawingWand) StrokeDashOffset() float64 { return DrawGetStrokeDashOffset(d) }

var DrawSetStrokeDashOffset func(d *DrawingWand, dashOffset float64)

func (d *DrawingWand) SetStrokeDashOffset(dashOffset float64) {
	DrawSetStrokeDashOffset(d, dashOffset)
}

var DrawGetStrokeLineCap func(d *DrawingWand) G.LineCap

func (d *DrawingWand) StrokeLineCap() G.LineCap { return DrawGetStrokeLineCap(d) }

var DrawSetStrokeLineCap func(d *DrawingWand, linecap G.LineCap)

func (d *DrawingWand) SetStrokeLineCap(linecap G.LineCap) { DrawSetStrokeLineCap(d, linecap) }

var DrawGetStrokeLineJoin func(d *DrawingWand) G.LineJoin

func (d *DrawingWand) StrokeLineJoin() G.LineJoin { return DrawGetStrokeLineJoin(d) }

var DrawSetStrokeLineJoin func(d *DrawingWand, linejoin G.LineJoin)

func (d *DrawingWand) SetStrokeLineJoin(linejoin G.LineJoin) { DrawSetStrokeLineJoin(d, linejoin) }

var DrawGetStrokeMiterLimit func(d *DrawingWand) uint32

func (d *DrawingWand) StrokeMiterLimit() uint32 { return DrawGetStrokeMiterLimit(d) }

var DrawSetStrokeMiterLimit func(d *DrawingWand, miterlimit uint32)

func (d *DrawingWand) SetStrokeMiterLimit(miterlimit uint32) { DrawSetStrokeMiterLimit(d, miterlimit) }

var DrawGetStrokeOpacity func(d *DrawingWand) float64

func (d *DrawingWand) StrokeOpacity() float64 { return DrawGetStrokeOpacity(d) }

var DrawSetStrokeOpacity func(d *DrawingWand, strokeOpacity float64)

func (d *DrawingWand) SetStrokeOpacity(strokeOpacity float64) {
	DrawSetStrokeOpacity(d, strokeOpacity)
}

var DrawGetStrokeWidth func(d *DrawingWand) float64

func (d *DrawingWand) StrokeWidth() float64 { return DrawGetStrokeWidth(d) }

var DrawSetStrokeWidth func(d *DrawingWand, strokeWidth float64)

func (d *DrawingWand) SetStrokeWidth(strokeWidth float64) { DrawSetStrokeWidth(d, strokeWidth) }

var DrawGetTextAntialias func(d *DrawingWand) bool

func (d *DrawingWand) TextAntialias() bool { return DrawGetTextAntialias(d) }

var DrawSetTextAntialias func(d *DrawingWand, textAntialias bool)

func (d *DrawingWand) SetTextAntialias(textAntialias bool) { DrawSetTextAntialias(d, textAntialias) }

var DrawGetTextDecoration func(d *DrawingWand) G.DecorationType

func (d *DrawingWand) TextDecoration() G.DecorationType { return DrawGetTextDecoration(d) }

var DrawSetTextDecoration func(d *DrawingWand, decoration G.DecorationType)

func (d *DrawingWand) SetTextDecoration(decoration G.DecorationType) {
	DrawSetTextDecoration(d, decoration)
}

var DrawGetTextEncoding func(d *DrawingWand) string

func (d *DrawingWand) TextEncoding() string { return DrawGetTextEncoding(d) }

var DrawSetTextEncoding func(d *DrawingWand, encoding string)

func (d *DrawingWand) SetTextEncoding(encoding string) { DrawSetTextEncoding(d, encoding) }

var DrawGetTextUnderColor func(d *DrawingWand, underColor *PixelWand)

func (d *DrawingWand) TextUnderColor(underColor *PixelWand) { DrawGetTextUnderColor(d, underColor) }

var DrawSetTextUnderColor func(d *DrawingWand, underWand *PixelWand)

func (d *DrawingWand) SetTextUnderColor(underWand *PixelWand) { DrawSetTextUnderColor(d, underWand) }

var DrawTranslate func(d *DrawingWand, x, y float64)

func (d *DrawingWand) Translate(x, y float64) { DrawTranslate(d, x, y) }

var DrawSetViewbox func(d *DrawingWand, x1, y1, x2, y2 int32)

func (d *DrawingWand) SetViewbox(x1, y1, x2, y2 int32) { DrawSetViewbox(d, x1, y1, x2, y2) }

var NewDrawingWand func() *DrawingWand

// Magick Wand
var CloneMagickWand func(m *MagickWand) *MagickWand

func (m *MagickWand) Clone() *MagickWand { return CloneMagickWand(m) }

var DestroyMagickWand func(m *MagickWand)

func (m *MagickWand) Destroy() (_ *MagickWand) { DestroyMagickWand(m); return }

var AdaptiveThresholdImage func(m *MagickWand, width, height uint32, offset int32) bool

func (m *MagickWand) AdaptiveThreshold(width, height uint32, offset int32) bool {
	return AdaptiveThresholdImage(m, width, height, offset)
}

var AddImage func(m *MagickWand, addWand *MagickWand) bool

func (m *MagickWand) Add(addWand *MagickWand) bool { return AddImage(m, addWand) }

var AddNoiseImage func(m *MagickWand, noiseType G.NoiseType) bool

func (m *MagickWand) AddNoise(noiseType G.NoiseType) bool {
	return AddNoiseImage(m, noiseType)
}

var AffineTransformImage func(m *MagickWand, d *DrawingWand) bool

func (m *MagickWand) AffineTransform(d *DrawingWand) bool {
	return AffineTransformImage(m, d)
}

var AnnotateImage func(m *MagickWand, d *DrawingWand, x, y, angle float64, text string) bool

func (m *MagickWand) Annotate(d *DrawingWand, x, y, angle float64, text string) bool {
	return AnnotateImage(m, d, x, y, angle, text)
}

var AnimateImages func(m *MagickWand, xServerName string) bool

func (m *MagickWand) Animate(xServerName string) bool {
	return AnimateImages(m, xServerName)
}

var AppendImages func(m *MagickWand, stack bool) *MagickWand

func (m *MagickWand) Append(stack bool) *MagickWand { return AppendImages(m, stack) }

var AverageImages func(m *MagickWand) *MagickWand

func (m *MagickWand) Average() *MagickWand { return AverageImages(m) }

var BlackThresholdImage func(m *MagickWand, threshold *PixelWand) bool

func (m *MagickWand) BlackThreshold(threshold *PixelWand) bool {
	return BlackThresholdImage(m, threshold)
}

var BlurImage func(m *MagickWand, radius, sigma float64) bool

func (m *MagickWand) Blur(radius, sigma float64) bool { return BlurImage(m, radius, sigma) }

var BorderImage func(m *MagickWand, bordercolor *PixelWand, width, height uint32) bool

func (m *MagickWand) Border(bordercolor *PixelWand, width, height uint32) bool {
	return BorderImage(m, bordercolor, width, height)
}

var CdlImage func(wand *MagickWand, cdl string) uint

func (m *MagickWand) CdlImage(cdl string) uint { return CdlImage(m, cdl) }

var CharcoalImage func(m *MagickWand, radius, sigma float64) bool

func (m *MagickWand) Charcoal(radius, sigma float64) bool {
	return CharcoalImage(m, radius, sigma)
}

var ChopImage func(m *MagickWand, width, height uint32, x, y int32) bool

func (m *MagickWand) Chop(width, height uint32, x, y int32) bool {
	return ChopImage(m, width, height, x, y)
}

var ClipImage func(m *MagickWand) bool

func (m *MagickWand) Clip() bool { return ClipImage(m) }

var ClipPathImage func(m *MagickWand, pathname string, inside bool) bool

func (m *MagickWand) ClipPath(pathname string, inside bool) bool {
	return ClipPathImage(m, pathname, inside)
}

var CoalesceImages func(m *MagickWand) *MagickWand

func (m *MagickWand) Coalesce() *MagickWand { return CoalesceImages(m) }

var ColorFloodfillImage func(m *MagickWand, fill *PixelWand, fuzz float64, bordercolor *PixelWand, x, y int32) bool

func (m *MagickWand) ColorFloodfill(fill *PixelWand, fuzz float64, bordercolor *PixelWand, x, y int32) bool {
	return ColorFloodfillImage(m, fill, fuzz, bordercolor, x, y)
}

var ColorizeImage func(m *MagickWand, colorize, opacity *PixelWand) bool

func (m *MagickWand) Colorize(colorize, opacity *PixelWand) bool {
	return ColorizeImage(m, colorize, opacity)
}

var CommentImage func(m *MagickWand, comment string) bool

func (m *MagickWand) Comment(comment string) bool { return CommentImage(m, comment) }

var CompareImageChannels func(m *MagickWand, reference *MagickWand, channel G.ChannelType, metric G.MetricType, distortion *float64) *MagickWand

func (m *MagickWand) CompareChannels(reference *MagickWand, channel G.ChannelType, metric G.MetricType, distortion *float64) *MagickWand {
	return CompareImageChannels(m, reference, channel, metric, distortion)
}

var CompareImages func(m *MagickWand, reference *MagickWand, metric G.MetricType, distortion *float64) *MagickWand

func (m *MagickWand) Compare(reference *MagickWand, metric G.MetricType, distortion *float64) *MagickWand {
	return CompareImages(m, reference, metric, distortion)
}

var CompositeImage func(m *MagickWand, sourceWand *MagickWand, compose G.CompositeOperator, x, y int32) bool

func (m *MagickWand) Composite(sourceWand *MagickWand, compose G.CompositeOperator, x, y int32) bool {
	return CompositeImage(m, sourceWand, compose, x, y)
}

var ContrastImage func(m *MagickWand, sharpen bool) bool

func (m *MagickWand) Contrast(sharpen bool) bool { return ContrastImage(m, sharpen) }

var ConvolveImage func(m *MagickWand, order uint32, kernel []float64) bool

func (m *MagickWand) Convolve(order uint32, kernel []float64) bool {
	return ConvolveImage(m, order, kernel)
}

var CropImage func(m *MagickWand, width, height uint32, x, y int32) bool

func (m *MagickWand) Crop(width, height uint32, x, y int32) bool {
	return CropImage(m, width, height, x, y)
}

var CycleColormapImage func(m *MagickWand, displace int32) bool

func (m *MagickWand) CycleColormap(displace int32) bool {
	return CycleColormapImage(m, displace)
}

var DeconstructImages func(m *MagickWand) *MagickWand

func (m *MagickWand) Deconstruct() *MagickWand { return DeconstructImages(m) }

var DescribeImage func(m *MagickWand) string

func (m *MagickWand) Describe() string { return DescribeImage(m) }

var DespeckleImage func(m *MagickWand) bool

func (m *MagickWand) Despeckle() bool { return DespeckleImage(m) }

var DisplayImage func(m *MagickWand, xServerName string) bool

func (m *MagickWand) Display(xServerName string) bool { return DisplayImage(m, xServerName) }

var DisplayImages func(m *MagickWand, xServerName string) bool

func (m *MagickWand) DisplayImages(xServerName string) bool {
	return DisplayImages(m, xServerName)
}

var DrawImage func(m *MagickWand, d *DrawingWand) bool

func (m *MagickWand) Draw(d *DrawingWand) bool { return DrawImage(m, d) }

var EdgeImage func(m *MagickWand, radius float64) bool

func (m *MagickWand) Edge(radius float64) bool { return EdgeImage(m, radius) }

var EmbossImage func(m *MagickWand, radius, sigma float64) bool

func (m *MagickWand) Emboss(radius, sigma float64) bool {
	return EmbossImage(m, radius, sigma)
}

var EnhanceImage func(m *MagickWand) bool

func (m *MagickWand) Enhance() bool { return EnhanceImage(m) }

var EqualizeImage func(m *MagickWand) bool

func (m *MagickWand) Equalize() bool { return EqualizeImage(m) }

var ExtentImage func(m *MagickWand, width, height uint32, x, y int32) bool

func (m *MagickWand) Extent(width, height uint32, x, y int32) bool {
	return ExtentImage(m, width, height, x, y)
}

var FlattenImages func(m *MagickWand) *MagickWand

func (m *MagickWand) Flatten() *MagickWand { return FlattenImages(m) }

var FlipImage func(m *MagickWand) bool

func (m *MagickWand) Flip() bool { return FlipImage(m) }

var FlopImage func(m *MagickWand) bool

func (m *MagickWand) Flop() bool { return FlopImage(m) }

var FrameImage func(m *MagickWand, matteColor *PixelWand, width, height uint32, innerBevel, outerBevel int32) bool

func (m *MagickWand) Frame(matteColor *PixelWand, width, height uint32, innerBevel, outerBevel int32) bool {
	return FrameImage(m, matteColor, width, height, innerBevel, outerBevel)
}

var FxImage func(m *MagickWand, expression string) *MagickWand

func (m *MagickWand) Fx(expression string) *MagickWand { return FxImage(m, expression) }

var FxImageChannel func(m *MagickWand, channel G.ChannelType, expression string) *MagickWand

func (m *MagickWand) FxChannel(channel G.ChannelType, expression string) *MagickWand {
	return FxImageChannel(m, channel, expression)
}

var GammaImage func(m *MagickWand, gamma float64) bool

func (m *MagickWand) Gamma(gamma float64) bool { return GammaImage(m, gamma) }

var GammaImageChannel func(m *MagickWand, channel G.ChannelType, gamma float64) bool

func (m *MagickWand) GammaChannel(channel G.ChannelType, gamma float64) bool {
	return GammaImageChannel(m, channel, gamma)
}

var GetConfigureInfo func(m *MagickWand, attribute string) string

func (m *MagickWand) ConfigureInfo(attribute string) string {
	return GetConfigureInfo(m, attribute)
}

var GetCopyright func() string

var GetException func(m *MagickWand, severity *G.ExceptionType) string

func (m *MagickWand) Exception(severity *G.ExceptionType) string {
	return GetException(m, severity)
}

var GetFilename func(m *MagickWand) string

func (m *MagickWand) Filename() string { return GetFilename(m) }

var GetHomeURL func() string
var GetImage func(m *MagickWand) *MagickWand

func (m *MagickWand) Image() *MagickWand { return GetImage(m) }

var GetImageAttribute func(m *MagickWand, property string) string

func (m *MagickWand) Attribute(property string) string {
	return GetImageAttribute(m, property)
}

var GetImageBackgroundColor func(m *MagickWand, backgroundColor *PixelWand) bool

func (m *MagickWand) ImageBackgroundColor(backgroundColor *PixelWand) bool {
	return GetImageBackgroundColor(m, backgroundColor)
}

var GetImageBluePrimary func(m *MagickWand, x, y *float64) bool

func (m *MagickWand) BluePrimary(x, y *float64) bool { return GetImageBluePrimary(m, x, y) }

var GetImageBorderColor func(m *MagickWand, borderColor *PixelWand) bool

func (m *MagickWand) BorderColor(borderColor *PixelWand) bool {
	return GetImageBorderColor(m, borderColor)
}

var GetImageBoundingBox func(wand *MagickWand, fuzz float64, width, height *G.Size, x, y *G.SSize) uint

func (m *MagickWand) BoundingBox(fuzz float64, width, height *G.Size, x, y *G.SSize) uint {
	return GetImageBoundingBox(m, fuzz, width, height, x, y)
}

var GetImageChannelDepth func(m *MagickWand, channel G.ChannelType) uint32

func (m *MagickWand) ChannelDepth(channel G.ChannelType) uint32 {
	return GetImageChannelDepth(m, channel)
}

var GetImageChannelExtrema func(m *MagickWand, c G.ChannelType, minima, maxima *uint32) bool

func (m *MagickWand) GetImageChannelExtrema(c G.ChannelType, minima, maxima *uint32) bool {
	return GetImageChannelExtrema(m, c, minima, maxima)
}

var GetImageChannelMean func(m *MagickWand, channel G.ChannelType, mean, standardDeviation *float64) bool

func (m *MagickWand) ChannelMean(channel G.ChannelType, mean, standardDeviation *float64) bool {
	return GetImageChannelMean(m, channel, mean, standardDeviation)
}

var GetImageColormapColor func(m *MagickWand, index uint32, color *PixelWand) bool

func (m *MagickWand) ColormapColor(index uint32, color *PixelWand) bool {
	return GetImageColormapColor(m, index, color)
}

var GetImageColors func(m *MagickWand) uint32

func (m *MagickWand) Colors() uint32 { return GetImageColors(m) }

var GetImageColorspace func(m *MagickWand) G.ColorspaceType

func (m *MagickWand) ImageColorspace() G.ColorspaceType { return GetImageColorspace(m) }

var GetImageCompose func(m *MagickWand) G.CompositeOperator

func (m *MagickWand) Compose() G.CompositeOperator { return GetImageCompose(m) }

var GetImageCompression func(m *MagickWand) G.CompressionType

func (m *MagickWand) ImageCompression() G.CompressionType { return GetImageCompression(m) }

var GetImageDelay func(m *MagickWand) uint32

func (m *MagickWand) Delay() uint32 { return GetImageDelay(m) }

var GetImageDepth func(m *MagickWand) uint32

func (m *MagickWand) Depth() uint32 { return GetImageDepth(m) }

var GetImageExtrema func(m *MagickWand, min, max *uint32) bool

func (m *MagickWand) Extrema(min, max *uint32) bool { return GetImageExtrema(m, min, max) }

var GetImageDispose func(m *MagickWand) G.DisposeType

func (m *MagickWand) Dispose() G.DisposeType { return GetImageDispose(m) }

var GetImageFilename func(m *MagickWand) string

func (m *MagickWand) ImageFilename() string { return GetImageFilename(m) }

var GetImageFormat func(m *MagickWand) string

func (m *MagickWand) ImageFormat() string { return GetImageFormat(m) }

var GetImageFuzz func(m *MagickWand) float64

func (m *MagickWand) Fuzz() float64 { return GetImageFuzz(m) }

var GetImageGamma func(m *MagickWand) float64

func (m *MagickWand) ImageGamma() float64 { return GetImageGamma(m) }

var GetImageGreenPrimary func(m *MagickWand, x, y *float64) bool

func (m *MagickWand) GreenPrimary(x, y *float64) bool { return GetImageGreenPrimary(m, x, y) }

var GetImageHeight func(m *MagickWand) uint32

func (m *MagickWand) Height() uint32 { return GetImageHeight(m) }

var GetImageHistogram func(m *MagickWand, numberColors *uint32) []*PixelWand

func (m *MagickWand) Histogram(numberColors *uint32) []*PixelWand {
	return GetImageHistogram(m, numberColors)
}

var GetImageIndex func(m *MagickWand) int32

func (m *MagickWand) GetImageIndex() int32 { return GetImageIndex(m) }

var GetImageInterlaceScheme func(m *MagickWand) G.InterlaceType

func (m *MagickWand) ImageInterlaceScheme() G.InterlaceType { return GetImageInterlaceScheme(m) }

var GetImageIterations func(m *MagickWand) uint32

func (m *MagickWand) Iterations() uint32 { return GetImageIterations(m) }

var GetImageMatteColor func(m *MagickWand, matteColor *PixelWand) bool

func (m *MagickWand) MatteColor(matteColor *PixelWand) bool {
	return GetImageMatteColor(m, matteColor)
}

var GetImagePage func(m *MagickWand, width, height *uint32, x, y *int32) bool

func (m *MagickWand) ImagePage(width, height *uint32, x, y *int32) bool {
	return GetImagePage(m, width, height, x, y)
}

var GetImagePixels func(m *MagickWand, x, y int32, columns, rows uint32, map_ string, storage G.StorageType, pixels *G.Void) bool

func (m *MagickWand) Pixels(x, y int32, columns, rows uint32, map_ string, storage G.StorageType, pixels *G.Void) bool {
	return GetImagePixels(m, x, y, columns, rows, map_, storage, pixels)
}

var GetImageProfile func(m *MagickWand, name string, length *uint32) *byte

func (m *MagickWand) Profile(name string, length *uint32) *byte {
	return GetImageProfile(m, name, length)
}

var GetImageRedPrimary func(m *MagickWand, x, y *float64) bool

func (m *MagickWand) RedPrimary(x, y *float64) bool { return GetImageRedPrimary(m, x, y) }

var GetImageRenderingIntent func(m *MagickWand) G.RenderingIntent

func (m *MagickWand) RenderingIntent() G.RenderingIntent { return GetImageRenderingIntent(m) }

var GetImageResolution func(m *MagickWand, x, y *float64) bool

func (m *MagickWand) ImageResolution(x, y *float64) bool { return GetImageResolution(m, x, y) }

var GetImageScene func(m *MagickWand) uint32

func (m *MagickWand) ImageScene() uint32 { return GetImageScene(m) }

var GetImageSignature func(m *MagickWand) string

func (m *MagickWand) ImageSignature() string { return GetImageSignature(m) }

var GetImageSize func(m *MagickWand, length *G.MagickSizeType) bool

func (m *MagickWand) ImageSize(length *G.MagickSizeType) bool { return GetImageSize(m, length) }

var GetImageType func(m *MagickWand) G.ImageType

func (m *MagickWand) ImageType() G.ImageType { return GetImageType(m) }

var GetImageSavedType func(*MagickWand) G.ImageType

func (m *MagickWand) SavedType() G.ImageType { return GetImageSavedType(m) }

var GetImageUnits func(m *MagickWand) G.ResolutionType

func (m *MagickWand) ImageUnits() G.ResolutionType { return GetImageUnits(m) }

var GetImageVirtualPixelMethod func(m *MagickWand) G.VirtualPixelMethod

func (m *MagickWand) ImageVirtualPixelMethod() G.VirtualPixelMethod {
	return GetImageVirtualPixelMethod(m)
}

var GetImageWhitePoint func(m *MagickWand, x, y *float64) bool

func (m *MagickWand) ImageWhitePoint(x, y *float64) bool { return GetImageWhitePoint(m, x, y) }

var GetImageWidth func(m *MagickWand) uint32

func (m *MagickWand) ImageWidth() uint32 { return GetImageWidth(m) }

var GetNumberImages func(m *MagickWand) uint32

func (m *MagickWand) NumberImages() uint32 { return GetNumberImages(m) }

var GetPackageName func() string
var GetQuantumDepth func(depth *uint32) string
var GetReleaseDate func() string
var ResourceLimit func(r G.ResourceType) G.MagickSizeType
var GetSamplingFactors func(m *MagickWand, numberFactors *uint32) []float64

func (m *MagickWand) SamplingFactors(numberFactors *uint32) []float64 {
	return GetSamplingFactors(m, numberFactors)
}

var GetSize func(m *MagickWand, columns, rows *uint32) bool

func (m *MagickWand) Size(columns, rows *uint32) bool { return GetSize(m, columns, rows) }

var GetVersion func(version *uint32) string
var HaldClutImage func(m *MagickWand, haldWand *MagickWand) bool

func (m *MagickWand) HaldClut(haldWand *MagickWand) bool {
	return HaldClutImage(m, haldWand)
}

var HasNextImage func(m *MagickWand) bool

func (m *MagickWand) HasNext() bool { return HasNextImage(m) }

var HasPreviousImage func(m *MagickWand) bool

func (m *MagickWand) HasPrevious() bool { return HasPreviousImage(m) }

var ImplodeImage func(m *MagickWand, radius float64) bool

func (m *MagickWand) Implode(radius float64) bool { return ImplodeImage(m, radius) }

var LabelImage func(m *MagickWand, label string) bool

func (m *MagickWand) Label(label string) bool { return LabelImage(m, label) }

var LevelImage func(m *MagickWand, blackPoint, gamma, whitePoint float64) bool

func (m *MagickWand) Level(blackPoint, gamma, whitePoint float64) bool {
	return LevelImage(m, blackPoint, gamma, whitePoint)
}

var LevelImageChannel func(m *MagickWand, channel G.ChannelType, blackPoint, gamma, whitePoint float64) bool

func (m *MagickWand) LevelChannel(channel G.ChannelType, blackPoint, gamma, whitePoint float64) bool {
	return LevelImageChannel(m, channel, blackPoint, gamma, whitePoint)
}

var MagnifyImage func(m *MagickWand) bool

func (m *MagickWand) Magnify() bool { return MagnifyImage(m) }

var MapImage func(m *MagickWand, mapWand *MagickWand, dither bool) bool

func (m *MagickWand) Map(mapWand *MagickWand, dither bool) bool {
	return MapImage(m, mapWand, dither)
}

var MatteFloodfillImage func(m *MagickWand, alpha, fuzz float64, bordercolor *PixelWand, x, y int32) bool

func (m *MagickWand) MatteFloodfill(alpha, fuzz float64, bordercolor *PixelWand, x, y int32) bool {
	return MatteFloodfillImage(m, alpha, fuzz, bordercolor, x, y)
}

var MedianFilterImage func(m *MagickWand, radius float64) bool

func (m *MagickWand) MedianFilter(radius float64) bool {
	return MedianFilterImage(m, radius)
}

var MinifyImage func(m *MagickWand) bool

func (m *MagickWand) Minify() bool { return MinifyImage(m) }

var ModulateImage func(m *MagickWand, brightness, saturation, hue float64) bool

func (m *MagickWand) Modulate(brightness, saturation, hue float64) bool {
	return ModulateImage(m, brightness, saturation, hue)
}

type MontageMode G.Enum

const (
	UndefinedMode MontageMode = iota
	FrameMode
	UnframeMode
	ConcatenateMode
)

var MontageImage func(m *MagickWand, drawingWand DrawingWand, tileGeometry, thumbnailGeometry string, mode MontageMode, frame string) *MagickWand

func (m *MagickWand) Montage(drawingWand DrawingWand, tileGeometry, thumbnailGeometry string, mode MontageMode, frame string) *MagickWand {
	return MontageImage(m, drawingWand, tileGeometry, thumbnailGeometry, mode, frame)
}

var MorphImages func(m *MagickWand, numberFrames uint32) *MagickWand

func (m *MagickWand) Morph(numberFrames uint32) *MagickWand {
	return MorphImages(m, numberFrames)
}

var MosaicImages func(m *MagickWand) *MagickWand

func (m *MagickWand) Mosaic() *MagickWand { return MosaicImages(m) }

var MotionBlurImage func(m *MagickWand, radius, sigma, angle float64) bool

func (m *MagickWand) MotionBlur(radius, sigma, angle float64) bool {
	return MotionBlurImage(m, radius, sigma, angle)
}

var NegateImage func(m *MagickWand, gray bool) bool

func (m *MagickWand) Negate(gray bool) bool { return NegateImage(m, gray) }

var NegateImageChannel func(m *MagickWand, channel G.ChannelType, gray bool) bool

func (m *MagickWand) NegateChannel(channel G.ChannelType, gray bool) bool {
	return NegateImageChannel(m, channel, gray)
}

var NextImage func(m *MagickWand) bool

func (m *MagickWand) Next() bool { return NextImage(m) }

var NormalizeImage func(m *MagickWand) bool

func (m *MagickWand) Normalize() bool { return NormalizeImage(m) }

var OilPaintImage func(m *MagickWand, radius float64) bool

func (m *MagickWand) OilPaint(radius float64) bool { return OilPaintImage(m, radius) }

var OpaqueImage func(m *MagickWand, target, fill *PixelWand, fuzz float64) bool

func (m *MagickWand) Opaque(target, fill *PixelWand, fuzz float64) bool {
	return OpaqueImage(m, target, fill, fuzz)
}

var PingImage func(m *MagickWand, filename string) bool

func (m *MagickWand) Ping(filename string) bool { return PingImage(m, filename) }

var PreviewImages func(m *MagickWand, preview G.PreviewType) *MagickWand

func (m *MagickWand) PreviewImages(preview G.PreviewType) *MagickWand {
	return PreviewImages(m, preview)
}

var PreviousImage func(m *MagickWand) bool

func (m *MagickWand) Previous() bool { return PreviousImage(m) }

var ProfileImage func(m *MagickWand, name string, profile *G.Void, length uint32) bool

func (m *MagickWand) ProfileImage(name string, profile *G.Void, length uint32) bool {
	return ProfileImage(m, name, profile, length)
}

var QuantizeImage func(m *MagickWand, numberColors uint32, colorspace G.ColorspaceType, treedepth uint32, dither, measureError bool) bool

func (m *MagickWand) Quantize(numberColors uint32, colorspace G.ColorspaceType, treedepth uint32, dither, measureError bool) bool {
	return QuantizeImage(m, numberColors, colorspace, treedepth, dither, measureError)
}

var QuantizeImages func(m *MagickWand, numberColors uint32, colorspace G.ColorspaceType, treedepth uint32, dither, measureError bool) bool

func (m *MagickWand) QuantizeImages(numberColors uint32, colorspace G.ColorspaceType, treedepth uint32, dither, measureError bool) bool {
	return QuantizeImages(m, numberColors, colorspace, treedepth, dither, measureError)
}

var QueryFontMetrics func(m *MagickWand, drawingWand *DrawingWand, text string) []float64

func (m *MagickWand) QueryFontMetrics(drawingWand *DrawingWand, text string) []float64 {
	return QueryFontMetrics(m, drawingWand, text)
}

type StrSlice []string

func (StrSlice) SizeArg() int { return 2 }

var QueryFonts func(pattern string, numberFonts *uint32) StrSlice
var QueryFormats func(pattern string, numberFormats *uint32) StrSlice
var RadialBlurImage func(m *MagickWand, angle float64) bool

func (m *MagickWand) RadialBlur(angle float64) bool { return RadialBlurImage(m, angle) }

var RaiseImage func(m *MagickWand, width, height uint32, x, y int32, raise bool) bool

func (m *MagickWand) Raise(width, height uint32, x, y int32, raise bool) bool {
	return RaiseImage(m, width, height, x, y, raise)
}

var ReadImage func(m *MagickWand, filename string) bool

func (m *MagickWand) Read(filename string) bool { return ReadImage(m, filename) }

var ReadImageBlob func(m *MagickWand, blob *byte, length uint32) bool

func (m *MagickWand) ReadBlob(blob *byte, length uint32) bool {
	return ReadImageBlob(m, blob, length)
}

var ReadImageFile func(m *MagickWand, file *G.FILE) bool

func (m *MagickWand) ReadFile(file *G.FILE) bool { return ReadImageFile(m, file) }

var ReduceNoiseImage func(m *MagickWand, radius float64) bool

func (m *MagickWand) ReduceNoise(radius float64) bool { return ReduceNoiseImage(m, radius) }

var RelinquishMemory func(resource *G.Void) *G.Void

var RemoveImage func(m *MagickWand) bool

func (m *MagickWand) Remove() bool { return RemoveImage(m) }

var RemoveImageProfile func(m *MagickWand, name string, length *uint32) []byte

func (m *MagickWand) RemoveProfile(name string, length *uint32) []byte {
	return RemoveImageProfile(m, name, length)
}

var ResetIterator func(m *MagickWand)

func (m *MagickWand) ResetIterator() { ResetIterator(m) }

var ResampleImage func(m *MagickWand, xResolution, yResolution float64, filter G.FilterTypes, blur float64) bool

func (m *MagickWand) Resample(xResolution, yResolution float64, filter G.FilterTypes, blur float64) bool {
	return ResampleImage(m, xResolution, yResolution, filter, blur)
}

var ResizeImage func(m *MagickWand, columns, rows uint32, filter G.FilterTypes, blur float64) bool

func (m *MagickWand) Resize(columns, rows uint32, filter G.FilterTypes, blur float64) bool {
	return ResizeImage(m, columns, rows, filter, blur)
}

var RollImage func(m *MagickWand, x int32, y uint32) bool

func (m *MagickWand) Roll(x int32, y uint32) bool { return RollImage(m, x, y) }

var RotateImage func(m *MagickWand, background *PixelWand, degrees float64) bool

func (m *MagickWand) Rotate(background *PixelWand, degrees float64) bool {
	return RotateImage(m, background, degrees)
}

var SampleImage func(m *MagickWand, columns, rows uint32) bool

func (m *MagickWand) Sample(columns, rows uint32) bool {
	return SampleImage(m, columns, rows)
}

var ScaleImage func(m *MagickWand, columns, rows uint32) bool

func (m *MagickWand) Scale(columns, rows uint32) bool { return ScaleImage(m, columns, rows) }

var SeparateImageChannel func(m *MagickWand, channel G.ChannelType) bool

func (m *MagickWand) SeparateChannel(channel G.ChannelType) bool {
	return SeparateImageChannel(m, channel)
}

var SetCompressionQuality func(m *MagickWand, quality uint32) bool

func (m *MagickWand) SetCompressionQuality(quality uint32) bool {
	return SetCompressionQuality(m, quality)
}

var SetDepth func(m *MagickWand, depth uint32) bool

func (m *MagickWand) SetDepth(depth uint32) bool { return SetDepth(m, depth) }

var SetFilename func(m *MagickWand, filename string) bool

func (m *MagickWand) SetFilename(filename string) bool { return SetFilename(m, filename) }

var SetFormat func(m *MagickWand, format string) bool

func (m *MagickWand) SetFormat(format string) bool { return SetFormat(m, format) }

var SetImage func(m *MagickWand, setWand *MagickWand) bool

func (m *MagickWand) SetImage(setWand *MagickWand) bool { return SetImage(m, setWand) }

var SetImageAttribute func(m *MagickWand, property, value string) bool

func (m *MagickWand) SetAttribute(property, value string) bool {
	return SetImageAttribute(m, property, value)
}

var SetImageBackgroundColor func(m *MagickWand, background *PixelWand) bool

func (m *MagickWand) SetImageBackgroundColor(background *PixelWand) bool {
	return SetImageBackgroundColor(m, background)
}

var SetImageBluePrimary func(m *MagickWand, x, y float64) bool

func (m *MagickWand) SetBluePrimary(x, y float64) bool {
	return SetImageBluePrimary(m, x, y)
}

var SetImageBorderColor func(m *MagickWand, border *PixelWand) bool

func (m *MagickWand) SetBorderColor(border *PixelWand) bool {
	return SetImageBorderColor(m, border)
}

var SetImageColormapColor func(m *MagickWand, index uint32, color *PixelWand) bool

func (m *MagickWand) SetColormapColor(index uint32, color *PixelWand) bool {
	return SetImageColormapColor(m, index, color)
}

var SetImageColorspace func(m *MagickWand, colorspace G.ColorspaceType) bool

func (m *MagickWand) SetImageColorspace(colorspace G.ColorspaceType) bool {
	return SetImageColorspace(m, colorspace)
}

var SetImageCompose func(m *MagickWand, compose G.CompositeOperator) bool

func (m *MagickWand) SetCompose(compose G.CompositeOperator) bool {
	return SetImageCompose(m, compose)
}

var SetImageCompression func(m *MagickWand, compression G.CompressionType) bool

func (m *MagickWand) SetImageCompression(compression G.CompressionType) bool {
	return SetImageCompression(m, compression)
}

var SetImageDelay func(m *MagickWand, delay uint32) bool

func (m *MagickWand) SetDelay(delay uint32) bool { return SetImageDelay(m, delay) }

var SetImageChannelDepth func(m *MagickWand, channel G.ChannelType, depth uint32) bool

func (m *MagickWand) SetChannelDepth(channel G.ChannelType, depth uint32) bool {
	return SetImageChannelDepth(m, channel, depth)
}

var SetImageDepth func(m *MagickWand, depth uint32) bool

func (m *MagickWand) SetImageDepth(depth uint32) bool { return SetImageDepth(m, depth) }

var SetImageDispose func(m *MagickWand, dispose G.DisposeType) bool

func (m *MagickWand) SetDispose(dispose G.DisposeType) bool {
	return SetImageDispose(m, dispose)
}

var SetImageFilename func(m *MagickWand, filename string) bool

func (m *MagickWand) SetImageFilename(filename string) bool {
	return SetImageFilename(m, filename)
}

var SetImageFormat func(m *MagickWand, format string) bool

func (m *MagickWand) SetImageFormat(format string) bool { return SetImageFormat(m, format) }

var SetImageFuzz func(m *MagickWand, fuzz float64) bool

func (m *MagickWand) SetFuzz(fuzz float64) bool { return SetImageFuzz(m, fuzz) }

var SetImageGamma func(m *MagickWand, gamma float64) bool

func (m *MagickWand) SetGamma(gamma float64) bool { return SetImageGamma(m, gamma) }

var SetImageGreenPrimary func(m *MagickWand, x, y float64) bool

func (m *MagickWand) SetGreenPrimary(x, y float64) bool {
	return SetImageGreenPrimary(m, x, y)
}

var SetImageIndex func(m *MagickWand, index int32) bool

func (m *MagickWand) SetIndex(index int32) bool { return SetImageIndex(m, index) }

var SetImageInterlaceScheme func(m *MagickWand, interlace G.InterlaceType) bool

func (m *MagickWand) SetImageInterlaceScheme(interlace G.InterlaceType) bool {
	return SetImageInterlaceScheme(m, interlace)
}

var SetImageIterations func(m *MagickWand, iterations uint32) bool

func (m *MagickWand) SetIterations(iterations uint32) bool {
	return SetImageIterations(m, iterations)
}

var SetImageMatteColor func(m *MagickWand, matte *PixelWand) bool

func (m *MagickWand) SetMatteColor(matte *PixelWand) bool {
	return SetImageMatteColor(m, matte)
}

var SetImageOption func(m *MagickWand, format, key, value string) bool

func (m *MagickWand) SetImageOption(format, key, value string) bool {
	return SetImageOption(m, format, key, value)
}

var SetImagePage func(m *MagickWand, width, height uint32, x, y int32) bool

func (m *MagickWand) SetImagePage(width, height uint32, x, y int32) bool {
	return SetImagePage(m, width, height, x, y)
}

var SetImagePixels func(m *MagickWand, x, y int32, columns, rows uint32, map_ string, storage G.StorageType, pixels *G.Void) bool

func (m *MagickWand) SetImagePixels(x, y int32, columns, rows uint32, map_ string, storage G.StorageType, pixels *G.Void) bool {
	return SetImagePixels(m, x, y, columns, rows, map_, storage, pixels)
}

var SetImageProfile func(m *MagickWand, name string, profile *G.Void, length uint32) bool

func (m *MagickWand) SetImageProfile(name string, profile *G.Void, length uint32) bool {
	return SetImageProfile(m, name, profile, length)
}

var SetImageRedPrimary func(m *MagickWand, x, y float64) bool

func (m *MagickWand) SetRedPrimary(x, y float64) bool { return SetImageRedPrimary(m, x, y) }

var SetImageRenderingIntent func(m *MagickWand, renderingIntent G.RenderingIntent) bool

func (m *MagickWand) SetRenderingIntent(renderingIntent G.RenderingIntent) bool {
	return SetImageRenderingIntent(m, renderingIntent)
}

var SetImageResolution func(m *MagickWand, xResolution, yResolution float64) bool

func (m *MagickWand) SetImageResolution(xResolution, yResolution float64) bool {
	return SetResolution(m, xResolution, yResolution)
}

var SetImageScene func(m *MagickWand, scene uint32) bool

func (m *MagickWand) SetScene(scene uint32) bool { return SetImageScene(m, scene) }

var SetImageType func(m *MagickWand, imageType G.ImageType) bool

func (m *MagickWand) SetImageType(imageType G.ImageType) bool {
	return SetImageType(m, imageType)
}

var SetImageSavedType func(*MagickWand, G.ImageType) bool

func (m *MagickWand) SetSavedType(t G.ImageType) { SetImageSavedType(m, t) }

var SetImageUnits func(m *MagickWand, units G.ResolutionType) bool

func (m *MagickWand) SetUnits(units G.ResolutionType) bool { return SetImageUnits(m, units) }

var SetImageVirtualPixelMethod func(m *MagickWand, method G.VirtualPixelMethod) G.VirtualPixelMethod

func (m *MagickWand) SetVirtualPixelMethod(method G.VirtualPixelMethod) G.VirtualPixelMethod {
	return SetImageVirtualPixelMethod(m, method)
}

var SetInterlaceScheme func(m *MagickWand, interlaceScheme G.InterlaceType) bool

func (m *MagickWand) SetInterlaceScheme(interlaceScheme G.InterlaceType) bool {
	return SetInterlaceScheme(m, interlaceScheme)
}

var SetResolution func(m *MagickWand, xResolution, yResolution float64) bool

func (m *MagickWand) SetResolution(xResolution, yResolution float64) bool {
	return SetResolution(m, xResolution, yResolution)
}

var SetResolutionUnits func(m *MagickWand, units G.ResolutionType) uint

func (m *MagickWand) SetResolutionUnits(units G.ResolutionType) uint {
	return SetResolutionUnits(m, units)
}

var SetResourceLimit func(r G.ResourceType, limit G.MagickSizeType) bool

var SetSamplingFactors func(m *MagickWand, numberFactors uint32, samplingFactors *float64) bool

func (m *MagickWand) SetSamplingFactors(numberFactors uint32, samplingFactors *float64) bool {
	return SetSamplingFactors(m, numberFactors, samplingFactors)
}

var SetSize func(m *MagickWand, columns, rows uint32) bool

func (m *MagickWand) SetSize(columns, rows uint32) bool { return SetSize(m, columns, rows) }

var SetImageWhitePoint func(m *MagickWand, x, y float64) bool

func (m *MagickWand) SetImageWhitePoint(x, y float64) bool { return SetImageWhitePoint(m, x, y) }

var SetPassphrase func(m *MagickWand, passphrase string) bool

func (m *MagickWand) SetPassphrase(passphrase string) bool { return SetPassphrase(m, passphrase) }

var SharpenImage func(m *MagickWand, radius, sigma float64) bool

func (m *MagickWand) Sharpen(radius, sigma float64) bool {
	return SharpenImage(m, radius, sigma)
}

var ShaveImage func(m *MagickWand, columns, rows uint32) bool

func (m *MagickWand) Shave(columns, rows uint32) bool { return ShaveImage(m, columns, rows) }

var ShearImage func(m *MagickWand, background *PixelWand, xShear, yShear float64) bool

func (m *MagickWand) Shear(background *PixelWand, xShear, yShear float64) bool {
	return ShearImage(m, background, xShear, yShear)
}

var SolarizeImage func(m *MagickWand, threshold float64) bool

func (m *MagickWand) Solarize(threshold float64) bool { return SolarizeImage(m, threshold) }

var SpreadImage func(m *MagickWand, radius float64) bool

func (m *MagickWand) Spread(radius float64) bool { return SpreadImage(m, radius) }

var SteganoImage func(m *MagickWand, watermarkWand, offset int32) *MagickWand

func (m *MagickWand) Stegano(watermarkWand, offset int32) *MagickWand {
	return SteganoImage(m, watermarkWand, offset)
}

var StereoImage func(m *MagickWand, offsetWand *MagickWand) *MagickWand

func (m *MagickWand) Stereo(offsetWand *MagickWand) *MagickWand {
	return StereoImage(m, offsetWand)
}

var StripImage func(m *MagickWand) bool

func (m *MagickWand) Strip() bool { return StripImage(m) }

var SwirlImage func(m *MagickWand, degrees float64) bool

func (m *MagickWand) Swirl(degrees float64) bool { return SwirlImage(m, degrees) }

var TextureImage func(m *MagickWand, textureWand *MagickWand) *MagickWand

func (m *MagickWand) Texture(textureWand *MagickWand) *MagickWand {
	return TextureImage(m, textureWand)
}

var ThresholdImage func(m *MagickWand, threshold float64) bool

func (m *MagickWand) Threshold(threshold float64) bool {
	return ThresholdImage(m, threshold)
}

var ThresholdImageChannel func(m *MagickWand, channel G.ChannelType, threshold float64) bool

func (m *MagickWand) ThresholdChannel(channel G.ChannelType, threshold float64) bool {
	return ThresholdImageChannel(m, channel, threshold)
}

var TintImage func(m *MagickWand, tint, opacity *PixelWand) bool

func (m *MagickWand) Tint(tint, opacity *PixelWand) bool {
	return TintImage(m, tint, opacity)
}

var TransformImage func(m *MagickWand, crop, geometry string) *MagickWand

func (m *MagickWand) Transform(crop, geometry string) *MagickWand {
	return TransformImage(m, crop, geometry)
}

var TransparentImage func(m *MagickWand, target *PixelWand, alpha, fuzz float64) bool

func (m *MagickWand) Transparent(target *PixelWand, alpha, fuzz float64) bool {
	return TransparentImage(m, target, alpha, fuzz)
}

var TrimImage func(m *MagickWand, fuzz float64) bool

func (m *MagickWand) Trim(fuzz float64) bool { return TrimImage(m, fuzz) }

var UnsharpMaskImage func(m *MagickWand, radius, sigma, amount, threshold float64) bool

func (m *MagickWand) UnsharpMask(radius, sigma, amount, threshold float64) bool {
	return UnsharpMaskImage(m, radius, sigma, amount, threshold)
}

var WaveImage func(m *MagickWand, amplitude, waveLength float64) bool

func (m *MagickWand) Wave(amplitude, waveLength float64) bool {
	return WaveImage(m, amplitude, waveLength)
}

var WhiteThresholdImage func(m *MagickWand, threshold *PixelWand) bool

func (m *MagickWand) WhiteThreshold(threshold *PixelWand) bool {
	return WhiteThresholdImage(m, threshold)
}

var WriteImage func(m *MagickWand, filename string) bool

func (m *MagickWand) Write(filename string) bool { return WriteImage(m, filename) }

var WriteImagesFile func(m *MagickWand, file *G.FILE) bool

func (m *MagickWand) WriteImagesFile(file *G.FILE) bool { return WriteImagesFile(m, file) }

var WriteImageBlob func(m *MagickWand, length *uint32) *byte

func (m *MagickWand) WriteBlob(length *uint32) *byte {
	return WriteImageBlob(m, length)
}

var WriteImageFile func(m *MagickWand, file *G.FILE) bool

func (m *MagickWand) WriteFile(file *G.FILE) bool { return WriteImageFile(m, file) }

var WriteImages func(m *MagickWand, filename string, adjoin bool) bool

func (m *MagickWand) WriteImages(filename string, adjoin bool) bool {
	return WriteImages(m, filename, adjoin)
}

var NewMagickWand func() *MagickWand

// Pixel Wand

var ClonePixelWand func(p *PixelWand) *PixelWand

func (p *PixelWand) Clone() *PixelWand { return ClonePixelWand(p) }

var ClonePixelWands func(wands []*PixelWand, numberWands uint32) []*PixelWand

var DestroyPixelWand func(p *PixelWand) *PixelWand

func (p *PixelWand) Destroy() *PixelWand { return DestroyPixelWand(p) }

var NewPixelWand func() *PixelWand

var NewPixelWands func(numberWands uint32) []*PixelWand

var PixelGetException func(p *PixelWand, severity *G.ExceptionType) string

func (p *PixelWand) Exception(severity *G.ExceptionType) string { return PixelGetException(p, severity) }

var PixelGetBlack func(p *PixelWand) float64

func (p *PixelWand) Black() float64 { return PixelGetBlack(p) }

var PixelGetBlackQuantum func(p *PixelWand) G.Quantum

func (p *PixelWand) BlackQuantum() G.Quantum { return PixelGetBlackQuantum(p) }

var PixelGetBlue func(p *PixelWand) float64

func (p *PixelWand) Blue() float64 { return PixelGetBlue(p) }

var PixelGetBlueQuantum func(p *PixelWand) G.Quantum

func (p *PixelWand) BlueQuantum() G.Quantum { return PixelGetBlueQuantum(p) }

var PixelGetColorAsString func(p *PixelWand) string

func (p *PixelWand) ColorAsString() string { return PixelGetColorAsString(p) }

var PixelGetColorCount func(p *PixelWand) uint32

func (p *PixelWand) ColorCount() uint32 { return PixelGetColorCount(p) }

var PixelGetCyan func(p *PixelWand) float64

func (p *PixelWand) Cyan() float64 { return PixelGetCyan(p) }

var PixelGetCyanQuantum func(p *PixelWand) G.Quantum

func (p *PixelWand) CyanQuantum() G.Quantum { return PixelGetCyanQuantum(p) }

var PixelGetGreen func(p *PixelWand) float64

func (p *PixelWand) Green() float64 { return PixelGetGreen(p) }

var PixelGetGreenQuantum func(p *PixelWand) G.Quantum

func (p *PixelWand) GreenQuantum() G.Quantum { return PixelGetGreenQuantum(p) }

var PixelGetMagenta func(p *PixelWand) float64

func (p *PixelWand) Magenta() float64 { return PixelGetMagenta(p) }

var PixelGetMagentaQuantum func(p *PixelWand) G.Quantum

func (p *PixelWand) MagentaQuantum() G.Quantum { return PixelGetMagentaQuantum(p) }

var PixelGetOpacity func(p *PixelWand) float64

func (p *PixelWand) Opacity() float64 { return PixelGetOpacity(p) }

var PixelGetOpacityQuantum func(p *PixelWand) G.Quantum

func (p *PixelWand) OpacityQuantum() G.Quantum { return PixelGetOpacityQuantum(p) }

var PixelGetQuantumColor func(p *PixelWand, color *G.PixelPacket)

func (p *PixelWand) QuantumColor(color *G.PixelPacket) { PixelGetQuantumColor(p, color) }

var PixelGetRed func(p *PixelWand) float64

func (p *PixelWand) Red() float64 { return PixelGetRed(p) }

var PixelGetRedQuantum func(p *PixelWand) G.Quantum

func (p *PixelWand) RedQuantum() G.Quantum { return PixelGetRedQuantum(p) }

var PixelGetYellow func(p *PixelWand) float64

func (p *PixelWand) Yellow() float64 { return PixelGetYellow(p) }

var PixelGetYellowQuantum func(p *PixelWand) G.Quantum

func (p *PixelWand) YellowQuantum() G.Quantum { return PixelGetYellowQuantum(p) }

var PixelSetBlack func(p *PixelWand, black float64)

func (p *PixelWand) SetBlack(black float64) { PixelSetBlack(p, black) }

var PixelSetBlackQuantum func(p *PixelWand, black G.Quantum)

func (p *PixelWand) SetBlackQuantum(black G.Quantum) { PixelSetBlackQuantum(p, black) }

var PixelSetBlue func(p *PixelWand, blue float64)

func (p *PixelWand) SetBlue(blue float64) { PixelSetBlue(p, blue) }

var PixelSetBlueQuantum func(p *PixelWand, blue G.Quantum)

func (p *PixelWand) SetBlueQuantum(blue G.Quantum) { PixelSetBlueQuantum(p, blue) }

var PixelSetColor func(p *PixelWand, color string) bool

func (p *PixelWand) SetColor(color string) bool { return PixelSetColor(p, color) }

var PixelSetColorCount func(p *PixelWand, count uint32)

func (p *PixelWand) SetColorCount(count uint32) { PixelSetColorCount(p, count) }

var PixelSetCyan func(p *PixelWand, cyan float64)

func (p *PixelWand) SetCyan(cyan float64) { PixelSetCyan(p, cyan) }

var PixelSetCyanQuantum func(p *PixelWand, cyan G.Quantum)

func (p *PixelWand) SetCyanQuantum(cyan G.Quantum) { PixelSetCyanQuantum(p, cyan) }

var PixelSetGreen func(p *PixelWand, green float64)

func (p *PixelWand) SetGreen(green float64) { PixelSetGreen(p, green) }

var PixelSetGreenQuantum func(p *PixelWand, green G.Quantum)

func (p *PixelWand) SetGreenQuantum(green G.Quantum) { PixelSetGreenQuantum(p, green) }

var PixelSetMagenta func(p *PixelWand, magenta float64)

func (p *PixelWand) SetMagenta(magenta float64) { PixelSetMagenta(p, magenta) }

var PixelSetMagentaQuantum func(p *PixelWand, magenta G.Quantum)

func (p *PixelWand) SetMagentaQuantum(magenta G.Quantum) { PixelSetMagentaQuantum(p, magenta) }

var PixelSetOpacity func(p *PixelWand, opacity float64)

func (p *PixelWand) SetOpacity(opacity float64) { PixelSetOpacity(p, opacity) }

var PixelSetOpacityQuantum func(p *PixelWand, opacity G.Quantum)

func (p *PixelWand) SetOpacityQuantum(opacity G.Quantum) { PixelSetOpacityQuantum(p, opacity) }

var PixelSetQuantumColor func(p *PixelWand, color *G.PixelPacket)

func (p *PixelWand) SetQuantumColor(color *G.PixelPacket) { PixelSetQuantumColor(p, color) }

var PixelSetRed func(p *PixelWand, red float64)

func (p *PixelWand) SetRed(red float64) { PixelSetRed(p, red) }

var PixelSetRedQuantum func(p *PixelWand, red G.Quantum)

func (p *PixelWand) SetRedQuantum(red G.Quantum) { PixelSetRedQuantum(p, red) }

var PixelSetYellow func(p *PixelWand, yellow float64)

func (p *PixelWand) SetYellow(yellow float64) { PixelSetYellow(p, yellow) }

var PixelSetYellowQuantum func(p *PixelWand, yellow G.Quantum)

func (p *PixelWand) SetYellowQuantum(yellow G.Quantum) { PixelSetYellowQuantum(p, yellow) }

// Undocumented

var CopyMagickString func(*G.Char, string, uint32) uint32
var FormatMagickString func(*G.Char, uint32, string, ...VArg) int
var FormatMagickStringList func(char *G.Char, length uint32, string, operands VAList) int
var ConcatenateMagickString func(*G.Char, string, uint32) uint32
var ImportImagePixels func(image *G.Image, xOffset, yOffset G.SSize, columns, rows G.Size, map_ string, type_ G.StorageType, pixels *G.Void) uint
var ParseAbsoluteGeometry func(geometry string, region_info *G.RectangleInfo) uint

type GeometryInfo struct{ Rho, Sigma, Xi, Psi, Chi float64 }

var ParseGeometry func(geometry *string, geometryInfo *GeometryInfo) uint
var RelinquishMagickMemory func(*G.Void) *G.Void
var ResizeMagickMemory func(memory *G.Void, size uint32) *G.Void

type MagickPixelPacket struct {
	Colorspace                G.ColorspaceType
	Matte                     uint
	Red, Green, Blue, Opacity G.Quantum
	Index                     G.IndexPacket
}

var QueryMagickColor func(string, *MagickPixelPacket, *G.ExceptionInfo) uint

var DLL = "CORE_RL_wand_.dll"
var allApis = Apis{
	{"CloneMagickWand", &CloneMagickWand},
	{"ClonePixelWand", &ClonePixelWand},
	{"ClonePixelWands", &ClonePixelWands},
	{"CopyMagickString", &CopyMagickString},
	{"DestroyMagickWand", &DestroyMagickWand},
	{"DestroyPixelWand", &DestroyPixelWand},
	{"FormatMagickString", &FormatMagickString},
	{"FormatMagickStringList", &FormatMagickStringList},
	{"GMPrivateConcatenateMagickString", &ConcatenateMagickString},
	{"GMPrivateImportImagePixels", &ImportImagePixels},
	{"GMPrivateParseAbsoluteGeometry", &ParseAbsoluteGeometry},
	{"GMPrivateParseGeometry", &ParseGeometry},
	{"GMPrivateRelinquishMagickMemory", &RelinquishMagickMemory},
	{"GMPrivateResizeMagickMemory", &ResizeMagickMemory},
	{"MagickAdaptiveThresholdImage", &AdaptiveThresholdImage},
	{"MagickAddImage", &AddImage},
	{"MagickAddNoiseImage", &AddNoiseImage},
	{"MagickAffineTransformImage", &AffineTransformImage},
	{"MagickAnimateImages", &AnimateImages},
	{"MagickAnnotateImage", &AnnotateImage},
	{"MagickAppendImages", &AppendImages},
	{"MagickAverageImages", &AverageImages},
	{"MagickBlackThresholdImage", &BlackThresholdImage},
	{"MagickBlurImage", &BlurImage},
	{"MagickBorderImage", &BorderImage},
	{"MagickCdlImage", &CdlImage},
	{"MagickCharcoalImage", &CharcoalImage},
	{"MagickChopImage", &ChopImage},
	{"MagickClipImage", &ClipImage},
	{"MagickClipPathImage", &ClipPathImage},
	{"MagickCloneDrawingWand", &CloneDrawingWand},
	{"MagickCoalesceImages", &CoalesceImages},
	{"MagickColorFloodfillImage", &ColorFloodfillImage},
	{"MagickColorizeImage", &ColorizeImage},
	{"MagickCommentImage", &CommentImage},
	{"MagickCompareImageChannels", &CompareImageChannels},
	{"MagickCompareImages", &CompareImages},
	{"MagickCompositeImage", &CompositeImage},
	{"MagickContrastImage", &ContrastImage},
	{"MagickConvolveImage", &ConvolveImage},
	{"MagickCropImage", &CropImage},
	{"MagickCycleColormapImage", &CycleColormapImage},
	{"MagickDeconstructImages", &DeconstructImages},
	{"MagickDescribeImage", &DescribeImage},
	{"MagickDespeckleImage", &DespeckleImage},
	{"MagickDestroyDrawingWand", &DestroyDrawingWand},
	{"MagickDisplayImage", &DisplayImage},
	{"MagickDisplayImages", &DisplayImages},
	{"MagickDrawAffine", &DrawAffine},
	{"MagickDrawAllocateWand", &DrawAllocateWand},
	{"MagickDrawAnnotation", &DrawAnnotation},
	{"MagickDrawArc", &DrawArc},
	{"MagickDrawBezier", &DrawBezier},
	{"MagickDrawCircle", &DrawCircle},
	{"MagickDrawClearException", &DrawClearException},
	{"MagickDrawColor", &DrawColor},
	{"MagickDrawComment", &DrawComment},
	{"MagickDrawComposite", &DrawComposite},
	{"MagickDrawEllipse", &DrawEllipse},
	{"MagickDrawGetClipPath", &DrawGetClipPath},
	{"MagickDrawGetClipRule", &DrawGetClipRule},
	{"MagickDrawGetClipUnits", &DrawGetClipUnits},
	{"MagickDrawGetException", &DrawGetException},
	{"MagickDrawGetFillColor", &DrawGetFillColor},
	{"MagickDrawGetFillOpacity", &DrawGetFillOpacity},
	{"MagickDrawGetFillRule", &DrawGetFillRule},
	{"MagickDrawGetFont", &DrawGetFont},
	{"MagickDrawGetFontFamily", &DrawGetFontFamily},
	{"MagickDrawGetFontSize", &DrawGetFontSize},
	{"MagickDrawGetFontStretch", &DrawGetFontStretch},
	{"MagickDrawGetFontStyle", &DrawGetFontStyle},
	{"MagickDrawGetFontWeight", &DrawGetFontWeight},
	{"MagickDrawGetGravity", &DrawGetGravity},
	{"MagickDrawGetStrokeAntialias", &DrawGetStrokeAntialias},
	{"MagickDrawGetStrokeColor", &DrawGetStrokeColor},
	{"MagickDrawGetStrokeDashArray", &DrawGetStrokeDashArray},
	{"MagickDrawGetStrokeDashOffset", &DrawGetStrokeDashOffset},
	{"MagickDrawGetStrokeLineCap", &DrawGetStrokeLineCap},
	{"MagickDrawGetStrokeLineJoin", &DrawGetStrokeLineJoin},
	{"MagickDrawGetStrokeMiterLimit", &DrawGetStrokeMiterLimit},
	{"MagickDrawGetStrokeOpacity", &DrawGetStrokeOpacity},
	{"MagickDrawGetStrokeWidth", &DrawGetStrokeWidth},
	{"MagickDrawGetTextAntialias", &DrawGetTextAntialias},
	{"MagickDrawGetTextDecoration", &DrawGetTextDecoration},
	{"MagickDrawGetTextEncoding", &DrawGetTextEncoding},
	{"MagickDrawGetTextUnderColor", &DrawGetTextUnderColor}, // CHECK ^^^
	{"MagickDrawImage", &DrawImage},
	{"MagickDrawLine", &DrawLine},
	{"MagickDrawMatte", &DrawMatte},
	{"MagickDrawPathClose", &DrawPathClose},
	{"MagickDrawPathCurveToAbsolute", &DrawPathCurveToAbsolute},
	{"MagickDrawPathCurveToQuadraticBezierAbsolute", &DrawPathCurveToQuadraticBezierAbsolute},
	{"MagickDrawPathCurveToQuadraticBezierRelative", &DrawPathCurveToQuadraticBezierRelative},
	{"MagickDrawPathCurveToQuadraticBezierSmoothAbsolute", &DrawPathCurveToQuadraticBezierSmoothAbsolute},
	{"MagickDrawPathCurveToQuadraticBezierSmoothRelative", &DrawPathCurveToQuadraticBezierSmoothRelative},
	{"MagickDrawPathCurveToRelative", &DrawPathCurveToRelative},
	{"MagickDrawPathCurveToSmoothAbsolute", &DrawPathCurveToSmoothAbsolute},
	{"MagickDrawPathCurveToSmoothRelative", &DrawPathCurveToSmoothRelative},
	{"MagickDrawPathEllipticArcAbsolute", &DrawPathEllipticArcAbsolute},
	{"MagickDrawPathEllipticArcRelative", &DrawPathEllipticArcRelative},
	{"MagickDrawPathFinish", &DrawPathFinish},
	{"MagickDrawPathLineToAbsolute", &DrawPathLineToAbsolute},
	{"MagickDrawPathLineToHorizontalAbsolute", &DrawPathLineToHorizontalAbsolute},
	{"MagickDrawPathLineToHorizontalRelative", &DrawPathLineToHorizontalRelative},
	{"MagickDrawPathLineToRelative", &DrawPathLineToRelative},
	{"MagickDrawPathLineToVerticalAbsolute", &DrawPathLineToVerticalAbsolute},
	{"MagickDrawPathLineToVerticalRelative", &DrawPathLineToVerticalRelative},
	{"MagickDrawPathMoveToAbsolute", &DrawPathMoveToAbsolute},
	{"MagickDrawPathMoveToRelative", &DrawPathMoveToRelative},
	{"MagickDrawPathStart", &DrawPathStart},
	{"MagickDrawPeekGraphicContext", &DrawPeekGraphicContext},
	{"MagickDrawPoint", &DrawPoint},
	{"MagickDrawPolygon", &DrawPolygon},
	{"MagickDrawPolyline", &DrawPolyline},
	{"MagickDrawPopClipPath", &DrawPopClipPath},
	{"MagickDrawPopDefs", &DrawPopDefs},
	{"MagickDrawPopGraphicContext", &DrawPopGraphicContext},
	{"MagickDrawPopPattern", &DrawPopPattern},
	{"MagickDrawPushClipPath", &DrawPushClipPath},
	{"MagickDrawPushDefs", &DrawPushDefs},
	{"MagickDrawPushGraphicContext", &DrawPushGraphicContext},
	{"MagickDrawPushPattern", &DrawPushPattern},
	{"MagickDrawRectangle", &DrawRectangle},
	{"MagickDrawRender", &DrawRender},
	{"MagickDrawRotate", &DrawRotate},
	{"MagickDrawRoundRectangle", &DrawRoundRectangle},
	{"MagickDrawScale", &DrawScale},
	{"MagickDrawSetClipPath", &DrawSetClipPath},
	{"MagickDrawSetClipRule", &DrawSetClipRule},
	{"MagickDrawSetClipUnits", &DrawSetClipUnits},
	{"MagickDrawSetFillColor", &DrawSetFillColor},
	{"MagickDrawSetFillOpacity", &DrawSetFillOpacity},
	{"MagickDrawSetFillPatternURL", &DrawSetFillPatternURL},
	{"MagickDrawSetFillRule", &DrawSetFillRule},
	{"MagickDrawSetFont", &DrawSetFont},
	{"MagickDrawSetFontFamily", &DrawSetFontFamily},
	{"MagickDrawSetFontSize", &DrawSetFontSize},
	{"MagickDrawSetFontStretch", &DrawSetFontStretch},
	{"MagickDrawSetFontStyle", &DrawSetFontStyle},
	{"MagickDrawSetFontWeight", &DrawSetFontWeight},
	{"MagickDrawSetGravity", &DrawSetGravity},
	{"MagickDrawSetStrokeAntialias", &DrawSetStrokeAntialias},
	{"MagickDrawSetStrokeColor", &DrawSetStrokeColor},
	{"MagickDrawSetStrokeDashArray", &DrawSetStrokeDashArray},
	{"MagickDrawSetStrokeDashOffset", &DrawSetStrokeDashOffset},
	{"MagickDrawSetStrokeLineCap", &DrawSetStrokeLineCap},
	{"MagickDrawSetStrokeLineJoin", &DrawSetStrokeLineJoin},
	{"MagickDrawSetStrokeMiterLimit", &DrawSetStrokeMiterLimit},
	{"MagickDrawSetStrokeOpacity", &DrawSetStrokeOpacity},
	{"MagickDrawSetStrokePatternURL", &DrawSetStrokePatternURL},
	{"MagickDrawSetStrokeWidth", &DrawSetStrokeWidth},
	{"MagickDrawSetTextAntialias", &DrawSetTextAntialias},
	{"MagickDrawSetTextDecoration", &DrawSetTextDecoration},
	{"MagickDrawSetTextEncoding", &DrawSetTextEncoding},
	{"MagickDrawSetTextUnderColor", &DrawSetTextUnderColor},
	{"MagickDrawSetViewbox", &DrawSetViewbox},
	{"MagickDrawSkewX", &DrawSkewX},
	{"MagickDrawSkewY", &DrawSkewY},
	{"MagickDrawTranslate", &DrawTranslate},
	{"MagickEdgeImage", &EdgeImage},
	{"MagickEmbossImage", &EmbossImage},
	{"MagickEnhanceImage", &EnhanceImage},
	{"MagickEqualizeImage", &EqualizeImage},
	{"MagickExtentImage", &ExtentImage},
	{"MagickFlattenImages", &FlattenImages},
	{"MagickFlipImage", &FlipImage},
	{"MagickFlopImage", &FlopImage},
	{"MagickFrameImage", &FrameImage},
	{"MagickFxImage", &FxImage},
	{"MagickFxImageChannel", &FxImageChannel},
	{"MagickGammaImage", &GammaImage},
	{"MagickGammaImageChannel", &GammaImageChannel},
	{"MagickGetConfigureInfo", &GetConfigureInfo},
	{"MagickGetCopyright", &GetCopyright},
	{"MagickGetException", &GetException},
	{"MagickGetFilename", &GetFilename},
	{"MagickGetHomeURL", &GetHomeURL},
	{"MagickGetImage", &GetImage},
	{"MagickGetImageAttribute", &GetImageAttribute},
	{"MagickGetImageBackgroundColor", &GetImageBackgroundColor},
	{"MagickGetImageBluePrimary", &GetImageBluePrimary},
	{"MagickGetImageBorderColor", &GetImageBorderColor},
	{"MagickGetImageBoundingBox", &GetImageBoundingBox},
	{"MagickGetImageChannelDepth", &GetImageChannelDepth},
	{"MagickGetImageChannelExtrema", &GetImageChannelExtrema},
	{"MagickGetImageChannelMean", &GetImageChannelMean},
	{"MagickGetImageColormapColor", &GetImageColormapColor},
	{"MagickGetImageColors", &GetImageColors},
	{"MagickGetImageColorspace", &GetImageColorspace},
	{"MagickGetImageCompose", &GetImageCompose},
	{"MagickGetImageCompression", &GetImageCompression},
	{"MagickGetImageDelay", &GetImageDelay},
	{"MagickGetImageDepth", &GetImageDepth},
	{"MagickGetImageDispose", &GetImageDispose},
	{"MagickGetImageExtrema", &GetImageExtrema},
	{"MagickGetImageFilename", &GetImageFilename},
	{"MagickGetImageFormat", &GetImageFormat},
	{"MagickGetImageFuzz", &GetImageFuzz},
	{"MagickGetImageGamma", &GetImageGamma},
	{"MagickGetImageGreenPrimary", &GetImageGreenPrimary},
	{"MagickGetImageHeight", &GetImageHeight},
	{"MagickGetImageHistogram", &GetImageHistogram},
	{"MagickGetImageIndex", &GetImageIndex},
	{"MagickGetImageInterlaceScheme", &GetImageInterlaceScheme},
	{"MagickGetImageIterations", &GetImageIterations},
	{"MagickGetImageMatteColor", &GetImageMatteColor},
	{"MagickGetImagePage", &GetImagePage},
	{"MagickGetImagePixels", &GetImagePixels},
	{"MagickGetImageProfile", &GetImageProfile},
	{"MagickGetImageRedPrimary", &GetImageRedPrimary},
	{"MagickGetImageRenderingIntent", &GetImageRenderingIntent},
	{"MagickGetImageResolution", &GetImageResolution},
	{"MagickGetImageSavedType", &GetImageSavedType},
	{"MagickGetImageScene", &GetImageScene},
	{"MagickGetImageSignature", &GetImageSignature},
	{"MagickGetImageSize", &GetImageSize},
	{"MagickGetImageType", &GetImageType},
	{"MagickGetImageUnits", &GetImageUnits},
	{"MagickGetImageVirtualPixelMethod", &GetImageVirtualPixelMethod},
	{"MagickGetImageWhitePoint", &GetImageWhitePoint},
	{"MagickGetImageWidth", &GetImageWidth},
	{"MagickGetNumberImages", &GetNumberImages},
	{"MagickGetPackageName", &GetPackageName},
	{"MagickGetQuantumDepth", &GetQuantumDepth},
	{"MagickGetReleaseDate", &GetReleaseDate},
	{"MagickGetResourceLimit", &ResourceLimit},
	{"MagickGetSamplingFactors", &GetSamplingFactors},
	{"MagickGetSize", &GetSize},
	{"MagickGetVersion", &GetVersion},
	{"MagickHaldClutImage", &HaldClutImage},
	{"MagickHasNextImage", &HasNextImage},
	{"MagickHasPreviousImage", &HasPreviousImage},
	{"MagickImplodeImage", &ImplodeImage},
	{"MagickLabelImage", &LabelImage},
	{"MagickLevelImage", &LevelImage},
	{"MagickLevelImageChannel", &LevelImageChannel},
	{"MagickMagnifyImage", &MagnifyImage},
	{"MagickMapImage", &MapImage},
	{"MagickMatteFloodfillImage", &MatteFloodfillImage},
	{"MagickMedianFilterImage", &MedianFilterImage},
	{"MagickMinifyImage", &MinifyImage},
	{"MagickModulateImage", &ModulateImage},
	{"MagickMontageImage", &MontageImage},
	{"MagickMorphImages", &MorphImages},
	{"MagickMosaicImages", &MosaicImages},
	{"MagickMotionBlurImage", &MotionBlurImage},
	{"MagickNegateImage", &NegateImage},
	{"MagickNegateImageChannel", &NegateImageChannel},
	{"MagickNewDrawingWand", &NewDrawingWand}, // CHECk
	{"MagickNextImage", &NextImage},
	{"MagickNormalizeImage", &NormalizeImage},
	{"MagickOilPaintImage", &OilPaintImage},
	{"MagickOpaqueImage", &OpaqueImage},
	{"MagickPingImage", &PingImage},
	{"MagickPreviewImages", &PreviewImages},
	{"MagickPreviousImage", &PreviousImage},
	{"MagickProfileImage", &ProfileImage},
	{"MagickQuantizeImage", &QuantizeImage},
	{"MagickQuantizeImages", &QuantizeImages},
	{"MagickQueryFontMetrics", &QueryFontMetrics},
	{"MagickQueryFonts", &QueryFonts},
	{"MagickQueryFormats", &QueryFormats},
	{"MagickRadialBlurImage", &RadialBlurImage},
	{"MagickRaiseImage", &RaiseImage},
	{"MagickReadImage", &ReadImage},
	{"MagickReadImageBlob", &ReadImageBlob},
	{"MagickReadImageFile", &ReadImageFile},
	{"MagickReduceNoiseImage", &ReduceNoiseImage},
	{"MagickRelinquishMemory", &RelinquishMemory},
	{"MagickRemoveImage", &RemoveImage},
	{"MagickRemoveImageProfile", &RemoveImageProfile},
	{"MagickResampleImage", &ResampleImage},
	{"MagickResetIterator", &ResetIterator},
	{"MagickResizeImage", &ResizeImage},
	{"MagickRollImage", &RollImage},
	{"MagickRotateImage", &RotateImage},
	{"MagickSampleImage", &SampleImage},
	{"MagickScaleImage", &ScaleImage},
	{"MagickSeparateImageChannel", &SeparateImageChannel},
	{"MagickSetCompressionQuality", &SetCompressionQuality},
	{"MagickSetDepth", &SetDepth},
	{"MagickSetFilename", &SetFilename},
	{"MagickSetFormat", &SetFormat},
	{"MagickSetImage", &SetImage},
	{"MagickSetImageAttribute", &SetImageAttribute},
	{"MagickSetImageBackgroundColor", &SetImageBackgroundColor},
	{"MagickSetImageBluePrimary", &SetImageBluePrimary},
	{"MagickSetImageBorderColor", &SetImageBorderColor},
	{"MagickSetImageChannelDepth", &SetImageChannelDepth},
	{"MagickSetImageColormapColor", &SetImageColormapColor},
	{"MagickSetImageColorspace", &SetImageColorspace},
	{"MagickSetImageCompose", &SetImageCompose},
	{"MagickSetImageCompression", &SetImageCompression},
	{"MagickSetImageDelay", &SetImageDelay},
	{"MagickSetImageDepth", &SetImageDepth},
	{"MagickSetImageDispose", &SetImageDispose},
	{"MagickSetImageFilename", &SetImageFilename},
	{"MagickSetImageFormat", &SetImageFormat},
	{"MagickSetImageFuzz", &SetImageFuzz},
	{"MagickSetImageGamma", &SetImageGamma},
	{"MagickSetImageGreenPrimary", &SetImageGreenPrimary},
	{"MagickSetImageIndex", &SetImageIndex},
	{"MagickSetImageInterlaceScheme", &SetImageInterlaceScheme},
	{"MagickSetImageIterations", &SetImageIterations},
	{"MagickSetImageMatteColor", &SetImageMatteColor},
	{"MagickSetImageOption", &SetImageOption},
	{"MagickSetImagePage", &SetImagePage},
	{"MagickSetImagePixels", &SetImagePixels},
	{"MagickSetImageProfile", &SetImageProfile},
	{"MagickSetImageRedPrimary", &SetImageRedPrimary},
	{"MagickSetImageRenderingIntent", &SetImageRenderingIntent},
	{"MagickSetImageResolution", &SetImageResolution},
	{"MagickSetImageSavedType", &SetImageSavedType},
	{"MagickSetImageScene", &SetImageScene},
	{"MagickSetImageType", &SetImageType},
	{"MagickSetImageUnits", &SetImageUnits},
	{"MagickSetImageVirtualPixelMethod", &SetImageVirtualPixelMethod},
	{"MagickSetImageWhitePoint", &SetImageWhitePoint},
	{"MagickSetInterlaceScheme", &SetInterlaceScheme},
	{"MagickSetPassphrase", &SetPassphrase},
	{"MagickSetResolution", &SetResolution},
	{"MagickSetResolutionUnits", &SetResolutionUnits},
	{"MagickSetResourceLimit", &SetResourceLimit},
	{"MagickSetSamplingFactors", &SetSamplingFactors},
	{"MagickSetSize", &SetSize},
	{"MagickSharpenImage", &SharpenImage},
	{"MagickShaveImage", &ShaveImage},
	{"MagickShearImage", &ShearImage},
	{"MagickSolarizeImage", &SolarizeImage},
	{"MagickSpreadImage", &SpreadImage},
	{"MagickSteganoImage", &SteganoImage},
	{"MagickStereoImage", &StereoImage},
	{"MagickStripImage", &StripImage},
	{"MagickSwirlImage", &SwirlImage},
	{"MagickTextureImage", &TextureImage},
	{"MagickThresholdImage", &ThresholdImage},
	{"MagickThresholdImageChannel", &ThresholdImageChannel},
	{"MagickTintImage", &TintImage},
	{"MagickTransformImage", &TransformImage},
	{"MagickTransparentImage", &TransparentImage},
	{"MagickTrimImage", &TrimImage},
	{"MagickUnsharpMaskImage", &UnsharpMaskImage},
	{"MagickWaveImage", &WaveImage},
	{"MagickWhiteThresholdImage", &WhiteThresholdImage},
	{"MagickWriteImage", &WriteImage},
	{"MagickWriteImageBlob", &WriteImageBlob},
	{"MagickWriteImageFile", &WriteImageFile},
	{"MagickWriteImages", &WriteImages},
	{"MagickWriteImagesFile", &WriteImagesFile},
	{"NewMagickWand", &NewMagickWand},
	{"NewPixelWand", &NewPixelWand},
	{"NewPixelWands", &NewPixelWands},
	{"PixelGetBlack", &PixelGetBlack},
	{"PixelGetBlackQuantum", &PixelGetBlackQuantum},
	{"PixelGetBlue", &PixelGetBlue},
	{"PixelGetBlueQuantum", &PixelGetBlueQuantum},
	{"PixelGetColorAsString", &PixelGetColorAsString},
	{"PixelGetColorCount", &PixelGetColorCount},
	{"PixelGetCyan", &PixelGetCyan},
	{"PixelGetCyanQuantum", &PixelGetCyanQuantum},
	{"PixelGetException", &PixelGetException},
	{"PixelGetGreen", &PixelGetGreen},
	{"PixelGetGreenQuantum", &PixelGetGreenQuantum},
	{"PixelGetMagenta", &PixelGetMagenta},
	{"PixelGetMagentaQuantum", &PixelGetMagentaQuantum},
	{"PixelGetOpacity", &PixelGetOpacity},
	{"PixelGetOpacityQuantum", &PixelGetOpacityQuantum},
	{"PixelGetQuantumColor", &PixelGetQuantumColor},
	{"PixelGetRed", &PixelGetRed},
	{"PixelGetRedQuantum", &PixelGetRedQuantum},
	{"PixelGetYellow", &PixelGetYellow},
	{"PixelGetYellowQuantum", &PixelGetYellowQuantum},
	{"PixelSetBlack", &PixelSetBlack},
	{"PixelSetBlackQuantum", &PixelSetBlackQuantum},
	{"PixelSetBlue", &PixelSetBlue},
	{"PixelSetBlueQuantum", &PixelSetBlueQuantum},
	{"PixelSetColor", &PixelSetColor},
	{"PixelSetColorCount", &PixelSetColorCount},
	{"PixelSetCyan", &PixelSetCyan},
	{"PixelSetCyanQuantum", &PixelSetCyanQuantum},
	{"PixelSetGreen", &PixelSetGreen},
	{"PixelSetGreenQuantum", &PixelSetGreenQuantum},
	{"PixelSetMagenta", &PixelSetMagenta},
	{"PixelSetMagentaQuantum", &PixelSetMagentaQuantum},
	{"PixelSetOpacity", &PixelSetOpacity},
	{"PixelSetOpacityQuantum", &PixelSetOpacityQuantum},
	{"PixelSetQuantumColor", &PixelSetQuantumColor},
	{"PixelSetRed", &PixelSetRed},
	{"PixelSetRedQuantum", &PixelSetRedQuantum},
	{"PixelSetYellow", &PixelSetYellow},
	{"PixelSetYellowQuantum", &PixelSetYellowQuantum},
	{"QueryMagickColor", &QueryMagickColor},
}
