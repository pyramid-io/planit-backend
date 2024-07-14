package session

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

type FileSystemSessionDriver struct {
	dir string
	defaultTTL time.Duration
}

func NewFileSystemSessionDriver(config map[string]interface{}) (SessionDriverInterface, error) {
	dir, ok := config["dir"]
	if (!ok) {
		log.Fatalf("parameter `dir` is missing for file system session driver construction")
	}
	directory, ok := dir.(string)
	if (!ok) {
		log.Fatalf("parameter `dir` wrong format")
	}


	defaultTTL, ok := config["defaultTTL"]
	if (!ok) {
		log.Fatalf("parameter `default TTL` is missing for file system session driver construction")
	}
	TTL, ok := defaultTTL.(time.Duration)
	if (!ok) {
		log.Fatalf("parameter `default TTL` worng format")
	}

	if err := os.MkdirAll(directory, 0755); err != nil {
		log.Fatalf("Error creating directory: %v", err)
	}
	return &FileSystemSessionDriver{
		dir: directory,
		defaultTTL: TTL,
	}, nil
}

func (driver *FileSystemSessionDriver) Create(key string, data map[string]interface{}, expiresAt *time.Time) (*Session, error) {

	if (expiresAt == nil) {
		ttlBasedExpiresAt := time.Now().Add(driver.defaultTTL)
		expiresAt = &ttlBasedExpiresAt
	}
	
	session := &Session{
		ID: key, 
		Data: data,
		ExpiresAt: expiresAt,
	}
	
	jsonData, err := json.Marshal(session)
	if err != nil {
		log.Fatalf("Error marshaling struct to JSON: %v", err)
	}
	fmt.Println(driver.getPath(key))
    os.WriteFile(driver.getPath(key), jsonData, 0644)


	_, err = os.Stat(driver.getPath(key))
	if err != nil {
		log.Fatalf("Error stating file: %v", err)
	}
	
	return session, nil
}

func (driver *FileSystemSessionDriver) Get(key string) (*Session, error) {

    data, err := os.ReadFile(driver.getPath(key))
    if err != nil {
        return nil, err
    }

	session := &Session{}
	if err := json.Unmarshal(data, session); err != nil {
        return nil, err
    }

	now := time.Now().UTC()

	if session.ExpiresAt.Before(now) {
		return nil, errors.New("session is expired") 
	}    
    return session, nil
}

func (driver *FileSystemSessionDriver) Delete(key string) (error) {
	err := os.Remove(driver.getPath(key))
    if err != nil {
        return err
    }

	return nil
}

func (driver *FileSystemSessionDriver) getPath(key string) string {
	return filepath.Join(driver.dir, key+".json")
}
