package client

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"sync"
)

type APIVersion struct {
	major uint8
	date  string
}

func GetAPIVersion() APIVersion {
	return APIVersion{7, "20230615"}
}

func GetMinAPIVersion() APIVersion {
	return APIVersion{5, "20191215"}
}

var DEFAULT_CHUNK_SIZE = uint32(16 * math.Pow(2, 20)) // 16 MiB
const MAX_INFLIGHT_CHUNKS = 4

func ParseAPIVersion(value string) (APIVersion, error) {
	// Define the regular expression pattern
	re := regexp.MustCompile(`v(?P<major>\d+)\.(?P<date>\d{8})$`)

	// Find submatches in the input string
	matches := re.FindStringSubmatch(value)
	if matches == nil {
		return APIVersion{}, fmt.Errorf("Could not parse the given API version string: %s", value)
	}

	// Extract named submatches
	majorStr := matches[re.SubexpIndex("major")]
	date := matches[re.SubexpIndex("date")]

	// Convert major version to an integer
	major, err := strconv.Atoi(majorStr)
	if err != nil {
		return APIVersion{}, fmt.Errorf("Failed to convert major version to integer: %v", err)
	}

	return APIVersion{uint8(major), date}, nil
}

type APIConfig struct{}

var mu = &sync.Mutex{}

var config *APIConfig

func GetConfig() *APIConfig {
	if config == nil {
		mu.Lock()
		defer mu.Unlock()
		if config == nil {
			config = &APIConfig{}
		}
	}
	return config
}

// var once sync.Once
// func GetLazyInitConfig() *APIConfig {
// 	once.Do(func() {
// 		// Initialize the singleton only once
// 		config = &APIConfig{}
// 	})
// 	return config
// }

func SetConfig(newValue APIConfig) {
	mu.Lock()
	defer mu.Unlock()

	// Critical section
	config = &newValue
}

/**
var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

func getInstance() *single {
    if singleInstance == nil {
        lock.Lock()
        defer lock.Unlock()
        if singleInstance == nil {
            fmt.Println("Creating single instance now.")
            singleInstance = &single{}
        } else {
            fmt.Println("Single instance already created.")
        }
    } else {
        fmt.Println("Single instance already created.")
    }

    return singleInstance
}
*/
