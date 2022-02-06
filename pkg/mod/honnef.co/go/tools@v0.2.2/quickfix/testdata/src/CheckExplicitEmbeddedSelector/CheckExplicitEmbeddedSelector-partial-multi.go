package pkg

type PartialMultiLevel struct{ F3 MultiLevel }
type PartialMultiLevel2Outer struct{ PartialMultiLevel2Inner }
type PartialMultiLevel2Inner struct{ F6 PartialMultiLevel2Outer2 }
type PartialMultiLevel2Outer2 struct{ PartialMultiLevel2Inner2 }
type PartialMultiLevel2Inner2 struct{ F7 int }

func fnPartialMulti() {
	var partialMulti PartialMultiLevel
	_ = partialMulti.F3.BasicOuter.F1            // want `could remove embedded field "BasicOuter" from selector`
	_ = partialMulti.F3.BasicOuter.BasicInner.F1 // want `could remove embedded field "BasicOuter" from selector` `could remove embedded field "BasicInner" from selector` `could simplify selectors`
	_ = partialMulti.F3.F1                       // minimal form

	var partialMulti2 PartialMultiLevel2Outer
	_ = partialMulti2.PartialMultiLevel2Inner.F6.PartialMultiLevel2Inner2.F7 // want `could remove embedded field "PartialMultiLevel2Inner2" from selector` `could remove embedded field "PartialMultiLevel2Inner" from selector` `could simplify selectors`
	_ = partialMulti2.F6.F7                                                  // minimal form
}
