package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
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
	pxSourceUrl := "http://192.168.86.58/adsbx/data/aircraft.json"
	//-----------------------------------------------------------------------

	//-----------------------------------------------------------------------
	//--- Load data from the API
	resp, err := http.Get(pxSourceUrl)
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("No response from request")
	}

	//--- Unmarshal the JSON into the "result" value
	var apiResult Response
	if err := json.Unmarshal(body, &apiResult); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	//-----------------------------------------------------------------------
	//-----------------------------------------------------------------------
	//--- Start iterating through the records
	for _, rec := range apiResult.Aircraft {

		hex := rec.Hex
		ttype := rec.Type
		flight := rec.Flight
		alt_baro := rec.AltBaro
		alt_geom := rec.AltGeom
		gs := rec.Gs
		track := rec.Track
		baro_rate := rec.BaroRate
		squawk := rec.Squawk
		emergency := rec.Emergency
		category := rec.Category
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
		sil_type := rec.SilType
		gva := rec.Gva
		sda := rec.Sda
		alert := rec.Alert
		spi := rec.Spi
		mlat := rec.Mlat
		tisb := rec.Tisb
		messages := rec.Messages
		seen := rec.Seen
		rssi := rec.Rssi

		var sb strings.Builder
		sb.WriteString("{")
		sb.WriteString("\"hex\": \"" + hex + "\", ")
		sb.WriteString("\"ttype\": \"" + ttype + "\", ")
		sb.WriteString("\"flight\": \"" + flight + "\", ")
		sb.WriteString("\"alt_baro\": \"" + fmt.Sprintf("%d", alt_baro) + "\", ")
		sb.WriteString("\"alt_geom\": \"" + fmt.Sprintf("%d", alt_geom) + "\", ")
		sb.WriteString("\"gs\": \"" + fmt.Sprintf("%g", gs) + "\", ")
		sb.WriteString("\"track\": \"" + fmt.Sprintf("%g", track) + "\", ")
		sb.WriteString("\"baro_rate\": \"" + fmt.Sprintf("%d", baro_rate) + "\", ")
		sb.WriteString("\"squawk\": \"" + squawk + "\", ")
		sb.WriteString("\"emergency\": \"" + emergency + "\", ")
		sb.WriteString("\"category\": \"" + category + "\", ")
		sb.WriteString("\"lat\": \"" + fmt.Sprintf("%g", lat) + "\", ")
		sb.WriteString("\"lon\": \"" + fmt.Sprintf("%g", lon) + "\", ")
		sb.WriteString("\"nic\": \"" + fmt.Sprintf("%d", nic) + "\", ")
		sb.WriteString("\"rc\": \"" + fmt.Sprintf("%d", rc) + "\", ")
		sb.WriteString("\"seen_pos\": \"" + fmt.Sprintf("%g", seen_pos) + "\", ")
		sb.WriteString("\"r_dst\": \"" + fmt.Sprintf("%g", r_dst) + "\", ")
		sb.WriteString("\"r_dir\": \"" + fmt.Sprintf("%g", r_dir) + "\", ")
		sb.WriteString("\"version\": \"" + fmt.Sprintf("%d", version) + "\", ")
		sb.WriteString("\"nic_baro\": \"" + fmt.Sprintf("%d", nic_baro) + "\", ")
		sb.WriteString("\"nac_p\": \"" + fmt.Sprintf("%d", nac_p) + "\", ")
		sb.WriteString("\"nac_v\": \"" + fmt.Sprintf("%d", nac_v) + "\", ")
		sb.WriteString("\"sil\": \"" + fmt.Sprintf("%d", sil) + "\", ")
		sb.WriteString("\"sil_type\": \"" + sil_type + "\", ")
		sb.WriteString("\"gva\": \"" + fmt.Sprintf("%d", gva) + "\", ")
		sb.WriteString("\"sda\": \"" + fmt.Sprintf("%d", sda) + "\", ")
		sb.WriteString("\"alert\": \"" + fmt.Sprintf("%d", alert) + "\", ")
		sb.WriteString("\"spi\": \"" + fmt.Sprintf("%d", spi) + "\", ")
		sb.WriteString("\"mlat\": \"" + fmt.Sprintf("%g", mlat) + "\", ")
		sb.WriteString("\"tisb\": \"" + fmt.Sprintf("%g", tisb) + "\", ")
		sb.WriteString("\"messages\": \"" + fmt.Sprintf("%d", messages) + "\", ")
		sb.WriteString("\"seen\": \"" + fmt.Sprintf("%g", seen) + "\", ")
		sb.WriteString("\"rssi\": \"" + fmt.Sprintf("%g", rssi) + "\", ")
		sb.WriteString("}")

		temp := sb.String()
		fmt.Println(temp)

	}
}
