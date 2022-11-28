package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	//didefinisikan contruct dan beri error jika ada error
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindByID(ID int) (User, error)
	Update(user User) (User, error)
}

// lihat "r" nya kecil karena tidak bersifat public / tidak bisa diakses oleh yg lain
type repository struct {
	db *gorm.DB
}

// jangan lupa tambahin untuk memanggil koneksi DB dan membuat object baru
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User

	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByID(ID int) (User, error) {
	var user User

	err := r.db.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) Update(user User) (User, error) {
	//Save adalah menyimpan data yg sudah ada, beda sama create yg belum ada
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
