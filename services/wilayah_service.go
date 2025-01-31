package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Struktur untuk menyimpan data provinsi
type Provinsi struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Struktur untuk menyimpan data kota/kabupaten
type Kota struct {
	ID         string `json:"id"`
	ProvinceID string `json:"province_id"`
	Name       string `json:"name"`
}

type WilayahService struct {
	BaseURL string
}

func NewWilayahService() *WilayahService {
	return &WilayahService{
		BaseURL: "https://emsifa.github.io/api-wilayah-indonesia/api",
	}
}

// Fetch daftar provinsi
func (ws *WilayahService) GetProvinsi() ([]Provinsi, error) {
	url := fmt.Sprintf("%s/provinces.json", ws.BaseURL)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("gagal mengambil data provinsi")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var provinces []Provinsi
	if err := json.Unmarshal(body, &provinces); err != nil {
		return nil, err
	}

	return provinces, nil
}

// Fetch daftar kota berdasarkan ID Provinsi
func (ws *WilayahService) GetKotaByProvinsi(provinceID string) ([]Kota, error) {
	url := fmt.Sprintf("%s/regencies/%s.json", ws.BaseURL, provinceID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("gagal mengambil data kota/kabupaten")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var cities []Kota
	if err := json.Unmarshal(body, &cities); err != nil {
		return nil, err
	}

	return cities, nil
}
