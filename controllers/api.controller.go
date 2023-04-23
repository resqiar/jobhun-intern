package controllers

import (
	"fmt"
	"jobhun-intern/database"
	"jobhun-intern/inputs"
	"jobhun-intern/libs"
	"jobhun-intern/model"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func SendHello(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func SendAllMahasiswa(c *fiber.Ctx) error {
	// SQL Explanation
	// 1. Select all from table mahasiswa
	// 2. Select nama_jurusan from table jurusan
	// 3. Select nama_hobi from table hobi
	// 4. JOIN with JURUSAN table with alias (jrs)
	// 5. JOIN with MAHASISWA_HOBI table with alias (mh)
	// 6. JOIN with HOBI table with alias (h)
	sql := "SELECT mhs.*, jrs.nama_jurusan, h.nama_hobi " +
		"FROM mahasiswa mhs " +
		"LEFT JOIN jurusan jrs ON mhs.jurusan = jrs.id " +
		"LEFT JOIN mahasiswa_hobi mh ON mhs.id = mh.id_mahasiswa " +
		"LEFT JOIN hobi h ON h.id = mh.id_hobi"

	rows, err := database.DB.Query(sql)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Loop through the rows and bind the result to struct
	var result []model.Mahasiswa
	for rows.Next() {
		var model model.Mahasiswa

		// Scan through all the data returned from the database, bind them into struct
		err := rows.Scan(&model.Id, &model.Nama, &model.Usia, &model.Gender, &model.TanggalRegistrasi, &model.Jurusan, &model.NamaJurusan, &model.NamaHobi)
		if err != nil {
			panic(err.Error())
		}

		// append individual data to the array of result
		result = append(result, model)
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": fiber.StatusOK,
		"result": result,
	})
}

func CreateMahasiswa(c *fiber.Ctx) error {
	var payload inputs.CreateInput
	// Bind request body to payload struct
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}

	// Validate the input to follow the format of
	// the defined struct. see inputs/create.input.go
	err := libs.InputValidator(payload)
	if err != "" {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err,
		})
	}

	// CHECK THE JURUSAN
	sql := fmt.Sprintf("SELECT id FROM jurusan WHERE nama_jurusan = '%s'", strings.ToLower(payload.NamaJurusan))
	rows, sqlErr := database.DB.Query(sql)
	if sqlErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": sqlErr,
		})
	}

	// IF JURUSAN IS ALREADY AVAILABLE, USE THAT JURUSAN ID AND
	var jurusanID int
	for rows.Next() {
		var id int

		// Scan through all the data returned from the database, bind them into struct
		err := rows.Scan(&id)

		if err != nil {
			panic(err.Error())
		}

		// update the jurusan id
		jurusanID = id
	}

	// If jurusan is not available in the database
	// then create new jurusan with the provided name,
	// then return the created ID.
	if jurusanID == 0 {
		createJurusanSQL := fmt.Sprintf("INSERT INTO jurusan (nama_jurusan) VALUES ('%s')", strings.ToLower(payload.NamaJurusan))

		raw, sqlErr := database.DB.Exec(createJurusanSQL)
		if sqlErr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"status":  fiber.StatusInternalServerError,
				"message": sqlErr,
			})
		}

		// get the last inserted ID
		lastInsertID, _ := raw.LastInsertId()
		// SAVE IT TO MAHASISWA FIELD
		jurusanID = int(lastInsertID)
	}

	// CHECK IF HOBBY IS AVAILABLE
	checkHobiSQL := fmt.Sprintf("SELECT id FROM hobi WHERE nama_hobi = '%s'", strings.ToLower(payload.NamaHobi))
	rows, sqlErr = database.DB.Query(checkHobiSQL)
	if sqlErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": sqlErr,
		})
	}

	// IF HOBBY IS ALREADY AVAILABLE, USE THAT HOBBY ID
	var hobbyID int
	for rows.Next() {
		var id int

		// Scan through all the data returned from the database, bind them into struct
		err := rows.Scan(&id)

		if err != nil {
			panic(err.Error())
		}

		// update the jurusan id
		hobbyID = id
	}

	// If hobby is not available in the database
	// then create new hobby with the provided name,
	// then return the created ID.
	if hobbyID == 0 {
		createHobiSQL := fmt.Sprintf("INSERT INTO hobi (nama_hobi) VALUES ('%s')", strings.ToLower(payload.NamaHobi))

		raw, sqlErr := database.DB.Exec(createHobiSQL)
		if sqlErr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"status":  fiber.StatusInternalServerError,
				"message": sqlErr,
			})
		}

		// get the last inserted ID
		lastInsertID, _ := raw.LastInsertId()
		// SAVE IT TO MAHASISWA FIELD
		hobbyID = int(lastInsertID)
	}

	// INSERT MAHASISWA DATA
	createMahasiswaSQL := fmt.Sprintf("INSERT INTO mahasiswa (nama, usia, gender, jurusan) VALUES ('%s', %d, %d, %d)", payload.Nama, payload.Usia, payload.Gender, jurusanID)
	raw, createErr := database.DB.Exec(createMahasiswaSQL)
	if createErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": createErr.Error(),
		})
	}

	// ID OF INSERTED MAHASISWA
	mahasiswaID, _ := raw.LastInsertId()

	// FINALLY SAVE MAHASISWA ID AND HOBI ID INTO mahasiswa_hobi
	createMHSQL := fmt.Sprintf("INSERT INTO mahasiswa_hobi (id_mahasiswa, id_hobi) VALUES (%d, %d)", mahasiswaID, hobbyID)
	_, createErr = database.DB.Query(createMHSQL)
	if createErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": createErr.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": fiber.StatusOK,
	})
}
