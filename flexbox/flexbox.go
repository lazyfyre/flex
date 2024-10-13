package flexbox

import "github.com/lazyfyre/flex"

type Prop = any
type Component = []Prop

type Id struct {
	Value string
}

type AlignContent struct {
	Value flex.Align
}

type AlignItems struct {
	Value flex.Align
}

type AlignSelf struct {
	Value flex.Align
}

type Flex struct {
	Value float64
}

type FlexBasis struct {
	Value float64
}

type FlexBasisPercent struct {
	Value float64
}

type FlexDirection struct {
	Value flex.FlexDirection
}

type FlexGrow struct {
	Value float64
}

type FlexShrink struct {
	Value float64
}

type FlexWrap struct {
	Value flex.Wrap
}

type JustifyContent struct {
	Value flex.Justify
}

type Height struct {
	Value float64
}

type HeightPercent struct {
	Value float64
}

type HeightAuto struct {
}

type Width struct {
	Value float64
}

type WidthPercent struct {
	Value float64
}

type WidthAuto struct {
}

type MaxHeight struct {
	Value float64
}

type MaxHeightPercent struct {
	Value float64
}

type MaxWidth struct {
	Value float64
}

type MaxWidthPercent struct {
	Value float64
}

type MinHeight struct {
	Value float64
}

type MinHeightPercent struct {
	Value float64
}

type MinWidth struct {
	Value float64
}

type MinWidthPercent struct {
	Value float64
}

type MarginTop struct {
	Value float64
}

type MarginTopPercent struct {
	Value float64
}

type MarginRight struct {
	Value float64
}

type MarginRightPercent struct {
	Value float64
}

type MarginBottom struct {
	Value float64
}

type MarginBottomPercent struct {
	Value float64
}
type MarginLeft struct {
	Value float64
}

type MarginLeftPercent struct {
	Value float64
}

type MarginUniform struct {
	Value float64
}

type MarginUniformPercent struct {
	Value float64
}

type PaddingTop struct {
	Value float64
}

type PaddingTopPercent struct {
	Value float64
}

type PaddingRight struct {
	Value float64
}

type PaddingRightPercent struct {
	Value float64
}

type PaddingBottom struct {
	Value float64
}

type PaddingBottomPercent struct {
	Value float64
}

type PaddingLeft struct {
	Value float64
}

type PaddingLeftPercent struct {
	Value float64
}

type PaddingUniform struct {
	Value float64
}

type PaddingUniformPercent struct {
	Value float64
}

type BorderWidthLeft struct {
	Value float64
}

type BorderWidthRight struct {
	Value float64
}

type BorderWidthTop struct {
	Value float64
}

type BorderWidthBottom struct {
	Value float64
}

type BorderWidthUniform struct {
	Value float64
}

type AspectRatio struct {
	Value float64
}

type Display struct {
	Value flex.Display
}

type Overflow struct {
	Value flex.Overflow
}

type Position struct {
	Value flex.PositionType
}

type Top struct {
	Value float64
}

type TopPercent struct {
	Value float64
}

type Right struct {
	Value float64
}

type RightPercent struct {
	Value float64
}

type Bottom struct {
	Value float64
}

type BottomPercent struct {
	Value float64
}

type Left struct {
	Value float64
}

type LeftPercent struct {
	Value float64
}

type Children struct {
	Value []Component
}

func Flexbox(props Component) *flex.Node {
	config := flex.NewConfig()
	node := flex.NewNodeWithConfig(config)
	node.Props = props
	for _, prop := range props {
		switch p := prop.(type) {
		case AlignContent:
			node.StyleSetAlignContent(p.Value)
		case AlignItems:
			node.StyleSetAlignItems(p.Value)
		case AlignSelf:

			node.StyleSetAlignSelf(p.Value)
		case Flex:
			node.StyleSetFlex(p.Value)
		case FlexBasis:
			node.StyleSetFlexBasis(p.Value)
		case FlexBasisPercent:
			node.StyleSetFlexBasisPercent(p.Value)
		case FlexDirection:

			node.StyleSetFlexDirection(p.Value)
		case FlexGrow:

			node.StyleSetFlexGrow(p.Value)
		case FlexShrink:
			node.StyleSetFlexShrink(p.Value)
		case FlexWrap:
			node.StyleSetFlexWrap(p.Value)
		case JustifyContent:
			node.StyleSetJustifyContent(p.Value)
		case Height:
			node.StyleSetHeight(p.Value)
		case HeightPercent:
			node.StyleSetHeightPercent(p.Value)
		case HeightAuto:
			node.StyleSetHeightAuto()
		case Width:
			node.StyleSetWidth(p.Value)
		case WidthPercent:
			node.StyleSetWidthPercent(p.Value)
		case WidthAuto:
			node.StyleSetWidthAuto()
		case MaxHeight:
			node.StyleSetMaxHeight(p.Value)
		case MaxHeightPercent:
			node.StyleSetMaxHeightPercent(p.Value)
		case MaxWidth:
			node.StyleSetMaxWidth(p.Value)
		case MaxWidthPercent:
			node.StyleSetMaxWidthPercent(p.Value)
		case MinHeight:
			node.StyleSetMinHeight(p.Value)
		case MinHeightPercent:
			node.StyleSetMinHeightPercent(p.Value)
		case MinWidth:
			node.StyleSetMinWidth(p.Value)
		case MinWidthPercent:
			node.StyleSetMinWidthPercent(p.Value)
		case MarginTop:
			node.StyleSetMargin(flex.EdgeTop, p.Value)
		case MarginTopPercent:
			node.StyleSetMarginPercent(flex.EdgeTop, p.Value)
		case MarginRight:
			node.StyleSetMargin(flex.EdgeRight, p.Value)
		case MarginRightPercent:
			node.StyleSetMarginPercent(flex.EdgeRight, p.Value)
		case MarginBottom:
			node.StyleSetMargin(flex.EdgeBottom, p.Value)
		case MarginBottomPercent:
			node.StyleSetMarginPercent(flex.EdgeBottom, p.Value)
		case MarginLeft:
			node.StyleSetMargin(flex.EdgeLeft, p.Value)
		case MarginLeftPercent:
			node.StyleSetMarginPercent(flex.EdgeLeft, p.Value)
		case MarginUniform:
			node.StyleSetMargin(flex.EdgeAll, p.Value)
		case MarginUniformPercent:
			node.StyleSetMarginPercent(flex.EdgeAll, p.Value)
		case PaddingTop:
			node.StyleSetPadding(flex.EdgeTop, p.Value)
		case PaddingTopPercent:
			node.StyleSetPaddingPercent(flex.EdgeTop, p.Value)
		case PaddingRight:
			node.StyleSetPadding(flex.EdgeRight, p.Value)
		case PaddingRightPercent:
			node.StyleSetPaddingPercent(flex.EdgeRight, p.Value)
		case PaddingBottom:
			node.StyleSetPadding(flex.EdgeBottom, p.Value)
		case PaddingBottomPercent:
			node.StyleSetPaddingPercent(flex.EdgeBottom, p.Value)
		case PaddingLeft:
			node.StyleSetPadding(flex.EdgeLeft, p.Value)
		case PaddingLeftPercent:
			node.StyleSetPaddingPercent(flex.EdgeLeft, p.Value)
		case PaddingUniform:
			node.StyleSetPadding(flex.EdgeAll, p.Value)
		case PaddingUniformPercent:
			node.StyleSetPaddingPercent(flex.EdgeAll, p.Value)
		case BorderWidthLeft:
			node.StyleSetBorder(flex.EdgeLeft, p.Value)
		case BorderWidthRight:
			node.StyleSetBorder(flex.EdgeRight, p.Value)
		case BorderWidthTop:
			node.StyleSetBorder(flex.EdgeTop, p.Value)
		case BorderWidthBottom:
			node.StyleSetBorder(flex.EdgeBottom, p.Value)
		case BorderWidthUniform:
			node.StyleSetBorder(flex.EdgeAll, p.Value)
		case AspectRatio:
			node.StyleSetAspectRatio(p.Value)
		case Display:
			node.StyleSetDisplay(p.Value)
		case Overflow:

			node.StyleSetOverflow(p.Value)
		case Position:

			node.StyleSetPositionType(p.Value)
		case Top:
			node.StyleSetPosition(flex.EdgeTop, p.Value)
		case TopPercent:
			node.StyleSetPositionPercent(flex.EdgeTop, p.Value)
		case Right:
			node.StyleSetPosition(flex.EdgeRight, p.Value)
		case RightPercent:
			node.StyleSetPositionPercent(flex.EdgeRight, p.Value)
		case Bottom:
			node.StyleSetPosition(flex.EdgeBottom, p.Value)
		case BottomPercent:
			node.StyleSetPositionPercent(flex.EdgeBottom, p.Value)
		case Left:
			node.StyleSetPosition(flex.EdgeLeft, p.Value)
		case LeftPercent:
			node.StyleSetPositionPercent(flex.EdgeLeft, p.Value)

		case Children:
			for i, child := range p.Value {
				node.InsertChild(Flexbox(child), i)
			}
		}
	}
	return node
}

type LayoutInterface struct {
	Left          float64
	Right         float64
	Top           float64
	Bottom        float64
	Height        float64
	Width         float64
	BorderTop     float64
	BorderRight   float64
	BorderBottom  float64
	BorderLeft    float64
	PaddingTop    float64
	PaddingRight  float64
	PaddingBottom float64
	PaddingLeft   float64
	MarginTop     float64
	MarginRight   float64
	MarginBottom  float64
	MarginLeft    float64
	Children      []LayoutInterface
	Props         []Prop
}

func Layout(node *flex.Node) LayoutInterface {
	childrenLayouts := make([]LayoutInterface, len(node.Children))
	for i, child := range node.Children {
		childrenLayouts[i] = Layout(child)
	}

	return LayoutInterface{
		Left:          node.LayoutGetLeft(),
		Right:         node.LayoutGetLeft(),
		Top:           node.LayoutGetTop(),
		Bottom:        node.LayoutGetBottom(),
		Height:        node.LayoutGetHeight(),
		Width:         node.LayoutGetWidth(),
		BorderTop:     node.LayoutGetBorder(flex.EdgeTop),
		BorderRight:   node.LayoutGetBorder(flex.EdgeRight),
		BorderBottom:  node.LayoutGetBorder(flex.EdgeBottom),
		BorderLeft:    node.LayoutGetBorder(flex.EdgeLeft),
		PaddingTop:    node.LayoutGetPadding(flex.EdgeTop),
		PaddingRight:  node.LayoutGetPadding(flex.EdgeRight),
		PaddingBottom: node.LayoutGetPadding(flex.EdgeBottom),
		PaddingLeft:   node.LayoutGetPadding(flex.EdgeLeft),
		MarginTop:     node.LayoutGetMargin(flex.EdgeTop),
		MarginRight:   node.LayoutGetMargin(flex.EdgeRight),
		MarginBottom:  node.LayoutGetMargin(flex.EdgeBottom),
		MarginLeft:    node.LayoutGetMargin(flex.EdgeLeft),
		Children:      childrenLayouts,
		Props:         node.Props,
	}
}

func NewFlexbox(width, height float64, props Component) LayoutInterface {
	root := Flexbox(props)
	if width == -1 {
		width = flex.NAN
	}
	if height == -1 {
		height = flex.NAN
	}
	flex.CalculateLayout(root, width, height, flex.DirectionLTR)
	return Layout(root)
}
