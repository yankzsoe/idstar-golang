package services

import (
	"idstar.com/session7/app/models"
	"idstar.com/session7/app/repositories"
	"idstar.com/session7/app/tools"
)

type UserService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *UserService {
	return &UserService{userRepository}
}

func (u *UserService) CreateUser(user *models.UserModel) (*models.UserModel, error) {
	// Enkripsi password
	aes128 := tools.Aes128{}
	encryptedPassword, err := aes128.Encrypt(user.Password)
	if err != nil {
		return nil, err
	}

	// Simpan user ke database
	user.Password = *encryptedPassword
	return u.userRepository.Create(user)
}

func (u *UserService) GetUserByID(userID string) (*models.UserModel, error) {
	return u.userRepository.FindByID(userID)
}

/*
func (u *UserService) UpdateUser(user models.UserModel) (*models.UserModel, error) {
	// Lakukan validasi input
	if user.ID == 0 {
		return nil, errors.New("ID user tidak valid")
	}

	// Verifikasi user
	currentUser, err := u.GetUserByID(user.ID)
	if err != nil {
		return nil, err
	}

	// Pemrosesan hak akses
	if !currentUser.IsAdmin {
		// Hanya admin yang boleh mengubah data user lain
		if currentUser.ID != user.ID {
			return nil, errors.New("Tidak diizinkan mengubah data user lain")
		}

		// Non-admin tidak bisa mengubah hak akses
		user.IsAdmin = false
	}

	// Update user di database
	return u.userRepository.UpdateUser(user)
}

func (u *UserService) DeleteUser(userID uint) error {
	// Lakukan validasi input
	if userID == 0 {
		return errors.New("ID user tidak valid")
	}

	// Verifikasi user
	currentUser, err := u.GetUserByID(userID)
	if err != nil {
		return err
	}

	// Pemrosesan hak akses
	if !currentUser.IsAdmin {
		return errors.New("Tidak diizinkan menghapus data user")
	}

	// Hapus user dari database
	return u.userRepository.DeleteUser(userID)
}
*/
