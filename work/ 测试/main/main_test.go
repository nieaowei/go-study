/*******************************************************
 *  File        :   main_test.go.go
 *  Author      :   nieaowei
 *  Date        :   2019/11/19 11:47 上午
 *  Notes       :
 *******************************************************/
package main

import "testing"

func TestJudgetDate(t *testing.T) {
	type args struct {
		year   int
		month  int
		date   int
		option int
		mode   int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		want1   int
		want2   int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "覆盖测试1",
			args: args{
				year:   2019,
				month:  12,
				date:   31,
				option: 1,
				mode:   1,
			},
			want:    2019,
			want1:   12,
			want2:   31,
			wantErr: false,
		},
		{
			name: "覆盖测试2",
			args: args{
				year:   2019,
				month:  12,
				date:   31,
				option: 2,
				mode:   1,
			},
			want:    2019,
			want1:   12,
			want2:   31,
			wantErr: true,
		},
		{
			name: "覆盖测试3",
			args: args{
				year:   2019,
				month:  12,
				date:   31,
				option: 3,
				mode:   1,
			},
			want:    2019,
			want1:   12,
			want2:   31,
			wantErr: false,
		},
		{
			name: "覆盖测试4",
			args: args{
				year:   2019,
				month:  12,
				date:   31,
				option: 4,
				mode:   1,
			},
			want:    2019,
			want1:   12,
			want2:   31,
			wantErr: true,
		},
		{
			name: "年份超限错误",
			args: args{
				year:   2111,
				month:  12,
				date:   31,
				option: 1,
				mode:   0,
			},
			want:    -1,
			want1:   12,
			want2:   31,
			wantErr: false,
		},
		{
			name: "月份上限错误测试",
			args: args{
				year:   2019,
				month:  13,
				date:   31,
				option: 1,
				mode:   0,
			},
			want:    2019,
			want1:   13,
			want2:   31,
			wantErr: true,
		},
		{
			name: "月份下限错误测试",
			args: args{
				year:   2019,
				month:  0,
				date:   31,
				option: 1,
				mode:   0,
			},
			want:    2019,
			want1:   0,
			want2:   31,
			wantErr: true,
		},
		{
			name: "12月日期上限错误测试",
			args: args{
				year:   2019,
				month:  12,
				date:   32,
				option: 1,
				mode:   0,
			},
			want:    2019,
			want1:   12,
			want2:   32,
			wantErr: true,
		},
		{
			name: "12月日期下限错误测试",
			args: args{
				year:   2019,
				month:  12,
				date:   0,
				option: 1,
				mode:   0,
			},
			want:    2019,
			want1:   12,
			want2:   0,
			wantErr: true,
		},
		{
			name: "11月日期上限错误测试",
			args: args{
				year:   2019,
				month:  11,
				date:   31,
				option: 1,
				mode:   0,
			},
			want:    2019,
			want1:   11,
			want2:   31,
			wantErr: true,
		},
		{
			name: "11月日期下限错误测试",
			args: args{
				year:   2019,
				month:  11,
				date:   0,
				option: 1,
				mode:   0,
			},
			want:    2019,
			want1:   11,
			want2:   0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, err := JudgetDate(tt.args.year, tt.args.month, tt.args.date, tt.args.option, tt.args.mode)
			if (err != nil) != tt.wantErr {
				t.Errorf("JudgetDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("JudgetDate() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("JudgetDate() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("JudgetDate() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
