package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "192.168.86.2"
	port     = 5432
	user     = "pi"
	password = "Boomer2025"
	dbname   = "adsb"
)

type Response struct {
	Now      float64 `json:"now"`
	Messages int     `json:"messages"`
	Aircraft []struct {
		Hex            string  `json:"hex"`
		Type           string  `json:"type"`
		Flight         string  `json:"flight,omitempty"`
		AltBaro        int     `json:"alt_baro,omitempty"`
		AltGeom        int     `json:"alt_geom"`
		Gs             float64 `json:"gs"`
		Track          float64 `json:"track"`
		BaroRate       int     `json:"baro_rate,omitempty"`
		Squawk         string  `json:"squawk,omitempty"`
		Emergency      string  `json:"emergency,omitempty"`
		Category       string  `json:"category,omitempty"`
		Lat            float64 `json:"lat,omitempty"`
		Lon            float64 `json:"lon,omitempty"`
		Nic            int     `json:"nic,omitempty"`
		Rc             int     `json:"rc,omitempty"`
		SeenPos        float64 `json:"seen_pos,omitempty"`
		RDst           float64 `json:"r_dst,omitempty"`
		RDir           float64 `json:"r_dir,omitempty"`
		Version        int     `json:"version"`
		NicBaro        int     `json:"nic_baro"`
		NacP           int     `json:"nac_p"`
		NacV           int     `json:"nac_v"`
		Sil            int     `json:"sil"`
		SilType        string  `json:"sil_type"`
		Gva            int     `json:"gva,omitempty"`
		Sda            int     `json:"sda,omitempty"`
		Alert          int     `json:"alert,omitempty"`
		Spi            int     `json:"spi,omitempty"`
		Mlat           []any   `json:"mlat"`
		Tisb           []any   `json:"tisb"`
		Messages       int     `json:"messages"`
		Seen           float64 `json:"seen"`
		Rssi           float64 `json:"rssi"`
		NavQnh         float64 `json:"nav_qnh,omitempty"`
		NavAltitudeMcp int     `json:"nav_altitude_mcp,omitempty"`
		NavHeading     float64 `json:"nav_heading,omitempty"`
		GeomRate       int     `json:"geom_rate,omitempty"`
		LastPosition   struct {
			Lat     float64 `json:"lat"`
			Lon     float64 `json:"lon"`
			Nic     int     `json:"nic"`
			Rc      int     `json:"rc"`
			SeenPos float64 `json:"seen_pos"`
		} `json:"lastPosition,omitempty"`
	} `json:"aircraft"`
}

func main() {

	//-----------------------------------------------------------------------
	//--- Set the data url
	fxSourceUrl := "http://192.168.86.58/tar1090/data/aircraft.json"
	//-----------------------------------------------------------------------

	//-----------------------------------------------------------------------
	//--- Load data from the API
	resp, err := http.Get(fxSourceUrl)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("No response from request")
	}
	var apiResult Response
	if err := json.Unmarshal(body, &apiResult); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	//-----------------------------------------------------------------------

	//-----------------------------------------------------------------------
	//--- Make and open the database connection
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//--- Are we good?
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	//-----------------------------------------------------------------------

	//-----------------------------------------------------------------------
	//--- Start iterating through the records
	for _, rec := range apiResult.Aircraft {
		flight_time := time.Now()
		hex_code := strings.ToUpper(strings.TrimSpace(rec.Hex))
		type_code := strings.TrimSpace(rec.Type)
		flight_code := strings.TrimSpace(rec.Flight)
		alt_baro := rec.AltBaro
		category := strings.TrimSpace(rec.Category)
		lat := rec.Lat
		lon := rec.Lon
		squawk := rec.Squawk

		sql := "CALL add_flight ($1, $2, $3, $4, $5, $6, $7, $8, $9);"

		_, err := db.Exec(sql, hex_code, type_code, flight_code, flight_time, alt_baro, category, lat, lon, squawk)

		if err != nil {
			panic(err)
		}

	}
}
