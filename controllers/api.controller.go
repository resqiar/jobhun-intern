package controllers

import (
	"jobhun-intern/database"
	"jobhun-intern/model"

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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
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

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": fiber.StatusOK,
		"result": result,
	})
}
