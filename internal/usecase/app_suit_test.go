package usecase

import (
	"context"
	"testing"

	"github.com/dostonshernazarov/doctor-appointment/config"
	"github.com/dostonshernazarov/doctor-appointment/internal/entity"
	"github.com/dostonshernazarov/doctor-appointment/internal/repo/persistent"
	"github.com/dostonshernazarov/doctor-appointment/internal/usecase/common"
	"github.com/dostonshernazarov/doctor-appointment/pkg/postgres"
	"github.com/stretchr/testify/assert"
)

func TestUserPostgres(t *testing.T) {
	// Load config
	config, err := config.NewConfig()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	// Connect to DB
	pg, err := postgres.New(config.PG.URL, postgres.MaxPoolSize(config.PG.PoolMax))
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer pg.Close()

	// Create usecase
	usecase := common.NewUseCase(
		persistent.NewUser(pg),
		persistent.NewDoctor(pg),
		persistent.NewAppointment(pg),
	)

	// Test create user
	user := entity.User{
		FullName: "John Doe",
		Email:    "john.doe@example.com",
		Phone:    "1234567890",
		Password: "hashedpassword",
		Role:     entity.RoleUser,
	}

	id, err := usecase.CreateUser(context.Background(), user)
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	assert.NotZero(t, id)

	// Test get user by id
	user, err = usecase.GetUserByID(context.Background(), id)
	if err != nil {
		t.Fatalf("Failed to get user by id: %v", err)
	}

	assert.Equal(t, user.ID, id)
	assert.Equal(t, user.FullName, "John Doe")
	assert.Equal(t, user.Email, "john.doe@example.com")
	assert.Equal(t, user.Phone, "1234567890")
	assert.Equal(t, user.Password, "hashedpassword")
	assert.Equal(t, user.Role, entity.RoleUser)

	// Test get user by email
	user, err = usecase.GetUserByEmail(context.Background(), "john.doe@example.com")
	if err != nil {
		t.Fatalf("Failed to get user by email: %v", err)
	}

	assert.Equal(t, user.ID, id)
	assert.Equal(t, user.FullName, "John Doe")
	assert.Equal(t, user.Email, "john.doe@example.com")
	assert.Equal(t, user.Phone, "1234567890")
	assert.Equal(t, user.Password, "hashedpassword")
	assert.Equal(t, user.Role, entity.RoleUser)

	// Test list users
	users, err := usecase.ListUsers(context.Background())
	if err != nil {
		t.Fatalf("Failed to list users: %v", err)
	}

	assert.NotEmpty(t, users)

	// Test update user
	user.FullName = "Jane Doe"
	err = usecase.UpdateUser(context.Background(), user)
	if err != nil {
		t.Fatalf("Failed to update user: %v", err)
	}

	// Test delete user
	err = usecase.DeleteUser(context.Background(), id)
	if err != nil {
		t.Fatalf("Failed to delete user: %v", err)
	}

}
