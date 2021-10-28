package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // Import go-sql library
)

func main() {
	var d = false
	lvl, ok := os.LookupEnv("LOG_LEVEL")
	// LOG_LEVEL not set, let's default to debug
	if ok {
		d = "debug" == lvl
	}

	outlog, err := os.Create("backupfiles.txt")

	if err != nil {
		log.Fatal(err.Error())
	}

	db, _ := sql.Open("mysql", os.Args[2])
	defer db.Close()

	out := bufio.NewWriter(outlog)

	taggedPhotoSQL := `select f.file_root, f.file_name from files f, labels l, photos_labels pl where f.photo_id = pl.photo_id and pl.label_id = l.id and f.file_root = '/' and l.label_slug = ?`
	generalAlbumSQL := `select f.file_root, f.file_name from files f, albums a, photos_albums pa where f.photo_uid = pa.photo_uid and pa.album_uid = a.album_uid and f.file_root = '/' and a.album_slug = ?`

	var op = os.Args[3]
	var sql = taggedPhotoSQL
	if op == "tag" {
		sql = taggedPhotoSQL
	} else if op == "album" {
		sql = generalAlbumSQL
	}

	for i := 4; i < len(os.Args); i++ {
		findPhotos(db, sql, os.Args[i], out, d)
	}
	out.Flush()
}

func findPhotos(db *sql.DB, sql string, arg string, out *bufio.Writer, d bool) {
	log.Println("Finding Photos")

	var count = 0

	row, err := db.Query(sql, arg)
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer row.Close()
	for row.Next() {
		var file_root string
		var file_name string

		row.Scan(&file_root, &file_name)

		if d {
			fmt.Printf("Photo: %s, %s\n", file_root, file_name)
		}

		var pathStr = os.Args[1] + file_root + file_name

		if d {
			fmt.Printf("Photo with path %s\n", pathStr)
		}

		out.WriteString(pathStr + "\n")
		count++
	}

	log.Println("Found ", count, " photos")
}
