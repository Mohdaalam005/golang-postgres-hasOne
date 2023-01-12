package controller

import (
	"github.com/gofiber/fiber"
	"github.com/mohdaalam005/one-to-one/database"
	"github.com/mohdaalam005/one-to-one/models"
)

// Retrieve user list with eager loading credit card

func GetAllUser(c *fiber.Ctx) {
	var users []models.User
	database.Db.Model(&models.User{}).Preload("CreditCard").Find(&users)
	c.JSON(users)
}

func CreateUser(ctx *fiber.Ctx) {
	//var credit models.CreditCard
	var user models.User
	ctx.BodyParser(&user)
	database.Db.Create(&user)
	ctx.JSON(user)
}

/*
func GetUser(ctx *fiber.Ctx) {
	var credit models.CreditCard
	database.Db.Find(&credit)
	var user = models.User{
		CreditCard: credit,
	}
	database.Db.Find(&user, ctx.Params("id"))
	ctx.JSON(user)
}
*/

func GetUser(ctx *fiber.Ctx) {
	var users models.User
	database.Db.Model(&models.User{}).Preload("CreditCard").Find(&users, ctx.Params("id"))
	ctx.JSON(users)

}

/*
func UpdateUser(ctx *fiber.Ctx) {
	var credit models.CreditCard

	log.Println(credit)
	var user models.User
	database.Db.Find(&user, ctx.Params("id"))
	database.Db.Find(&credit, user.CreditCardId)
	log.Println(user)
	ctx.BodyParser(&credit)
	credit.ID = user.CreditCardId
	database.Db.Save(&credit)
	fmt.Println(credit.Number)
	ctx.BodyParser(&user)
	fmt.Println(credit.Number)

	user.CreditCard.Number = credit.Number
	fmt.Println(user.CreditCard.Number)
	database.Db.Save(&user)
	ctx.JSON(user)

}*/

/*
	func UpdateUser(ctx *fiber.Ctx) {
		var user models.CreditCard
		ctx.BodyParser(&user)
		database.Db.Model(&models.CreditCard{}).Association("CreditCard").Replace(&user)
		ctx.JSON(user)
	}
*/
func DeleteUser(ctx *fiber.Ctx) {
	id := ctx.Params("id")
	var user models.User
	database.Db.Delete(&user, id)
	database.Db.Delete(&user, user.CreditCard.ID)
	ctx.Send("successfully deleted !")

}

func UpdateUsers(c *fiber.Ctx) {
	var user models.User
	var creditCard models.CreditCard
	database.Db.Find(&user, c.Params("id"))
	database.Db.Find(&creditCard, user.CreditCardId)
	c.BodyParser(&user)
	//creditCard.ID = user.CreditCardId
	user.CreditCard.ID = creditCard.ID
	creditCard.Number = user.CreditCard.Number
	database.Db.Save(&user)
	database.Db.Save(&creditCard)
	c.JSON(user)

}
