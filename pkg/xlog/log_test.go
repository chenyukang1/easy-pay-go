package xlog

import "testing"

func TestDebug(t *testing.T) {
	type args struct {
		args []any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testDebug",
			args: args{args: []any{"Hello,", "It's a debug"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetLevel(DebugLevel)
			Debug(tt.args.args...)
		})
	}
}

func TestError(t *testing.T) {
	type args struct {
		args []any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testError",
			args: args{args: []any{"Hello,", "It's a error"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Error(tt.args.args...)
		})
	}
}

func TestInfo(t *testing.T) {
	type args struct {
		args []any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testInfo",
			args: args{args: []any{"Hello,", "It's a info"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Info(tt.args.args...)
		})
	}
}

func TestWarn(t *testing.T) {
	type args struct {
		args []any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testWarn",
			args: args{args: []any{"Hello,", "It's a warn"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Warn(tt.args.args...)
		})
	}
}

func TestDebugf(t *testing.T) {
	type args struct {
		fmt  string
		args []any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testDebugf",
			args: args{
				fmt:  "Hello, It's a %s",
				args: []any{"debug"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetLevel(DebugLevel)
			Debugf(tt.args.fmt, tt.args.args...)
		})
	}
}

func TestInfof(t *testing.T) {
	type args struct {
		fmt  string
		args []any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testInfof",
			args: args{
				fmt:  "Hello, It's a %s",
				args: []any{"info"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Infof(tt.args.fmt, tt.args.args...)
		})
	}
}

func TestWarnf(t *testing.T) {
	type args struct {
		fmt  string
		args []any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testWarnf",
			args: args{
				fmt:  "Hello, It's a %s",
				args: []any{"warn"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Warnf(tt.args.fmt, tt.args.args...)
		})
	}
}

func TestErrorf(t *testing.T) {
	type args struct {
		fmt  string
		args []any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testErrorf",
			args: args{
				fmt:  "Hello, It's a %s",
				args: []any{"error"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Errorf(tt.args.fmt, tt.args.args...)
		})
	}
}
