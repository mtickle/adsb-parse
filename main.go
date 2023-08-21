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

		var sb strings.Builder

		sb.WriteString("{")
		sb.WriteString("\"hex\": \"" + rec.Hex + "\", ")
		sb.WriteString("\"ttype\": \"" + rec.Type + "\", ")
		sb.WriteString("\"flight\": \"" + rec.Flight + "\", ")
		sb.WriteString("\"alt_baro\": \"" + fmt.Sprintf("%d", rec.AltBaro) + "\", ")
		sb.WriteString("\"alt_geom\": \"" + fmt.Sprintf("%d", rec.AltGeom) + "\", ")
		sb.WriteString("\"gs\": \"" + fmt.Sprintf("%g", rec.Gs) + "\", ")
		sb.WriteString("\"track\": \"" + fmt.Sprintf("%g", rec.Track) + "\", ")
		sb.WriteString("\"baro_rate\": \"" + fmt.Sprintf("%d", rec.BaroRate) + "\", ")
		sb.WriteString("\"squawk\": \"" + rec.Squawk + "\", ")
		sb.WriteString("\"emergency\": \"" + rec.Emergency + "\", ")
		sb.WriteString("\"category\": \"" + rec.Category + "\", ")
		sb.WriteString("\"lat\": \"" + fmt.Sprintf("%g", rec.Lat) + "\", ")
		sb.WriteString("\"lon\": \"" + fmt.Sprintf("%g", rec.Lon) + "\", ")
		sb.WriteString("\"nic\": \"" + fmt.Sprintf("%d", rec.Nic) + "\", ")
		sb.WriteString("\"rc\": \"" + fmt.Sprintf("%d", rec.Rc) + "\", ")
		sb.WriteString("\"seen_pos\": \"" + fmt.Sprintf("%g", rec.SeenPos) + "\", ")
		sb.WriteString("\"r_dst\": \"" + fmt.Sprintf("%g", rec.RDst) + "\", ")
		sb.WriteString("\"r_dir\": \"" + fmt.Sprintf("%g", rec.RDir) + "\", ")
		sb.WriteString("\"version\": \"" + fmt.Sprintf("%d", rec.Version) + "\", ")
		sb.WriteString("\"nic_baro\": \"" + fmt.Sprintf("%d", rec.NicBaro) + "\", ")
		sb.WriteString("\"nac_p\": \"" + fmt.Sprintf("%d", rec.NacP) + "\", ")
		sb.WriteString("\"nac_v\": \"" + fmt.Sprintf("%d", rec.NacV) + "\", ")
		sb.WriteString("\"sil\": \"" + fmt.Sprintf("%d", rec.Sil) + "\", ")
		sb.WriteString("\"sil_type\": \"" + rec.SilType + "\", ")
		sb.WriteString("\"gva\": \"" + fmt.Sprintf("%d", rec.Gva) + "\", ")
		sb.WriteString("\"sda\": \"" + fmt.Sprintf("%d", rec.Sda) + "\", ")
		sb.WriteString("\"alert\": \"" + fmt.Sprintf("%d", rec.Alert) + "\", ")
		sb.WriteString("\"spi\": \"" + fmt.Sprintf("%d", rec.Spi) + "\", ")
		sb.WriteString("\"mlat\": \"" + fmt.Sprintf("%g", rec.Mlat) + "\", ")
		sb.WriteString("\"tisb\": \"" + fmt.Sprintf("%g", rec.Tisb) + "\", ")
		sb.WriteString("\"messages\": \"" + fmt.Sprintf("%d", rec.Messages) + "\", ")
		sb.WriteString("\"seen\": \"" + fmt.Sprintf("%g", rec.Seen) + "\", ")
		sb.WriteString("\"rssi\": \"" + fmt.Sprintf("%g", rec.Rssi) + "\", ")
		sb.WriteString("}")

		temp := sb.String()
		fmt.Println(temp)

	}
}
