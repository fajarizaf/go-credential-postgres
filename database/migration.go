package database

import "go-auth/models"

func MigrateDatabase() {

	// migration schema
	DB.AutoMigrate((&models.Site_department{}))
	DB.AutoMigrate((&models.Site_role{}))
	DB.AutoMigrate((&models.Site_branch{}))
	DB.AutoMigrate((&models.User{}))
	DB.AutoMigrate((&models.Corporation{}))

	// created data seeder
	if result := DB.First(&models.User{}); result.RowsAffected == 0 {

		site_department := models.Site_department{
			Status:     "active",
			Department: "Research And Development",
		}

		site_role := models.Site_role{
			Role_Name:         "Administrator",
			Role_Type:         "admin",
			Site_departmentID: 1,
		}

		site_branch := models.Site_branch{
			Name:    "Jababeka",
			Address: "Bekasi, Cikarang district 88",
			Status:  "active",
		}

		corporation := models.Corporation{
			Parent_ID:   1,
			CompanyName: "PT. Your Company",
			Address:     "Jakarta, Jl KH. Agus Salim",
			City:        "Jakarta",
			State:       "Indonesia",
			Phone:       "082125649665",
			Email:       "rangga@gmail.com",
		}

		user := models.User{
			Site_branchID: 1,
			Site_roleID:   1,
			User_Login:    "rangga",
			User_Password: "Ada",
			ContactName:   "Rangga Aditya",
			Address:       "Jakarta, Jl KH. Agus Salim",
			City:          "Jakarta",
			State:         "Indonesia",
			Phone:         "082125649665",
			Isactive:      true,
			Email:         "rangga@gmail.com",
		}

		user_company := models.User_corporation{
			Corporation_id: 1,
			User_id:        1,
		}

		DB.Create(&site_department) // pass pointer of data to Create
		DB.Create(&site_role)       // pass pointer of data to Create
		DB.Create(&site_branch)     // pass pointer of data to Create
		DB.Create(&corporation)     // pass pointer of data to Create
		DB.Create(&user)            // pass pointer of data to Create
		DB.Create(&user_company)    // pass pointer of data to Create

	}

}
