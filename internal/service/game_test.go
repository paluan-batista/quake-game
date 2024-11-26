package service_test

import (
	"fmt"
	"quake-log-parser/internal/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type LogParserTestSuite struct {
	suite.Suite
}

func (suite *LogParserTestSuite) TestParseLogFile() {
	games, err := service.ParseLogFile("../mocks/qgames.log")
	if err != nil {
		fmt.Printf("error on read logs: %s", err)
	}
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), games)
	assert.Equal(suite.T(), 1, len(games))

	game1 := games["jogo_1"]
	assert.Equal(suite.T(), 7, game1.TotalKills)
	assert.Equal(suite.T(), 7, len(game1.Kills))
	assert.Equal(suite.T(), 1, game1.Deaths["Dano por Trigger"])
	assert.Equal(suite.T(), 4, game1.Deaths["Explosão de Foguete"])
	assert.Equal(suite.T(), 1, game1.Deaths["Foguete"])
	assert.Equal(suite.T(), 1, game1.Deaths["Queda"])
}

func TestLogParserTestSuite(t *testing.T) {
	suite.Run(t, new(LogParserTestSuite))
}
