package models

import (
	"errors"
	"fmt"
)

var (
	Objects map[string]*Object
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
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}


func isValidEmail(s string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(s)
}


func isValidPhone(s string) bool {
	return true
	//phoneRegExp := regexp.MustCompile(`^(?:\+?88)?01[15-9]\d{8}$`)
	//return phoneRegExp.MatchString(s)
}


func isValidFirstName(s string) bool {
	firstNameRegExp := regexp.MustCompile(`([a-zA-Z',.-]+( [a-zA-Z',.-]+)*){2,30}`)
	return firstNameRegExp.MatchString(s)
}

func isValidLastName(s string) bool {
	lastNameRegExp := regexp.MustCompile(`([a-zA-Z',.-]+( [a-zA-Z',.-]+)*){2,30}`)
	return lastNameRegExp.MatchString(s)
}

func isValidBirthDate(s string) bool {
	re := regexp.MustCompile("(0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[012])/((19|20)\\d\\d)")
	fmt.Println("test", re.MatchString("31/07/2010"))
	return re.MatchString(s)
}
*/
func AddOne(object Object) (ObjectId string) {

	// connection string
	//psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// open database
	//db, err := sql.Open("postgres", psqlconn)
	//fmt.Println("Connected!", err)
	/*
	extract user data
	first_name := object.FirstName
	last_name := object.LastName
	phone_num := object.Phone
	email := object.Email
	//password := object.Password
	birth_date := object.DoB
	*/
	/*
		Inserting user data into databse
	*/
	//hash user password
	/*
	hashed_password, _ := HashPassword(password)
	fmt.Println(hashed_password)
	fmt.Println(birth_date)
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
	*/
	/*
		if data are valid then insert into the database

	if isOkBirthDate && isOkEmail && isOkPhone {
		value := fmt.Sprintf("'%s','%s','%s','%s','%s','%s'", first_name, last_name, phone_num, email, hashed_password, birth_date)
		add_user_query := "INSERT INTO user_info_table (first_name,last_name,phone_no,email,password,birth_date) VALUES (" + value + ");"
		fmt.Println("Data Insertion Query", add_user_query)
		_, e := db.Exec(add_user_query)
		fmt.Println(e)
		// check db
		e = db.Ping()
		CheckError(e)
		fmt.Println("Data Inserted Successfully")
	}else {
		fmt.Println("Data is not Valid")
	}
	*/
	//defer db.Close()
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
		fmt.Println("Database error is ", err)

		panic(err)

	}

}
