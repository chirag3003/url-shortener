package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/chirag3003/go-backend-template/dto/response"
	mw "github.com/chirag3003/go-backend-template/middleware"
	"github.com/chirag3003/go-backend-template/models"
	"github.com/chirag3003/go-backend-template/pkg/auth"
	"github.com/chirag3003/go-backend-template/pkg/idgen"
	"github.com/chirag3003/go-backend-template/repository/mock"
	"github.com/chirag3003/go-backend-template/service"
	"github.com/gofiber/fiber/v3"
	"github.com/rs/zerolog"
	"go.uber.org/mock/gomock"
)

type testIDGen struct {
	next int64
}

func (g *testIDGen) NewID() (int64, error) {
	id := g.next
	g.next++
	return id, nil
}

var _ idgen.Generator = (*testIDGen)(nil)

// testApp creates a Fiber app wired with gomock repositories for integration testing.
func testApp(t *testing.T) (*fiber.App, *mock.MockUserRepository, *auth.JWTService) {
	ctrl := gomock.NewController(t)
	log := zerolog.New(os.Stderr).Level(zerolog.Disabled)
	jwtService := auth.NewJWTService("test-secret", 1*time.Hour)
	userRepo := mock.NewMockUserRepository(ctrl)

	idgen.SetDefault(&testIDGen{next: 1000})
	authService := service.NewAuthService(userRepo, jwtService, log)
	userService := service.NewUserService(userRepo, log)

	authCtrl := NewAuthController(authService)
	userCtrl := NewUserController(userService)

	app := fiber.New(fiber.Config{
		ErrorHandler: mw.ErrorHandler(log),
	})

	// Auth routes (public)
	authGroup := app.Group("/api/v1/auth")
	authGroup.Post("/login", authCtrl.Login)
	authGroup.Post("/register", authCtrl.Register)

	// User routes (protected)
	authMiddleware := mw.Auth(jwtService)
	userGroup := app.Group("/api/v1/user", authMiddleware)
	userGroup.Get("/me", userCtrl.GetMe)

	return app, userRepo, jwtService
}

// readJSON reads and parses the JSON response body.
func readJSON(t *testing.T, resp *http.Response) *response.APIResponse {
	t.Helper()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}
	defer resp.Body.Close()

	var apiResp response.APIResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		t.Fatalf("failed to parse JSON response: %v\nbody: %s", err, string(body))
	}
	return &apiResp
}

func jsonBody(t *testing.T, v interface{}) io.Reader {
	t.Helper()
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("failed to marshal JSON: %v", err)
	}
	return bytes.NewReader(b)
}

// --- Register Endpoint Tests ---

func TestRegisterEndpoint_Success(t *testing.T) {
	app, userRepo, _ := testApp(t)

	userRepo.EXPECT().
		GetUserByEmail(gomock.Any(), "alice@example.com").
		Return(nil, nil)
	userRepo.EXPECT().
		CreateUser(gomock.Any(), gomock.Any()).
		Return(nil)

	req, _ := http.NewRequest("POST", "/api/v1/auth/register", jsonBody(t, map[string]string{
		"name":     "Alice",
		"email":    "alice@example.com",
		"password": "securepass123",
	}))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		t.Fatalf("expected status 201, got %d, body: %s", resp.StatusCode, string(body))
	}

	apiResp := readJSON(t, resp)
	if !apiResp.Success {
		t.Fatal("expected success to be true")
	}
}

func TestRegisterEndpoint_DuplicateEmail(t *testing.T) {
	app, userRepo, _ := testApp(t)

	userRepo.EXPECT().
		GetUserByEmail(gomock.Any(), "alice@example.com").
		Return(&models.User{
			ID:    int64(1),
			Name:  "Alice",
			Email: "alice@example.com",
		}, nil)

	req, _ := http.NewRequest("POST", "/api/v1/auth/register", jsonBody(t, map[string]string{
		"name":     "Alice 2",
		"email":    "alice@example.com",
		"password": "anotherpass123",
	}))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	if resp.StatusCode != http.StatusConflict {
		t.Fatalf("expected status 409, got %d", resp.StatusCode)
	}

	apiResp := readJSON(t, resp)
	if apiResp.Success {
		t.Fatal("expected success to be false")
	}
	if apiResp.Error == nil || apiResp.Error.Code != "USER_ALREADY_EXISTS" {
		t.Fatalf("expected error code USER_ALREADY_EXISTS, got %+v", apiResp.Error)
	}
}

func TestRegisterEndpoint_ValidationError(t *testing.T) {
	app, _, _ := testApp(t)

	req, _ := http.NewRequest("POST", "/api/v1/auth/register", jsonBody(t, map[string]string{
		"name":     "",
		"email":    "bad-email",
		"password": "short",
	}))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)
	if resp.StatusCode != http.StatusUnprocessableEntity {
		body, _ := io.ReadAll(resp.Body)
		t.Fatalf("expected status 422, got %d, body: %s", resp.StatusCode, string(body))
	}

	apiResp := readJSON(t, resp)
	if apiResp.Success {
		t.Fatal("expected success to be false")
	}
	if apiResp.Error == nil || apiResp.Error.Code != "VALIDATION_ERROR" {
		t.Fatalf("expected error code VALIDATION_ERROR, got %+v", apiResp.Error)
	}
}

func TestRegisterEndpoint_InvalidJSON(t *testing.T) {
	app, _, _ := testApp(t)

	req, _ := http.NewRequest("POST", "/api/v1/auth/register", bytes.NewReader([]byte("not json")))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)
	if resp.StatusCode != http.StatusBadRequest {
		body, _ := io.ReadAll(resp.Body)
		t.Fatalf("expected status 400, got %d, body: %s", resp.StatusCode, string(body))
	}
}

// --- Login Endpoint Tests ---

func TestLoginEndpoint_Success(t *testing.T) {
	app, userRepo, _ := testApp(t)

	hash, _ := auth.HashPassword("securepass123")
	userID := int64(2001)

	userRepo.EXPECT().
		GetUserByEmail(gomock.Any(), "bob@example.com").
		Return(&models.User{
			ID:    userID,
			Name:  "Bob",
			Email: "bob@example.com",
			Hash:  hash,
		}, nil)

	req, _ := http.NewRequest("POST", "/api/v1/auth/login", jsonBody(t, map[string]string{
		"email":    "bob@example.com",
		"password": "securepass123",
	}))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		t.Fatalf("expected status 200, got %d, body: %s", resp.StatusCode, string(body))
	}

	apiResp := readJSON(t, resp)
	if !apiResp.Success {
		t.Fatal("expected success to be true")
	}

	data, ok := apiResp.Data.(map[string]interface{})
	if !ok {
		t.Fatal("expected data to be a map")
	}
	if data["token"] == nil || data["token"] == "" {
		t.Fatal("expected non-empty token in response")
	}
	user, ok := data["user"].(map[string]interface{})
	if !ok {
		t.Fatal("expected user object in response")
	}
	if user["email"] != "bob@example.com" {
		t.Fatalf("expected email bob@example.com, got %v", user["email"])
	}
}

func TestLoginEndpoint_WrongPassword(t *testing.T) {
	app, userRepo, _ := testApp(t)

	hash, _ := auth.HashPassword("securepass123")

	userRepo.EXPECT().
		GetUserByEmail(gomock.Any(), "bob@example.com").
		Return(&models.User{
			ID:    int64(2002),
			Name:  "Bob",
			Email: "bob@example.com",
			Hash:  hash,
		}, nil)

	req, _ := http.NewRequest("POST", "/api/v1/auth/login", jsonBody(t, map[string]string{
		"email":    "bob@example.com",
		"password": "wrongpassword",
	}))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	if resp.StatusCode != http.StatusUnauthorized {
		t.Fatalf("expected status 401, got %d", resp.StatusCode)
	}

	apiResp := readJSON(t, resp)
	if apiResp.Success {
		t.Fatal("expected success to be false")
	}
	if apiResp.Error == nil || apiResp.Error.Code != "INVALID_CREDENTIALS" {
		t.Fatalf("expected error code INVALID_CREDENTIALS, got %+v", apiResp.Error)
	}
}

func TestLoginEndpoint_NonExistentUser(t *testing.T) {
	app, userRepo, _ := testApp(t)

	userRepo.EXPECT().
		GetUserByEmail(gomock.Any(), "nobody@example.com").
		Return(nil, nil)

	req, _ := http.NewRequest("POST", "/api/v1/auth/login", jsonBody(t, map[string]string{
		"email":    "nobody@example.com",
		"password": "securepass123",
	}))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	if resp.StatusCode != http.StatusUnauthorized {
		t.Fatalf("expected status 401, got %d", resp.StatusCode)
	}
}

// --- User Endpoint Tests ---

func TestGetMeEndpoint_Success(t *testing.T) {
	app, userRepo, jwtService := testApp(t)

	userID := int64(3001)

	// Generate a valid token
	token, err := jwtService.GenerateToken(strconv.FormatInt(userID, 10), "Charlie", "charlie@example.com", "")
	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}

	// Expect the GetByID call from the user service
	userRepo.EXPECT().
		GetUserByID(gomock.Any(), strconv.FormatInt(userID, 10)).
		Return(&models.User{
			ID:    userID,
			Name:  "Charlie",
			Email: "charlie@example.com",
		}, nil)

	req, _ := http.NewRequest("GET", "/api/v1/user/me", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		t.Fatalf("expected status 200, got %d, body: %s", resp.StatusCode, string(body))
	}

	apiResp := readJSON(t, resp)
	if !apiResp.Success {
		t.Fatal("expected success to be true")
	}

	data := apiResp.Data.(map[string]interface{})
	if data["name"] != "Charlie" {
		t.Fatalf("expected name Charlie, got %v", data["name"])
	}
	if data["email"] != "charlie@example.com" {
		t.Fatalf("expected email charlie@example.com, got %v", data["email"])
	}
}

func TestGetMeEndpoint_NoAuthHeader(t *testing.T) {
	app, _, _ := testApp(t)

	req, _ := http.NewRequest("GET", "/api/v1/user/me", nil)
	resp, _ := app.Test(req)

	if resp.StatusCode != http.StatusUnauthorized {
		t.Fatalf("expected status 401, got %d", resp.StatusCode)
	}

	apiResp := readJSON(t, resp)
	if apiResp.Success {
		t.Fatal("expected success to be false")
	}
	if apiResp.Error == nil || apiResp.Error.Code != "UNAUTHORIZED" {
		t.Fatalf("expected error code UNAUTHORIZED, got %+v", apiResp.Error)
	}
}

func TestGetMeEndpoint_InvalidToken(t *testing.T) {
	app, _, _ := testApp(t)

	req, _ := http.NewRequest("GET", "/api/v1/user/me", nil)
	req.Header.Set("Authorization", "Bearer invalid-token-here")
	resp, _ := app.Test(req)

	if resp.StatusCode != http.StatusUnauthorized {
		t.Fatalf("expected status 401, got %d", resp.StatusCode)
	}
}

func TestGetMeEndpoint_ExpiredOrWrongSecretToken(t *testing.T) {
	app, _, _ := testApp(t)

	// Generate a token with a different secret
	wrongJWT := auth.NewJWTService("wrong-secret", 1*time.Hour)
	token, _ := wrongJWT.GenerateToken("some-id", "name", "email@test.com", "")

	req, _ := http.NewRequest("GET", "/api/v1/user/me", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	resp, _ := app.Test(req)

	if resp.StatusCode != http.StatusUnauthorized {
		t.Fatalf("expected status 401, got %d", resp.StatusCode)
	}
}

// --- 404 Test ---

func TestNotFoundRoute(t *testing.T) {
	app, _, _ := testApp(t)

	req, _ := http.NewRequest("GET", "/api/v1/nonexistent", nil)
	resp, _ := app.Test(req)

	if resp.StatusCode != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d", resp.StatusCode)
	}
}
