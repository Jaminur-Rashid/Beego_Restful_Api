package controllers

import (
	"Beego_Restful_Api/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"regexp"
	_ "github.com/lib/pq"
	beego "github.com/beego/beego/v2/server/web"
	"golang.org/x/crypto/bcrypt"
)

// Operations about object
type ObjectController struct {
	beego.Controller
}

/*
Initialize database
*/
const (
	host = "localhost"

	port = 5432

	user = "postgres"

	password = "admin"

	dbname = "user_database"
)

/*
function that returns the hased password
*/
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
/*
validate email
*/
func isValidEmail(s string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(s)
}

/*
function that validates phone number
*/
func isValidPhone(s string) bool {
	return true
	//phoneRegExp := regexp.MustCompile(`^(?:\+?88)?01[15-9]\d{8}$`)
	//return phoneRegExp.MatchString(s)
}

/*
function that validates first name
*/
func isValidFirstName(s string) bool {
	firstNameRegExp := regexp.MustCompile(`([a-zA-Z',.-]+( [a-zA-Z',.-]+)*){2,30}`)
	return firstNameRegExp.MatchString(s)
}

/*
function that validates last name
*/
func isValidLastName(s string) bool {
	lastNameRegExp := regexp.MustCompile(`([a-zA-Z',.-]+( [a-zA-Z',.-]+)*){2,30}`)
	return lastNameRegExp.MatchString(s)
}

/*
function that validates birth date
can validate dd/mm/yy format
*/
func isValidBirthDate(s string) bool {
	re := regexp.MustCompile("(0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[012])/((19|20)\\d\\d)")
	fmt.Println("test", re.MatchString("31/07/2010"))
	return re.MatchString(s)
}

// @Title Create
// @Description add new user to the database
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [post]
func (o *ObjectController) Post() {
	var ob models.Object
	fmt.Println("Test from Controller", &ob)
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	fmt.Println("Test 2 ", &ob)
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// open database
	db, err := sql.Open("postgres", psqlconn)
	fmt.Println("Connected!", err)
	first_name := ob.FirstName
	last_name := ob.LastName
	phone_num := ob.Phone
	email := ob.Email
	password := ob.Password
	birth_date := ob.DoB
	fmt.Println(first_name, " Ok", last_name, phone_num, email, password, birth_date, db)

	/*
		Inserting user data into databse
	*/
	//hash user password

	hashed_password, _ := HashPassword(password)
	fmt.Println(hashed_password)
	isOkEmail := isValidEmail(email)
	isOkPhone := isValidPhone(phone_num)
	isOkFirstName := isValidFirstName(phone_num)
	isOkLastName := isValidLastName(last_name)
	isOkBirthDate := isValidBirthDate(birth_date)
	fmt.Println("Is valid Email : ", isOkEmail)
	fmt.Println("Valid phone ? ", isOkPhone)
	fmt.Println("Is valid Firstname : ", isOkFirstName)
	fmt.Println("Is valid Last Name : ", isOkLastName)
	fmt.Println("Is valid Birth Date : ", isOkBirthDate)
	/*
		if data are valid then insert into the database
	*/
	if isOkBirthDate && isOkEmail && isOkPhone {
		fmt.Println("From Controller")
		value := fmt.Sprintf("'%s','%s','%s','%s','%s','%s'", first_name, last_name, phone_num, email, hashed_password, birth_date)
		add_user_query := "INSERT INTO user_info_table (first_name,last_name,phone_no,email,password,birth_date) VALUES (" + value + ");"
		fmt.Println("Data Insertion Query", add_user_query)
		_, err := db.Exec(add_user_query)
		fmt.Println(err)
		CheckError(err)
		fmt.Println("Data Inserted Successfully")
	} else {
		fmt.Println("Data is not Valid")
	}
	defer db.Close()
	/*
		test end
	*/
	objectid := models.AddOne(ob)
	fmt.Println("Id is : ", objectid)
	o.Data["json"] = map[string]string{"ObjectId": objectid}
	o.ServeJSON()

}

// @Title Get
// @Description find object by objectid
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [get]
func (o *ObjectController) Get() {
	objectId := o.Ctx.Input.Param(":objectId")
	if objectId != "" {
		ob, err := models.GetOne(objectId)
		if err != nil {
			o.Data["json"] = err.Error()
		} else {
			o.Data["json"] = ob
		}
	}
	o.ServeJSON()
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router / [get]
func (o *ObjectController) GetAll() {
	obs := models.GetAll()
	o.Data["json"] = obs
	o.ServeJSON()
}
/*
func (o *ObjectController) Delete() {
	objectId := o.Ctx.Input.Param(":objectId")
	models.Delete(objectId)
	o.Data["json"] = "delete success!"
	o.ServeJSON()
}
*/

/*
check error
*/
func CheckError(err error) {

	if err != nil {
		fmt.Println("Database error is ", err)

		panic(err)

	}

}
