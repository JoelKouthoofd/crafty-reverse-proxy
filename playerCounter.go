package main

import (
	"strconv"
	"time"
)

var playerMap *map[string]int

func indexFromServer(server ServerType) string {
	return server.InternalIp + ":" + server.InternalPort
}

func getPlayerMap() *map[string]int {
	if playerMap == nil {
		intermediaryPlayerMap := make(map[string]int)
		for _, elem := range getConfig().Addresses {
			intermediaryPlayerMap[indexFromServer(elem)] = 0
		}

		playerMap = &intermediaryPlayerMap
	}

	return playerMap
}
