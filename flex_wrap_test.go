package flex

import (
	"testing"
)

func TestWrap_column(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 30)
	NodeStyleSetHeight(rootChild0, 30)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 30)
	NodeStyleSetHeight(rootChild1, 30)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 30)
	NodeStyleSetHeight(rootChild2, 30)
	YGNodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild3, 30)
	NodeStyleSetHeight(rootChild3, 30)
	YGNodeInsertChild(root, rootChild3, 3)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 60, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild3))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 60, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 60, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild3))
}

func TestWrap_row(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 30)
	NodeStyleSetHeight(rootChild0, 30)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 30)
	NodeStyleSetHeight(rootChild1, 30)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 30)
	NodeStyleSetHeight(rootChild2, 30)
	YGNodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild3, 30)
	NodeStyleSetHeight(rootChild3, 30)
	YGNodeInsertChild(root, rootChild3, 3)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 60, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild3))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 70, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 70, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild3))
}

func TestWrap_row_align_items_flex_end(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetAlignItems(root, AlignFlexEnd)
	NodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 30)
	NodeStyleSetHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 30)
	NodeStyleSetHeight(rootChild1, 20)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 30)
	NodeStyleSetHeight(rootChild2, 30)
	YGNodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild3, 30)
	NodeStyleSetHeight(rootChild3, 30)
	YGNodeInsertChild(root, rootChild3, 3)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 60, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild3))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 70, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 70, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild3))
}

func TestWrap_row_align_items_center(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetAlignItems(root, AlignCenter)
	NodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 30)
	NodeStyleSetHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 30)
	NodeStyleSetHeight(rootChild1, 20)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 30)
	NodeStyleSetHeight(rootChild2, 30)
	YGNodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild3, 30)
	NodeStyleSetHeight(rootChild3, 30)
	YGNodeInsertChild(root, rootChild3, 3)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 5, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 60, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild3))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 60, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 70, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 5, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 70, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild3))
}

func TestFlex_wrap_children_with_min_main_overriding_flex_basis(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexBasis(rootChild0, 50)
	NodeStyleSetMinWidth(rootChild0, 55)
	NodeStyleSetHeight(rootChild0, 50)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetFlexBasis(rootChild1, 50)
	NodeStyleSetMinWidth(rootChild1, 55)
	NodeStyleSetHeight(rootChild1, 50)
	YGNodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 55, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 55, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild1))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 45, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 55, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 45, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 55, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild1))
}

func TestFlex_wrap_wrap_to_child_height(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(rootChild0, FlexDirectionRow)
	NodeStyleSetAlignItems(rootChild0, AlignFlexStart)
	NodeStyleSetFlexWrap(rootChild0, WrapWrap)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0Child0, 100)
	YGNodeInsertChild(rootChild0, rootChild0Child0, 0)

	rootChild0Child0_child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0Child0_child0, 100)
	NodeStyleSetHeight(rootChild0Child0_child0, 100)
	YGNodeInsertChild(rootChild0Child0, rootChild0Child0_child0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 100)
	NodeStyleSetHeight(rootChild1, 100)
	YGNodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0Child0_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0Child0_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild1))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0Child0_child0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0Child0_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild1))
}

func TestFlex_wrap_align_stretch_fits_one_row(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 150)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 50)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 50)
	YGNodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild1))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 100, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild1))
}

func TestWrap_reverse_row_align_content_flex_start(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetFlexWrap(root, WrapWrapReverse)
	NodeStyleSetWidth(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 30)
	NodeStyleSetHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 30)
	NodeStyleSetHeight(rootChild1, 20)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 30)
	NodeStyleSetHeight(rootChild2, 30)
	YGNodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild3, 30)
	NodeStyleSetHeight(rootChild3, 40)
	YGNodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild4, 30)
	NodeStyleSetHeight(rootChild4, 50)
	YGNodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 70, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 60, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 60, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild4))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 70, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 70, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 60, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 70, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild4))
}

func TestWrap_reverse_row_align_content_center(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetAlignContent(root, AlignCenter)
	NodeStyleSetFlexWrap(root, WrapWrapReverse)
	NodeStyleSetWidth(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 30)
	NodeStyleSetHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 30)
	NodeStyleSetHeight(rootChild1, 20)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 30)
	NodeStyleSetHeight(rootChild2, 30)
	YGNodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild3, 30)
	NodeStyleSetHeight(rootChild3, 40)
	YGNodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild4, 30)
	NodeStyleSetHeight(rootChild4, 50)
	YGNodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 70, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 60, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 60, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild4))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 70, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 70, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 60, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 70, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild4))
}

func TestWrap_reverse_row_single_line_different_size(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetFlexWrap(root, WrapWrapReverse)
	NodeStyleSetWidth(root, 300)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 30)
	NodeStyleSetHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 30)
	NodeStyleSetHeight(rootChild1, 20)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 30)
	NodeStyleSetHeight(rootChild2, 30)
	YGNodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild3, 30)
	NodeStyleSetHeight(rootChild3, 40)
	YGNodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild4, 30)
	NodeStyleSetHeight(rootChild4, 50)
	YGNodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 300, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 60, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 90, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 120, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild4))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 300, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 270, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 240, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 210, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 20, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 180, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 150, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild4))
}

func TestWrap_reverse_row_align_content_stretch(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetAlignContent(root, AlignStretch)
	NodeStyleSetFlexWrap(root, WrapWrapReverse)
	NodeStyleSetWidth(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 30)
	NodeStyleSetHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 30)
	NodeStyleSetHeight(rootChild1, 20)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 30)
	NodeStyleSetHeight(rootChild2, 30)
	YGNodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild3, 30)
	NodeStyleSetHeight(rootChild3, 40)
	YGNodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild4, 30)
	NodeStyleSetHeight(rootChild4, 50)
	YGNodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 70, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 60, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 60, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild4))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 70, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 70, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 60, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 70, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild4))
}

func TestWrap_reverse_row_align_content_space_around(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetAlignContent(root, AlignSpaceAround)
	NodeStyleSetFlexWrap(root, WrapWrapReverse)
	NodeStyleSetWidth(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 30)
	NodeStyleSetHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 30)
	NodeStyleSetHeight(rootChild1, 20)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 30)
	NodeStyleSetHeight(rootChild2, 30)
	YGNodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild3, 30)
	NodeStyleSetHeight(rootChild3, 40)
	YGNodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild4, 30)
	NodeStyleSetHeight(rootChild4, 50)
	YGNodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 70, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 60, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 60, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild4))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 70, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 70, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 60, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 10, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 70, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 40, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild4))
}

func TestWrap_reverse_column_fixed_size(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetAlignItems(root, AlignCenter)
	NodeStyleSetFlexWrap(root, WrapWrapReverse)
	NodeStyleSetWidth(root, 200)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 30)
	NodeStyleSetHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 30)
	NodeStyleSetHeight(rootChild1, 20)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 30)
	NodeStyleSetHeight(rootChild2, 30)
	YGNodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild3, 30)
	NodeStyleSetHeight(rootChild3, 40)
	YGNodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild4, 30)
	NodeStyleSetHeight(rootChild4, 50)
	YGNodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 170, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 170, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 170, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 170, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 60, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 140, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild4))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 10, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 30, YGNodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild3))
	assertFloatEqual(t, 60, YGNodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(rootChild4))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 30, YGNodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, YGNodeLayoutGetHeight(rootChild4))
}

func TestWrapped_row_within_align_items_center(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetAlignItems(root, AlignCenter)
	NodeStyleSetWidth(root, 200)
	NodeStyleSetHeight(root, 200)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(rootChild0, FlexDirectionRow)
	NodeStyleSetFlexWrap(rootChild0, WrapWrap)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0Child0, 150)
	NodeStyleSetHeight(rootChild0Child0, 80)
	YGNodeInsertChild(rootChild0, rootChild0Child0, 0)

	rootChild0child1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0child1, 80)
	NodeStyleSetHeight(rootChild0child1, 80)
	YGNodeInsertChild(rootChild0, rootChild0child1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 160, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(rootChild0child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetWidth(rootChild0child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(rootChild0child1))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 160, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 120, YGNodeLayoutGetLeft(rootChild0child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(rootChild0child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetWidth(rootChild0child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(rootChild0child1))
}

func TestWrapped_row_within_align_items_flex_start(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetAlignItems(root, AlignFlexStart)
	NodeStyleSetWidth(root, 200)
	NodeStyleSetHeight(root, 200)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(rootChild0, FlexDirectionRow)
	NodeStyleSetFlexWrap(rootChild0, WrapWrap)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0Child0, 150)
	NodeStyleSetHeight(rootChild0Child0, 80)
	YGNodeInsertChild(rootChild0, rootChild0Child0, 0)

	rootChild0child1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0child1, 80)
	NodeStyleSetHeight(rootChild0child1, 80)
	YGNodeInsertChild(rootChild0, rootChild0child1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 160, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(rootChild0child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetWidth(rootChild0child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(rootChild0child1))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 160, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 120, YGNodeLayoutGetLeft(rootChild0child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(rootChild0child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetWidth(rootChild0child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(rootChild0child1))
}

func TestWrapped_row_within_align_items_flex_end(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetAlignItems(root, AlignFlexEnd)
	NodeStyleSetWidth(root, 200)
	NodeStyleSetHeight(root, 200)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(rootChild0, FlexDirectionRow)
	NodeStyleSetFlexWrap(rootChild0, WrapWrap)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0Child0, 150)
	NodeStyleSetHeight(rootChild0Child0, 80)
	YGNodeInsertChild(rootChild0, rootChild0Child0, 0)

	rootChild0child1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0child1, 80)
	NodeStyleSetHeight(rootChild0child1, 80)
	YGNodeInsertChild(rootChild0, rootChild0child1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 160, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(rootChild0child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetWidth(rootChild0child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(rootChild0child1))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 160, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 150, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 120, YGNodeLayoutGetLeft(rootChild0child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetTop(rootChild0child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetWidth(rootChild0child1))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(rootChild0child1))
}

func TestWrapped_column_max_height(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetJustifyContent(root, JustifyCenter)
	NodeStyleSetAlignContent(root, AlignCenter)
	NodeStyleSetAlignItems(root, AlignCenter)
	NodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 700)
	NodeStyleSetHeight(root, 500)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 100)
	NodeStyleSetHeight(rootChild0, 500)
	NodeStyleSetMaxHeight(rootChild0, 200)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetMargin(rootChild1, EdgeLeft, 20)
	NodeStyleSetMargin(rootChild1, EdgeTop, 20)
	NodeStyleSetMargin(rootChild1, EdgeRight, 20)
	NodeStyleSetMargin(rootChild1, EdgeBottom, 20)
	NodeStyleSetWidth(rootChild1, 200)
	NodeStyleSetHeight(rootChild1, 200)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 100)
	NodeStyleSetHeight(rootChild2, 100)
	YGNodeInsertChild(root, rootChild2, 2)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 700, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 500, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 250, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 200, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 250, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 420, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 200, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild2))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 700, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 500, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 350, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 30, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 300, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 250, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 200, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 180, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 200, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild2))
}

func TestWrapped_column_max_height_flex(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetJustifyContent(root, JustifyCenter)
	NodeStyleSetAlignContent(root, AlignCenter)
	NodeStyleSetAlignItems(root, AlignCenter)
	NodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 700)
	NodeStyleSetHeight(root, 500)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetFlexShrink(rootChild0, 1)
	NodeStyleSetFlexBasisPercent(rootChild0, 0)
	NodeStyleSetWidth(rootChild0, 100)
	NodeStyleSetHeight(rootChild0, 500)
	NodeStyleSetMaxHeight(rootChild0, 200)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild1, 1)
	NodeStyleSetFlexShrink(rootChild1, 1)
	NodeStyleSetFlexBasisPercent(rootChild1, 0)
	NodeStyleSetMargin(rootChild1, EdgeLeft, 20)
	NodeStyleSetMargin(rootChild1, EdgeTop, 20)
	NodeStyleSetMargin(rootChild1, EdgeRight, 20)
	NodeStyleSetMargin(rootChild1, EdgeBottom, 20)
	NodeStyleSetWidth(rootChild1, 200)
	NodeStyleSetHeight(rootChild1, 200)
	YGNodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 100)
	NodeStyleSetHeight(rootChild2, 100)
	YGNodeInsertChild(root, rootChild2, 2)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 700, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 500, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 300, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 180, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 250, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 200, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 180, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 300, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 400, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild2))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 700, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 500, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 300, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 180, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 250, YGNodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 200, YGNodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 200, YGNodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 180, YGNodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 300, YGNodeLayoutGetLeft(rootChild2))
	assertFloatEqual(t, 400, YGNodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 100, YGNodeLayoutGetHeight(rootChild2))
}

func TestWrap_nodes_with_content_sizing_overflowing_margin(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 500)
	NodeStyleSetHeight(root, 500)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(rootChild0, FlexDirectionRow)
	NodeStyleSetFlexWrap(rootChild0, WrapWrap)
	NodeStyleSetWidth(rootChild0, 85)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	YGNodeInsertChild(rootChild0, rootChild0Child0, 0)

	rootChild0Child0_child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0Child0_child0, 40)
	NodeStyleSetHeight(rootChild0Child0_child0, 40)
	YGNodeInsertChild(rootChild0Child0, rootChild0Child0_child0, 0)

	rootChild0child1 := NewNodeWithConfig(config)
	NodeStyleSetMargin(rootChild0child1, EdgeRight, 10)
	YGNodeInsertChild(rootChild0, rootChild0child1, 1)

	rootChild0child1Child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0child1Child0, 40)
	NodeStyleSetHeight(rootChild0child1Child0, 40)
	YGNodeInsertChild(rootChild0child1, rootChild0child1Child0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 500, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 500, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 85, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(rootChild0Child0_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0Child0_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(rootChild0child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(rootChild0child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0child1Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0child1Child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(rootChild0child1Child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0child1Child0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 500, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 500, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 415, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 85, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 80, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 45, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(rootChild0Child0_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0Child0_child0))

	assertFloatEqual(t, 35, YGNodeLayoutGetLeft(rootChild0child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetTop(rootChild0child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(rootChild0child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0child1Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0child1Child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(rootChild0child1Child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0child1Child0))
}

func TestWrap_nodes_with_content_sizing_margin_cross(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 500)
	NodeStyleSetHeight(root, 500)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(rootChild0, FlexDirectionRow)
	NodeStyleSetFlexWrap(rootChild0, WrapWrap)
	NodeStyleSetWidth(rootChild0, 70)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	YGNodeInsertChild(rootChild0, rootChild0Child0, 0)

	rootChild0Child0_child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0Child0_child0, 40)
	NodeStyleSetHeight(rootChild0Child0_child0, 40)
	YGNodeInsertChild(rootChild0Child0, rootChild0Child0_child0, 0)

	rootChild0child1 := NewNodeWithConfig(config)
	NodeStyleSetMargin(rootChild0child1, EdgeTop, 10)
	YGNodeInsertChild(rootChild0, rootChild0child1, 1)

	rootChild0child1Child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0child1Child0, 40)
	NodeStyleSetHeight(rootChild0child1Child0, 40)
	YGNodeInsertChild(rootChild0child1, rootChild0child1Child0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 500, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 500, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 70, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 90, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(rootChild0Child0_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0Child0_child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild0child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(rootChild0child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0child1Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0child1Child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(rootChild0child1Child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0child1Child0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 500, YGNodeLayoutGetWidth(root))
	assertFloatEqual(t, 500, YGNodeLayoutGetHeight(root))

	assertFloatEqual(t, 430, YGNodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 70, YGNodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 90, YGNodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0Child0_child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0Child0_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(rootChild0Child0_child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0Child0_child0))

	assertFloatEqual(t, 30, YGNodeLayoutGetLeft(rootChild0child1))
	assertFloatEqual(t, 50, YGNodeLayoutGetTop(rootChild0child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(rootChild0child1))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0child1))

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(rootChild0child1Child0))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(rootChild0child1Child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetWidth(rootChild0child1Child0))
	assertFloatEqual(t, 40, YGNodeLayoutGetHeight(rootChild0child1Child0))
}
