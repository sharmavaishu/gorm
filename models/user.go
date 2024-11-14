package models

import "gorm.io/gorm"

type User struct {
	gorm.Model  // for models metadata -> CreatedAt , UpdatedAt , DeletedAt 
	ID int 
	Name string
	Email string
}

// Create user
func CreateUser(db *gorm.DB, user *User) (err error){
    err = db.Create(user).Error
	if err != nil{
		return err
	}
    return nil
}

// List users
func GetUsers(db *gorm.DB,user *User) (err error){
	err = db.Find(user).Error  // .Error will contain any error that occurs during the query execution.
	if err != nil{
		
		return err
	}
	return nil
}

// Get user by id 
func GetUser(db *gorm.DB,user *User, id int) (err error){
	err = db.Where("id = ?",id).First(user).Error
	if err != nil{
		
		return err
	}
	return nil
}


// Update user
func UpdateUser(db *gorm.DB, user *User) (err error) {
   
    err = db.Save(user).Error
    if err != nil {
        return err 
    }
    return nil 
}

// Delete user
func DeleteUser(db *gorm.DB, user *User, id int) (err error) {
    
    err = db.Where("id = ?", id).Delete(user).Error
    if err != nil {
        return err
    }
    return nil 
}




