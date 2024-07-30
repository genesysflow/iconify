package iconify

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func GetCollections(api string) ([]Collection, error) {
	var result map[string]Collection

	resp, err := http.Get(api + "/collections?prefixes=mdi,fa")

	if err != nil {
		log.Fatal("Error getting collections:", err)
		return nil, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatal("Error decoding collections:", err)
		return nil, err
	}

	ret := make([]Collection, 0)
	for key, value := range result {
		value.Key = key
		ret = append(ret, value)
	}

	return ret, nil
}

func GetIconCollection(api string, key string) (*IconCollection, error) {
	var result IconCollection

	resp, err := http.Get(api + "/collection?prefix=" + key)
	if err != nil {
		log.Fatal("Error getting collections:", err)
		return nil, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatal("Error decoding collections:", err)
		return nil, err
	}

	return &result, nil
}

func GetIcon(api, prefix, icon string) (string, error) {
	resp, err := http.Get(api + "/" + prefix + "/" + icon + ".svg")
	if err != nil {
		log.Fatal("Error getting icon:", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading icon:", err)
		return "", err
	}

	return string(body), nil
}
