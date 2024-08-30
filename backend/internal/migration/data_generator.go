package migration

import (
	"github.com/Blxssy/Golang-React-Ecommerce/internal/container"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/models"
	"log"
)

func InitData(container container.Container) {
	rep := container.GetRepository()

	u := models.NewUserWithPlainPassword("test", "test@test.com", "test")
	u.Create(rep)

	u = models.NewUserWithPlainPassword("test2", "test2@test.com", "test2")
	u.Create(rep)

	// Сначала создаем категории
	category := models.Category{Name: "Sneakers", Slug: "sneakers"}
	rep.Create(&category)

	// Сохраняем категории в базе данных
	//for _, category := range categories {
	//	if err := rep.Create(category).Error; err != nil {
	//		log.Fatalf("could not create category: %v", err)
	//	}
	//}

	products := []models.Product{
		{Name: "Nike Air Max 270", Price: 120, Description: "Comfortable and stylish Nike sneakers.", Slug: "nike-air-max-270", Image: "https://static.nike.com/a/images/t_default/66989868-c275-494d-8647-5f8ddf1fd814/W+AIR+MAX+270.png", CategoryID: category.ID},
		{Name: "Nike Air Zoom Pegasus 39", Price: 130, Description: "Versatile and lightweight running shoes.", Slug: "nike-air-zoom-pegasus-39", Image: "https://imagedelivery.net/2DfovxNet9Syc-4xYpcsGg/2f474dce-47a9-42ae-621b-e677602d1200/products", CategoryID: category.ID},
		{Name: "Nike Air Force 1", Price: 100, Description: "Classic Nike sneakers with a modern twist.", Slug: "nike-air-force-1", Image: "https://static.nike.com/a/images/t_default/57558712-5ebe-4abb-9984-879f9e896b4c/W+AIR+FORCE+1+%2707+FLYEASE.png", CategoryID: category.ID},
		{Name: "Nike Blazer Mid '77", Price: 110, Description: "Retro basketball shoes with a vintage look.", Slug: "nike-blazer-mid-77", Image: "https://static.nike.com/a/images/t_default/859575fc-6a8c-4904-848e-c43391630b14/W+BLAZER+MID+%2777.png", CategoryID: category.ID},
		{Name: "Nike Dunk Low", Price: 120, Description: "Classic low-top basketball shoes.", Slug: "nike-dunk-low", Image: "https://static.nike.com/a/images/t_default/0f76f73e-2578-4d62-abab-c5563ea4f78c/NIKE+DUNK+LOW+RETRO.png", CategoryID: category.ID},
		{Name: "Nike Air Huarache", Price: 125, Description: "Comfortable and stylish shoes with a snug fit.", Slug: "nike-air-huarache", Image: "https://static.nike.com/a/images/t_default/6c7b04c0-c1a3-44e3-95b9-271627a92e00/NIKE+AIR+HUARACHE.png", CategoryID: category.ID},
		{Name: "Nike Air Max 90", Price: 140, Description: "Iconic design with classic Nike Air cushioning.", Slug: "nike-air-max-90", Image: "https://static.nike.com/a/images/t_default/w2ldynwtyuspv6r5rffj/AIR+MAX+90.png", CategoryID: category.ID},
		{Name: "Nike LeBron 21", Price: 200, Description: "High-performance basketball shoes endorsed by LeBron James.", Slug: "nike-lebron-19", Image: "https://static.nike.com/a/images/t_default/eef94a3f-2db3-4bd5-a36c-7f19758b02ce/LEBRON+XXI.png", CategoryID: category.ID},
	}

	// Сохраняем продукты в базе данных
	for _, product := range products {
		if err := rep.Create(&product).Error; err != nil {
			log.Fatalf("could not create product: %v", err)
		}
	}
}
