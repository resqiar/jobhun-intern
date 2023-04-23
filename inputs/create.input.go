package inputs

type CreateInput struct {
	Nama string `json:"nama" validate:"required"`

	// usia needs to be greater than or equal (gte) 0
	// and lower than or equal (lte) 120
	Usia int `json:"usia" validate:"required,gte=0,lte=120"`

	// gender needs to be one of 0 or 1
	// 0 means male
	// 1 means female
	// IMPORTANT NOTE:
	// Here i am using a package of "class-validator" in which validate
	// the input from the client. However, when using required in int data,
	// the package has some weird behavior which always return false if the input is
	// 0, this has an implication in this dataset for sure, so I omit the required field.
	Gender int `json:"gender" validate:"oneof=0 1"`

	NamaJurusan string `json:"namaJurusan" validate:"required"`
	NamaHobi    string `json:"namaHobi" validate:"required"`
}
