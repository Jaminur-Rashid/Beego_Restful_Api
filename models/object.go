package models

import (
	"database/sql"
	"errors"

	"fmt"
	"regexp"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

var (
	Objects map[string]*Object
)

const (
	host = "localhost"

	port = 5432

	user = "postgres"

	password = "admin"

	dbname = "user_database"
)

type Object struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Phone     string `json:"Phone"`
	Email     string `json:"Email"`
	Password  string `json:"Password"`
	DoB       string `json:"DoB"`
}

func init() {
	Objects = make(map[string]*Object)

}

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
func isValidEmail(s string) bool{
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
    return emailRegex.MatchString(s)
}
/*
function that validates phone number
*/
func isValidPhone (s string) bool{
	return true
}
/*
function that validates first name
*/
func isValidFirstName(s string) bool{
	firstNameRegExp := regexp.MustCompile(`([a-zA-Z',.-]+( [a-zA-Z',.-]+)*){2,30}`)
    return firstNameRegExp.MatchString(s)
}
/*
function that validates last name
*/
func isValidLastName (s string) bool{
	lastNameRegExp := regexp.MustCompile(`([a-zA-Z',.-]+( [a-zA-Z',.-]+)*){2,30}`)
    return lastNameRegExp.MatchString(s)
}
func AddOne(object Object) (ObjectId string) {

	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	if err != nil {
		panic(err)
	}
	// close database
	fmt.Println("Connected!")
	//ashkas
	first_name := object.FirstName
	last_name := object.LastName
	phone_num := object.Phone
	email := object.Email
	password := object.Password
	birth_date := object.DoB
	/*
		fmt.Println(first_name)
		fmt.Println(last_name)
		fmt.Println(phone_num)
		fmt.Println(email)
		fmt.Println(password)
		fmt.Println(birth_date)
	*/
	/*
		Inserting user data into databse
	*/
	//hash user password
	hashed_password, _ := HashPassword(password)
	fmt.Println(hashed_password)
	isOkEmail :=isValidEmail(email)
	isOkPhone:=isValidPhone(phone_num)
	isOkFirstName := isValidFirstName(phone_num)
	isOkLastName := isValidLastName(last_name)
	fmt.Println(isOkEmail)
	fmt.Println(isOkPhone)
	fmt.Println(isOkFirstName)
	fmt.Println(isValidEmail(email))
	fmt.Println(isOkLastName)
	value := fmt.Sprintf("'%s','%s','%s','%s','%s','%s'", first_name, last_name, phone_num, email, hashed_password, birth_date)
	add_user_query := "INSERT INTO user_info_table (first_name,last_name,phone_no,email,password,birth_date) VALUES (" + value + ");"
	fmt.Println("Data Insertion Query", add_user_query)
	_, e := db.Exec(add_user_query)
	fmt.Println("Data Inserted Successfully")
	defer db.Close()
	// check db
	err = db.Ping()
	CheckError(err)
	CheckError(e)
	Objects[object.Email] = &object
	return object.Email
}

func GetOne(ObjectId string) (object *Object, err error) {
	if v, ok := Objects[ObjectId]; ok {
		return v, nil
	}
	return nil, errors.New("ObjectId Not Exist")
}

func GetAll() map[string]*Object {
	return Objects
}

func Delete(ObjectId string) {
	delete(Objects, ObjectId)
}

func CheckError(err error) {

	if err != nil {

		panic(err)

	}

}
