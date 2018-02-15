package score

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var filePath = "scores.json"

type Score struct {
	Value int `json:"value"`
}

func init() {
	_, err := os.Stat(filePath)
	if err != nil {
		scores := []Score{}
		os.Create(filePath)
		data, _ := json.Marshal(&scores)
		ioutil.WriteFile(filePath, data, 0644)
    }
}

func Create(value int) {
	score := Score{value}
	scores := append(All(), score)
	data, _ := json.Marshal(&scores)
	ioutil.WriteFile(filePath, data, 0644)
}

func All() (scores []Score) {
	file, _ := os.Open(filePath)
	defer file.Close()
	data, _ := ioutil.ReadAll(file)
	json.Unmarshal(data, &scores)
	return
}

func Delete(num int) {
	scores := []Score{}
	for i, score := range All() {
		if i != num - 1 {
			scores = append(scores, score)
		}
	}
	data, _ := json.Marshal(&scores)
	ioutil.WriteFile(filePath, data, 0644)
}

func Average() (average int) {
	scores := All()
	len := len(scores)
	if len == 0 {
		return
	}
	sum := 0
	for _, score := range scores {
		sum += score.Value
	}
	average = sum / len
	return
}
