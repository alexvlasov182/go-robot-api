package models

// Robot represent a robot for testing payment terminals
type Robot struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Status      string   `json:"status"`
	LastTested  string   `json:"last_tested"`
	TestResults []string `json:"test_results"`
}
