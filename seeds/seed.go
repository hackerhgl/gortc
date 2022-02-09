package main

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	"strconv"
	"strings"

	"github.com/jaswdr/faker"
	"gopkg.in/nullbio/null.v4"

	models "gortc/models"
	mysql "gortc/services/mysql"
	utils "gortc/utils"
)

func Seed() {
	faker := faker.New()
	usersCount := 200
	users := make([]models.User, usersCount)
	invites := make([]models.InviteCode, usersCount)

	for i := 0; i < usersCount; i++ {
		role := models.RoleUser
		no := i + 1
		if i < 4 {
			role = models.RoleSuperAdmin
		} else if i >= 4 && i < 15 {
			no = i - 3
			role = models.RoleAdmin
		} else {
			no = i - 14
		}
		email := string(role) + strconv.Itoa(no) + "@gortc.com"
		password := strings.Split(email, "@")[0]
		hash, salt := utils.SaltNHash(password)

		users[i] = models.User{
			Name:       faker.Person().Name(),
			IsVerified: true,
			Email:      email,
			Password:   hash,
			Salt:       salt,
			Role:       role,
		}
	}
	for i := 0; i < usersCount; i++ {
		userId := uint(1 + i)
		bytes := make([]byte, 3)

		io.ReadFull(rand.Reader, bytes)
		invites[i] = models.InviteCode{
			RedeemedBy: null.UintFrom(userId),
			Code:       hex.EncodeToString(bytes),
			CreatedBy:  1,
			IsActive:   false,
		}
	}

	mysql.Ins().CreateInBatches(&users, usersCount)
	mysql.Ins().CreateInBatches(&invites, usersCount)

}
