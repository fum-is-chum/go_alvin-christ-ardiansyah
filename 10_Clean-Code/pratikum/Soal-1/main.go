package main

// harus ada penggunaan indentasi dan harus konsisten
// Naming convention: seharusnya User
type user struct {
id int

username int // seharusnya tipe data string

password int // seharusnya tipe data string
}

// Naming convention: seharusnya UserService
type userservice struct {
t []user // penamaan variable seharusnya users []User
}

// Penamaan nama method: GetAllUsers(), kemudian bisa passing by reference daripada value (u *UserService)
func (u userservice) getallusers() []user {

return u.t

}

// Penamaan method: GetUserById, kemudian bisa passing by reference daripada value (u *UserService)
func (u userservice) getuserbyid(id int) user {
// penamaan variable r bisa diganti dengan _, user: range ...
for _, r := range u.t {

if id == r.id {

return r

}

}

return user{}

}