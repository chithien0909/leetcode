package main

import (
    "reflect"
    "testing"
)

func TestMySqrt(t *testing.T) {
    type args struct {
        x int
    }
    tests := []struct {
        name string
        args args
        want int
    }{
        {
            name: "test1",
            args: args{
                x: 4,
            },
            want: 2,
        },
        {
            name: "test2",
            args: args{
                x: 8,
            },
            want: 2,
        },
        {
            name: "test3",
            args: args{
                x: 9,
            },
            want: 3,
        },
        {
            name: "test4",
            args: args{
                x: 0,
            },
            want: 0,
        },
        {
            name: "test5",
            args: args{
                x: 1,
            },
            want: 1,
        },
        {
            name: "test6",
            args: args{
                x: 234,
            },
            want: 15,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := mySqrt(tt.args.x); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("mySqrt() = %v, want %v", got, tt.want)
            }
        })
    }
}