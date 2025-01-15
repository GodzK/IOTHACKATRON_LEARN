package ultrasonic
import (
"database/sql"
)
type UltrasonicRepo struct {
db *sql.DB
}
func NewUltrasonicRepo(db *sql.DB) *UltrasonicRepo {
return &UltrasonicRepo{
db: db,
}
}
func (u *UltrasonicRepo) GetAll() ([]Ultrasonic, error) {
rows, err := u.db.Query("SELECT * FROM ultrasonic")
if err != nil {
return nil, err
}
defer rows.Close()
ultrasonics := []Ultrasonic{}
for rows.Next() {
var ultrasonic Ultrasonic
err := rows.Scan(&ultrasonic.DateTimestamp,
&ultrasonic.Value)
if err != nil {
return nil, err
}
ultrasonics = append(ultrasonics, ultrasonic)
}
return ultrasonics, nil
}
func (u *UltrasonicRepo) Insert(data Ultrasonic) error {
	_, err := u.db.Exec("INSERT INTO ultrasonic (date_time, value)VALUES (?, ?)", data.DateTimestamp, data.Value)
	if err != nil {
	return err
	}
	return nil
	}