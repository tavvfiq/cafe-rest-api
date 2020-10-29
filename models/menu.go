package models

import (
	"context"

	"fmt"
	"log"

	"strconv"

	d "github.com/tavvfiq/cafe-rest-api/database"

	tb "github.com/tavvfiq/cafe-rest-api/interfaces"
)

// GetMenu get menu from database with pagination and filter
func GetMenu(c context.Context, query map[string]string) ([]tb.Menu, error) {
	am := []tb.Menu{}
	var err error

	// convert page int
	page, err := strconv.ParseInt(query["page"], 10, 32)
	if err != nil {
		log.Fatal("convert to int error")
		return []tb.Menu{}, err
	}
	// convert limit to int
	limit, err := strconv.ParseInt(query["limit"], 10, 32)
	if err != nil {
		log.Fatal("convert to int error")
		return []tb.Menu{}, err
	}
	// offset calculation based on page and limit
	offset := (page - 1) * limit
	// concat search query  with '%query["search"]%'
	search := "'%" + query["search"] + "%'"

	rows, err := d.DB.QueryContext(c, fmt.Sprintf("SELECT menu.id, menu.name, menu.image_path, menu.price, menu.quantity, category.category, menu.added_at, menu.updated_at FROM menu JOIN category ON menu.category_id = category.id WHERE menu.name LIKE %s AND category_id= ? ORDER BY menu.%s %s LIMIT ? OFFSET ?", search, query["sortby"], query["order"]), query["filter"], limit, offset)

	defer rows.Close()
	if err != nil {
		log.Fatal(err)
		return []tb.Menu{}, err
	}

	for rows.Next() {
		m := tb.Menu{}
		err := rows.Scan(&m.ID, &m.Name, &m.Image, &m.Price, &m.Quantity, &m.Category, &m.AddedAt, &m.UpdatedAt)
		if err != nil {
			log.Fatal("error in sql data acquisition")
			return []tb.Menu{}, err
		}
		am = append(am, m)
	}

	return am, nil
}
