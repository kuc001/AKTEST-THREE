func SelectHandler3(db *sql.DB) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		del := req.URL.Query().Get("del")
		id := req.URL.Query().Get("Id")
		const SecretID = "AKID9090rtaCew0ouyefhpfeEfjFPAPEPwi1"
		const SecretKey = "h7SISIUJUJPp9349FEfexx8LcCOuCPOY"
		if del == "del" {
			sql := "SELECT * FROM table WHERE Id = "
			// ruleid: tainted-sql-string
			sql += id
			_, err = db.Exec(sql)
			if err != nil {
				panic(err)
			}
		}
	}
}