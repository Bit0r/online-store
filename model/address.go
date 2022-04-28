package model

type Address struct {
	ID        uint64
	UserID    uint64
	Name      string
	Recipient string
	Phone     string
	Province  string
	City      string
	County    string
	TownShip  string
	Detail    string
}

func GetAddresses(userID uint64) {

}

func GetAddress(id uint64) (address Address, err error) {
	address.ID = id

	query := `select user_id, name, recipient, phone, province, city, county, township, detail
	from address
	where id = ?`

	err = db.QueryRow(query, id).
		Scan(&address.UserID,
			&address.Name,
			&address.Recipient,
			&address.Phone,
			&address.Province,
			&address.City,
			&address.County,
			&address.TownShip,
			&address.Detail)
	return
}
