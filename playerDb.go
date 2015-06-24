package main

import (
	"database/sql"
	"errors"
)

type PlayerDbRepo struct {
	Conn *sql.DB
}

// Save one or more players to the database
func (p *PlayerDbRepo) Save(players ...Player) ([]Player, []error) {
	savedPlayers := make([]Player, 0)
	errorsFound := make([]error, 0)

	query := "INSERT INTO `players` (`name`,`is_host`) VALUES (?,?)"
	for _, player := range players {

		res, err := p.Conn.Exec(query, player.Name, player.IsHost)
		if err != nil {
			errorsFound = append(errorsFound, err)

		} else {
			lastId, err := res.LastInsertId()

			if err != nil {
				errorsFound = append(errorsFound, err)

			} else {
				savedPlayers = append(savedPlayers, Player{Id: int(lastId), Name: player.Name, IsHost: player.IsHost})
			}
		}
	}

	return savedPlayers, errorsFound
}

// FindById pulls a player out of the database based on their ID
func (p *PlayerDbRepo) FindById(id int) (Player, error) {
	if id == 0 {
		return Player{}, errors.New("Player ID will never be zero")
	}

	result := Player{}
	query := "SELECT name, is_host FROM `players` WHERE id = ?"

	err := p.Conn.QueryRow(query, id).Scan(&result.Name, &result.IsHost)
	if err != nil {
		return Player{}, err
	}

	return result, nil
}
