package config

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
)

type PostgresInfo struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
}

type HttpServer struct {
	Port    string `json:"port"`
	Timeout int    `json:"timeout"`
}

type Config struct {
	PostgresInfo `json:"database"`
	HttpServer   `json:"server"`
}

func Init(path string) (*Config, error) {
	file, err := os.OpenFile(path, os.O_RDWR, 0o644)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(file)
	var cfg Config
	err = json.NewDecoder(bytes.NewBuffer(body)).Decode(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
