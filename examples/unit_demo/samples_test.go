package unit_demo

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert" // Testify 断言库，主要在提供断言功能之外，提供mock功能
)

func TestHello(t *testing.T) {
	// 得到的值与预期的值不相同
	got := Hello()
	want := "Hello, world"

	// 在测试代码中直接加上main()的调用
	main()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestAdd(t *testing.T) {
	sum := Add(5, 5)
	if sum == 10 {
		t.Log("the result is ok")
	} else {
		t.Fatal("the result is wrong")
	}
}

func TestCheckUrl(t *testing.T) {
	// 定义测试用例名称，
	convey.Convey("TestCheckTeachUrl true", t, func() {
		ok := CheckUrl("learnku.com")
		convey.So(ok, convey.ShouldBeTrue)
	})

	convey.Convey("TestCheckTeachUrl false", t, func() {
		ok := CheckUrl("xxxxxxxx.com")
		convey.So(ok, convey.ShouldBeFalse)
	})

	convey.Convey("TestCheckTeachUrl null", t, func() {
		ok := CheckUrl("")
		convey.So(ok, convey.ShouldBeFalse)
	})

	// 测试用例嵌套执行,可以更好的将测试用例组织起来.
	// 外层在套一个convey,需要传t指针，里面的covey都不需要t指针
	convey.Convey("TestCheckTeachUrl", t, func() {
		convey.Convey("TestCheckTeachUrl true", func() {
			ok := CheckUrl("learnku.com")
			convey.So(ok, convey.ShouldBeTrue)
		})
	})
}

func TestCheckUtl2(t *testing.T) {
	ok := CheckUrl("learnku.com")
	assert.True(t, ok)
}
