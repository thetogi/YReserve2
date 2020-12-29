package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/thetogi/YReserve2/model"
)

func TestCreateUser(t *testing.T) {
	t.Log("Starting create user test case")
	apiTest := GetApiTest()

	user := GetTestUser()
	userAuth := apiTest.CreateUserAuthFromTestAPI(t, apiTest.API, user)
	apiTest.CheckValidTestUser(t, user, userAuth.User)

	if userAuth.Token == "" {
		t.Errorf("handler returned wrong token: got %v",
			userAuth.Token)
	}
}

func TestUpdateUser(t *testing.T) {
	t.Log("Starting create user test case")
	apiTest := GetApiTest()

	user := GetTestUser()
	userAuth := apiTest.CreateUserAuthFromTestAPI(t, apiTest.API, user)
	apiTest.CheckValidTestUser(t, user, userAuth.User)

	if userAuth.Token == "" {
		t.Errorf("handler returned wrong token: got %v",
			userAuth.Token)
	}

	expectedUser := userAuth.User
	t.Log(expectedUser)
	expectedUser.FirstName = "NewTest"
	expectedUser.LastName = "NewTest"
	jsonUser, _ := json.Marshal(expectedUser)
	req, err := http.NewRequest("PUT", fmt.Sprintf("/api/v1/user/%d", expectedUser.UserID), bytes.NewBuffer(jsonUser))
	req.Header.Set(model.AUTHENTICATION, userAuth.Token)
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()
	handler := apiTest.API.requestWithAuthHandler(apiTest.API.updateUser)
	handler.ServeHTTP(res, req)

	CheckOkStatus(t, res.Code)

	t.Log(res.Body.String())
	receivedUser := model.UserFromString(res.Body.String())

	if receivedUser.UserID != expectedUser.UserID {
		t.Errorf("handler returned wrong name: got %d expected %d",
			receivedUser.UserID, expectedUser.UserID)
	}

	if receivedUser.FirstName != expectedUser.FirstName {
		t.Errorf("handler returned wrong name: got %s expected %s",
			receivedUser.FirstName, expectedUser.FirstName)
	}

	if receivedUser.LastName != expectedUser.LastName {
		t.Errorf("handler returned wrong name: got %s expected %s",
			receivedUser.LastName, expectedUser.LastName)
	}
}

func TestGetUser(t *testing.T) {
	t.Log("Starting get user test case")
	apiTest := GetApiTest()

	user := GetTestUser()
	userAuth := apiTest.CreateUserAuthFromTestAPI(t, apiTest.API, user)
	apiTest.CheckValidTestUser(t, user, userAuth.User)

	if userAuth.Token == "" {
		t.Errorf("handler returned wrong token: got %v",
			userAuth.Token)
	}

	expectedUser := userAuth.User

	t.Log(expectedUser)
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/v1/user/%d", expectedUser.UserID), nil)
	req.Header.Set(model.AUTHENTICATION, userAuth.Token)
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()
	handler := apiTest.API.requestWithAuthHandler(apiTest.API.getUser)
	handler.ServeHTTP(res, req)

	CheckOkStatus(t, res.Code)
	t.Log(res.Body.String())

	receivedUser := model.UserFromString(res.Body.String())

	if receivedUser.UserID != expectedUser.UserID {
		t.Errorf("handler returned wrong name: got %d expected %d",
			receivedUser.UserID, user.UserID)
	}
}

func TestDeleteUser(t *testing.T) {
	t.Log("Starting delete user test case")
	apiTest := GetApiTest()

	user := GetTestUser()
	userAuth := apiTest.CreateUserAuthFromTestAPI(t, apiTest.API, user)
	apiTest.CheckValidTestUser(t, user, userAuth.User)

	if userAuth.Token == "" {
		t.Errorf("handler returned wrong token: got %v",
			userAuth.Token)
	}

	expectedUser := userAuth.User

	t.Log(expectedUser)
	req, err := http.NewRequest("DELETE", fmt.Sprintf("/api/v1/user/%d", expectedUser.UserID), nil)
	req.Header.Set(model.AUTHENTICATION, userAuth.Token)
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()
	handler := apiTest.API.requestWithAuthHandler(apiTest.API.deleteUser)
	handler.ServeHTTP(res, req)

	CheckOkStatus(t, res.Code)

	if res.Body.String() != "{'response': 'OK'}" {
		t.Errorf("handler returned wrong name: got %s expected %s",
			res.Body.String(), "{'response': 'OK'}")
	}
}
