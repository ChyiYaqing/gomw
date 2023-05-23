package base_demo

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	// 表格驱动测试
	type args struct {
		s   string
		sep string
	}
	tests := []struct {
		name       string
		args       args
		wantResult []string
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				s:   "chyi.ya.qing",
				sep: ".",
			},
			wantResult: []string{"chyi", "ya", "qing"}, // 期望的结果
		},
	}
	for _, tt := range tests {
		tt := tt // 这里重新声明tt变量，避免多个goroutine使用相同的变量
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()                                                                                  // 将每个测试用例标记为能够彼此并行运行
			if gotResult := Split(tt.args.s, tt.args.sep); !reflect.DeepEqual(gotResult, tt.wantResult) { // 借助反射包中的方法比较
				t.Errorf("Split() = %v, want %v", gotResult, tt.wantResult) // 测试失败输出错误提示
			}
		})
	}
}

func TestSplitWithComplexSep(t *testing.T) {
	got := Split("abcd", "bc")
	want := []string{"a", "d"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

func TestTimeConsuming(t *testing.T) {
	if testing.Short() {
		t.Skip("short 模式下会跳过改测试用例")
	}
}

/**

➜ go test ./...
PASS
ok      github.com/chyiyaqing/gomw/examples/base_demo   0.116s


➜ go test -v ./...	# 输出完整的测试结果
=== RUN   TestSplit
=== RUN   TestSplit/test1
--- PASS: TestSplit (0.00s)
    --- PASS: TestSplit/test1 (0.00s)
=== RUN   TestSplitWithComplexSep
--- PASS: TestSplitWithComplexSep (0.00s)
PASS
ok      github.com/chyiyaqing/gomw/examples/base_demo   0.107s

➜ go test -run=Sep -v ./... # -run 参数，对应一个正则表达式,只有函数名匹配的测试函数才会被go test命令执行 只运行TestSplitWithComplexSep
=== RUN   TestSplitWithComplexSep
--- PASS: TestSplitWithComplexSep (0.00s)
PASS
ok      github.com/chyiyaqing/gomw/examples/base_demo   0.313s

➜ go test --short -v ./...
=== RUN   TestSplit
=== RUN   TestSplit/test1
--- PASS: TestSplit (0.00s)
    --- PASS: TestSplit/test1 (0.00s)
=== RUN   TestSplitWithComplexSep
--- PASS: TestSplitWithComplexSep (0.00s)
=== RUN   TestTimeConsuming
    split_test.go:47: short 模式下会跳过改测试用例
--- SKIP: TestTimeConsuming (0.00s)
PASS
ok      github.com/chyiyaqing/gomw/examples/base_demo   0.103s



*/
