package controller

import (
	"fmt"
	"net/http"
	"github.com/gofiber/fiber/v2"
	inimodel "github.com/jul003/BE_Tb/model"
	cek "github.com/jul003/BE_Tb/module"
	"github.com/jul003/Boiler_Tb/config"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	
)
// GetAllGadget godoc
// @Summary Get All Data Gadget.
// @Description Mengambil semua data gadget.
// @Tags Gadget
// @Accept json
// @Produce json
// @Success 200 {object} Gadget
// @Router /gadget [get]
func GetAllGagdet(c *fiber.Ctx) error {
	ps := cek.GetDataGadget(config.Ulbimongoconn, "gadget2024")
	return c.JSON(ps)
}

// GetAllReview godoc
// @Summary Get All Data Review.
// @Description Mengambil semua data review.
// @Tags Review
// @Accept json
// @Produce json
// @Success 200 {object} Gadget
// @Router /review [get]
func GetAllReview(c *fiber.Ctx) error {
	ps := cek.GetDataReview(config.Ulbimongoconn, "review2024")
	return c.JSON(ps)
}

// InsertDataGadget godoc
// @Summary Insert data Gadget.
// @Description Input data gadget.
// @Tags Gadget
// @Accept json
// @Produce json
// @Param request body ReqGadget true "Payload Body [RAW]"
// @Success 200 {object} Gadget
// @Failure 400
// @Failure 500
// @Router /insertgadget [post]
func InsertDataGadget(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var gadgets inimodel.Gadget
	if err := c.BodyParser(&gadgets); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	if gadgets.Nama == "" || gadgets.Merk == "" || gadgets.Harga == 0 ||
		gadgets.Deskripsi == "" || gadgets.Spesifikasi.Prosesor == "" || gadgets.Spesifikasi.RAM == 0 || gadgets.Spesifikasi.Storage == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Please fill all the required fields.",
		})
	}

	insertedID, err := cek.InsertGadget(db, "gadget2024",
		gadgets.Nama,
		gadgets.Merk,
		gadgets.Harga,
		gadgets.Spesifikasi,
		gadgets.Deskripsi)
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

// InsertDataReview godoc
// @Summary Insert data Review.
// @Description Input data review.
// @Tags Review
// @Accept json
// @Produce json
// @Param request body ReqReview true "Payload Body [RAW]"
// @Success 200 {object} Review
// @Failure 400
// @Failure 500
// @Router /insertreview [post]
func InsertDataReview(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var reviews inimodel.Review
	if err := c.BodyParser(&reviews); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	if reviews.Rating == 0 || reviews.Review == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Please fill all the required fields.",
		})
	}

	insertedID, err := cek.InsertReview(db, "review2024",
		reviews.Rating,
		reviews.Review)
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

// UpdateDataGadget godoc
// @Summary Update data Gadegt.
// @Description Ubah data gadget.
// @Tags Gadget
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Param request body ReqGadget true "Payload Body [RAW]"
// @Success 200 {object} Gadget
// @Failure 400
// @Failure 500
// @Router /updategadget/{id} [put]
func UpdateDataGadget(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	var gadgets inimodel.Gadget
	if err := c.BodyParser(&gadgets); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	err = cek.UpdateGadget(db, "gadget2024",
		objectID,
		gadgets.Nama,
		gadgets.Merk,
		gadgets.Harga,
		gadgets.Spesifikasi,
		gadgets.Deskripsi)
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

// DeleteGadgetByID godoc
// @Summary Delete data Gadget.
// @Description Hapus data gadget.
// @Tags Gadget
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /deletegadget/{id} [delete]
func DeleteGadgetByID(c *fiber.Ctx) error {
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

	err = cek.DeleteGadgetID(objID, config.Ulbimongoconn, "gadget2024")
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
// GetGadgetByID godoc
// @Summary Get By ID Data Gadgets.
// @Description Ambil per ID data gadget.
// @Tags Gadget
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200 {object} Gadget
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /gadget/{id} [get]
func GetGadgetByID(c *fiber.Ctx) error {
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
	ps, err := cek.GetGadgetByID(objID, config.Ulbimongoconn, "gadget2024")
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