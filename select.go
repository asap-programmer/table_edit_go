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
		sQuery := "INSERT INTO individuals (text, description, keywords) VALUES ('"+val1+"', '"+val2+"', '"+val3+"')"

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
		sQuery := "UPDATE individuals SET text = '" + val1 + "', description = '" + val2 + "', keywords = '" + val3 + "' WHERE id = " + id

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
			viewHeadQuery(w, db, "select COLUMN_NAME AS clnme from information_schema.COLUMNS where TABLE_NAME='individuals'")
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
     for i < 4 {
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
		text string
		description string
		keywords string
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
        err := rows.Scan(&p.id, &p.text, &p.description, &p.keywords)
        if err != nil{
            fmt.Println(err)
            continue
        }
        banks = append(banks, p)
    }

	// перебор массива из БД.
	for _, p := range banks {
		fmt.Fprintf(w, "<tr><td>"+strconv.Itoa(p.id)+"</td><td>"+p.text+"</td><td>"+p.description+"</td><td>"+p.keywords+"</td>")
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
		text string
		description string
		keywords string
	}
	var p bank
	err := db.QueryRow("SELECT id, text, description, keywords FROM individuals WHERE id = ?", id).Scan(&p.id, &p.text, &p.description, &p.keywords)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fmt.Fprintf(w, `<form action="/update" method="post">`)
	fmt.Fprintf(w, `<input type="hidden" name="id" value="%d">`, p.id)
	fmt.Fprintf(w, `Text: <input type="text" name="col1" value="%s"><br>`, p.text)
	fmt.Fprintf(w, `Description: <input type="text" name="col2" value="%s"><br>`, p.description)
	fmt.Fprintf(w, `Keywords: <input type="text" name="col3" value="%s"><br>`, p.keywords)
	fmt.Fprintf(w, `<input type="submit" value="Update">`)
	fmt.Fprintf(w, `</form>`)
	fmt.Fprintf(w, `<br><a style="background-color: #33bee1; color: #fff; padding: 2px 3px; text-decoration: none" href="/">Return to main menu</a>`)
}