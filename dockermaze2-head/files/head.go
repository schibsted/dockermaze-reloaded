package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

type Challenge struct {
	ID        int    `json:"id"`
	Challenge string `json:"challenge"`
}

type Response struct {
	ID       int    `json:"id"`
	Response string `json:"response"`
}

type Result struct {
	Success   bool    `json:"success"`
	ScoreRate float64 `json:"score_rate"`
	Message   string  `json:"message"`
}

const (
	port = ":7777"
)

var token string
var endpoint string

func init() {
	endpoint = os.Getenv("DM2_ENDPOINT")
}

func getChallenge() (Challenge, error) {
	client := http.Client{}
	req, _ := http.NewRequest(
		"GET",
		"http://"+endpoint+"/challenge/genByGame/Arms",
		nil,
	)
	req.Header.Set("User-Token", token)
	resp, err := client.Do(req)
	if err != nil {
		return Challenge{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return Challenge{}, errors.New("The backend returned an error.")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Challenge{}, err
	}
	var challenge Challenge
	err = json.Unmarshal(body, &challenge)
	if err != nil {
		return Challenge{}, err
	}
	return challenge, nil
}

func postChallenge(challenge Challenge, responseBytes []byte) (Result, error) {
	response := Response{
		challenge.ID,
		base64.StdEncoding.EncodeToString(responseBytes),
	}
	data, _ := json.Marshal(response)
	client := http.Client{}
	req, _ := http.NewRequest(
		"POST",
		"http://"+endpoint+"/challenge/result",
		bytes.NewReader(data),
	)
	req.Header.Set("User-Token", token)
	resp, err := client.Do(req)
	if err != nil {
		return Result{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return Result{}, errors.New("The backend returned an error.")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Result{}, err
	}
	var result Result
	if err = json.Unmarshal(body, &result); err != nil {
		return Result{}, err
	}
	return result, nil
}

func handleRequest(conn net.Conn) {
	challenge, err := getChallenge()
	if err != nil {
		conn.Write([]byte("ERROR"))
		log.Println(err)
		return
	}
	challengePlain, _ := base64.StdEncoding.DecodeString(challenge.Challenge)
	conn.Write(challengePlain)
	buf := bufio.NewReader(conn)
	var responses []byte
	for {
		response, err := buf.ReadBytes('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			conn.Write([]byte("ERROR"))
			log.Println(err)
			return
		} else if strings.Contains(string(response), "EOF") {
			break
		} else {
			responses = append(responses, response...)
		}
	}
	if err != nil {
		conn.Write([]byte("ERROR"))
		log.Println(err)
		return
	}
	result, err := postChallenge(challenge, responses)
	if err != nil {
		conn.Write([]byte("ERROR"))
		log.Println(err)
		return
	}
	message, _ := base64.StdEncoding.DecodeString(result.Message)
	conn.Write(append(message, '\n'))
}

func main() {
	token = os.Getenv("DM2_TOKEN")
	srv, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	defer srv.Close()
	for {
		conn, err := srv.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleRequest(conn)
	}
}
