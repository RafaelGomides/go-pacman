package main

import (
	"reflect"
	"testing"
)

func TestGameStatus_checkIfHasDot(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		g    *GameStatus
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.g.checkIfHasDot(tt.args.x, tt.args.y)
		})
	}
}

func TestGameStatus_Up(t *testing.T) {
	tests := []struct {
		name string
		g    *GameStatus
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.g.Up()
		})
	}
}

func TestGameStatus_Down(t *testing.T) {
	tests := []struct {
		name string
		g    *GameStatus
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.g.Down()
		})
	}
}

func TestGameStatus_Left(t *testing.T) {
	tests := []struct {
		name string
		g    *GameStatus
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.g.Left()
		})
	}
}

func TestGameStatus_Right(t *testing.T) {
	tests := []struct {
		name string
		g    *GameStatus
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.g.Right()
		})
	}
}

func TestGameStatus_generateManStartPosition(t *testing.T) {
	tests := []struct {
		name string
		g    *GameStatus
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.g.generateManStartPosition()
		})
	}
}

func Test_getGridMap(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want *GameStatus
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getGridMap(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getGridMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_playGame(t *testing.T) {
	type args struct {
		game *GameStatus
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := playGame(tt.args.game); got != tt.want {
				t.Errorf("playGame() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
