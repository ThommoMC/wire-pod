package pluginmanager

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func PluginapiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	switch strings.TrimPrefix(r.URL.Path, "/plugin-api/") {
	case "get_plugin_list":
		handleGetPluginList(w)
	}

}

func handleGetPluginList(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	// TODO: Allow user to set custom repositorys
	resp, err := http.Get("https://raw.githubusercontent.com/ThommoMC/wirepod-plugins/refs/heads/main/main.json")
	if err != nil {
		http.Error(w, "error communicating with repository: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	type plugin struct {
		Name            string `json:"name"`
		Author          string `json:"author"`
		Description     string `json:"description"`
		Image           string `json:"image"`
		Version         string `json:"version"`
		Minidescription string `json:"minidescription"`
		Download        string `json:"download"`
	}

	type pluginEntrys map[string]plugin

	var pluginentry pluginEntrys
	if err := json.NewDecoder(resp.Body).Decode(&pluginentry); err != nil {
		http.Error(w, "error while decoding json: "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(pluginentry)
}
