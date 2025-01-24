package main

import (
	golog "github.com/Vladroon22/GoLog"
	"github.com/Vladroon22/TestTask-Bank-Operation/config"
	"github.com/Vladroon22/TestTask-Bank-Operation/internal/database"
)

func main() {
	cnf := config.CreateConfig()
	logger := golog.New()

	conn, err := database.DBConnection(cnf.DB)
	if err != nil {
		logger.Fatalln(err)
	}

	database.CloseConn(conn)
}
