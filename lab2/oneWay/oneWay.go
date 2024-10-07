package oneway

import (
	"yanandco/lab1/crypto"
)

type SBlock = crypto.SBlock
type TelegraphChar = crypto.TelegraphChar

func OneWayEcrypt(data *SBlock, key []*TelegraphChar, iterations int) (*SBlock, error) {
	return nil, nil
}
