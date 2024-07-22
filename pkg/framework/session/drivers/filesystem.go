package drivers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/pyramid.io/planit-backend/pkg/framework/session/session_interfaces"
)

type FileSystemSessionDriver struct {
	dir string
	defaultTTL time.Duration
}

type Session struct {
	id        string
	data      map[string]interface{}
	expiresAt *time.Time
}

func (s *Session) GetID() string {
	return s.id
}

func (s *Session) GetData() map[string]interface{} {
	return s.data
}

func (s *Session) GetExpiresAt() *time.Time {
	return s.expiresAt
}

func NewFileSystemSessionDriver(config map[string]interface{}) (session_interfaces.SessionDriverInterface, error) {
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

func (driver *FileSystemSessionDriver) Create(key string, data map[string]interface{}, expiresAt *time.Time) (session_interfaces.SessionInterface, error) {

	if (expiresAt == nil) {
		ttlBasedExpiresAt := time.Now().Add(driver.defaultTTL)
		expiresAt = &ttlBasedExpiresAt
	}
	
	session := &Session{
		id: key, 
		data: data,
		expiresAt: expiresAt,
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

func (driver *FileSystemSessionDriver) Get(key string) (session_interfaces.SessionInterface, error) {

    data, err := os.ReadFile(driver.getPath(key))
    if err != nil {
        return nil, err
    }

	session := &Session{}
	if err := json.Unmarshal(data, session); err != nil {
        return nil, err
    }

	now := time.Now().UTC()

	if session.expiresAt.Before(now) {
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
