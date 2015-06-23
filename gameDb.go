package main

import (
	"database/sql"
)

type GameDbRepo struct {
	Conn *sql.DB
}

func NewGameDbRepo(Connection *sql.DB) GameDbRepo {
	return GameDbRepo{
		Conn: Connection,
	}
}

func (g *GameDbRepo) Save(game Game) (int, error) {

	if game.Id == 0 {
		res, err := g.Conn.Exec("INSERT INTO games (name, theme) VALUES (?, ?)", game.Name, game.Theme)
		if err != nil {
			return 0, err
		}

		lastId, err := res.LastInsertId()
		if err != nil {
			return 0, err
		}

		return int(lastId), nil
	} else {
		_, err := g.Conn.Exec("UPDATE games SET name = ?, theme = ? WHERE id = ?", game.Name, game.Theme, game.Id)
		if err != nil {
			return 0, err
		}

		return game.Id, nil
	}

}

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
