package main

import (
	"database/sql"
	"errors"
)

type GameDbRepo struct {
	Conn *sql.DB
}

func NewGameDbRepo(Connection *sql.DB) GameDbRepo {
	return GameDbRepo{
		Conn: Connection,
	}
}

// Save saves a game to the database
func (g *GameDbRepo) Save(game Game) (int, error) {

	gameId := 0

	if game.Id == 0 {
		res, err := g.Conn.Exec("INSERT INTO games (name, theme) VALUES (?, ?)", game.Name, game.Theme)
		if err != nil {
			return 0, err
		}

		lastId, err := res.LastInsertId()
		if err != nil {
			return 0, err
		}

		gameId = int(lastId)

	} else {
		_, err := g.Conn.Exec("UPDATE games SET name = ?, theme = ? WHERE id = ?", game.Name, game.Theme, game.Id)
		if err != nil {
			return 0, err
		}

		gameId = game.Id

	}

	_, errsEncountered := g.savePlayers(gameId, game.Players...)
	if len(errsEncountered) > 0 {
		return 0, errors.New("Unable to save players to database")
	}

	return gameId, nil

}

// savePlayers saves all the players to a map associated with a game
func (g *GameDbRepo) savePlayers(gameId int, players ...Player) ([]Player, []error) {
	savedPlayers := make([]Player, 0)
	errorsFound := make([]error, 0)

	query := "INSERT INTO `games_players` (`game_id`,`player_id`) VALUES (?,?)"
	for _, player := range players {

		_, err := g.Conn.Exec(query, gameId, player.Id)
		if err != nil {
			errorsFound = append(errorsFound, err)
		} else {
			savedPlayers = append(savedPlayers, player)
		}
	}

	return savedPlayers, errorsFound
}

// FindById locates a game based on a given ID
func (g *GameDbRepo) FindById(id int) (Game, error) {

	rows, err := g.Conn.Query(
		"SELECT id, name, theme FROM games WHERE id = ?",
		id)
	if err != nil {
		return Game{}, err
	}
	defer rows.Close()

	foundGame := Game{}
	for rows.Next() {
		if foundGame.Id == 0 {
			err := rows.Scan(&foundGame.Id, &foundGame.Name, &foundGame.Theme)
			if err != nil {
				return Game{}, err
			}
		}
	}

	err = rows.Err()
	if err != nil {
		return Game{}, err
	}

	return foundGame, nil
}
