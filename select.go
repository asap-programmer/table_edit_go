package main

import (
	"bufio"
	"fmt"
	"strconv"

	//"reflect"
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
    db, err := sql.Open("mysql", "root:qwerty123)@/bank")

    if err != nil {
		fmt.Println("error")
        panic(err)
    }
    defer db.Close()

	// открытие из браузера корневого каталога.
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		viewSelect(w, db)
    })

	// сохранение отправленных значений через поля формы.
	http.HandleFunc("/postform", func(w http.ResponseWriter, r *http.Request){
        val1 := r.FormValue("col1")
        val2 := r.FormValue("col2")
        val3 := r.FormValue("col3")
        val4 := r.FormValue("col4")
        val5 := r.FormValue("col5")
        val6 := r.FormValue("col6")
        val7 := r.FormValue("col7")
        val8 := r.FormValue("col8")
        val9 := r.FormValue("col9")
		sQuery := "INSERT INTO individuals (first_name, surname, patronymic, passport, inn, snils, driver_license, add_documents, notice) VALUES ('"+val1+"', '"+val2+"', '"+val3+"', '"+val4+"', '"+val5+"', '"+val6+"', '"+val7+"', '"+val8+"', '"+val9+"')"

		fmt.Println(sQuery)

		rows, err := db.Query(sQuery)
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		viewSelect(w, db)
    })

	// отображение формы редактирования
	http.HandleFunc("/edit", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		renderEditForm(w, db, id)
	})

	// обработка обновления данных
	http.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		id := r.FormValue("id")
		val1 := r.FormValue("col1")
		val2 := r.FormValue("col2")
		val3 := r.FormValue("col3")
		val4 := r.FormValue("col4")
        val5 := r.FormValue("col5")
        val6 := r.FormValue("col6")
        val7 := r.FormValue("col7")
        val8 := r.FormValue("col8")
        val9 := r.FormValue("col9")
		sQuery := "UPDATE individuals SET first_name = '" + val1 + "', surname = '" + val2 + "', patronymic = '" + val3 + "' , passport = '" + val4 + "', inn = '" + val5 + "', snils = '" + val6 + "', driver_license = '" + val7 + "', add_documents = '" + val8 + "', notice = '" + val9 + "' WHERE id = " + id

		fmt.Println(sQuery)

		_, err := db.Exec(sQuery)
		if err != nil {
			panic(err)
		}

		viewSelect(w, db)
	})

    fmt.Println("Server is listening on http://localhost:8181/")
    http.ListenAndServe(":8181", nil)
}

// главная функция для показа таблицы в браузере, которая показывается при любом запросе.
func viewSelect(w http.ResponseWriter, db *sql.DB) {
	// чтение шаблона.
	file, err := os.Open("select.html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//	кодовая фраза для вставки значений из БД.
		if scanner.Text() != "@tr" && scanner.Text() != "@ver" {
			fmt.Fprintf(w, scanner.Text())
		}
		if scanner.Text() == "@tr" {
			viewHeadQuery(w, db, "SELECT COLUMN_NAME FROM information_schema.COLUMNS WHERE TABLE_NAME = 'individuals' ORDER BY ORDINAL_POSITION")
			viewSelectQuery(w, db, "SELECT * FROM individuals WHERE id>14 ORDER BY id DESC")
		}
		if scanner.Text() == "@ver" {
			viewSelectVerQuery(w, db, "SELECT VERSION() AS ver")
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// отправка в браузер заголовка таблицы.
func viewHeadQuery(w http.ResponseWriter, db *sql.DB, sShow string) {
	type sHead struct {
		clnme string
	}
    rows, err := db.Query(sShow)
    if err != nil {
        panic(err)
    }
    defer rows.Close()

	fmt.Fprintf(w, "<tr>")
	i := 0
     for i < 10 {
		rows.Next()
        p := sHead{}
        err := rows.Scan(&p.clnme)
        if err != nil{
            fmt.Println(err)
            continue
        }
		fmt.Fprintf(w, "<td>"+p.clnme+"</td>")
		i++
    }
	fmt.Fprintf(w, "</tr>")
}

// отправка в браузер строк из таблицы.
func viewSelectQuery(w http.ResponseWriter, db *sql.DB, sSelect string) {
	type bank struct {
		id int
		first_name string
		surname string
		patronymic string
		passport string
		inn string
		snils string
		driver_license string
		add_documents string
		notice string
	}
	banks := []bank{}
	//fmt.Println(reflect.TypeOf(banks))

	// получение значений в массив banks из струкрур типа bank.
    rows, err := db.Query(sSelect)
    if err != nil {
        panic(err)
    }
    defer rows.Close()

    for rows.Next(){
        p := bank{}
        err := rows.Scan(&p.id, &p.first_name, &p.surname, &p.patronymic, &p.passport, &p.inn, &p.snils, &p.driver_license, &p.add_documents, &p.notice)
        if err != nil{
            fmt.Println(err)
            continue
        }
        banks = append(banks, p)
    }

	// перебор массива из БД.
	for _, p := range banks {
		fmt.Fprintf(w, "<tr><td>"+strconv.Itoa(p.id)+"</td><td>"+p.first_name+"</td><td>"+p.patronymic+"</td><td>"+p.passport+"</td><td>"+p.inn+"</td><td>"+p.snils+"</td><td>"+p.driver_license+"</td><td>"+p.add_documents+"</td><td>"+p.notice+"</td>")
		fmt.Fprintf(w, "<td><a href=\"/edit?id="+strconv.Itoa(p.id)+"\">Edit</a></td></tr>")
	}
}

// отправка в браузер версии базы данных.
func viewSelectVerQuery (w http.ResponseWriter, db *sql.DB, sSelect string) {
	type sVer struct {
		ver string
	}
    rows, err := db.Query(sSelect)
    if err != nil {
        panic(err)
    }
    defer rows.Close()
     for rows.Next() {
        p := sVer{}
        err := rows.Scan(&p.ver)
        if err != nil{
            fmt.Println(err)
            continue
        }
		fmt.Fprintf(w, p.ver)
    }
}

// отображение формы редактирования
func renderEditForm(w http.ResponseWriter, db *sql.DB, id string) {
	type bank struct {
		id int
		first_name string
		surname string
		patronymic string
		passport string
		inn string
		snils string
		driver_license string
		add_documents string
		notice string
	}
	var p bank
	err := db.QueryRow("SELECT id, first_name, surname, patronymic, passport, inn, snils, driver_license, add_documents, notice FROM individuals WHERE id = ?", id).Scan(&p.id, &p.first_name, &p.surname, &p.patronymic, &p.passport, &p.inn, &p.snils, &p.driver_license, &p.add_documents, &p.notice)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fmt.Fprintf(w, `<form action="/update" method="post">`)
	fmt.Fprintf(w, `<input type="hidden" name="id" value="%d">`, p.id)
	fmt.Fprintf(w, `first_name: <input type="text" name="col1" value="%s"><br>`, p.first_name)
	fmt.Fprintf(w, `surname: <input type="text" name="col2" value="%s"><br>`, p.surname)
	fmt.Fprintf(w, `patronymic: <input type="text" name="col3" value="%s"><br>`, p.patronymic)
	fmt.Fprintf(w, `passport: <input type="text" name="col4" value="%s"><br>`, p.passport)
	fmt.Fprintf(w, `inn: <input type="text" name="col5" value="%s"><br>`, p.inn)
	fmt.Fprintf(w, `snils: <input type="text" name="col6" value="%s"><br>`, p.snils)
	fmt.Fprintf(w, `driver_license: <input type="text" name="col7" value="%s"><br>`, p.driver_license)
	fmt.Fprintf(w, `add_documents: <input type="text" name="col8" value="%s"><br>`, p.add_documents)
	fmt.Fprintf(w, `notice: <input type="text" name="col9" value="%s"><br>`, p.notice)
	fmt.Fprintf(w, `<input type="submit" value="Update">`)
	fmt.Fprintf(w, `</form>`)
	fmt.Fprintf(w, `<br><a style="background-color: #33bee1; color: #fff; padding: 2px 3px; text-decoration: none" href="/">Return to main menu</a>`)
}