package database

import (
	"log"
	"os"
	"testing"

	"github.com/harisfi/alterra-agmc/day4/submission/configs"
	"github.com/harisfi/alterra-agmc/day4/submission/models"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestGetAllUsers(t *testing.T) {
	var testCases = []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "success_get_all_users",
			wantErr: false,
		},
		{
			name:    "failed_get_all_users",
			wantErr: true,
		},
	}

	if err := godotenv.Load("../../.env"); err != nil {
		log.Println(err)
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.wantErr {
				os.Clearenv()
				assert.Panics(t, configs.InitDB)
			} else {
				configs.InitDB()
			}

			u, e := GetAllUsers()

			if tc.wantErr {
				assert.Error(t, e)
			} else {
				assert.NoError(t, e)
				assert.IsType(t, []models.User{}, u)
			}
		})
	}
}

func TestGetUserById(t *testing.T) {
	var testCases = []struct {
		name    string
		userId  uint
		wantErr bool
		errDB   bool
	}{
		{
			name:    "success_get_user_by_id",
			userId:  1,
			wantErr: false,
		},
		{
			name:    "failed_get_user_by_id_false_id",
			userId:  2,
			wantErr: true,
		},
		{
			name:    "failed_get_user_by_id_db_err",
			userId:  1,
			wantErr: true,
			errDB:   true,
		},
	}
	userLoaded := false

	if err := godotenv.Load("../../.env"); err != nil {
		log.Println(err)
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.wantErr && tc.errDB {
				os.Clearenv()
				assert.Panics(t, configs.InitDB)
			} else {
				configs.InitDB()
				if !userLoaded {
					configs.DB.Exec("TRUNCATE TABLE users;")
					configs.DB.Exec("INSERT INTO users (name,email,password) VALUES ('a','b','c');")
					userLoaded = true
				}
			}

			u, e := GetUserById(tc.userId)

			if tc.wantErr {
				assert.Error(t, e)
			} else {
				assert.NoError(t, e)
				assert.IsType(t, models.User{}, u)
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	var testCases = []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "success_create_user",
			wantErr: false,
		},
		{
			name:    "failed_create_user",
			wantErr: true,
		},
	}

	if err := godotenv.Load("../../.env"); err != nil {
		log.Println(err)
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.wantErr {
				os.Clearenv()
				assert.Panics(t, configs.InitDB)
			} else {
				configs.InitDB()
			}

			u, e := CreateUser(models.User{})

			if tc.wantErr {
				assert.Error(t, e)
			} else {
				assert.NoError(t, e)
				assert.IsType(t, models.User{}, u)
			}
		})
	}
}

func TestUpdateUserById(t *testing.T) {
	var testCases = []struct {
		name    string
		userId  uint
		wantErr bool
		errDB   bool
	}{
		{
			name:    "success_update_user_by_id",
			userId:  1,
			wantErr: false,
		},
		{
			name:    "failed_update_user_by_id_false_id",
			userId:  2,
			wantErr: true,
		},
		{
			name:    "failed_update_user_by_id_db_err",
			userId:  1,
			wantErr: true,
			errDB:   true,
		},
	}
	userLoaded := false

	if err := godotenv.Load("../../.env"); err != nil {
		log.Println(err)
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.wantErr && tc.errDB {
				os.Clearenv()
				assert.Panics(t, configs.InitDB)
			} else {
				configs.InitDB()
				if !userLoaded {
					configs.DB.Exec("TRUNCATE TABLE users;")
					configs.DB.Exec("INSERT INTO users (name,email,password) VALUES ('a','b','c');")
					userLoaded = true
				}
			}

			u, e := UpdateUserById(tc.userId, models.User{})

			if tc.wantErr {
				assert.Error(t, e)
			} else {
				assert.NoError(t, e)
				assert.IsType(t, models.User{}, u)
			}
		})
	}
}

func TestDeleteUserById(t *testing.T) {
	var testCases = []struct {
		name    string
		userId  uint
		wantErr bool
		errDB   bool
	}{
		{
			name:    "success_delete_user_by_id",
			userId:  1,
			wantErr: false,
		},
		{
			name:    "failed_delete_user_by_id_false_id",
			userId:  2,
			wantErr: true,
		},
		{
			name:    "failed_delete_user_by_id_db_err",
			userId:  1,
			wantErr: true,
			errDB:   true,
		},
	}
	userLoaded := false

	if err := godotenv.Load("../../.env"); err != nil {
		log.Println(err)
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.wantErr && tc.errDB {
				os.Clearenv()
				assert.Panics(t, configs.InitDB)
			} else {
				configs.InitDB()
				if !userLoaded {
					configs.DB.Exec("TRUNCATE TABLE users;")
					configs.DB.Exec("INSERT INTO users (name,email,password) VALUES ('a','b','c');")
					userLoaded = true
				}
			}

			e := DeleteUserById(tc.userId)

			if tc.wantErr {
				assert.Error(t, e)
			} else {
				assert.NoError(t, e)
			}
		})
	}
}
