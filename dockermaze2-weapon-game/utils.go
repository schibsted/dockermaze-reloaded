package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
    "log"
)

type Challenge struct {
	Challenge string `json:"challenge"`
}

type Evaluation struct {
	Target string `json:"target"`
	Action string `json:"action"`
}

type Result struct {
	Success   bool    `json:"success"`
	ScoreRate float64 `json:"score_rate"`
	Message   string  `json:"message"`
}

type Message struct {
	Success          bool    `json:"success"`
	Score            float64 `json:"score"`
	EnemiesDestroyed int     `json:"enemies_destroyed"`
	EnemiesSpared    int     `json:"enemies_spared"`
	AlliesDestroyed  int     `json:"allies_destroyed"`
	AlliesSpared     int     `json:"allies_spared"`
	Message          string  `json:"message"`
}

type Response struct {
	Response string `json:"response"`
}

// Fetch target image.
func fetchImage(target string) ([]byte, error) {
	client := http.Client{}
	req, err := http.NewRequest(
		"GET",
		target,
		nil,
	)
	if err != nil {
		return []byte{}, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, nil
}

// Get targets from head.
func getTargets() ([]string, error) {
	client := http.Client{}
	req, err := http.NewRequest(
		"GET",
		"http://"+endpoint+"/targets",
		nil,
	)
	if err != nil {
		return []string{}, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return []string{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []string{}, err
	}
	var challenge Challenge
	err = json.Unmarshal(body, &challenge)
	if err != nil {
		return []string{}, err
	}
	data, err := base64.StdEncoding.DecodeString(challenge.Challenge)
	if err != nil {
		return []string{}, err
	}
	return strings.Split(string(data), "\n"), nil
}

// Send target evaluations to head.
func sendEvaluations(evaluations []Evaluation) (Message, error) {
	client := http.Client{}
	evaluationsJSON, err := json.Marshal(evaluations)
	if err != nil {
		return Message{}, err
	}
	response := Response{
		Response: base64.StdEncoding.EncodeToString(evaluationsJSON),
	}
	responseJSON, err := json.Marshal(response)
	if err != nil {
		return Message{}, err
	}
	req, err := http.NewRequest(
		"POST",
		"http://"+endpoint+"/evaluations",
		bytes.NewReader(responseJSON),
	)
	if err != nil {
		return Message{}, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return Message{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Message{}, err
	}
	var result Result
	err = json.Unmarshal(body, &result)
	if err != nil {
		return Message{}, err
	}
	messageBytes, err := base64.StdEncoding.DecodeString(result.Message)
	var message Message
	err = json.Unmarshal(messageBytes, &message)
	if err != nil {
		return Message{}, err
	}
	return message, nil
}

// Print evaluation result.
func printResult(result Message) {
	log.Printf("\n"+
		"+------------------------------+\n"+
		"|  WEAPON VERIFICATION RESULT  |\n"+
		"+---------+-----------+--------+\n"+
		"|         | DESTROYED |     %2d |\n"+
		"| ENEMIES +-----------+--------+\n"+
		"|         | SPARED    |     %2d |\n"+
		"+---------+-----------+--------+\n"+
		"|         | DESTROYED |     %2d |\n"+
		"| ALLIES  +-----------+--------+\n"+
		"|         | SPARED    |     %2d |\n"+
		"+---------+-----------+--------+\n"+
		"| SCORE   |               %3d%% |\n"+
		"+---------+--------------------+\n"+
		"| SUCCESS |              %5s |\n"+
		"+---------+--------------------+\n"+
		"Message: %s\n",
		result.EnemiesDestroyed,
		result.EnemiesSpared,
		result.AlliesDestroyed,
		result.AlliesSpared,
		int(result.Score*100),
		strings.ToUpper(
			strconv.FormatBool(
				result.Success,
			),
		),
		result.Message,
	)
}
