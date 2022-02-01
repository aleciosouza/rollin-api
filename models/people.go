package models

type People struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
}

type User struct {
	Id           uint      `json:"id"`
	Login        string    `json:"login"`
	Password     string    `json:"password"`
	User_type_id int       `json:"user_type_id"`
	User_type    User_type `json:"user_type" gorm:foreignKey:User_type_id`
	People_id    int       `json:"people_id"`
	Person       People    `json:"person" gorm:"foreignKey:People_id"`
}

type User_type struct {
	Id        uint   `json:"id"`
	User_type string `json:"user_type"`
}
