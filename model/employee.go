package model
type Employee struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
}

// git stash -> untuk menyimpan perubahan sementara
// git stash pop -> untuk mengambil kembali perubahan tersebut

// go get github.com/lib/pq -> driver database potsgresql
// go get github.com/joho/godotenv -> reading .env file
// go get github.com/google/uuid 
// generate uuid (Universally Unique Identifier) => random string (kombinasi angka dan huruf) 
// untuk kebutuhan primary key setiap data
// SECURE 
// jiak menggunakan auto increment itu mudah ditebak karena angkanya konsisten, contoh : 1,2,3,4...