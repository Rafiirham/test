package controllers

import (
	"crud-Articles/models"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetArticless(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		sqlQuery := "SELECT id, title, body from Articless order by id"
		rows, err := db.Query(sqlQuery)
		if err != nil {
			fmt.Println(err)
		}
		defer rows.Close()
		result := models.Articles{}

		for rows.Next() {
			article := models.Article{}
			err2 := rows.Scan(&article.Id, &article.Title, &article.Body)
			if err2 != nil {
				return err2
			}
			result.Articles = append(result.Articles, article)
		}
		return c.JSON(http.StatusCreated, result)
	}
}

func AddArticles(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		Articles := models.Articles{}
		if err := c.Bind(&Articles); err != nil {
			return err
		}
		Articles.Title = c.FormValue("title")
		Articles.Author_id, _ = strconv.Atoi(c.FormValue("author_id"))
		Articles.Body = c.FormValue("body")

		sqlQuery := "INSERT INTO Articless (title, author_id , body , created_at) VALUES ($1, $2, $3, $4)"
		res, err := db.Query(sqlQuery, Articles.Title, Articles.Author_id, Articles.Body, Articles.Created_at)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, Articles)
		}
		return c.String(http.StatusOK, "ok")
	}
}

func UpdateArticles(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		Articles := models.Articles{}
		if err := c.Bind(&Articles); err != nil {
			return err
		}
		id := c.Param("id")
		fmt.Println(id)
		Articles.Title = c.FormValue("title")
		Articles.Author_id, _ = strconv.Atoi(c.FormValue("author_id"))
		Articles.Body = c.FormValue("body")

		sqlQuery := "UPDATE Articless SET title=$1, rating=$2 WHERE id=$3"
		res, err := db.Query(sqlQuery, Articles.Title, Articles.Author_id, Articles.Body, id)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, Articles)
		}
		return c.String(http.StatusOK, id+" Updated")
	}
}

func DeleteArticles(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		sqlQuery := "DELETE FROM Articless WHERE id=$1"
		res, err := db.Query(sqlQuery, id)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, "Deleted")
		}
		return c.String(http.StatusOK, id+" Deleted")
	}
}
