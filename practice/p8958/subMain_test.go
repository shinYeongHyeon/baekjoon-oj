package p8958

import "testing"

func TestScoreIsOneWhenGradingStringWasO(t *testing.T) {
	score := CalculateScore("O")

	if score != 1 {
		t.Error("ScoreIsOneWhenGradingStringWasO")
	}
}

func TestScoreIsThreeWhenGradingStringWasOO(t *testing.T) {
	score := CalculateScore("OO")

	if score != 3 {
		t.Error("ScoreIsThreeWhenGradingStringWasOO")
	}
}

func TestScoreIsFourWhenGradingStringWasOOXO(t *testing.T) {
	score := CalculateScore("OOXO")

	if score != 4 {
		t.Error("ScoreIsFourWhenGradingStringWasOOXO")
	}
}
