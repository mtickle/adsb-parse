package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
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

	//--- Load up environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	//fxApiKey := os.Getenv("FX_API_KEY")

	//-----------------------------------------------------------------------
	//--- Set the data url
	fxSourceUrl := os.Getenv("FX_SOURCE_URL")
	//fxApiUrl := os.Getenv("FX_API_URL")
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
		hex_code := strings.ToUpper(strings.TrimSpace(rec.Hex))
		type_code := strings.TrimSpace(rec.Type)
		flight := strings.TrimSpace(rec.Flight)
		alt_baro := rec.AltBaro
		alt_geom := rec.AltGeom
		gs := rec.Gs
		track := rec.Track
		baro_rate := rec.BaroRate
		squawk := strings.TrimSpace(rec.Squawk)
		emergency := strings.TrimSpace(rec.Emergency)
		category := strings.TrimSpace(rec.Category)
		lat := rec.Lat
		lon := rec.Lon
		nic := rec.Nic
		rc := rec.Rc
		seen_pos := rec.SeenPos
		r_dst := rec.RDst
		r_dir := rec.RDir
		version := rec.Version
		nic_baro := rec.NicBaro
		nac_p := rec.NacP
		nac_v := rec.NacV
		sil := rec.Sil
		sil_type := strings.TrimSpace(rec.SilType)
		gva := rec.Gva
		sda := rec.Sda
		alert := rec.Alert
		spi := rec.Spi
		messages := rec.Messages
		seen := rec.Seen
		rssi := rec.Rssi

		sql := "CALL add_flight ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31);"

		_, err := db.Exec(sql, hex_code, type_code, flight, alt_baro, alt_geom, gs, track, baro_rate, squawk, emergency, category, lat, lon, nic, rc, seen_pos, r_dst, r_dir, version, nic_baro, nac_p, nac_v, sil, sil_type, gva, sda, alert, spi, messages, seen, rssi)
		if err != nil {
			panic(err.Error())
		}

		// var sb strings.Builder
		// sb.WriteString("{")
		// sb.WriteString("\"hex\": \"" + hex + "\", ")
		// sb.WriteString("\"ttype\": \"" + ttype + "\", ")
		// sb.WriteString("\"flight\": \"" + flight + "\", ")
		// sb.WriteString("\"alt_baro\": \"" + fmt.Sprintf("%d", alt_baro) + "\", ")
		// sb.WriteString("\"alt_geom\": \"" + fmt.Sprintf("%d", alt_geom) + "\", ")
		// sb.WriteString("\"gs\": \"" + fmt.Sprintf("%g", gs) + "\", ")
		// sb.WriteString("\"track\": \"" + fmt.Sprintf("%g", track) + "\", ")
		// sb.WriteString("\"baro_rate\": \"" + fmt.Sprintf("%d", baro_rate) + "\", ")
		// sb.WriteString("\"squawk\": \"" + squawk + "\", ")
		// sb.WriteString("\"emergency\": \"" + emergency + "\", ")
		// sb.WriteString("\"category\": \"" + category + "\", ")
		// sb.WriteString("\"lat\": \"" + fmt.Sprintf("%g", lat) + "\", ")
		// sb.WriteString("\"lon\": \"" + fmt.Sprintf("%g", lon) + "\", ")
		// sb.WriteString("\"nic\": \"" + fmt.Sprintf("%d", nic) + "\", ")
		// sb.WriteString("\"rc\": \"" + fmt.Sprintf("%d", rc) + "\", ")
		// sb.WriteString("\"seen_pos\": \"" + fmt.Sprintf("%g", seen_pos) + "\", ")
		// sb.WriteString("\"r_dst\": \"" + fmt.Sprintf("%g", r_dst) + "\", ")
		// sb.WriteString("\"r_dir\": \"" + fmt.Sprintf("%g", r_dir) + "\", ")
		// sb.WriteString("\"version\": \"" + fmt.Sprintf("%d", version) + "\", ")
		// sb.WriteString("\"nic_baro\": \"" + fmt.Sprintf("%d", nic_baro) + "\", ")
		// sb.WriteString("\"nac_p\": \"" + fmt.Sprintf("%d", nac_p) + "\", ")
		// sb.WriteString("\"nac_v\": \"" + fmt.Sprintf("%d", nac_v) + "\", ")
		// sb.WriteString("\"sil\": \"" + fmt.Sprintf("%d", sil) + "\", ")
		// sb.WriteString("\"sil_type\": \"" + sil_type + "\", ")
		// sb.WriteString("\"gva\": \"" + fmt.Sprintf("%d", gva) + "\", ")
		// sb.WriteString("\"sda\": \"" + fmt.Sprintf("%d", sda) + "\", ")
		// sb.WriteString("\"alert\": \"" + fmt.Sprintf("%d", alert) + "\", ")
		// sb.WriteString("\"spi\": \"" + fmt.Sprintf("%d", spi) + "\", ")
		// sb.WriteString("\"mlat\": \"" + fmt.Sprintf("%g", mlat) + "\", ")
		// sb.WriteString("\"tisb\": \"" + fmt.Sprintf("%g", tisb) + "\", ")
		// sb.WriteString("\"messages\": \"" + fmt.Sprintf("%d", messages) + "\", ")
		// sb.WriteString("\"seen\": \"" + fmt.Sprintf("%g", seen) + "\", ")
		// sb.WriteString("\"rssi\": \"" + fmt.Sprintf("%g", rssi) + "\"")
		// sb.WriteString("}")

		// //--- Just hold this here
		// temp := sb.String()
		// jsonBody := []byte(temp)
		// apiPayload := bytes.NewReader(jsonBody)

		// fmt.Println(temp)

		// //--- Send the request downrange to the API URL
		// req, err := http.NewRequest(http.MethodPost, fxApiUrl, apiPayload)
		// if err != nil {
		// 	panic(err)
		// }
		// req.Header.Set("Content-Type", "application/json; charset=UTF-8")
		// req.Header.Set("x-api-key", fxApiKey)

		// client := &http.Client{}
		// response, error := client.Do(req)
		// //fmt.Println(response)
		// if error != nil {
		// 	panic(error)
		// }
		// defer response.Body.Close()

		// //--- Logging
		// fmt.Println("\"hex\": \""+hex+"\" -  API response:", response.Status)
	}
}
