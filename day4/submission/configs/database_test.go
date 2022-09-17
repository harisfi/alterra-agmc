package configs

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestInitDB(t *testing.T) {
	var testCases = []struct {
		name         string
		expectOutput string
		wantErr      bool
	}{
		{
			name:         "success_connect_database",
			expectOutput: "connected to database\n",
			wantErr:      false,
		},
		{
			name:         "failed_connect_database",
			expectOutput: "failed to connect database",
			wantErr:      true,
		},
	}

	err := godotenv.Load("../.env")
	if err != nil {
		log.Println(err)
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var buff bytes.Buffer
			log.SetOutput(&buff)

			defer func() {
				log.SetOutput(os.Stdout)
			}()

			if tc.wantErr {
				os.Setenv("MYSQL_HOST", "127.0.0.0")
				assert.PanicsWithValue(t, tc.expectOutput, InitDB)
			} else {
				InitDB()
				assert.True(t, strings.HasSuffix(buff.String(), tc.expectOutput))
			}
		})
	}
}
