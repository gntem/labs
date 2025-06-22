package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type News struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"created_at"`
}

type NewsRepository interface {
	Connect() error
	GetAll() ([]News, error)
	GetByID(id int) (News, error)
	GetByCategory(category string) ([]News, error)
}

type MockNewsRepository struct {
	Name      string
	connected bool
	newsDB    []News
}

// Connect establishes connection to the database
func (repo *MockNewsRepository) Connect() error {
	if repo.connected {
		return nil
	}

	log.Debug().Str("db", repo.Name).Msg("connecting to database...")
	time.Sleep(1 * time.Second) // Simulate connection time

	// Initialize mock data
	repo.newsDB = []News{
		{
			ID:        1,
			Title:     "Title1",
			Content:   "Content1",
			Author:    "Author1",
			Category:  "technology",
			CreatedAt: time.Now().Add(-24 * time.Hour),
		},
		{
			ID:        2,
			Title:     "Title2",
			Content:   "Content2",
			Author:    "Author2",
			Category:  "programming",
			CreatedAt: time.Now().Add(-48 * time.Hour),
		},
		{
			ID:        3,
			Title:     "Title3",
			Content:   "Content3",
			Author:    "Author3",
			Category:  "technology",
			CreatedAt: time.Now().Add(-72 * time.Hour),
		},
	}

	repo.connected = true
	log.Info().Str("db", repo.Name).Msg("connected to database successfully")
	return nil
}

func (repo *MockNewsRepository) GetAll() ([]News, error) {
	if !repo.connected {
		return nil, fmt.Errorf("repository not connected")
	}
	return repo.newsDB, nil
}

func (repo *MockNewsRepository) GetByID(id int) (News, error) {
	if !repo.connected {
		return News{}, fmt.Errorf("repository not connected")
	}

	for _, news := range repo.newsDB {
		if news.ID == id {
			return news, nil
		}
	}

	return News{}, fmt.Errorf("news with ID %d not found", id)
}

func (repo *MockNewsRepository) GetByCategory(category string) ([]News, error) {
	if !repo.connected {
		return nil, fmt.Errorf("repository not connected")
	}

	var result []News
	for _, news := range repo.newsDB {
		if strings.EqualFold(news.Category, category) {
			result = append(result, news)
		}
	}

	return result, nil
}

type NewsService struct {
	Name string
	repo NewsRepository
}

// Start initializes the news service
func (s *NewsService) Start() error {
	log.Debug().Str("service", s.Name).Msg("starting service...")

	// Connect to the repository
	err := s.repo.Connect()
	if err != nil {
		return fmt.Errorf("failed to connect to repository: %w", err)
	}

	log.Info().Str("service", s.Name).Msg("service started successfully")
	return nil
}

func (s *NewsService) GetAllNews() ([]News, error) {
	resultChan := make(chan []News)
	errChan := make(chan error)

	go func() {
		news, err := s.repo.GetAll()
		if err != nil {
			errChan <- err
			return
		}
		resultChan <- news
	}()

	select {
	case news := <-resultChan:
		return news, nil
	case err := <-errChan:
		return nil, err
	case <-time.After(5 * time.Second):
		return nil, fmt.Errorf("operation timed out")
	}
}

func (s *NewsService) GetNewsByID(id int) (News, error) {
	resultChan := make(chan News)
	errChan := make(chan error)

	go func() {
		news, err := s.repo.GetByID(id)
		if err != nil {
			errChan <- err
			return
		}
		resultChan <- news
	}()

	select {
	case news := <-resultChan:
		return news, nil
	case err := <-errChan:
		return News{}, err
	case <-time.After(5 * time.Second):
		return News{}, fmt.Errorf("operation timed out")
	}
}

func (s *NewsService) GetNewsByCategory(category string) ([]News, error) {
	resultChan := make(chan []News)
	errChan := make(chan error)

	go func() {
		news, err := s.repo.GetByCategory(category)
		if err != nil {
			errChan <- err
			return
		}
		resultChan <- news
	}()

	select {
	case news := <-resultChan:
		return news, nil
	case err := <-errChan:
		return nil, err
	case <-time.After(5 * time.Second):
		return nil, fmt.Errorf("operation timed out")
	}
}

func main() {
	// Configure zerolog
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	// Initialize the service
	newsService := &NewsService{
		Name: "NewsService",
		repo: &MockNewsRepository{Name: "NewsDB"},
	}

	// Start the service
	err := newsService.Start()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start service")
	}

	// Set up API endpoints
	http.HandleFunc("/api/news", func(w http.ResponseWriter, r *http.Request) {
		news, err := newsService.GetAllNews()
		if err != nil {
			log.Error().Err(err).Msg("error getting all news")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(news)
	})

	http.HandleFunc("/api/news/", func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/api/news/")

		if strings.HasPrefix(path, "category/") {
			category := strings.TrimPrefix(path, "category/")
			news, err := newsService.GetNewsByCategory(category)
			if err != nil {
				log.Error().Err(err).Str("category", category).Msg("error getting news by category")
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(news)
			return
		}

		id, err := strconv.Atoi(path)
		if err != nil {
			log.Error().Err(err).Str("path", path).Msg("invalid news id")
			http.Error(w, "Invalid news ID", http.StatusBadRequest)
			return
		}

		news, err := newsService.GetNewsByID(id)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				log.Error().Err(err).Int("id", id).Msg("news not found")
				http.Error(w, err.Error(), http.StatusNotFound)
			} else {
				log.Error().Err(err).Int("id", id).Msg("error getting news by id")
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(news)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		fmt.Fprintln(w, "OK")
	})

	port := "8080"
	log.Info().Str("port", port).Msg("web server running on http://localhost:" + port)
	log.Fatal().Err(http.ListenAndServe(":"+port, nil)).Msg("server stopped")
}
