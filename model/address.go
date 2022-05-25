package model

import "log"

type Address struct {
	ID        uint64
	UserID    uint64
	Consignee string
	Phone     string
	Province  string
	City      string
	County    string
	TownShip  string
	Detail    string
}

type Addresses []Address

func GetAddresses(userID uint64) (addresses Addresses, err error) {
	query := `select id, consignee, phone, province, city, county, township, detail
		from address
		where user_id = ?`
	rows, err := db.Query(query, userID)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var address Address
		err := rows.Scan(&address.ID, &address.Consignee, &address.Phone, &address.Province, &address.City, &address.County, &address.TownShip, &address.Detail)
		if err != nil {
			log.Println(err)
		}
		addresses = append(addresses, address)
	}

	return
}

func GetAddress(id uint64) (address Address, err error) {
	address.ID = id

	query := `select user_id, consignee, phone, province, city, county, township, detail
	from address
	where id = ?`

	err = db.QueryRow(query, id).
		Scan(&address.UserID,
			&address.Consignee,
			&address.Phone,
			&address.Province,
			&address.City,
			&address.County,
			&address.TownShip,
			&address.Detail)
	return
}
