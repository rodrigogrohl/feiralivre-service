package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const (
	targetPath = "./assets/FEIRAS_LIVRES/CSV/DEINFO_DADOS_AB_FEIRASLIVRES/"
	targetFile = "DEINFO_AB_FEIRASLIVRES_2014.csv"
)

type StreetMarket struct {
	Id            int64   `csv:"ID"`
	Longitude     float64 `csv:"LONG"`
	Latitude      float64 `csv:"LAT"`
	SectorCense   int64   `csv:"SETCENS"`
	AreaPonderate int64   `csv:"AREAP"`
	DistrictCode  int64   `csv:"CODDIST"`
	District      string  `csv:"DISTRITO"`
	SubTownCode   int64   `csv:"CODSUBPREF"`
	SubTown       string  `csv:"SUBPREFE"`
	Region5       string  `csv:"REGIAO5"`
	Region8       string  `csv:"REGIAO8"`
	Name          string  `csv:"NOME_FEIRA"`
	Registry      string  `csv:"REGISTRO"`
	Address       string  `csv:"LOGRADOURO"`
	Number        string  `csv:"NUMERO"`
	Neighborhood  string  `csv:"BAIRRO"`
	Reference     string  `csv:"reference"`
}

func main() {

	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)

	// insert
	stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
	checkErr(err)

	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
	// update
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("astaxieupdate", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	// query
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)
	var uid int
	var username string
	var department string
	var created time.Time

	for rows.Next() {
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	rows.Close() //good habit to close

	// delete
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
