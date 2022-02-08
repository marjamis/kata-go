package serve

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func requestGeneration(method string, path string, body io.Reader) (*httptest.ResponseRecorder, error) {
	router := createMuxRouter()

	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		return nil, err
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	return rr, nil
}
func TestProductHandler(t *testing.T) {
	t.Run("Valid product key provided", func(t *testing.T) {
		rr, err := requestGeneration("Get", "/products/5", nil)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Contains(t, rr.Body.String(), "Product: 5\n")
	})

	t.Run("Product key not provided", func(t *testing.T) {
		rr, err := requestGeneration("Get", "/products", nil)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, http.StatusNotFound, rr.Code)
	})
}

func TestArticlesCategoryHandler(t *testing.T) {
	t.Run("Valid article category provided", func(t *testing.T) {
		rr, err := requestGeneration("Get", "/articles/5", nil)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Contains(t, rr.Body.String(), "Category: 5\n")
	})

	t.Run("Article category not provided", func(t *testing.T) {
		rr, err := requestGeneration("GET", "/articles", nil)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, http.StatusNotFound, rr.Code)
	})
}

func TestArticleHandler(t *testing.T) {
	t.Run("Valid article for category provided", func(t *testing.T) {
		rr, err := requestGeneration("GET", "/articles/5/2", nil)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Contains(t, rr.Body.String(), "Category: 5 - Article: 2\n")
	})

	t.Run("No article for category provided", func(t *testing.T) {
		rr, err := requestGeneration("GET", "/articles/5/", nil)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, http.StatusNotFound, rr.Code)
	})

	t.Run("Invalid article for category provided", func(t *testing.T) {
		rr, err := requestGeneration("GET", "/articles/5/a", nil)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, http.StatusNotFound, rr.Code)
	})
}
