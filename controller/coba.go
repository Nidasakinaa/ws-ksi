package controller

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	inimodel "github.com/Nidasakinaa/be_KaloriKu/model"
	cek "github.com/Nidasakinaa/be_KaloriKu/module"
	"github.com/Nidasakinaa/ws-kaloriku/config"
	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

// GetMenuID godoc
// @Summary Get By ID Data Menu.
// @Description Ambil per ID data menu.
// @Tags MenuItem
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200 {object} MenuItem
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /menu/{id} [get]
func GetMenuID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}
	ps, err := cek.GetMenuItemByID(objID, config.Ulbimongoconn, "Proyek3")
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  http.StatusNotFound,
				"message": fmt.Sprintf("No data found for id %s", id),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error retrieving data for id %s", id),
		})
	}
	return c.JSON(ps)
}

// GetMenuByCategory godoc
// @Summary Get By Category Data Menu.
// @Description Ambil per ID data menu.
// @Tags MenuItem
// @Accept json
// @Produce json
// @Param id path string true "Masukan category"
// @Success 200 {object} MenuItem
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /menu/{category} [get]
func GetMenuByCategory(c *fiber.Ctx) error {
	category := c.Params("category")
	if category == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}
	ps, err := cek.GetMenuItemByCategory(category, config.Ulbimongoconn, "Proyek3")
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  http.StatusNotFound,
				"message": fmt.Sprintf("No data found for category %s", category),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error retrieving data for category %s", category),
		})
	}
	return c.JSON(ps)
}

// GetMenu godoc
// @Summary Get All Data Menu.
// @Description Mengambil semua data menu.
// @Tags MenuItem
// @Accept json
// @Produce json
// @Success 200 {object} MenuItem
// @Router /menu [get]
// GetAllMenuItem retrieves all menu items from the database
func GetMenu(c *fiber.Ctx) error {
	ps := cek.GetAllMenuItem(config.Ulbimongoconn, "Menu")
	return c.JSON(ps)
}

// InsertDataMenu godoc
// @Summary Insert data menu.
// @Description Input data menu.
// @Tags MenuItem
// @Accept json
// @Produce json
// @Param request body ReqMenuItem true "Payload Body [RAW]"
// @Success 200 {object} MenuItem
// @Failure 400
// @Failure 500
// @Router /insert [post]
func InsertDataMenu(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var menuItem inimodel.MenuItem
	if err := c.BodyParser(&menuItem); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	insertedID, err := cek.InsertMenuItem(db, "Menu",
		menuItem.Name,
		menuItem.Ingredients,
		menuItem.Description,
		menuItem.Calories,
		menuItem.Category,
		menuItem.Image)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

// UpdateData godoc
// @Summary Update data menu.
// @Description Ubah data menu.
// @Tags MenuItem
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Param request body ReqMenuItem true "Payload Body [RAW]"
// @Success 200 {object} MenuItem
// @Failure 400
// @Failure 500
// @Router /update/{id} [put]
func UpdateDataMenuItem(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	id := c.Params("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	var menuItem inimodel.MenuItem
	if err := c.BodyParser(&menuItem); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	err = cek.UpdateMenuItem(context.Background(), db, "Menu",
		objectID,
		menuItem.Name,
		menuItem.Ingredients,
		menuItem.Description,
		menuItem.Calories,
		menuItem.Category,
		menuItem.Image)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Data successfully updated",
	})
}

// DeleteMenuItemByID godoc
// @Summary Delete data menuItem.
// @Description Hapus data menuItem.
// @Tags MenuItem
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /delete/{id} [delete]
func DeleteMenuItemByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}

	err = cek.DeleteMenuItemByID(objID, config.Ulbimongoconn, "Menu")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error deleting data for id %s", id),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("Data with id %s deleted successfully", id),
	})
}

// GetUser godoc
// @Summary Get All Data User.
// @Description Mengambil semua data user.
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} User
// @Router /user [get]
func GetUser(c *fiber.Ctx) error {
	ps := cek.GetAllUser(config.Ulbimongoconn, "User")
	return c.JSON(ps)
}

// GetUserID godoc
// @Summary Get By ID Data User.
// @Description Ambil per ID data USER.
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200 {object} User
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /user/{id} [get]
func GetUserID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}
	ps, err := cek.GetUserByID(objID, config.Ulbimongoconn, "User")
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  http.StatusNotFound,
				"message": fmt.Sprintf("No data found for id %s", id),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error retrieving data for id %s", id),
		})
	}
	return c.JSON(ps)
}

// InsertDataUser godoc
// @Summary Insert data user.
// @Description Input data user.
// @Tags User
// @Accept json
// @Produce json
// @Param request body ReqUser true "Payload Body [RAW]"
// @Success 200 {object} User
// @Failure 400
// @Failure 500
// @Router /insert [post]
func InsertDataUser(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var user inimodel.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	insertedID, err := cek.InsertUser(db, "User",
		user.FullName,
		user.Phone,
		user.Username,
		user.Password,
		user.Role)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

// UpdateData godoc
// @Summary Update data user.
// @Description Ubah data user.
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Param request body ReqUser true "Payload Body [RAW]"
// @Success 200 {object} User
// @Failure 400
// @Failure 500
// @Router /update/{id} [put]
func UpdateDataUser(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	id := c.Params("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	var user inimodel.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	err = cek.UpdateUser(context.Background(), db, "User",
		objectID,
		user.FullName,
		user.Phone,
		user.Username,
		user.Password,
		user.Role)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Data successfully updated",
	})
}

// DeleteUserByID godoc
// @Summary Delete data user.
// @Description Hapus data user.
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /delete/{id} [delete]
func DeleteUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}

	err = cek.DeleteUserByID(objID, config.Ulbimongoconn, "User")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error deleting data for id %s", id),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("Data with id %s deleted successfully", id),
	})
}
