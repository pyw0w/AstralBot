package commands

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type CoinFlipCommand struct{}

func (c *CoinFlipCommand) Name() string {
	return "coinflip"
}

func (c *CoinFlipCommand) Description() string {
	return "Симуляция подбрасывания монеты для двух игроков"
}

func (c *CoinFlipCommand) Execute(args []string) (interface{}, error) {
	if len(args) < 1 {
		return "необходимо указать второго игрока и выбор (орёл или решка)", nil
	}
	player1 := "отправитель" // Assuming the sender's name is "отправитель"
	player2 := args[0]
	choice := args[1]

	if choice != strings.ToLower("орёл") && choice != strings.ToLower("решка") {
		return "выбор должен быть 'орёл' или 'решка'", nil
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	flipResult := r.Intn(2) // 0 for heads, 1 for tails

	var winner string
	var result string
	if flipResult == 0 {
		result = "орёл"
	} else {
		result = "решка"
	}

	if result == choice {
		winner = player2
	} else {
		winner = player1
	}

	return fmt.Sprintf("Монета упала на %s! Победитель: %s", result, winner), nil
}

func RegisterCoinFlipCommand(cmdHandler *CommandHandler) {
	cmdHandler.RegisterCommand(&CoinFlipCommand{})
}
