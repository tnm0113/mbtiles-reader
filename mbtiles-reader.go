package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("Mbtiles reader")

	db, err := sql.Open("sqlite3", "trails.mbtiles")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM tiles")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var zoomLevel int
		var tileColumn int
		var tileRow int
		var tileData []byte

		err = rows.Scan(&zoomLevel, &tileColumn, &tileRow, &tileData)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("zoomLevel: %d tileColumn: %d tileRow: %d\n", zoomLevel, tileColumn, tileRow)
		folderName := fmt.Sprintf("tiles/%d/%d", zoomLevel, tileColumn)
		err = os.MkdirAll(folderName, 0750)
		if err != nil && !os.IsExist(err) {
			log.Fatal(err)
		}
		fileName := fmt.Sprintf("tiles/%d/%d/%d.png", zoomLevel, tileColumn, tileRow)
		err = os.WriteFile(fileName, tileData, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}
