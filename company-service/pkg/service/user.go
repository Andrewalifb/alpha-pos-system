package service

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	pb "github.com/Andrewalifb/alpha-pos-system/company-service/api/proto"
	"github.com/Andrewalifb/alpha-pos-system/company-service/entity"
	"github.com/Andrewalifb/alpha-pos-system/company-service/pkg/repository"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"

	"golang.org/x/crypto/bcrypt"
)

type PosUserService interface {
	CreatePosUser(ctx context.Context, req *pb.CreatePosUserRequest) (*pb.CreatePosUserResponse, error)
	ReadPosUser(ctx context.Context, req *pb.ReadPosUserRequest) (*pb.ReadPosUserResponse, error)
	UpdatePosUser(ctx context.Context, req *pb.UpdatePosUserRequest) (*pb.UpdatePosUserResponse, error)
	DeletePosUser(ctx context.Context, req *pb.DeletePosUserRequest) (*pb.DeletePosUserResponse, error)
	ReadAllPosUsers(ctx context.Context, req *pb.ReadAllPosUsersRequest) (*pb.ReadAllPosUsersResponse, error)
	Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error)
}

type PosUserServiceServer struct {
	pb.UnimplementedPosUserServiceServer
	userRepo repository.PosUserRepository
	roleRepo repository.PosRoleRepository
}

func NewPosUserServiceServer(userRepo repository.PosUserRepository, roleRepo repository.PosRoleRepository) *PosUserServiceServer {
	return &PosUserServiceServer{
		userRepo: userRepo,
		roleRepo: roleRepo,
	}
}

func (s *PosUserServiceServer) CreatePosUser(ctx context.Context, req *pb.CreatePosUserRequest) (*pb.CreatePosUserResponse, error) {
	// Extract role ID from JWT payload
	jwtRoleID := req.JwtPayload.Role

	loginRole, err := s.roleRepo.ReadPosRole(jwtRoleID)
	if err != nil {
		return nil, err
	}

	reqCreatedRole, err := s.roleRepo.ReadPosRole(req.PosUser.RoleId)
	if err != nil {
		return nil, err
	}

	// Check role ID to determine what kind of user can be created
	switch loginRole.RoleName {
	case "super user":
		// Can create users with role "Company", "Branch", and "Store"
		if reqCreatedRole.RoleName != "Company" && reqCreatedRole.RoleName != "Branch" && reqCreatedRole.RoleName != "Store" {
			return nil, errors.New("Invalid role for new user")
		}
	case "company":
		// Can create users with role "Branch" and "Store"
		if reqCreatedRole.RoleName != "Branch" && reqCreatedRole.RoleName != "Store" {
			return nil, errors.New("Invalid role for new user")
		}
	case "branch":
		// Can create users with role "Store"
		if reqCreatedRole.RoleName != "Store" {
			return nil, errors.New("Invalid role for new user")
		}
	case "store":
		// Cannot create any user
		return nil, errors.New("Cannot create new user")
	default:
		return nil, errors.New("Invalid role for current user")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.PosUser.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	req.PosUser.PasswordHash = string(hashedPassword)
	req.PosUser.UserId = uuid.New().String() // Generate a new UUID for the user_id

	now := timestamppb.New(time.Now())
	req.PosUser.CreatedAt = now
	req.PosUser.UpdatedAt = now

	fmt.Println("CREATED AT :", req.PosUser.CreatedAt)
	fmt.Println("UPDATED AT :", req.PosUser.UpdatedAt)

	// Convert pb.PosUser to entity.PosUser
	gormUser := &entity.PosUser{
		UserID:       uuid.MustParse(req.PosUser.UserId),
		Username:     req.PosUser.Username,
		PasswordHash: req.PosUser.PasswordHash,
		RoleID:       uuid.MustParse(req.PosUser.RoleId),
		CompanyID:    uuid.MustParse(req.PosUser.CompanyId),
		BranchID:     uuid.MustParse(req.PosUser.BranchId),
		StoreID:      uuid.MustParse(req.PosUser.StoreId),
		FirstName:    req.PosUser.FirstName,
		LastName:     req.PosUser.LastName,
		Email:        req.PosUser.Email,
		PhoneNumber:  req.PosUser.PhoneNumber,
		CreatedAt:    req.PosUser.CreatedAt.AsTime(),
		CreatedBy:    uuid.MustParse(req.PosUser.CreatedBy),
		UpdatedAt:    req.PosUser.UpdatedAt.AsTime(),
		UpdatedBy:    uuid.MustParse(req.PosUser.UpdatedBy),
	}

	err = s.userRepo.CreatePosUser(gormUser)
	if err != nil {
		return nil, err
	}

	return &pb.CreatePosUserResponse{
		PosUser: req.PosUser,
	}, nil
}

func (s *PosUserServiceServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var secretKey = os.Getenv("SECRET_KEY")
	fmt.Println("LOGIN SECRET : ", secretKey)
	user, err := s.userRepo.ReadPosUserByUsername(req.Username)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return nil, err
	}

	claims := &pb.JWTPayload{
		Name:      user.FirstName + " " + user.LastName,
		Role:      user.RoleId,
		CompanyId: user.CompanyId,
		BranchId:  user.BranchId,
		StoreId:   user.StoreId,
		UserId:    user.UserId,
		StandardClaims: &pb.StandardClaims{
			Audience:  "Alpha Pos System",
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Id:        user.CompanyId,
			IssuedAt:  time.Now().Unix(),
			Issuer:    "Alpha Company",
			NotBefore: time.Now().Unix(),
			Subject:   user.RoleId,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		JwtPayload: claims,
		JwtToken:   ss,
	}, nil
}
