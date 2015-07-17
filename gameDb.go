package main

import (
	"database/sql"
	"errors"
)

// GameDbRepo is a way of accessing the persistence layer for saving
// and loading of the game.
type GameDbRepo struct {
	db *sql.DB
}

// NewGameDbRepo creates a new struct that can be used to save and
// load game information.
func NewGameDbRepo(Connection *sql.DB) GameDbRepo {
	return GameDbRepo{
		db: Connection,
	}
}

// Save saves a game to the database.
func (g *GameDbRepo) Save(game Game) (int, error) {

	gameID := 0

	if game.Id == 0 {
		res, err := g.db.Exec("INSERT INTO games (name, theme) VALUES (?, ?)", game.Name, game.Theme)
		if err != nil {
			return 0, err
		}

		lastID, err := res.LastInsertId()
		if err != nil {
			return 0, err
		}

		gameID = int(lastID)

	} else {
		_, err := g.db.Exec("UPDATE games SET name = ?, theme = ? WHERE id = ?", game.Name, game.Theme, game.Id)
		if err != nil {
			return 0, err
		}

		gameID = game.Id

	}

	_, errsEncountered := g.savePlayers(gameID, game.Players()...)
	if len(errsEncountered) > 0 {
		return 0, errors.New("Unable to save players to database")
	}

	return gameID, nil

}

// savePlayers saves all the players to a map associated with a game
func (g *GameDbRepo) savePlayers(gameID int, players ...Player) ([]Player, []error) {
	savedPlayers := make([]Player, 0)
	errorsFound := make([]error, 0)

	query := "INSERT INTO `games_players` (`game_id`,`player_id`) VALUES (?,?)"
	for _, player := range players {

		_, err := g.db.Exec(query, gameID, player.Id)
		if err != nil {
			errorsFound = append(errorsFound, err)
		} else {
			savedPlayers = append(savedPlayers, player)
		}
	}

	return savedPlayers, errorsFound
}

// FindByID locates a game based on a given ID
func (g *GameDbRepo) FindByID(id int) (Game, error) {

	rows, err := g.db.Query(
		"SELECT g.id, g.name, g.theme, p.id, p.name, p.is_host FROM games g INNER JOIN `games_players` gp ON gp.game_id = g.id INNER JOIN players p ON p.id = gp.player_id WHERE g.id = ? ORDER BY g.id",
		id)
	if err != nil {
		return Game{}, err
	}
	defer rows.Close()

	foundGame := Game{}
	foundPlayers := make([]Player, 0)
	for rows.Next() {

		var gameID, playerID int
		var gameName, gameTheme, playerName string
		var isHost bool
		err := rows.Scan(&gameID, &gameName, &gameTheme, &playerID, &playerName, &isHost)
		if err != nil {
			return foundGame, err
		}

		// set game info if first record
		if foundGame.Id == 0 {
			foundGame.Id = gameID
			foundGame.Name = gameName
			foundGame.Theme = gameTheme
		}

		foundPlayers = append(foundPlayers, Player{Id: playerID, Name: playerName, IsHost: isHost})
	}

	//foundGame.AddPlayers(foundPlayers)

	err = rows.Err()
	if err != nil {
		return Game{}, err
	}

	return foundGame, nil
}

// FindPlayersByGame finds all players associated with the selected game
func (g *GameDbRepo) FindPlayersByGame(gameID int) ([]Player, error) {
	foundPlayers := make([]Player, 0)

	rows, err := g.db.Query(
		"SELECT p.id, p.name, p.is_host FROM `players` p INNER JOIN `games_players` gp ON gp.player_id = p.id WHERE gp.game_id = ?",
		gameID)
	if err != nil {
		return foundPlayers, err
	}
	defer rows.Close()

	for rows.Next() {
		player := Player{}
		err := rows.Scan(&player.Id, &player.Name, &player.IsHost)
		if err != nil {
			return foundPlayers, err
		}

		foundPlayers = append(foundPlayers, player)
	}

	err = rows.Err()
	if err != nil {
		return foundPlayers, err
	}

	return foundPlayers, nil
}
